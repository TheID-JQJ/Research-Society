import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/home',
    children: [{
      path: 'home',
      name: 'Home',
      component: () => import('@/views/home/index'),
      meta: { title: '首页', icon: 'el-icon-s-home' }
    }]
  },

  {
    path: '/roomManagement',
    component: Layout,
    redirect: '/roomManagement/seat',
    name: 'RoomManagement',
    meta: { title: '自习室管理', icon: 'el-icon-s-management' },
    children: [
      {
        path: 'seat',
        name: 'Seat',
        component: () => import('@/views/room/seat/index'),
        meta: { title: '座位管理' }
      },
      {
        path: 'resource',
        name: 'Resource',
        component: () => import('@/views/room/resource/index'),
        meta: { title: '公共资源管理' }
      },
      {
        path: 'sort',
        name: 'ResourceSort',
        component: () => import('@/views/room/sort/index'),
        meta: { title: '资源种类管理' }
      }
    ]
  },

  {
    path: '/userManagement',
    component: Layout,
    redirect: '/userManagement/account',
    name: 'UserManagement',
    meta: { title: '用户管理', icon: 'el-icon-user-solid' },
    children: [
      {
        path: 'account',
        name: 'Account',
        component: () => import('@/views/user/account/index'),
        meta: { title: '账户管理' }
      },
      {
        path: 'message',
        name: 'UserMessage',
        component: () => import('@/views/user/message/index'),
        meta: { title: '用户信息管理' }
      },
      {
        path: 'friend',
        name: 'Friend',
        component: () => import('@/views/user/friend/index'),
        meta: { title: '好友管理' }
      },
      {
        path: 'sort',
        name: 'TargetSort',
        component: () => import('@/views/user/sort/index'),
        meta: { title: '目标类别管理' }
      }
    ]
  },

  {
    path: '/commodityManagement',
    component: Layout,
    redirect: '/commodityManagement/message',
    name: 'CommodityManagement',
    meta: {
      title: '商品管理',
      icon: 'el-icon-s-goods'
    },
    children: [
      {
        path: 'message',
        name: 'CommodityMessage',
        component: () => import('@/views/commodity/message/index'),
        meta: { title: '商品信息管理' }
      },
      {
        path: 'sort',
        name: 'CommoditySort',
        component: () => import('@/views/commodity/sort/index'),
        meta: { title: '商品类别管理' }
      }
    ]
  },

  {
    path: '/moneyManagement',
    component: Layout,
    redirect: '/moneyManagement/money',
    name: 'MoneyManagement',
    meta: { title: '财务管理', icon: 'el-icon-s-finance' },
    children: [
      {
        path: 'money',
        name: 'Money',
        component: () => import('@/views/finance/money/index'),
        meta: { title: '财务管理' }
      }
    ]
  },

  // {
  //   path: '/form',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'index',
  //       name: 'Form',
  //       component: () => import('@/views/form/index'),
  //       meta: { title: 'Form', icon: 'form' }
  //     }
  //   ]
  // },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
