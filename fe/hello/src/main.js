import Vue from 'vue'
import App from './App.vue'
import router from './router'
import AxiosPlugin from './assets/AxiosPlugin.js'

Vue.config.productionTip = false
Vue.use(AxiosPlugin)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
