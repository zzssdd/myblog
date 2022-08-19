import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login'
import Admin from '../views/Admin'
import AddArt from '../components/article/AddArt'
import ArtList from '../components/article/ArtList'
import Category from '../components/cate/Category'
import Comment from '../components/comment/Comment'
import Lcomment from '../components/comment/Lcomment'
import Board from '../components/board/Board'
import LifeList from '../components/Life/LifeList'
import AddLife from '../components/Life/AddLife'
import Project from '../components/project/Project'
import AddProject from '../components/project/AddProject'

Vue.use(VueRouter)

const routes = [
  {
    path:'/',
    redirect:'/login',
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta:{
      title:'登录'
    }
  },
  {
    path:'/admin',
    name:'admin',
    component:Admin,
    children:[
      {
        path:'/addart',
        component:AddArt,
        meta:{
          title:'文章管理'
        }
      },
      {
        path:'/addart/:id',
        component:AddArt,
        props:true,
        meta:{
          title:'文章管理'
        }
      },
      {
        path:'/artlist',
        component:ArtList,
        meta:{
          title:'文章列表'
        }
      },
      {
        path:'/addlife/',
        component:AddLife,
        props:true,
        meta:{
          title:'动态管理'
        }
      },
      {
        path:'/addlife/:id',
        component:AddLife,
        props:true,
        meta:{
          title:'动态管理'
        }
      },
      {
        path:'/prolist',
        component:Project,
        meta:{
          title:'项目列表'
        }
      },
      {
        path:'/addpro/',
        component:AddProject,
        props:true,
        meta:{
          title:'项目管理'
        }
      },
      {
        path:'/addpro/:id',
        component:AddProject,
        props:true,
        meta:{
          title:'项目管理'
        }
      },
      {
        path:'/lifelist',
        component:LifeList,
        meta:{
          title:'编辑列表'
        }
      },
      {
        path:'/category',
        component:Category,
        meta:{
          title:'文章分类'
        }
      },
      {
        path:'/comment',
        component:Comment,
        meta:{
          title:'文章评论'
        }
      },
      {
        path:'/lifecomment',
        component:Lcomment,
        meta:{
          title:'生活评论'
        }
      },
      {
        path:'/board',
        component:Board,
        meta:{
          title:'留言板'
        }
      }
    ]
  }
]

const router = new VueRouter({
  mode:'history',
  routes
})

const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push (location) {

return originalPush.call(this, location).catch(err => err)

}


router.beforeEach((to, from, next) => {

  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()

  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    next('/login')
  } else {
    next()
  }
})



export default router
