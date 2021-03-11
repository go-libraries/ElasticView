<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入关键词</el-tag>
        <el-input v-model="input" class="filter-item" style="width: 300px" />
      </div>
      <el-table
        :header-cell-style="{background:'#eef1f6',color:'#606266'}"
        :data="filterData(list,input)"
        style="width: 100%;margin-top:30px;"
        border
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

        <el-table-column align="center" label="集群名称" width="120">
          <template slot-scope="scope">
            {{ scope.row.cluster }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="集群状态" width="100">
          <template slot-scope="scope">
            {{ scope.row.status }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点总数" width="100">
          <template slot-scope="scope">
            {{ scope.row["node.total"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="数据节点总数" width="100">
          <template slot-scope="scope">
            {{ scope.row["node.data"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="分片总数" width="100">
          <template slot-scope="scope">
            {{ scope.row.shards }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="主分片总数" width="100">
          <template slot-scope="scope">
            {{ scope.row.pri }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="复制节点总数" width="100">
          <template slot-scope="scope">
            {{ scope.row.relo }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="初始化节点总数" width="100">
          <template slot-scope="scope">
            {{ scope.row.init }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="未分配分片总数" width="100">
          <template slot-scope="scope">
            {{ scope.row.unassign }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="待定任务总数" width="100">
          <template slot-scope="scope">
            {{ scope.row.pending_tasks }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="等待最长任务的等待时间" width="100">
          <template slot-scope="scope">
            {{ scope.row.max_task_wait_time }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="活动分片百分比" width="100">
          <template slot-scope="scope">
            {{ scope.row.active_shards_percent }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="自标准时间（1970-01-01 00:00:00）以来的秒数" width="100">
          <template slot-scope="scope">
            {{ scope.row.epoch }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="时分秒,utc时区" width="100">
          <template slot-scope="scope">
            {{ scope.row.timestamp }}
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>

import { filterData } from '@/utils/table'
import { CatAction } from '@/api/es'
export default {
  name: 'CatHealth',
  data() {
    return {
      list: [],
      input: ''
    }
  },
  mounted() {
    this.searchData()
  },
  methods: {
    filterData(list, input) {
      return filterData(list, input)
    },
    searchData() {
      this.connectLoading = true
      const form = {
        cat: this.$options.name,
        es_connect: this.$store.state.baseData.EsConnect
      }
      CatAction(form).then(res => {
        if (res.code == 0) {
          this.list = res.data
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          })
        }
        this.connectLoading = false
      }).catch(err => {
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
