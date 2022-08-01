import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Layout from '@/layout'
import RBACRouter from './modules/rbac-manage'
import AccountRouter from './modules/account'

export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@v/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@v/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@v/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@v/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@v/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@v/dashboard/index'),
        meta: { title: '仪表盘', icon: 'el-icon-location-information', affix: true }
      }
    ]
  },
  {
    path: '/profile',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '',
        component: () => import('@v/profile/index'),
        name: 'Profile',
        meta: { title: '个人资料', icon: 'user', noCache: true }
      }
    ]
  },
  // http://localhost:19527/marketing/callback
  {
    path: '/marketing/callback',
    component: () => import('@v/marketing/callback'),
    hidden: true
  }
  /*{
    path: '/test',
    component: Layout,
    children: [
      {
        path: '',
        component: () => import('@v/test/test'),
        name: 'TestPage',
        meta: { title: '测试', icon: 'question', noCache: true }
      }
    ]
  }*/
]

// 路由规则：
// 权限控制为 meta 中的 auth 属性，填写规则：后端路由去掉前缀 [/api/]
// 如果未设置 auth 属性，表示无需权限都可以访问
export const asyncRoutes = [
  AccountRouter,
  RBACRouter,
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history',
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
