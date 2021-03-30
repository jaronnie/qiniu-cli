import Vue from 'vue'
import Router from 'vue-router'
import upload from '../components/upload'
import login from '../view/login/login'
import Dashboard from "../view/dashboard/Dashboard"

Vue.use(Router)


export default new Router ({
  routes: [
    {
      path: '/upload',
      name: 'upload',
      component: upload
    },
    {
      path: '/login',
      name: 'login',
      component: login
    },
    {
      path: '/Dashboard',
      name: 'Dashboard',
      component: Dashboard
    },
    {
      path: '',
      redirect: Dashboard
    },
  ]
  }
)
