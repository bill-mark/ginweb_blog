import Vue from 'vue'
import App from './App.vue'
import router from './router'
import AxiosPlugin from './assets/AxiosPlugin.js'
import {Pagination} from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false
Vue.use(AxiosPlugin)
Vue.use(Pagination)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
