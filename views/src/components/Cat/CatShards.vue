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

        <el-table-column align="center" label="索引名称" width="220">
          <template slot-scope="scope">
            {{ scope.row.index }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="分片序号" width="100" sortable prop="shard">
          <template slot-scope="scope">
            {{ scope.row.shard }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="分片类型" width="120">
          <template slot-scope="scope">
            <div v-if="scope.row.prirep == 'p'">主分片</div>
            <div v-if="scope.row.prirep == 'r'">复制分片</div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="分片状态" width="220">
          <template slot-scope="scope">
            {{ scope.row.state }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="该分片存放的文档数量" width="140" sortable prop="docs">
          <template slot-scope="scope">
            {{ scope.row.docs }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="该分片占用的存储空间大小" width="200">
          <template slot-scope="scope">
            {{ scope.row.store }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="该分片所在的服务器ip" width="200">
          <template slot-scope="scope">
            {{ scope.row.ip }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="该分片所在的节点名称" width="200">
          <template slot-scope="scope">
            {{ scope.row.node }}
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
  name: 'CatShards',
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
          for (const k in this.list) {
            this.list[k]['docs'] = Number(this.list[k]['docs'])
          }
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
