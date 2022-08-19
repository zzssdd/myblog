import Vue from "vue"
import axios from "axios"

let Url='http://43.139.46.112:3000/api/'

axios.defaults.baseURL=Url

axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
  })
  

Vue.prototype.$http=axios

export {Url}