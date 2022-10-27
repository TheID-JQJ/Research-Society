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
    redirect: '/userManagement/appoint',
    name: 'UserManagement',
    meta: { title: '用户管理', icon: 'el-icon-user-solid' },
    children: [
      {
        path: 'appoint',
        name: 'Appoint',
        component: () => import('@/views/user/appoint/index'),
        meta: { title: '用户预约管理' }
      },
      {
        path: 'message',
        name: 'UserMessage',
        component: () => import('@/views/user/message/index'),
        meta: { title: '用户信息管理' }
      },
      {
        path: 'target',
        name: 'Target',
        component: () => import('@/views/user/target/index'),
        meta: { title: '用户目标管理' }
      },
      {
        path: 'account',
        name: 'Account',
        component: () => import('@/views/user/account/index'),
        meta: { title: '用户账户管理' }
      },
      {
        path: 'password',
        name: 'Password',
        component: () => import('@/views/user/password/index'),
        meta: { title: '用户密码管理' }
      },
      {
        path: 'order',
        name: 'Order',
        component: () => import('@/views/user/order/index'),
        meta: { title: '用户订单管理' }
      },
      {
        path: 'friend',
        name: 'Friend',
        component: () => import('@/views/user/friend/index'),
        meta: { title: '用户好友管理' }
      },
      {
        path: 'group',
        name: 'Group',
        component: () => import('@/views/user/group/index'),
        meta: { title: '群组管理' }
      },
      {
        path: 'userGroup',
        name: 'UserGroup',
        component: () => import('@/views/user/userGroup/index'),
        meta: { title: '用户群组管理' }
      },
      {
        path: 'donate',
        name: 'Donate',
        component: () => import('@/views/user/donate/index'),
        meta: { title: '用户捐献管理' }
      },
      {
        path: 'role',
        name: 'Role',
        component: () => import('@/views/user/role/index'),
        meta: { title: '用户角色管理' }
      },
      {
        path: 'permission',
        name: 'Permission',
        component: () => import('@/views/user/permission/index'),
        meta: { title: '权限管理' }
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

  {
    path: '/systemManagement',
    component: Layout,
    redirect: '/systemManagement/notice',
    name: 'SystemManagement',
    meta: { title: '系统管理', icon: 'el-icon-s-tools' },
    children: [
      {
        path: 'notice',
        name: 'Notice',
        component: () => import('@/views/system/notice/index'),
        meta: { title: '公告管理' }
      },
      {
        path: 'putNotice',
        name: 'PutNotice',
        component: () => import('@/views/system/form/index'),
        meta: { title: '发布公告' }
      },
      {
        path: 'feedback',
        name: 'Feedback',
        component: () => import('@/views/system/feedback/index'),
        meta: { title: '反馈信息' }
      }
    ]
  },

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
