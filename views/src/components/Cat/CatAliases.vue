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
        <el-table-column align="center" label="别名" width="300">
          <template slot-scope="scope">
            {{ scope.row.alias }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引别名指向" width="250">
          <template slot-scope="scope">
            {{ scope.row.index }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="过滤器" width="220">
          <template slot-scope="scope">
            {{ scope.row.filter }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引路由" width="220">
          <template slot-scope="scope">
            {{ scope.row["routing.index"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="搜索路由" width="220">
          <template slot-scope="scope">
            {{ scope.row["routing.search"] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="写索引" width="220">
          <template slot-scope="scope">
            {{ scope.row.is_write_index }}
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
  name: 'CatAliases',
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
      console.log('filterData')
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
