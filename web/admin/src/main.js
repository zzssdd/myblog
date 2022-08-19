import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/ui'
import './plugin/http'
import './plugin/edit'
import tinymce from 'tinymce'
import VueTinymce from '@packy-tang/vue-tinymce'

Vue.config.productionTip = false

Vue.prototype.$tinymce = tinymce // 将全局tinymce对象指向给Vue作用域下
Vue.use(VueTinymce)              // 安装vue的tinymce组件

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
