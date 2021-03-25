<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请选择Http Method</el-tag>
        <el-select v-model="input.method" class="filter-item select-method" placeholder="请选择Http Method" filterable>
          <el-option label="PUT【修改】" value="PUT" />
          <el-option label="GET【查询】" value="GET" />
          <el-option label="DELETE【删除】" value="DELETE" />
          <el-option label="POST【新增】" value="POST" />
          <el-option label="HEAD【是否存在】" value="HEAD" />
        </el-select>
        <el-tag class="filter-item">请输入Path</el-tag>
        <el-autocomplete
          ref="autocomplete"
          v-model="input.path"
          clearable
          class="filter-item select-path autocomplete"
          placeholder="请输入内容"

          :fetch-suggestions="querySearch"
          @clear="clear"
          @keyup.enter.native="run"
          @select="mySelect"
        />

        <el-button
          class="filter-item go"
          style="display: inline;"
          :loading="loading"
          type="success"
          icon="el-icon-right"
          @click="go"
        >GO
        </el-button>
        <el-button
          v-show="input.method == 'GET'"
          class="filter-item sql-format"
          style="display: inline;"
          type="warning"
          icon="el-icon-refresh"
          @click="openDrag"
        >SQL转换
        </el-button>
        <el-button
          class="filter-item search-history"
          style="display: inline;"
          type="warning"
          icon="el-icon-search"
          @click.native="dialogVisible = true"
        >
          搜索历史
        </el-button>
        <download-excel
          v-if="canExport"
          ref="download"
          class="download"
          :fields="json_fields"
          :data="json_data"
          :name="String(this.input.path+'.xls')"
          :before-generate="startDownload"
          :before-finish="endDownload"
        >
          <el-button v-loading="downloadLoading" type="primary" icon="el-icon-download" class="filter-item download">
            下载
          </el-button>
        </download-excel>
      </div>

      <json-editor
        v-model="input.body"
        class="req-body"
        styles="width: 30%"
        :point-out="pointOut"
        :only-read="false"
        title="请求Body"
        @getValue="getBody"
      />
      <json-editor v-model="resData" class="res-body" styles="width: 70%" :only-read="true" title="返回信息" />
    </el-card>
    <el-drawer
      ref="drawer"
      title="Edit SQL"
      :before-close="drawerHandleClose"
      :visible.sync="drawerShow"
      direction="rtl"
      close-on-press-escape
      destroy-on-close
      size="50%"
    >
      <el-button style="margin: 20px" type="warning" icon="el-icon-refresh" @click="sqlToDsl">开始转换</el-button>
      <el-link type="success" disabled>表名可用索引名代替</el-link>
      <sql-editor
        v-model="sqlStr"
        styles="width: 100%"
        @getValue="getSql"
      />
    </el-drawer>
    <back-to-top />
    <history
      v-if="dialogVisible"
      :dialog-visible="dialogVisible"
      @getHistoryData="getHistoryData"
      @close="closeHistory"
    />
  </div>
</template>

<script>
import Driver from 'driver.js' // import driver.js
import 'driver.js/dist/driver.min.css' // import driver.js css
import steps from './guide'
import { clone } from '@/utils/index'
import { RunDslAction, SqlToDslAction } from '@/api/es'
import { ListAction } from '@/api/es-map'
import { Finish, IsFinish } from '@/api/guid'
import { filterData } from '@/utils/table'

import { esBodyKeyWords, esPathKeyWords } from '@/utils/base-data'

export default {
  name: 'Index',
  components: {
    'SqlEditor': () => import('@/components/SqlEditor/index'),
    'BackToTop': () => import('@/components/BackToTop/index'),
    'JsonEditor': () => import('@/components/JsonEditor/index'),
    'History': () => import('@/views/rest/components/history')
  },
  data() {
    return {
      dialogVisible: false,
      driver: null,
      modName: 'DSL面板',
      downloadLoading: false,
      sqlLoading: false,
      sqlStr: '',
      drawerShow: false,
      loading: false,
      json_fields: {},
      json_data: '',
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
  computed: {
    canExport() {
      this.json_data = ''
      this.json_fields = {}
      const resData = JSON.parse(this.resData)
      if (resData == null) {
        return false
      }
      if (Array.isArray(resData)) {
        // this.json_fields[defaultKey] = defaultKey

        if (resData.length <= 0) {
          return false
        }
        this.json_data = this.replaceArrSpece(resData)
        Object.keys(resData[0]).forEach((key, index) => {
          this.json_fields[key] = key
        })
        return true
      } else {
        if (resData.hasOwnProperty('hits')) {
          if (resData['hits']['hits'].length > 0) {
            const json_data = resData['hits']['hits']
            const defaultKeys = ['_index', '_type', '_id']
            for (const defaultKey of defaultKeys) {
              this.json_fields[defaultKey] = defaultKey
            }
            const arrayColumns = []
            for (const v of resData['hits']['hits']) {
              const sourceMap = v['_source']
              if (sourceMap == null) {
                continue
              }
              Object.getOwnPropertyNames(sourceMap).forEach((sourceVal, index) => {
                // 如果是对象
                if (Object.prototype.toString.call(sourceMap[sourceVal]) === '[object Object]') {
                  Object.keys(sourceMap[sourceVal]).map(key => {
                    this.json_fields[sourceVal + '->' + key] = '_source.' + sourceVal + '.' + key
                  })
                } else if (Array.isArray(sourceMap[sourceVal])) { // 如果是数组
                  if (Object.prototype.toString.call(sourceMap[sourceVal][0]) === '[object Object]') {
                    arrayColumns.push(sourceVal)
                  }
                  this.json_fields[sourceVal.toString()] = '_source.' + sourceVal.toString()
                } else {
                  this.json_fields[sourceVal.toString()] = '_source.' + sourceVal.toString()
                }
              })
            }
            arrayColumns.forEach((arrayColumn, index) => {
              for (const i in json_data) {
                for (const column in json_data[i]['_source']) {
                  if (column == arrayColumn) {
                    json_data[i]['_source'][column] = JSON.stringify(json_data[i]['_source'][column])
                  }
                }
              }
            })
            this.json_data = json_data
            return true
          }
        }
      }
      return false
    }
  },
  created() {
    this.mergeEsPathKeyWords()
    const resReqInfo = sessionStorage.getItem('resReqInfo')

    if (resReqInfo != null && resReqInfo != '' && resReqInfo != 'null') {
      this.input = JSON.parse(resReqInfo)
    }
    this.startGuid()
  },
  destroyed() {
    const input = this.input
    const resReqInfo = JSON.stringify(input)
    sessionStorage.setItem('resReqInfo', resReqInfo)
  },
  methods: {
    getHistoryData(v) {
      this.input.path = v['path']
      this.input.method = v['method']
      this.input.body = v['body']
    },
    closeHistory(v) {
      this.dialogVisible = v
    },
    async finishGuid() {
      const { data, code, msg } = await Finish({ 'guid_name': this.modName })
    },
    async startGuid() {
      const { data, code, msg } = await IsFinish({ 'guid_name': this.modName })

      if (!data) {
        console.log('开始新手引导')
        this.driver = new Driver({
          className: 'scoped-class',
          animate: true,
          opacity: 0.75,
          padding: 10,
          allowClose: false,
          overlayClickNext: false,
          doneBtnText: '完成',
          closeBtnText: '关闭',
          nextBtnText: '下一步',
          prevBtnText: '上一步',
          onNext: (Element) => {
            console.log('Element', Element)
          }
        })
        setTimeout(() => {
          this.driver.defineSteps(steps)
          this.driver.start()
          this.finishGuid()
        }, 500)
      }
    },
    startDownload() {
      this.$message({
        type: 'success',
        message: '开始下载'
      })
    },
    endDownload() {
      this.$message({
        type: 'success',
        message: '下载完毕'
      })
    },
    async sqlToDsl() {
      const { data, code, msg } = await SqlToDslAction({ sql: this.sqlStr })
      if (code == 0) {
        this.$message({
          type: 'success',
          message: msg
        })
        this.input.body = data.dsl
        this.input.path = data.tableName + '/_search'
        this.go()
      } else {
        this.$message({
          type: 'error',
          message: msg
        })
        return
      }
    },
    getSql(sql) {
      this.sqlStr = sql
    },
    openDrag() {
      this.drawerShow = true
    },
    drawerHandleClose(done) {
      this.indexName = ''
      done()
    },
    mergeEsPathKeyWords() {
      const input = this.$store.state.baseData.EsConnect
      ListAction(input).then(res => {
        if (res.code == 0) {
          const list = res.data
          const indices = Object.keys(list)

          const queryData = clone(this.queryData)
          for (const esPathKeyWord of queryData) {
            if (esPathKeyWord.value.indexOf('{indices}') !== -1) {
              for (const indice of indices) {
                const mappings = Object.keys(list[indice]['mappings'])
                if (mappings.length == 0) {
                  mappings[0] = '{type}'
                }
                const obj = {
                  'value': esPathKeyWord.value.replace(/{indices}/g, indice).replace(/{type}/g, mappings[0]),
                  'data': esPathKeyWord.data.replace(/{indices}/g, indice).replace(/{type}/g, mappings[0])
                }
                this.queryData.push(obj)
              }
            }
          }
        }
      }).catch(err => {
        console.log(err)
        this.$message({
          type: 'error',
          message: '网络异常'
        })
      })
    },
    replaceArrSpece(arr) {
      for (const index in arr) {
        for (const index2 in arr[index]) {
          if (index2.indexOf('.') != -1) {
            arr[index][index2.split('.').join('->')] = arr[index][index2]
            delete arr[index][index2]
          }
        }
      }
      return arr
    },
    getBody(v) {
      this.input.body = v
    },
    clear() {
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
    MeetingConfirmBox(title) {
      return new Promise((resolve, reject) => {
        this.$confirm(title, '警告', {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning'
        })
          .then(() => resolve(true))
          .catch(err => resolve(false))
      })
    },
    async go() {
      const input = clone(this.input)

      if (input['method'] == 'DELETE' || input['path'].indexOf('_delete_by_query') != -1) {
        const isFinish = await this.MeetingConfirmBox('确定执行删除操作吗')
        if (!isFinish) {
          this.$message({
            type: 'success',
            message: '已取消'
          })
          return
        }
      }

      if (input['path'].trim().length > 0) {
        if (input['path'].trim().substr(0, 1) != '/') {
          input['path'] = '/' + input['path'].trim()
        }
      }

      input['es_connect'] = this.$store.state.baseData.EsConnect
      this.loading = true
      RunDslAction(input).then(res => {
        this.loading = false
        if (res.code == 0 || res.code == 200 || res.code == 201) {
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
        console.log(err)
        this.loading = false
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
  .search-history {
    width: 100px;
    font-size: 8px;
  }

  .download {
    display: inline;
    width: 100px;
  }

  .sql-format {
    width: 100px;
    font-size: 8px;
  }

  .go {
    width: 100px;
    font-size: 8px;
  }

  .autocomplete {
    width: 600px;
    font-size: 8px;
  }

  .select-method {
    width: 180px;
    font-size: 8px;
  }

  /deep/ :focus {
    outline: 0;
  }
</style>
