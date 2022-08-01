/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const userRouter = {
  path: '/account',
  component: Layout,
  redirect: '/account/list',
  meta: {title: '用户权限', icon: 'el-icon-unlock'},
  children: [
    {
      path: 'list',
      name: 'UserList',
      component: () => import('@v/account/list'),
      meta: {title: '账户列表'}
    },
  ]
}

export default userRouter
