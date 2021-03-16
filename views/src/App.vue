<template>
  <div id="app">
    <keep-alive v-if="isRouterAlive">
      <router-view />
    </keep-alive>
  </div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      isRouterAlive: true // 控制视图是否显示的变量
    }
  },
  created() {
    // 在页面加载时读取sessionStorage里的状态信息
    if (sessionStorage.getItem('EsConnect')) {
      const EsConnect = JSON.parse(sessionStorage.getItem('EsConnect'))
      console.log(EsConnect)
      this.$store.dispatch('baseData/SetEsConnect', EsConnect)
    }

    // 在页面刷新时将vuex里的信息保存到sessionStorage里
    window.addEventListener('beforeunload', () => {
      const EsConnect = JSON.stringify(this.$store.state.baseData.EsConnect)
      sessionStorage.setItem('EsConnect', EsConnect)
    })
  },
  provide() { // 父组件中通过provide来提供变量，在子组件中通过inject来注入变量。
    return {
      reload: this.reload
    }
  },
  methods: {
    reload() {
      console.log('reload')
      this.isRouterAlive = false // 先关闭，
      this.$nextTick(() => {
        this.isRouterAlive = true
      })
    }
  }
}
</script>
