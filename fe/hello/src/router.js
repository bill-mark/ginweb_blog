import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Register from './views/register.vue'
import Login from './views/login.vue'
import uppassword from './views/up_password.vue'
import school from './views/school.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/update_password',
      name: 'update_password',
      component: uppassword,
    },
    {
      path: '/school',
      name: 'school',
      component: school,
    },
  ]
})
