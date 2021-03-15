<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入Http Method</el-tag>
        <el-select v-model="input.method" class="filter-item" placeholder="请选择版本" filterable>
          <el-option label="PUT【增加】" value="PUT" />
          <el-option label="GET【查询】" value="GET" />
          <el-option label="DELETE【删除】" value="DELETE" />
          <el-option label="POST【修改】" value="POST" />
          <el-option label="HEAD" value="HEAD" />
        </el-select>
        <el-tag class="filter-item">请输入Path</el-tag>
        <el-autocomplete
          ref="autocomplete"
          v-model="input.path"
          clearable
          class="filter-item"
          placeholder="请输入内容"
          style="width: 900px"
          :fetch-suggestions="querySearch"
          @clear="clear"
          @keyup.enter.native="run"
          @select="mySelect"
        />

        <el-button class="filter-item" type="success" @click="run">RUN-></el-button>
      </div>
      <json-editor v-model="input.body" styles="width: 30%" :point-out="pointOut" :read="false" title="请求Body" />
      <json-editor v-model="resData" styles="width: 70%" :read="true" title="返回信息" />
    </el-card>
  </div>
</template>

<script>
import { RunDslAction } from '@/api/es'
import { filterData } from '@/utils/table'
import { esPathKeyWords, esBodyKeyWords } from '@/utils/base-data'
export default {
  name: 'Index',
  components: {
    'JsonEditor': () => import('@/components/JsonEditor/index')
  },
  data() {
    return {
      queryData: esPathKeyWords,
      pointOut: esBodyKeyWords,
      address: 'test',
      input: {
        body: '{}',
        method: 'GET',
        path: ''
      },
      resData: '{}'
    }
  },
  methods: {
    clear() {
      console.log(111)
      this.$refs.autocomplete.activated = true
      this.$refs.autocomplete.handleFocus()
    },
    mySelect(obj) {
      this.input.path = obj.data
    },
    querySearch(queryString, cb) {
      if (queryString == '') {
        cb(this.queryData)
        return
      }

      const queryData = filterData(this.queryData, queryString)
      cb(queryData)
    },
    run() {
      const input = this.input

      if (input['path'].trim() != '') {
        input['path'] = '/' + input['path']
      }

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
