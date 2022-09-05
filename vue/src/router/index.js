import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

export const constantRoutes = [{
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/register',
    component: () => import('@/views/register/index'),
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
    redirect: '/stipend',
    children: [{
      path: 'stipend',
      name: 'Proofmaterial',
      component: () => import('@/views/proofmaterial/list/index'),
      meta: {
        title: '助学金申请名单',
        icon: '高校助学金'
      }
    }]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/applicationmaterials',
    component: Layout,
    redirect: '/applicationmaterials/text',
    name: 'Applicationmaterials',
    alwaysShow: true,
    meta: {
      roles: ['editor'],
      title: '助学金材料',
      icon: '助学金材料'
    },
    children: [{
      path: 'text',
      name: 'ApplicationmaterialsText',
      component: () => import('@/views/proofmaterial/add/text/index'),
      meta: {
        roles: ['editor'],
        title: '提交申请材料（文本）',
        icon: '提交数据'
      }
    },
    {
      path: 'file',
      name: 'ApplicationmaterialsFile',
      component: () => import('@/views/proofmaterial/add/file/index'),
      meta: {
        roles: ['editor'],
        title: '提交申请材料（文件）',
        icon: '提交文件'
      }
    }, {
      path: 'updatetext',
      name: 'ApplicationmaterialsUpdateText',
      component: () => import('@/views/proofmaterial/add/updatetext/index'),
      meta: {
        roles: ['editor'],
        title: '修改申请材料（文本）',
        icon: '修改数据'
      }
    }]
  },

  {
    path: '/assess',
    component: Layout,
    redirect: '/assess/all',
    name: 'Assess',
    alwaysShow: true,
    meta: {
      title: '助学金评定',
      icon: '助学金评定'
    },
    children: [{
        path: 'allrank',
        name: 'Allrank',
        component: () => import('@/views/assess/rank/index'),
        meta: {
          title: '申请名单详细信息',
          icon: '排名'
        }
      },
      {
        path: 'allaward',
        name: 'AllAward',
        component: () => import('@/views/assess/award/index'),
        meta: {
          title: '获助学金名单',
          icon: '奖杯'
        }
      },
    ]
  },
  {
    path: '/addblacklist',
    component: Layout,
    meta: {
      roles: ['admin']
    },
    children: [{
      path: '/addblacklist',
      name: 'AddBlacklist',
      component: () => import('@/views/assess/vote/index'),
      meta: {
        title: '助后评分',
        icon: '评分'
      }
    }]
  },


  // 404 page must be placed at the end !!!
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const createRouter = () => new Router({
  base: '/web',
  // mode: 'history', // require service support
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router