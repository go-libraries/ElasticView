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
    alwaysShow: true,
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
  }
]

// 路由组件 映射 map
export const RoutesComponentmaps = {
  'layout': Layout,
  'views/dashboard/index': () => import('@/views/dashboard/index'), // 主页
  'views/permission/role': () => import('@/views/permission/role'), // 角色管理
  'views/permission/user': () => import('@/views/permission/user'), // 用户管理
  'views/connect-tree/index': () => import('@/views/connect-tree/index')
}

