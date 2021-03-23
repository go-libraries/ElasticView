<template>
  <div>
    <el-button id="guide-menu">
      1
    </el-button>
    <el-button class="collapse-box">
      2
    </el-button>
    <el-button id="guide-breadcrumb">
      3
    </el-button>
    <el-button class="user-content">
      4
    </el-button>
    <el-button id="tagsView">
      5
    </el-button>
  </div>
</template>

<script>
import Driver from 'driver.js' // import driver.js
import 'driver.js/dist/driver.min.css' // import driver.js css
import steps from './guide'

export default {
  name: 'Guide',
  data() {
    return {
      driver: null
    }
  },
  mounted() {
    this.driver = new Driver({
      className: 'scoped-class', // className to wrap driver.js popover
      animate: true, // Animate while changing highlighted element
      opacity: 0.75, // Background opacity (0 means only popovers and without overlay)
      padding: 10, // Distance of element from around the edges
      allowClose: true, // Whether clicking on overlay should close or not
      overlayClickNext: false, // Should it move to next step on overlay click
      doneBtnText: '完成', // Text on the final button
      closeBtnText: '关闭', // Text on the close button for this step
      nextBtnText: '下一步', // Next button text for this step
      prevBtnText: '上一步' // Previous button text for this step
      // Called when moving to next step on any step
    })
    setTimeout(() => {
      this.guide()
    }, 1000)
    // this.driver = new Driver()
  },
  methods: {
    guide() {
      this.driver.defineSteps(steps)
      this.driver.start()
    }
  }
}
</script>

<style scoped>

</style>
