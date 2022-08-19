import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Vuesax from 'vuesax'
import 'vuesax/dist/vuesax.css'
import 'boxicons'
import 'boxicons/css/boxicons.min.css'
import 'font-awesome/css/font-awesome.min.css'
import axios from 'axios'

let Url='http://43.139.46.112:3000/api/'

axios.defaults.baseURL=Url


Vue.config.productionTip = false
Vue.use(Vuesax, {})

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios

Vue.mixin({
  methods: {
    changeTitle: function (title) {
      document.title = `${title} - Yogen的博客`
    },
  },
})

router.afterEach((to, from, next) => {
  window.scrollTo(0, 0)
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
