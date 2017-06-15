import Vue from 'vue'
import VueResource from 'vue-resource'
import App from './App.vue'
import HttpRequests from './components/HttpRequests.vue'

Vue.component('app-http-requests', HttpRequests)

Vue.use(VueResource)

new Vue({
  el: '#app',
  render: h => h(App)
})
