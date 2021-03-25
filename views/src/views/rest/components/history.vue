<template>
  <div>
    <el-dialog width="80%" :visible.sync="dialogVisible" title="历史记录" @close="close">
      <el-table
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

        <el-table-column align="center" label="Method" width="220">
          <template slot-scope="scope">
            {{ scope.row.method }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="Path" width="300">
          <template slot-scope="scope">
            {{ scope.row.path }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="Body" width="300">
          <template slot-scope="scope">
            {{ scope.row.body }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="创建时间" width="220">
          <template slot-scope="scope">
            {{ scope.row.created }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="操作" fixed="right" width="300">
          <template slot="header" slot-scope="scope">
            <el-button type="danger" size="small" icon="el-icon-delete" @click="clean">清空</el-button>
          </template>
          <template slot-scope="scope">
            <el-button type="success" size="small" icon="el-icon-search" @click="getHistoryData(scope)">搜索</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>

<script>
import { CleanAction, ListAction } from '@/api/dsl-history'

export default {
  name: 'History',
  props: {
    dialogVisible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      list: []
    }
  },
  mounted() {
    this.searchHistory()
  },
  methods: {
    getHistoryData(scope) {
      this.$emit('getHistoryData', scope.row)
      this.$emit('close', false)
    },
    async clean() {
      const { data, code, msg } = await CleanAction()
      console.log(data)
      if (code != 0) {
        this.$message({
          type: 'error',
          message: msg
        })
        return
      } else {
        this.$message({
          type: 'success',
          message: msg
        })
        this.searchHistory()
      }
    },
    close() {
      this.$emit('close', false)
    },
    async searchHistory() {
      const { data, code, msg } = await ListAction()
      console.log(data)
      if (code != 0) {
        this.$message({
          type: 'error',
          message: msg
        })
        return
      } else {
        this.list = data
      }
    }
  }
}
</script>

<style scoped>

</style>
