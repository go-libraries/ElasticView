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

        <el-table-column align="center" label="索引健康状态" width="100">
          <template slot-scope="scope">
            {{ scope.row.health }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="索引的开启状态" width="100">
          <template slot-scope="scope">
            {{ scope.row.status }}
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
        <el-table-column align="center" label="索引中文档总数" width="120" prop="docsCount" sortable>
          <template slot-scope="scope">
            {{ scope.row["docs.count"] }}
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
      </el-table>
    </el-card>
  </div>
</template>

<script>
import { filterData } from '@/utils/table'
import { CatAction } from '@/api/es'
export default {
  name: 'CatIndices',
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
            this.list[k]['docsCount'] = Number(this.list[k]['docs.count'])
            this.list[k]['docsDeleted'] = Number(this.list[k]['docs.deleted'])
            this.list[k]['storeSize'] = Number(this.list[k]['store.size'])
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
