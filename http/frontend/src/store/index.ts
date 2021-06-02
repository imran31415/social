import { createStore } from "vuex";
import createPersistedState from 'vuex-persistedstate'

export default createStore({
  state: {
    username: '',
    password: '',
    newPost: '',
    userId: 0,
    loggedIn: false,
    posts : [],
    feed: [],
  },
  mutations: {},
  actions: {},
  modules: {},
  plugins: [createPersistedState({
    storage: window.sessionStorage,
  })],
});
