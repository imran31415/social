<template>

  <div id="app" class="container" >
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Welcome to the social network</h1>
      </div>
    </div>
    <div class="row" v-if="loggedIn">
      <div class="col-md-6 offset-md-3 py-5">
        <h2>You are logged in, welcome {{this.username}}</h2>
      </div>
    </div>
    <div class="row" v-else>
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Login to Social Network</h1>
        <p>If you dont have a login, enter a username and password and one will be created</p>
        <form v-on:submit.prevent="login">
          <div class="form-group">
            <input v-model="username" type="text" id="username-input" placeholder="Enter a username" class="form-control">
            <input v-model="password" type="password" id="password-input" placeholder="Enter a password" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Login!</button>
          </div>
        </form>
      </div>
    </div>
    <div v-if="loggedIn">
      <div class="row">
        <div class="col-md-6 offset-md-3 py-5">
          <h3>{{this.username}} Create a Post: </h3>
          <p>Enter some text to create a post</p>
          <form v-on:submit.prevent="createPost">
            <div class="form-group">
              <input v-model="newPost" type="text" id="newPost-input" placeholder="Your Post" class="form-control">
            </div>
            <div class="form-group">
              <button class="btn btn-primary">Post!</button>
            </div>
          </form>
        </div>
      </div>
      <div class="row" >
        <div class="col-md-6 offset-md-3 py-5" v-if="posts" >
          <h2>{{this.username}}'s Posts: </h2>
          <div class="row" v-for="post in posts" :key="post.id">
            {{post.id}}: {{post.content}}
          </div>
        </div>
        <div class="col-md-6 offset-md-3 py-5" v-else>
          <h2> {{this.username}} has no posts in the social network </h2>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import {SocialPromiseClient} from './static/js/social_grpc_web_pb'
import {CreateUserReq, GetPostsReq, GETPOSTSIDTYPE_USER, CreatePostReq} from './static/js/social_pb'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
export default {
  name: 'App',

  created: function() {
    this.grpcClient = new SocialPromiseClient("http://localhost:8083", null, null)
  },

  data() { return {
    username: '',
    password: '',
    newPost: '',
    userId: 0,
    loggedIn: false,
    posts : [],
  } },

  methods: {
     async login() {
       try {
         const user = new CreateUserReq()
         user.setUserName(this.username)
         user.setPassword(this.password)
         const res = await this.grpcClient.createUser(user, {})
         console.log("Login Successful")
         this.loggedIn = true
         this.userId = res.toObject().id
         await this.getUserPosts()
         this.$forceUpdate();

       } catch (err) {
         console.error(err.message)
         console.log("err in grpc response: ", err.message)
         this.username = ""
         this.password = ""
         this.userId = 0
         throw err
       }
     },
    async createPost() {
      if (this.userId !== 0 && this.loggedIn && this.newPost !== '') {
        try {
          const post = new CreatePostReq()
          post.setContent("{\"body\":\""+  this.newPost +"\"}")
          post.setUserId(this.userId)
          const res = await this.grpcClient.createPost(post, {})
          console.log("createPost Successful")
          await this.getUserPosts();

        } catch (err) {
          console.error(err.message)
          console.log("err in grpc response: ", err.message)
          throw err
        }
      }
    },
    async getUserPosts() {
      if (this.userId !== 0 && this.loggedIn) {
        try {
          const postReq = new GetPostsReq()
          postReq.setGetBy(GETPOSTSIDTYPE_USER)
          let idList = []
          idList.push (this.userId)
          postReq.setIdsList(idList)
          const res = await this.grpcClient.getPosts(postReq, {})
          console.log("getUserPosts successful")
          this.posts = res.toObject().itemsList

        } catch (err) {
          console.error(err.message)
          console.log("err in grpc response: ", err.message)
          throw err
        }
      }
    }

  }
}
</script>