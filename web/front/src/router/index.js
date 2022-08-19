import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ViewPost from '../views/ViewPost.vue'
import ViewTag from '../views/ViewTag.vue'
import About from '../views/About.vue'
import Resume from '../views/Resume.vue'
import ViewProjects from '../views/ViewProjects.vue'
import Life from '../views/Life.vue'
import LifeInfo from '../views/LifeInfo.vue'
import ViewPosts from '../views/ViewPosts.vue'
import NotFound from '../views/errors/NotFound.vue'
import ViewTags from '../views/ViewTags.vue'
import Message from '../views/Message.vue'

// let Posts = null
// try {
//   Posts = require('@/../posts/data/posts.json')
// } catch (e) {
//   Posts = require('@/defaults/posts.json')
// }

// Posts.posts.map(post => {
//   children.push({
//     path: post.id,
//     component: async function () {
//       let value
//       // await import(`@/../posts/${post.id}.md`).then((val) => {
//       //   value = val
//       // })
//       return value.vue.component
//     }
//   })
// })
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/posts',
    name: 'ViewPosts',
    component: ViewPosts
  },
  {
    path:'/article/:id',
    name:'ArticleInfo',
    component:ViewPost
  },
  {
    path: '/posts',
    name: 'ViewPost',
    component: ViewPost
  },
  {
    path: '/tags/:tag',
    name: 'ViewTag',
    component: ViewTag
  },
  {
    path: '/tags',
    name: 'ViewTag',
    component: ViewTags
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/resume',
    name: 'Resume',
    component: Resume
  },
  {
    path: '/projects',
    name: 'ViewProjects',
    component: ViewProjects
  },
  {
    path: '/life',
    name: 'Life',
    component: Life
  },
  {
    path: '/life/:id',
    name: 'LifeInfo',
    component: LifeInfo
  },
  {
    path:'/board',
    name:'Board',
    component:Message
  },
  {
    path: '*',
    name: 'NotFound',
    component: NotFound
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
