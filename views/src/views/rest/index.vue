<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入Http Method</el-tag>
        <el-select v-model="input.method" class="filter-item" placeholder="请选择版本" filterable>
          <el-option label="GET" value="GET" />
          <el-option label="POST" value="POST" />
          <el-option label="PUT" value="PUT" />
          <el-option label="PATCH" value="PATCH" />
          <el-option label="DELETE" value="DELETE" />
          <el-option label="HEAD" value="HEAD" />
          <el-option label="OPTIONS" value="OPTIONS" />
        </el-select>
        <el-tag class="filter-item">请输入Path</el-tag>
        <el-input v-model="input.path" class="filter-item" style="width: 900px" @keyup.enter.native="run" />
        <el-button class="filter-item" type="success" @click="run">RUN-></el-button>
      </div>
      <json-editor v-model="input.body" :read="false" title="请求Body" />
      <json-editor v-model="resData" :read="true" title="返回信息" />
    </el-card>
  </div>
</template>

<script>
import { RunDslAction } from '@/api/es'
export default {
  name: 'Index',
  components: {
    'JsonEditor': () => import('@/components/JsonEditor/index')
  },
  data() {
    return {
      address: 'test',
      input: {
        body: '{}',
        method: 'GET',
        path: '/_cat/indices'
      },
      resData: '{}'
    }
  },
  methods: {
    run() {
      const input = this.input
      input['es_connect'] = this.$store.state.baseData.EsConnect
      RunDslAction(input).then(res => {
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: res.msg
          })
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          })
        }
        this.resData = JSON.stringify(res.data, null, '\t')
      }).catch(err => {
        this.$message({
          type: 'error',
          message: '网络异常'
        })
      })
    }

  }
}
</script>

<style scoped>

</style>
