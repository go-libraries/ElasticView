<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入Http Method</el-tag>
        <el-select v-model="input.method" class="filter-item" placeholder="请选择版本" filterable>
          <el-option label="PUT【修改】" value="PUT"/>
          <el-option label="GET【查询】" value="GET"/>
          <el-option label="DELETE【删除】" value="DELETE"/>
          <el-option label="POST【新增】" value="POST"/>
          <el-option label="HEAD【是否存在】" value="HEAD"/>
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

        <el-button class="filter-item" :loading="loading" type="success" @click="run">RUN-></el-button>
        <el-button class="filter-item" v-show="input.method == 'GET'" type="warning" @click="openDrag">SqlFormatDSl
        </el-button>
        <download-excel
          v-if="canExport"
          ref="download"
          :fields="json_fields"
          :data="json_data"
          :name="String(this.input.path+'.xls')"
        >
          <el-button type="primary" icon="el-icon-download" class="filter-item">导出结果</el-button>
        </download-excel>
      </div>
      <json-editor
        v-model="input.body"
        styles="width: 30%"
        :point-out="pointOut"
        :onlyRead="false"
        title="请求Body"
        @getValue="getBody"
      />
      <json-editor v-model="resData" styles="width: 70%" :onlyRead="true" title="返回信息"/>
    </el-card>
    <el-drawer
      title="Edit SQL"
      :before-close="drawerHandleClose"
      :visible.sync="drawerShow"
      direction="rtl"
      custom-class="demo-drawer"
      ref="drawer"
      close-on-press-escape
      destroy-on-close
      size="50%"
    >

      <el-button style="margin: 20px" :loading="sqlLoading" type="warning" @click="sqlToDsl">Convert</el-button>
      <sql-editor
        @getValue="getSql"
        v-model="sqlStr"
        styles="width: 100%"
      />
    </el-drawer>
    <back-to-top/>
  </div>
</template>

<script>
    import {clone} from '@/utils/index'
    import {RunDslAction,SqlToDslAction} from '@/api/es'
    import {ListAction} from '@/api/es-map'

    import {filterData} from '@/utils/table'

    import {esBodyKeyWords, esPathKeyWords} from '@/utils/base-data'

    export default {
        name: 'Index',
        components: {
            'SqlEditor': () => import('@/components/SqlEditor/index'),
            'BackToTop': () => import('@/components/BackToTop/index'),
            'JsonEditor': () => import('@/components/JsonEditor/index')
        },
        data() {
            return {
                sqlLoading: false,
                sqlStr: "",
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
        },
        destroyed() {
            const input = this.input
            const resReqInfo = JSON.stringify(input)
            sessionStorage.setItem('resReqInfo', resReqInfo)
        },
        methods: {
            async sqlToDsl() {
               const {data,code,msg } = await SqlToDslAction({sql:this.sqlStr})
                if (code == 0) {
                    this.$message({
                        type: 'success',
                        message: msg
                    })
                    this.input.body = data
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
            run() {
                this.loading = true

                const input = clone(this.input)

                if (input['path'].trim() != '') {
                    input['path'] = '/' + input['path'].trim()
                }

                input['es_connect'] = this.$store.state.baseData.EsConnect

                RunDslAction(input).then(res => {
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
                    this.loading = false
                    this.resData = JSON.stringify(res.data, null, '\t')
                }).catch(err => {
                    this.$message({
                        type: 'error',
                        message: '网络异常'
                    })
                    this.loading = false
                })
            }
        }
    }
</script>

<style scoped>
  /deep/ :focus {
    outline: 0;
  }
</style>
