<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-select v-model="snapshotNameList" clearable filterable multiple placeholder="请选择仓库名" :loading="loading" @change="search">
          <el-option
            v-for="(v,k,index) of resData"
            :key="index"
            :label="k"
            :value="k"
          />
        </el-select>
        <el-button type="warning">新建仓库</el-button>
      </div>
      <el-table
        :data="tableData"
        @row-click="openDetail"
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
        <el-table-column align="center" label="仓库名" prop="name" width="200" />
        <el-table-column align="center" label="仓库地址" prop="location" width="400" />
        <el-table-column align="center" label="类型" prop="type" width="200" />

        <el-table-column align="center" label="是否压缩" prop="compress" width="200" />
        <el-table-column align="center" label="制作快照的速度" width="200">
          <template slot-scope="scope">
            <div v-if="scope.row.max_snapshot_bytes_per_sec != ''">{{ scope.row.max_snapshot_bytes_per_sec }}/s</div>
            <div v-else>20mb/s</div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="快照恢复的速度" width="200">
          <template slot-scope="scope">
            <div v-if="scope.row.max_restore_bytes_per_sec != ''">{{ scope.row.max_restore_bytes_per_sec }}/s</div>
            <div v-else>20mb/s</div>
          </template>
        </el-table-column>

        <el-table-column align="center" label="操作" fixed="right" width="300">
          <template slot-scope="scope">
            <el-button type="success" size="small" icon="el-icon-search" @click="look(scope.row.name)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-drawer

        ref="drawer"
        title="JSON数据"
        :before-close="drawerHandleClose"
        :visible.sync="drawerShow"

        direction="rtl"
        close-on-press-escape
        destroy-on-close
        size="70%"
      >

        <json-editor
          v-if="!isArray"
          v-model="JSON.stringify(resData[name],null, '\t')"
          class="res-body"
          styles="width: 100%"
          :read="true"
          title="JSON数据"
        />
      </el-drawer>
    </el-card>
  </div>
</template>

<script>
import { SnapshotListAction } from '@/api/es-backup'

export default {
  name: 'ResTable',
  components: {
    'JsonEditor': () => import('@/components/JsonEditor/index')
  },
  data() {
    return {
      loading: false,
      snapshotNameList: [],
      resData: {},
      name: '',
      drawerShow: false,
      tableData: [],
      index: 0,
      isArray: false,
      tableHeader: []
    }
  },
  created() {
    this.initSnapshotName()
    this.search()
  },

  methods: {
    async initSnapshotName() {
      this.loading = true
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnectID
      const { data, code, msg } = await SnapshotListAction(input)
      if (code != 0) {
        this.$message({
          type: 'error',
          message: msg
        })
        return
      } else {
        this.resData = data.res
      }
      this.loading = false
    },
    async search() {
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnectID
      input['snapshot_info_list'] = this.snapshotNameList
      const { data, code, msg } = await SnapshotListAction(input)
      if (code != 0) {
        this.$message({
          type: 'error',
          message: msg
        })
        return
      } else {
        this.tableData = data.list
      }
    },

    openDetail(row, index, tmp) {
      this.name = row.name
      this.drawerShow = true
    },
    look(index) {
      this.name = index
      this.drawerShow = true
    },
    drawerHandleClose(done) {
      done()
    }
  }
}
</script>

<style scoped>

</style>
