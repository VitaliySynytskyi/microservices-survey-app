import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import SurveyForm from '../forms/Survey.vue'
import Survey from '../views/Survey.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { 
      title: 'Home - Survey App',
      transition: 'fade'
    }
  },
  {
    path: '/new',
    name: 'SurveyForm',
    component: SurveyForm,
    meta: {
      title: 'Create New Survey',
      transition: 'slide'
    }
  },
  {
    path: '/survey/:id',
    name: 'Survey',
    component: Survey,
    meta: {
      title: 'View Survey',
      transition: 'slide'
    }
  },
  {
    path: '*',
    redirect: '/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  }
})

// Update page title based on route
router.beforeEach((to, from, next) => {
  document.title = to.meta.title || 'Survey App'
  next()
})

export default router
