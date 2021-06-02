<template>
  <nav class="navbar navbar-dark bg-dark">
    <div class="col-md-2 offset-md-1"> <a class="navbar-brand" href="#">Social Network</a></div>
  </nav>
  <div id="app" class="container bg-light text-dark border">
    <div class="row" v-if="this.$store.state.loggedIn">
    </div>
    <div class="row" v-else>
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Login</h1>
        <p>if you are a new user, a login will be created.</p>
        <p>If you dont have a login, enter a username and password and one will be created</p>
        <form v-on:submit.prevent="login">
          <div class="form-group">
            <input v-model="this.$store.state.username" type="text" id="username-input" placeholder="Enter a username" class="form-control">
            <input v-model="this.$store.state.password" type="password" id="password-input" placeholder="Enter a password" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Login!</button>
          </div>
        </form>
      </div>
    </div>
    <div v-if="this.$store.state.loggedIn">
      <div class="row">
        <div class="col-md-6 offset-md-3 py-5">
          <h3>{{this.$store.state.username}} Create a Post: </h3>
          <p>Enter some text to create a post</p>
          <form v-on:submit.prevent="createPost">
            <div class="form-group">
              <textarea v-model="this.$store.state.newPost" type="text" id="newPost-input" placeholder="Your Post" class="form-control" rows="3"></textarea>
            </div>
            <div class="form-group">
              <button class="btn btn-primary">Post!</button>
            </div>
          </form>
        </div>
      </div>
      <div v-if="hasFeed()">
        <div class="row" v-for="feedItem in this.$store.state.feed" :key="feedItem.id">
          <div class="col-md-12">
            <div class="container">
              <div class="row d-flex justify-content-center">
                <div class="col-md-8">
                  <div class="d-flex flex-row"></div>
                  <div class="row news-card p-3 bg-white">
                    <div class="col-md-4">
                      <div class="feed-image"><img class="news-feed-image rounded img-fluid img-responsive" src="https://i.imgur.com/EGa6hnF.jpg"></div>
                    </div>
                    <div class="col-md-8">
                      <div class="news-feed-text">
                        <span>{{feedItem.id}} {{feedItem.postContent}}}</span>
                        <div class="d-flex flex-row justify-content-between align-items-center mt-2">
                          <div class="d-flex creator-profile">
                            <img class="rounded-circle" src="https://i.imgur.com/EGa6hnF.jpg" width="50" height="50">
                            <div class="d-flex flex-column ml-2">
                              <h6 class="username">{{ this.$store.state.username }}</h6>
                              <span class="date">Jan 20,2020</span>
                            </div>
                          </div>
                          <i class="fa fa-share share"></i>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import {SocialPromiseClient} from './static/js/social_grpc_web_pb'
import {CreateUserReq, GetPostsReq, GETPOSTSIDTYPE_USER, CreatePostReq, GetFeedReq} from './static/js/social_pb'


import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'


export default {
  name: 'App',
  created: async function() {
    this.grpcClient = new SocialPromiseClient("http://localhost:8083", null, null)
    if(localStorage.loggedIn) this.$store.state.loggedIn = localStorage.loggedIn;
    if(localStorage.userId) this.$store.state.userId = localStorage.userId;
    if(localStorage.loggedIn) await this.getFeed();

  },
  data() { return {} },
  methods: {
     async login() {
       try {
         const user = new CreateUserReq()
         user.setUserName(this.$store.state.username)
         user.setPassword(this.$store.state.password)
         const res = await this.grpcClient.createUser(user, {})
         console.log("Login Successful")
         this.$store.state.loggedIn = true
         this.$store.state.userId = res.toObject().id
         localStorage.loggedIn = this.$store.state.loggedIn;
         localStorage.userId = this.$store.state.userId;
         await this.getFeed()

       } catch (err) {
         console.error(err.message)
         console.log("err in grpc response: ", err.message)
         this.$store.state.username = ''
         this.$store.state.feed = ''
         this.$store.state.password = ''
         this.$store.state.userId = 0

         throw err
       }
     },
    async createPost() {
      if (this.$store.state.userId !== 0 && this.$store.state.loggedIn && this.$store.state.newPost !== '') {
        try {
          const post = new CreatePostReq()
          post.setContent("{\"body\":\""+  this.$store.state.newPost +"\"}")
          post.setUserId(this.$store.state.userId)
          const res = await this.grpcClient.createPost(post, {})
          console.log("createPost Successful")
          // wait 5 seconds for feed item to propagate to users feed
          await new Promise(r => setTimeout(r, 5000));
          await this.getFeed();

        } catch (err) {
          console.error(err.message)
          console.log("err in grpc response: ", err.message)
          throw err
        }
      }
    },
    async getUserPosts() {
      if (this.$store.state.userId !== 0 && this.$store.state.loggedIn) {
        try {
          const postReq = new GetPostsReq()
          postReq.setGetBy(GETPOSTSIDTYPE_USER)
          let idList = []
          idList.push (this.$store.state.userId)
          postReq.setIdsList(idList)
          const res = await this.grpcClient.getPosts(postReq, {})
          console.log("getUserPosts successful")
          this.$store.state.posts = res.toObject().itemsList

        } catch (err) {
          console.error(err.message)
          console.log("err in grpc response: ", err.message)
          throw err
        }
      }
    },
    async getFeed() {
      if (this.$store.state.userId !== 0 && this.$store.state.loggedIn) {
        try {
          const postReq = new GetFeedReq()
          postReq.setOwnerId(this.$store.state.userId)
          const res = await this.grpcClient.getFeed(postReq, {})
          console.log("getFeed successful")
          this.$store.state.feed = res.toObject().itemsList
        } catch (err) {
          console.error(err.message)
          console.log("err in grpc response: ", err.message)
          throw err
        }
      }
    },
    hasPosts() {
       return this.$store.state.posts.length > 0;
    },
    hasFeed() {
      return this.$store.state.feed.length > 0;
    },

  }
}
</script>


<style>
@import url('https://fonts.googleapis.com/css2?family=Manrope&display=swap');

body {
  background-color: #eee;
  font-family: 'Manrope', sans-serif
}

.news-card {
  border-radius: 8px
}

.news-feed-image {
  border-radius: 8px;
  width: 100%
}

.date {
  font-size: 12px
}

.username {
  color: blue
}

.share {
  color: blue
}
</style>