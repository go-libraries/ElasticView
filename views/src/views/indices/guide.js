// guide.js
const steps = [
  {
    element: '#guide-menu',
    popover: {
      title: '菜单导航',
      description: '点击菜单可切换右侧菜单内容',
      position: 'right'
    }
  },
  {
    element: '.collapse-box',
    popover: {
      title: '汉堡包',
      description: '点击收缩和展开菜单导航',
      position: 'bottom'
    },
    padding: 5
  },
  {
    element: '#guide-breadcrumb',
    popover: {
      title: '面包屑导航',
      description: '用于显示 当前导航菜单的位置',
      position: 'bottom'
    }
  },
  {
    element: '.user-content',
    popover: {
      title: '个人中心',
      description: '点击可操作',
      position: 'left'
    }
  },
  {
    element: '#tagsView',
    popover: {
      title: '最近打开任务',
      description: '最近打开任务菜单，点击可快速切换',
      position: 'bottom'
    }
  }
]

export default steps
