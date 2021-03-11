<template>
  <div id="app">
    <keep-alive>
      <router-view />
    </keep-alive>
  </div>
</template>

<script>

export default {
  name: 'App',
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
  }
}
</script>
