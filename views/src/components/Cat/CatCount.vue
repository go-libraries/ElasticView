<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入关键词</el-tag>
        <el-input v-model="input" class="filter-item" style="width: 300px" />
      </div>
      <el-table
        :loading="connectLoading"
        :header-cell-style="{background:'#eef1f6',color:'#606266'}"
        :data="filterData(list,input)"
        style="width: 100%;margin-top:30px;"
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

        <el-table-column align="center" label="自标准时间（1970-01-01 00:00:00）以来的秒数" width="500">
          <template slot-scope="scope">
            {{ scope.row.epoch }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="时分秒,utc时区" width="500">
          <template slot-scope="scope">
            {{ scope.row.timestamp }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="文档总数" width="500" sortable prop="count">
          <template slot-scope="scope">
            {{ scope.row.count }}
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
  name: 'CatCount',
  data() {
    return {
      list: [],
      input: '',
      connectLoading: false
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
