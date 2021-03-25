<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入关键词</el-tag>
        <el-input v-model="input" class="filter-item width300" clearable @input="search" />

        <el-select v-model="status" class="filter-item width300" clearable filterable @change="search">
          <el-option label="索引健康状态" value="" />
          <el-option label="green" value="green" />
          <el-option label="yellow" value="yellow" />
          <el-option label="red" value="red" />
        </el-select>

        <el-button type="success" class="filter-item" icon="el-icon-refresh" @click="search">刷新</el-button>
        <el-button type="success" class="filter-item" icon="el-icon-plus" @click="openSettingDialog('','add')">新建索引
        </el-button>
        <el-button
          type="success"
          class="filter-item"
          icon="el-icon-sort"
          :loading="readOnlyAllowDeleteLoading"
          @click="readOnlyAllowDelete()"
        >
          切换为可读写状态
        </el-button>

      </div>
      <back-to-top />
      <el-table
        :loading="connectLoading"
        :header-cell-style="{background:'#eef1f6',color:'#606266'}"
        :data="list"
      >
        <el-table-column
          label="序号"
          align="center"
          fixed
          width="50"
        >
          <template slot-scope="scope">
            {{ scope.$index+1 }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="索引健康状态" width="100">
          <template slot-scope="scope">
            <el-button v-if="scope.row.health == 'green'" type="success" circle />
            <el-button v-if="scope.row.health == 'yellow'" type="warning" circle />
            <el-button v-if="scope.row.health == 'red'" type="danger" circle />
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引的开启状态" width="100">
          <template slot-scope="scope">
            <el-button
              v-show="scope.row.status == 'open'"
              type="success"
              size="small"
              icon="el-icon-success"
              @click="OpenOrCloseIndex(scope.row.index,'close',1)"
            >开启
            </el-button>
            <el-button
              v-show="scope.row.status == 'close'"
              type="danger"
              size="small"
              icon="el-icon-circle-close"
              @click="OpenOrCloseIndex(scope.row.index,'open',1)"
            >关闭
            </el-button>
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引名称" width="180">
          <template slot-scope="scope">
            {{ scope.row.index }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引uuid" width="220">
          <template slot-scope="scope">
            {{ scope.row.uuid }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引主分片数" width="80" prop="pri" sortable>
          <template slot-scope="scope">
            {{ scope.row.pri }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引副本分片数量" width="80" prop="rep" sortable>
          <template slot-scope="scope">
            {{ scope.row.rep }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引文档总数" width="80" prop="docs->count" sortable>
          <template slot-scope="scope">
            {{ bigNumberTransform(scope.row["docs->count"]) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引中删除状态的文档" width="80" prop="docs->deleted" sortable>
          <template slot-scope="scope">
            {{ scope.row["docs->deleted"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="主分片+副本分分片的大小" width="120" prop="store->size" sortable>
          <template slot-scope="scope">
            {{ scope.row["store->size"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="主分片的大小" width="150" prop="pri->store->size" sortable>
          <template slot-scope="scope">
            {{ scope.row["pri->store->size"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="300">
          <template slot-scope="scope">
            <el-button-group>
              <el-button
                type="primary"
                size="small"
                icon="el-icon-setting"
                @click="openSettingDialog(scope.row.index,'update')"
              >修改配置
              </el-button>
              <el-button
                icon="el-icon-more"
                type="primary"
                size="small"
                @click="openDrawer(scope.row.index)"
              >更多
              </el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        v-if="pageshow"
        class="pagination-container"
        :current-page="page"
        :page-sizes="[10, 20, 30, 50]"
        :page-size="limit"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
      <el-drawer
        ref="drawer"
        :title="indexName"
        :before-close="drawerHandleClose"
        :visible.sync="drawerShow"
        direction="rtl"
        custom-class="demo-drawer"
        close-on-press-escape
        destroy-on-close
        size="50%"
      >
        <el-tabs v-model="activeName" class="margin-left-10" @tab-click="changeTab">
          <el-tab-pane label="设置" name="Settings">
            <json-editor
              v-if="activeName == 'Settings'"
              v-model="activeData"
              v-loading="tabLoading"
              styles="width: 100%"
              :only-read="false"
              title="设置"
              @get
            />

          </el-tab-pane>
          <el-tab-pane label="映射" name="Mapping">
            <json-editor
              v-if="activeName == 'Mapping'"
              v-model="activeData"
              v-loading="tabLoading"
              styles="width: 100%"
              :only-read="true"
              title="映射"
            />
          </el-tab-pane>
          <el-tab-pane label="Stats" name="Stats">
            <json-editor
              v-if="activeName == 'Stats'"
              v-model="activeData"
              v-loading="tabLoading"
              styles="width: 100%;"
              :only-read="true"
              title="Stats"
            />
          </el-tab-pane>
          <el-tab-pane label="编辑索引配置" name="editSettings">
            <el-form>
              <el-form-item label="编辑索引配置">
                <el-button type="primary" icon="el-icon-edit-outline" @click="submitSettings()">提交</el-button>
                <el-button icon="refresh" @click="resetSettings">重置</el-button>
              </el-form-item>
              <el-form-item>

                <json-editor
                  v-if="activeName == 'editSettings'"
                  v-model="activeData"
                  v-loading="tabLoading"
                  :point-out="pointOut"
                  styles="width: 100%;"
                  :only-read="false"
                  title="编辑配置"
                  @getValue="getSettings"
                />
              </el-form-item>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="修改别名" name="alias">
            <alias v-if="activeName == 'alias'" :index-name="indexName" />
          </el-tab-pane>
          <div class="filter-container operate">
            <el-button
              type="danger"
              :loading="loadingGroup[0]"
              size="small"
              icon="el-icon-circle-close"
              class="filter-item"
              @click="OpenOrCloseIndex(indexName,'close',0)"
            >关闭
            </el-button>

            <el-button
              type="success"
              :loading="loadingGroup[1]"
              size="small"
              icon="el-icon-success"
              class="filter-item"
              @click="OpenOrCloseIndex(indexName,'open',1)"
            >打开
            </el-button>
            <el-button
              size="small"
              icon="el-icon-connection"
              class="filter-item"
              :loading="loadingGroup[2]"
              @click="runCommand('/_forcemerge?max_num_segments=1',2)"
            >强制合并索引
            </el-button>

            <el-popover
              placement="top-start"
              title="提示"
              width="200"
              trigger="hover"
              content="为了让最新的数据可以立即被搜索到"
            >
              <el-button
                slot="reference"
                class="filter-item"
                size="small"
                :loading="loadingGroup[3]"
                type="primary"
                icon="el-icon-refresh"
                @click="runCommand('/_refresh',3)"
              >刷新索引
              </el-button>
            </el-popover>

            <el-popover
              placement="top-start"
              title="提示"
              width="200"
              trigger="hover"
              content="为了让数据持久化到磁盘中"
            >
              <el-button
                slot="reference"
                size="small"
                :loading="loadingGroup[4]"
                type="info"
                icon="el-icon-s-open"
                class="filter-item"
                @click="runCommand('/_flush',4)"
              >冲洗索引
              </el-button>
            </el-popover>

            <el-button
              class="filter-item"
              size="small"
              type="warning"
              :loading="loadingGroup[5]"
              icon="el-icon-toilet-paper"
              @click="runCommand('/_cache/clear',5)"
            >清理缓存
            </el-button>

            <el-button
              class="filter-item"
              type="danger"
              size="small"
              :loading="loadingGroup[6]"
              icon="el-icon-delete"
              @click="deleteIndex(indexName,6)"
            >删除索引
            </el-button>

          </div>
        </el-tabs>

      </el-drawer>

      <settings
        v-if="openSettings"
        :index-name="indexName"
        :settings-type="settingsType"
        :open="openSettings"
        @close="closeSettings"
      />
    </el-card>
  </div>
</template>

<script>

import { filterData } from '@/utils/table'
import { CatAction, RunDslAction } from '@/api/es'
import { bigNumberTransform } from '@/utils/format'
import { CreateAction, GetSettingsAction } from '@/api/es-index'
import { esSettingsWords } from '@/utils/base-data'

export default {
  name: 'CatIndices',
  components: {
    'Settings': () => import('@/views/indices/components/settings'),
    'BackToTop': () => import('@/components/BackToTop/index'),
    'JsonEditor': () => import('@/components/JsonEditor/index'),
    'Alias': () => import('@/views/indices/components/alias')
  },

  data() {
    return {
      aliasList: [],
      pointOut: esSettingsWords,
      settings: {},
      readOnlyAllowDeleteLoading: false,
      loadingGroup: [
        false, false, false, false, false, false, false
      ],
      forceMergeLoading: false,
      tabLoading: false,
      activeData: '{}',
      activeName: 'Settings',
      drawerShow: false,
      settingsType: 'add',
      indexName: '',
      openSettings: false,
      total: 0,
      connectLoading: false,
      page: 1,
      limit: 10,
      pageshow: true,
      list: [],
      input: '',
      status: ''
    }
  },
  mounted() {
    this.searchData()
  },
  methods: {
    removeAlias(aliasName) {

    },
    addAlias() {

    },
    submitSettings() {
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnect
      try {
        input['settings'] = JSON.parse(this.activeData)
      } catch (e) {
        console.log(e)
        this.$message({
          type: 'error',
          message: 'JSON 解析异常'
        })
        return
      }
      input['index_name'] = this.indexName
      input['types'] = 'update'
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      CreateAction(input).then(res => {
        console.log(input['settings'])
        if (res.code == 0 || res.code == 200) {
          this.$message({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          })
        }
        loading.close()
      }).catch(err => {
        loading.close()
        this.$message({
          type: 'error',
          message: '网络异常'
        })
      })
    },
    getSettings(value) {
      this.activeData = value
    },
    resetSettings() {
      this.changeTab()
    },
    runCommand(command, loadingType) {
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnect
      input['method'] = 'POST'
      input['path'] = '/' + this.indexName + command
      this.loadingGroup[loadingType] = true

      RunDslAction(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          this.$message({
            type: 'success',
            message: res.msg
          })
          this.search()
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          })
        }
        this.loadingGroup[loadingType] = false
      }).catch(err => {
        this.loadingGroup[loadingType] = false
        this.$message({
          type: 'error',
          message: '网络异常'
        })
      })
    },
    async changeTab() {
      let path = ''
      switch (this.activeName) {
        case 'Settings':
          path = '/' + this.indexName + '/_settings'
          break
        case 'Mapping':
          path = '/' + this.indexName + '/_mapping'
          break
        case 'Stats':
          path = '/' + this.indexName + '/_stats'
          break
        case 'editSettings':
          const input = {}
          input['es_connect'] = this.$store.state.baseData.EsConnect
          input['index_name'] = this.indexName
          const { data } = await GetSettingsAction(input)

          const deleteKeyArr = [
            'creation_date', 'version', 'provided_name', 'uuid', 'format', 'number_of_shards'
          ]

          for (const key of deleteKeyArr) {
            if (data['index'].hasOwnProperty(key)) {
              delete data['index'][key]
            }
          }

          this.activeData = JSON.stringify(data, null, '\t')
          return
          break
        case 'alias':
          return
          break
        default:
          this.activeData = '{}'
          return
      }
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnect
      input['method'] = 'GET'
      input['path'] = path
      this.tabLoading = true
      RunDslAction(input).then(res => {
        if (res.code == 0 || res.code == 200) {
          this.activeData = JSON.stringify(res.data, null, '\t')
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          })
        }
        this.tabLoading = false
      }).catch(err => {
        this.tabLoading = false
        this.$message({
          type: 'error',
          message: '网络异常'
        })
      })
    },
    openDrawer(indexName) {
      this.indexName = indexName
      this.changeTab()
      this.drawerShow = true
    },
    drawerHandleClose(done) {
      this.indexName = ''
      done()
    },
    bigNumberTransform(value) {
      return bigNumberTransform(value)
    },
    OpenOrCloseIndex(indexName, types, loadingType) {
      if (types == '') {
        return
      }
      const typesMap = {
        'open': '开启',
        'close': '关闭'
      }

      this.$confirm('确定' + typesMap[types] + '该索引吗?', '警告', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const input = {}
          input['es_connect'] = this.$store.state.baseData.EsConnect
          input['method'] = 'POST'
          input['path'] = '/' + indexName + '/_' + types
          this.loadingGroup[loadingType] = true
          RunDslAction(input).then(res => {
            if (res.code == 0 || res.code == 200) {
              this.$message({
                type: 'success',
                message: res.msg
              })
              this.searchData()
            } else {
              this.$message({
                type: 'error',
                message: res.msg
              })
            }
            this.loadingGroup[loadingType] = false
          }).catch(err => {
            this.loadingGroup[loadingType] = false
            this.$message({
              type: 'error',
              message: '网络异常'
            })
          })
        })
        .catch(err => {
          console.error(err)
        })
    },
    openSettingDialog(indexName, settingsType) {
      this.indexName = indexName
      this.settingsType = settingsType
      this.openSettings = true
    },
    readOnlyAllowDelete() {
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnect
      input['method'] = 'PUT'
      input['path'] = '/_settings'
      input['body'] = `{
                    "index": {
                        "blocks": {
                            "read_only_allow_delete": "false"
                        }
                    }
                }`
      this.readOnlyAllowDeleteLoading = true
      RunDslAction(input).then(res => {
        if (res.code == 0 || res.code == 200) {
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
        this.readOnlyAllowDeleteLoading = false
      }).catch(err => {
        this.readOnlyAllowDeleteLoading = false
        this.$message({
          type: 'error',
          message: '网络异常'
        })
      })
    },
    editIndex(row) {
      console.log(row)
    },
    deleteIndex(indexName, loadingType) {
      this.$confirm('确定删除该索引吗?', '警告', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const input = {}
          input['es_connect'] = this.$store.state.baseData.EsConnect
          input['method'] = 'DELETE'
          input['path'] = '/' + indexName
          this.loadingGroup[loadingType] = true
          RunDslAction(input).then(res => {
            if (res.code == 0 || res.code == 200) {
              this.$message({
                type: 'success',
                message: indexName + '已删除'
              })
              this.searchData()
              this.loadingGroup[loadingType] = false
            } else {
              this.$message({
                type: 'error',
                message: res.msg
              })
              this.loadingGroup[loadingType] = false
            }
          }).catch(err => {
            this.$message({
              type: 'error',
              message: '网络异常'
            })
          })
        })
        .catch(err => {
          console.error(err)
        })
    },
    closeSettings() {
      this.indexName = ''
      this.settingsType = 'add'
      this.openSettings = false
    },
    search() {
      this.page = 1
      this.pageshow = false
      this.searchData()
      this.$nextTick(() => {
        this.pageshow = true
      })
    },
    filterData(list, input) {
      return filterData(list, input)
    },
    // 当每页数量改变
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`)
      this.limit = val
      this.searchData()
    },
    // 当当前页改变
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`)
      this.page = val
      this.searchData()
    },
    searchData() {
      this.connectLoading = true
      const form = {
        cat: this.$options.name,
        es_connect: this.$store.state.baseData.EsConnect
      }
      CatAction(form).then(res => {
        if (res.code == 0) {
          let list = res.data

          for (const index in list) {
            const obj = list[index]
            // 把 . 转成 ->
            for (const key in obj) {
              let value = parseInt(obj[key])
              if (isNaN(value)) {
                value = obj[key]
              }
              list[index][key.split('.').join('->')] = value
            }
          }

          list = filterData(list, this.status.trim())
          list = filterData(list, this.input.trim())
          this.list = list.filter((item, index) =>
            index < this.page * this.limit && index >= this.limit * (this.page - 1)
          )
          this.total = list.length
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          })
        }
        this.connectLoading = false
      }).catch(err => {
        console.log(err)
        this.$message({
          type: 'error',
          message: '网络异常'
        })
        this.connectLoading = false
      })
    }
  }
}
</script>

<style scoped>
  .operate {

  }

  .aliasName {
    width: 400px;
  }

  .margin-left-10 {
    margin-left: 10px
  }

  .width300 {
    width: 300px;
  }

  /deep/ :focus {
    outline: 0;
  }
</style>
