<template>
  <div class="header-search">
    <el-select
      v-model="link"
      filterable
      default-first-option
      placeholder="请选择ES连接"
      @change="change"
    >

      <el-option :value="Number(0)" label="请选择ES连接" />
      <el-option v-for="item in opt" :key="item.id" :value="Number(item.id)" :label="item.remark" />
    </el-select>
  </div>
</template>

<script>
// fuse is a lightweight fuzzy-search module
// make search results more in line with expectations
import { ListAction } from '@/api/es-link'

export default {
  inject: ['reload'],
  name: 'HeaderSearch',
  data() {
    return {
      opt: [],
      link: '',
      time: null,
      timeSecend: 45
    }
  },
  computed: {},
  watch: {},
  mounted() {
    const obj = this.$store.state.baseData.EsConnectID
    this.link = Number(obj)
    this.getEsOpt()
    this.startLoop()
  },

  beforeDestroy() {
    // 清除定时器
    clearInterval(this.time)
    this.time = null
    console.log('beforeDestroy')
  },

  methods: {
    startLoop() {
      this.time = setInterval(() => {
        this.getEsOpt()
      }, this.timeSecend * 1000)
    },
    async getEsOpt() {
      const res = await ListAction()
      this.opt = res.data
    },
    change(link) {
      this.$store.dispatch('baseData/SetEsConnect', link)
      this.reload()
      // console.log(this.$store.state.baseData.EsConnectID)
    }
  }

}
</script>

<style lang="scss" scoped>
  .header-search {
    font-size: 0 !important;

    .search-icon {
      cursor: pointer;
      font-size: 18px;
      vertical-align: middle;
    }

    .header-search-select {
      font-size: 18px;
      transition: width 0.2s;
      width: 0;
      overflow: hidden;
      background: transparent;
      border-radius: 0;
      display: inline-block;
      vertical-align: middle;

      /deep/ .el-input__inner {
        border-radius: 0;
        border: 0;
        padding-left: 0;
        padding-right: 0;
        box-shadow: none !important;
        border-bottom: 1px solid #d9d9d9;
        vertical-align: middle;
      }
    }

    &.show {
      .header-search-select {
        width: 210px;
        margin-left: 10px;
      }
    }
  }
</style>
