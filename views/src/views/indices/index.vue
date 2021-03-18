<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入关键词</el-tag>
        <el-input v-model="input" class="filter-item" style="width: 300px" clearable @input="search" />

        <el-select v-model="status" style="width: 300px" class="filter-item" clearable filterable @change="search">
          <el-option label="索引健康状态" value="" />
          <el-option label="green" value="green" />
          <el-option label="yellow" value="yellow" />
          <el-option label="red" value="red" />
        </el-select>

        <el-button type="success" class="filter-item" icon="el-icon-refresh" @click="search">刷新</el-button>
        <el-button type="success" class="filter-item" icon="el-icon-plus" @click="openSettingDialog('','add')">新建索引
        </el-button>
      </div>
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
            <el-button v-show="scope.row.status == 'open'" type="success" size="small" icon="el-icon-success">开启
            </el-button>
            <el-button v-show="scope.row.status == 'close'" type="danger" size="small" icon="el-icon-circle-close">关闭
            </el-button>
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引名称" width="200">
          <template slot-scope="scope">
            {{ scope.row.index }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引uuid" width="220">
          <template slot-scope="scope">
            {{ scope.row.uuid }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引主分片数" width="100" prop="pri" sortable>
          <template slot-scope="scope">
            {{ scope.row.pri }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引副本分片数量" width="100" prop="rep" sortable>
          <template slot-scope="scope">
            {{ scope.row.rep }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引文档总数" width="120" prop="docsCount" sortable>
          <template slot-scope="scope">
            {{ bigNumberTransform(scope.row["docs.count"]) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引中删除状态的文档" width="100" prop="docsDeleted" sortable>
          <template slot-scope="scope">
            {{ scope.row["docs.deleted"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="主分片+副本分分片的大小" width="180">
          <template slot-scope="scope">
            {{ scope.row["store.size"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="主分片的大小" width="200">
          <template slot-scope="scope">
            {{ scope.row["pri.store.size"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="200">
          <template slot-scope="scope">
            <el-button
              style="margin-top: 10px;"
              type="primary"
              size="small"
              icon="el-icon-setting"
              @click="openSettingDialog(scope.row.index,'update')"
            >配置
            </el-button>
            <el-button
              style="margin-top: 10px"
              type="danger"
              size="small"
              icon="el-icon-delete"
              @click="deleteIndex(scope.row.index)"
            >删除
            </el-button>
            <el-button
              style="margin-top: 10px"
              type="success"
              size="small"
              icon="el-icon-success"
              @click="OpenOrCloseIndex(scope.row.index,'open')"
            >开启
            </el-button>
            <el-button
              style="margin-top: 10px"
              type="danger"
              size="small"
              icon="el-icon-circle-close"
              @click="OpenOrCloseIndex(scope.row.index,'close')"
            >关闭
            </el-button>
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

export default {
  name: 'CatIndices',
  components: {
    'Settings': () => import('@/views/indices/settings')
  },
  data() {
    return {
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
    bigNumberTransform(value) {
      return bigNumberTransform(value)
    },
    OpenOrCloseIndex(indexName, types) {
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
          const loading = this.$loading({
            lock: true,
            text: 'Loading',
            spinner: 'el-icon-loading',
            background: 'rgba(0, 0, 0, 0.7)'
          })
          RunDslAction(input).then(res => {
            if (res.code == 0 || res.code == 200) {
              console.log(res)
              this.searchData()
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
    editIndex(row) {
      console.log(row)
    },
    deleteIndex(indexName) {
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
          RunDslAction(input).then(res => {
            if (res.code == 0 || res.code == 200) {
              this.$message({
                type: 'success',
                message: indexName + '已删除'
              })
              this.searchData()
            } else {
              this.$message({
                type: 'error',
                message: res.msg
              })
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
          for (const k in list) {
            list[k]['docsCount'] = Number(list[k]['docs.count'])
            list[k]['docsDeleted'] = Number(list[k]['docs.deleted'])
            list[k]['storeSize'] = Number(list[k]['store.size'])
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

</style>
