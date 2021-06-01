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
        <form v-on:submit.prevent="createUser">
          <div class="form-group">
            <input v-model="username" type="text" id="username-input" placeholder="Enter a username" class="form-control">
            <input v-model="password" type="password" id="password-input" placeholder="Enter a password" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Create!</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import {SocialPromiseClient} from './static/js/social_grpc_web_pb'
import {CreateUserReq, User} from './static/js/social_pb'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
export default {
  name: 'App',

  created: function() {
    this.grpcClient = new SocialPromiseClient("http://localhost:8083", null, null);
  },

  data() { return {
    username: '',
    password: '',
    userId: 0,
    loggedIn: false,
  } },

  methods: {
     async createUser() {
       console.log("Creating user: " + this.username)

       try {
         const user = new CreateUserReq()
         user.setUserName(this.username)
         user.setPassword(this.password)
         const res = await this.grpcClient.createUser(user, {})
         console.log("Successfully received GRPC response, object returned is: " + JSON.stringify(res))
         this.loggedIn = true
         this.$forceUpdate()

       } catch (err) {
         console.error(err.message)
         console.log("err in grpc response: ", err.message);
         throw err
       }
     }
  }
}
</script>