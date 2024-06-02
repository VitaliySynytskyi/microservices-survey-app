import Vue from 'vue'
import App from './App.vue'
import router from './router'

// Disable production tip in the console
Vue.config.productionTip = false

// Create a new Vue instance, inject the router, and render the App component
new Vue({
  router,
  render: h => h(App)
}).$mount('#app') // Mount the app to the element with the id "app"
