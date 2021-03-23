/* Layout */
import Layout from '@/layout'

// 动态路由列表
export const asyncRoutes = [
  {
    path: '/permission',
    component: 'layout',
    redirect: '/permission/index',
    alwaysShow: true,
    meta: {
      title: '权限',
      icon: 'el-icon-user-solid'

    },
    children: [
      {
        path: 'role',
        component: 'views/permission/role',
        name: 'RolePermission',
        meta: {
          title: '角色管理',

          icon: 'el-icon-s-check'
        }
      },
      {
        path: 'user',
        component: 'views/permission/user',
        name: 'user',
        meta: {
          title: '用户管理',

          icon: 'el-icon-user'
        }
      }
    ]
  },
  {
    path: '/connect-tree',
    component: 'layout',
    redirect: '/connect-tree/index',
    alwaysShow: false,
    meta: {
      title: '连接树管理',
      icon: 'el-icon-link'
    },
    children: [
      {
        path: 'index',
        component: 'views/connect-tree/index',
        name: 'index',
        meta: {
          title: '连接树管理',
          icon: 'el-icon-link'
        }
      }
    ]
  },
  {
    path: '/cat',
    component: 'layout',
    redirect: '/cat/index',
    alwaysShow: false,
    meta: {
      title: 'ES状态',
      icon: 'el-icon-pie-chart'
    },
    children: [
      {
        path: 'index',
        component: 'views/cat/index',
        name: 'index',
        meta: {
          title: 'ES状态',
          icon: 'el-icon-pie-chart'
        }
      }
    ]
  },
  {
    path: '/rest',
    component: 'layout',
    redirect: '/rest/index',
    alwaysShow: false,
    meta: {
      title: 'dev Tools',
      icon: 'el-icon-edit'
    },
    children: [
      {
        path: 'index',
        component: 'views/rest/index',
        name: 'index',
        meta: {
          title: 'DSL面板',
          icon: 'el-icon-edit'
        }
      }
    ]
  },
  {
    path: '/indices',
    component: 'layout',
    redirect: '/indices/index',
    alwaysShow: true,
    meta: {
      title: '索引管理',
      icon: 'el-icon-coin'
    },
    children: [
      {
        path: 'index',
        component: 'views/indices/index',
        name: 'index',
        meta: {
          title: '索引管理',
          icon: 'el-icon-coin'
        }
      },
      {
        path: 'reindex',
        component: 'views/indices/reindex',
        name: 'reindex',
        meta: {
          title: '重建索引',
          icon: 'el-icon-document-copy'
        }
      },
      {
        path: 'sql',
        component: 'views/indices/sql',
        name: 'sql',
        meta: {
          title: 'sql',
          icon: 'el-icon-document-copy'
        }
      }
    ]
  }
]

// 路由组件 映射 map
export const RoutesComponentmaps = {
  'layout': Layout,
  'views/dashboard/index': () => import('@/views/dashboard/index'), // 主页
  'views/permission/role': () => import('@/views/permission/role'), // 角色管理
  'views/permission/user': () => import('@/views/permission/user'), // 用户管理
  'views/connect-tree/index': () => import('@/views/connect-tree/index'),
  'views/cat/index': () => import('@/views/cat/index'),
  'views/rest/index': () => import('@/views/rest/index'),
  'views/indices/index': () => import('@/views/indices/index'),
  'views/indices/reindex': () => import('@/views/indices/reindex'),
  'views/indices/sql': () => import('@/views/indices/sql')
}

