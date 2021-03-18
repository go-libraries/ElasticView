<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <div style="margin-left: 30px;">
          <el-button type="success" class="filter-item" icon="el-icon-refresh" @click="refresh">刷新</el-button>
          <el-button class="filter-item" @click="commit">确认</el-button>
        </div>
      </div>
      <el-card class="box-card card">
        <el-tag effect="dark" type="success">url参数</el-tag>
        <div style="margin: 20px;" />
        <el-form label-position="right" label-width="400px">
          <el-form-item label="配置">
            <el-checkbox-group v-model="urlParmasConfig">
              <el-checkbox v-for="(v,k,index) in form.urlParmas" :label="k" />
            </el-checkbox-group>
          </el-form-item>
          <el-form-item
            v-if="urlParmasConfig.indexOf('requests_per_second') != -1"
            label="requests_per_second (每秒数据量阈值控制)"
          >
            <el-input v-model="form.urlParmas.requests_per_second" type="number" />
          </el-form-item>
          <el-form-item
            v-if="urlParmasConfig.indexOf('wait_for_active_shards') != -1"
            label="wait_for_active_shards (重建索引分片响应设置)"
          >
            <el-input v-model="form.urlParmas.wait_for_active_shards" type="number" />
          </el-form-item>
          <el-form-item v-if="urlParmasConfig.indexOf('scroll') != -1" label="scroll (快照查询时间)">
            <el-input v-model="form.urlParmas.scroll" clearable />
          </el-form-item>
          <el-form-item v-if="urlParmasConfig.indexOf('slices') != -1" label="slices (重建并行任务切片)">
            <el-input v-model="form.urlParmas.slices" type="number" />
          </el-form-item>
          <el-form-item v-if="urlParmasConfig.indexOf('max_docs') != -1" label="max_docs (要处理的最大文档数)">
            <el-input v-model="form.urlParmas.max_docs" type="number" clearable />
          </el-form-item>
          <el-form-item v-if="urlParmasConfig.indexOf('refresh') != -1" label="refresh (目标索引是否立即刷新)">
            <el-select v-model="form.urlParmas.refresh" filterable>
              <el-option label="true" :value="true" />
              <el-option label="false" :value="false" />
              <el-option label="wait_for" value="wait_for" />

            </el-select>
          </el-form-item>
          <el-form-item v-if="urlParmasConfig.indexOf('master_timeout') != -1" label="master_timeout (指定等待连接到主节点的时间段)">
            <el-input v-model="form.urlParmas.master_timeout" clearable />
          </el-form-item>
          <el-form-item v-if="urlParmasConfig.indexOf('timeout') != -1" label="timeout (指定等待响应的时间段)">
            <el-input v-model="form.urlParmas.timeout" clearable />
          </el-form-item>
          <el-form-item
            v-if="urlParmasConfig.indexOf('wait_for_completion') != -1"
            label="wait_for_completion (如果为true，则请求将阻塞，直到操作完成为止)"
          >
            <el-select v-model="form.urlParmas.wait_for_completion" filterable>
              <el-option label="true" :value="true" />
              <el-option label="false" :value="false" />
            </el-select>
          </el-form-item>

        </el-form>
      </el-card>
      <el-card class="box-card card" style="width: 45%;">
        <el-tag effect="dark">源索引</el-tag>
        <div style="margin: 20px;" />
        <el-form>
          <el-form-item label="配置">
            <el-checkbox-group v-model="sourceConfig">
              <el-checkbox v-for="(v,k,index) in form.source" v-if="k != 'index'" :label="k" />
            </el-checkbox-group>
          </el-form-item>
          <el-form-item v-if="sourceConfig.indexOf('remote') != -1">
            <el-form>
              <el-form-item label="host (远程连接信息)">
                <el-input v-model="form.source.remote.host" placeholder="远程连接信息" clearable />
              </el-form-item>
              <el-form-item label="username (用户名)">
                <el-input v-model="form.source.remote.username" placeholder="用户名" clearable />
              </el-form-item>
              <el-form-item label="password (密码)">
                <el-input v-model="form.source.remote.password" placeholder="密码" clearable />
              </el-form-item>
              <el-form-item label="socket_timeout">
                <el-input v-model="form.source.remote.socket_timeout" placeholder="socket_timeout" clearable />
              </el-form-item>
              <el-form-item label="connect_timeout">
                <el-input
                  v-model="form.source.remote.connect_timeout"
                  placeholder="connect_timeout"
                  clearable
                />
              </el-form-item>
            </el-form>
          </el-form-item>
          <el-form-item label="索引名">
            <el-select
              v-model="form.source.index"
              placeholder="请选择索引名"
              filterable
              allow-create
              multiple
              clearable
              @change="changeSourceIndex"
            >
              <el-option
                v-for="(item,index) of indexNameList"
                :key="index"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item v-if="sourceConfig.indexOf('type') != -1" label="类型名">
            <el-select v-model="form.source.type" placeholder="请选择索引名" filterable allow-create multiple clearable>
              <el-option
                v-for="(item,index) of mappingConfig"
                :key="index"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item v-if="sourceConfig.indexOf('_source') != -1" label="_source (字段列表)">
            <el-select v-model="form.source._source" placeholder="字段列表" filterable allow-create multiple clearable />
          </el-form-item>
          <el-form-item v-if="sourceConfig.indexOf('sort') != -1" label="sort (排序)">
            <el-tag effect="dark">字段</el-tag>
            <el-input v-model="form.source.sort.key" style="width: 200px" />
            <el-tag effect="dark">是否正序</el-tag>
            <el-select v-model="form.source.sort.sortType" placeholder="是否正序" filterable filterable>
              <el-option
                key="index"
                label="是"
                value="asc"
              />
              <el-option
                key="index"
                label="否"
                value="desc"
              />
            </el-select>
          </el-form-item>
          <el-form-item v-if="sourceConfig.indexOf('size') != -1" label="size (每批要编制索引的文档数)">
            <el-input v-model="form.source.size" type="number" style="width: 200px" />
          </el-form-item>
          <el-form-item v-if="urlParmasConfig.indexOf('max_docs') != -1" label="max_docs (要重新编制索引的最大文档数)">
            <el-input v-model="form.urlParmas.max_docs" type="number" clearable />
          </el-form-item>
          <el-form-item v-if="sourceConfig.indexOf('query') != -1" label="QUERY DSL">
            <json-editor
              v-model="form.source.query"
              styles="width: 100%"
              :point-out="pointOut"
              :read="false"
              title="QUERY"
              @getValue="getBody"
            />
          </el-form-item>
        </el-form>
      </el-card>
      <el-card class="box-card card" style="width: 45%;">
        <el-tag effect="dark" type="danger">目标索引</el-tag>
        <div style="margin: 20px;" />
        <el-form>
          <el-form-item label="配置">
            <el-checkbox-group v-model="destConfig">
              <el-checkbox v-for="(v,k,index) in form.dest" v-if="k != 'index'" :label="k" />
            </el-checkbox-group>
          </el-form-item>

          <el-form-item label="索引名">
            <el-select v-model="form.dest.index" placeholder="请选择索引名" clearable filterable>
              <el-option
                v-for="(item,index) of indexNameList"
                :key="index"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item v-if="destConfig.indexOf('version_type') != -1" label="version_type(用于索引操作的版本控制)">
            <el-select v-model="form.dest.version_type" placeholder="请选择version_type" clearable filterable>
              <el-option label="external" value="external" />
              <el-option label="internal" value="internal" />
              <el-option label="external_gt" value="external_gt" />
              <el-option label="external_gte" value="external_gte" />
            </el-select>
          </el-form-item>
          <el-form-item v-if="destConfig.indexOf('op_type') != -1" label="op_type(仅创建不存在的索引文档)">
            <el-select v-model="form.dest.op_type" placeholder="op_type(仅创建不存在的索引文档)" clearable filterable>
              <el-option
                label="create"
                value="create"
              />
              <el-option
                label="index"
                value="index"
              />
            </el-select>
          </el-form-item>
          <el-form-item v-if="destConfig.indexOf('routing') != -1" label="routing">
            <el-input v-model="form.dest.routing" style="width: 300px" clearable />
          </el-form-item>
          <el-form-item v-if="destConfig.indexOf('pipeline') != -1" label="pipeline">
            <el-input v-model="form.dest.pipeline" style="width: 300px" clearable />
          </el-form-item>
        </el-form>
      </el-card>
      <el-card class="box-card card">
        <el-tag effect="dark" type="warning">脚本</el-tag>
        <div style="margin: 20px;" />
        <el-form>
          <el-form-item label="配置">
            <el-checkbox-group v-model="scriptConfig">
              <el-checkbox v-for="(v,k,index) in form.script" v-if="k != 'index'" :label="k" />
            </el-checkbox-group>
          </el-form-item>
          <el-form-item v-if="scriptConfig.indexOf('source') != -1" label="script.source 定义脚本语言">
            <el-input v-model="form.script.source" clearable />
          </el-form-item>
          <el-form-item v-if="scriptConfig.indexOf('lang') != -1" label="script.lang 定义脚本实现的代码">
            <el-input v-model="form.script.lang" clearable />
          </el-form-item>
          <el-form-item v-if="scriptConfig.indexOf('params') != -1" label="定义脚本语言的参数">
            <json-editor
              v-model="form.script.params"
              style="height: 300px"
              styles="width: 100%"
              :read="false"
              title="定义脚本语言的参数"
              @getValue="getParams"
            />
          </el-form-item>
        </el-form>
      </el-card>
      <el-card class="box-card card">
        <el-tag effect="dark" type="info">额外参数</el-tag>
        <div style="margin: 20px;" />
        <el-form>
          <el-form-item label="配置">
            <el-checkbox-group v-model="extendConfig">
              <el-checkbox v-for="(v,k,index) in form.extend" v-if="k != 'index'" :label="k" />
            </el-checkbox-group>
          </el-form-item>
          <el-form-item v-if="extendConfig.indexOf('conflicts') != -1" label="conflicts">
            <el-input v-model="form.extend.conflicts" clearable />
          </el-form-item>
          <el-form-item v-if="extendConfig.indexOf('max_docs') != -1" label="max_docs">
            <el-input v-model="form.extend.max_docs" clearable />
          </el-form-item>
        </el-form>

      </el-card>
    </el-card>
  </div>
</template>

<script>
import { esBodyKeyWords } from '@/utils/base-data'
import { ListAction } from '@/api/es-map'

export default {
  name: 'Reindex',
  components: {
    'JsonEditor': () => import('@/components/JsonEditor/index')
  },
  data() {
    return {
      extendConfig: [],
      typeList: [],
      pointOut: esBodyKeyWords,
      urlParmasConfig: [],
      sourceConfig: [],
      destConfig: [],
      scriptConfig: [],
      indexNameList: [],
      form: {
        urlParmas: {
          timeout: '30s',
          master_timeout: '30s',
          requests_per_second: -1,
          max_docs: 1,
          slices: 5,
          scroll: '',
          wait_for_active_shards: 0,
          refresh: '',
          wait_for_completion: true
        },
        source: {
          max_docs: 1,
          index: [],
          type: [],
          query: '{}',
          _source: [],
          sort: {
            key: '',
            sortType: 'asc'
          },
          remote: {
            host: 'http://127.0.0.1:9200', // 远程es的ip和port列表
            username: 'elastic',
            password: '',
            socket_timeout: '',
            connect_timeout: '' // 超时时间设置
          },
          size: 0
        },
        dest: {
          index: '',
          version_type: '',
          op_type: 'index',
          routing: '=cat',
          pipeline: 'some_ingest_pipeline'
        },
        script: {
          source: "if (ctx._source.foo == 'bar') {ctx._version++; ctx._source.remove('foo')}",
          lang: 'painless',
          params: ''
        },
        extend: {
          conflicts: 'proceed',
          max_docs: 100
        }
      },
      mappings: {},
      mappingConfig: []
    }
  },
  mounted() {
    this.getIndexList()
  },
  methods: {
    commit() {
      const body = {
        source: {

        },
        dest: {

        },
        script: {

        }
      }
      const urlParmas = {}

      for (const extendConfig of this.extendConfig) {
        body[extendConfig] = this.form.extend[extendConfig]
      }
      for (const urlParma of this.urlParmasConfig) {
        urlParmas[urlParma] = this.form.urlParmas[urlParma]
      }
      for (const source of this.sourceConfig) {
        if (source == 'sort') {
          body['source']['sort'] = {}
          body['source']['sort'][this.form.source.sort.key] = this.form.source.sort.sortType
        } else if (source == 'query') {
          try {
            body['source']['query'] = JSON.parse(this.form.source.query)
          } catch (e) {
            console.log(e)
            this.$message({
              type: 'error',
              message: 'query json解析失败'
            })
            return
          }
        } else {
          body['source'][source] = this.form.source[source]
        }
      }
      for (const dest of this.destConfig) {
        body['dest'][dest] = this.form.dest[dest]
      }
      for (const script of this.scriptConfig) {
        if (script == 'params') {
          try {
            body['script']['params'] = JSON.parse(this.form.script.params)
          } catch (e) {
            console.log(e)
            this.$message({
              type: 'error',
              message: '脚本语言参数 json解析失败'
            })
            return
          }
        } else {
          body['script'][script] = this.form.script[script]
        }
      }
      console.log(urlParmas, 'urlParmas')
      console.log(body, 'body')
    },
    refresh() {
      console.log(11)
      this.form = this.$options.data().form
    },
    changeSourceIndex() {
      if (this.sourceConfig.indexOf('type') == -1) {
        return
      }
      this.mappingConfig = []
      const mappingConfig = []
      for (const indexName of this.form.source.index) {
        if (this.mappings.hasOwnProperty(indexName)) {
          const mappings = Object.keys(this.mappings[indexName].mappings)
          if (mappings.length > 0) {
            mappingConfig.push(mappings[0])
          }
        }
      }
      this.mappingConfig = mappingConfig
    },
    getParams(v) {
      console.log(v, typeof v)
      this.form.script.params = v
    },
    getBody(v) {
      this.form.source.query = v
    },
    changeSourceConfig() {
      console.log(111)
    },
    async getIndexList() {
      const { data, code, msg } = await ListAction(this.$store.state.baseData.EsConnect)
      if (code == 0) {
        this.indexNameList = Object.keys(data)
        this.mappings = data
      } else {
        this.$message({
          type: 'error',
          message: msg
        })
        return
      }
    }
  }
}
</script>

<style scoped>
  .card {
    width: 95%;
    float: left;
    margin: 30px
  }
</style>
