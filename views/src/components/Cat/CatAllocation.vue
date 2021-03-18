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
        <el-table-column align="center" label="节点host" width="150">
          <template slot-scope="scope">
            {{ scope.row.host }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点ip" width="150">
          <template slot-scope="scope">
            {{ scope.row.ip }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点名称" width="170">
          <template slot-scope="scope">
            {{ scope.row.node }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点承载分片数量" width="150" sortable prop="shards">
          <template slot-scope="scope">
            {{ scope.row.shards }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引占用空间大小" width="150">
          <template slot-scope="scope">
            {{ scope.row["disk.indices"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点所在机器已使用的磁盘空间大小" width="150">
          <template slot-scope="scope">
            {{ scope.row["disk.used"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点可用空间大小" width="170">
          <template slot-scope="scope">
            {{ scope.row["disk.avail"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点总空间大小" width="170">
          <template slot-scope="scope">
            {{ scope.row["disk.total"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="节点磁盘占用百分比" width="170">
          <template slot-scope="scope">
            {{ scope.row["disk.percent"] }}%
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
  name: 'CatAllocation',
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
