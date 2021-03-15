<template>
  <div class="app-container">
    <el-card class="box-card">
      <div class="filter-container">
        <el-tag class="filter-item">请输入关键词</el-tag>
        <el-input v-model="input" class="filter-item" style="width: 300px" clearable @input="search" />
      </div>
      <el-table
        :loading="connectLoading"
        :header-cell-style="{background:'#eef1f6',color:'#606266'}"
        :data="list"
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
      <el-pagination
        v-if="pageshow"
        class="pagination-container"
        :current-page="page"
        :page-sizes=" [10, 20, 30, 50]"
        :page-size="limit"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
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
      connectLoading: false,
      total: 0,
      page: 1,
      limit: 10,
      pageshow: true,
      list: [],
      input: ''
    }
  },
  mounted() {
    this.searchData()
  },
  methods: {
    search() {
      this.page = 1
      this.pageshow = false
      this.searchData()
      this.$nextTick(() => {
        this.pageshow = true
      })
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
          let list = res.data
          for (const k in this.list) {
            list[k]['docs'] = Number(list[k]['docs'])
          }
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
