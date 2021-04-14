<template>
  <div>
    <el-dialog width="80%" :visible.sync="dialogVisible" title="查询结果" @close="close">
      <pl-table
        :data="tableData"
        highlight-current-row
        use-virtual
        :row-height="30"
        @row-click="openDetail"
      >
        <pl-table-column
          label="序号"
          align="center"
          fixed
          width="50"
        >
          <template slot-scope="scope">
            {{ scope.$index+1 }}
          </template>
        </pl-table-column>
        <pl-table-column
          v-for="(val,key) in tableHeader"
          v-if="val != 'xwl_index'"
          :key="key"
          :prop="val"
          :sortable="true"
          align="center"
          :label="val"
        >
          <template slot-scope="scope">
            <el-popover

                          placement="top-start"
                          trigger="hover"
            >
              <div>{{ scope.row[val] }}</div>
               <span v-if="scope.row[val].length>=50"  slot="reference">{{ scope.row[val].toString().substr(0, 50) + "..." }}
              </span>
               <span v-else slot="reference">{{ scope.row[val].toString() }}
              </span>
            </el-popover>
          </template>
        </pl-table-column>
        <pl-table-column align="center" label="操作" fixed="right" width="300">
          <template slot-scope="scope">
            <el-button type="success" size="small" icon="el-icon-search" @click="look(scope.$index)">查看</el-button>
          </template>
        </pl-table-column>

      </pl-table>

    </el-dialog>
    <el-drawer
      ref="drawer"
      title="详细数据"
      :before-close="drawerHandleClose"
      :visible.sync="drawerShow"

      direction="rtl"
      close-on-press-escape
      destroy-on-close
      size="70%"
    >

      <json-editor
        v-if="isArray"
        v-model="JSON.stringify(json_data[index],null, '\t')"
        class="res-body"
        styles="width: 100%"
        :read="true"
        title="详细数据"
      />
      <json-editor
        v-if="!isArray"
        v-model="JSON.stringify(json_data['hits']['hits'][index],null, '\t')"
        class="res-body"
        styles="width: 100%"
        :read="true"
        title="详细数据"
      />
    </el-drawer>
  </div>
</template>

<script>
import { clone } from '@/utils/index'

export default {
  name: 'ResTable',
  components: {
    'JsonEditor': () => import('@/components/JsonEditor/index')
  },
  props: {
    dialogVisible: {
      type: Boolean,
      default: false
    },
    jsonData: {
      type: Array,
      default: []
    }
  },
  data() {
    return {
      drawerShow: false,
      tableData: [],
      index: 0,
      isArray: false,
      tableHeader: []
    }
  },
  created() {
    const resData = clone(this.json_data)
    console.log(resData, 'resData', this.json_data)
    if (Array.isArray(resData)) {
      if (resData.length > 500) {
        this.$message({
          type: 'error',
          message: '请减少数据条数'
        })
        this.$emit('close', false)
        return
      }
      this.tableData = resData

      for (const index in this.tableData) {
        const obj = this.tableData[index]

        for (const key in obj) {
          let value = this.strToNum(obj[key])
          if (value == false) {
            value = obj[key]
          }
          if (key.indexOf('.') != -1) {
            this.tableData[index][key.split('.').join('->')] = value
            delete this.tableData[index][key]
          } else {
            this.tableData[index][key] = value
          }

          this.tableData[index]['xwl_index'] = index
        }
      }
      console.log(this.tableData)

      this.isArray = true
    } else {
      if (resData.hasOwnProperty('hits')) {
        if (resData['hits']['hits'].length > 0) {
          if (resData['hits']['hits'].length > 500) {
            this.$message({
              type: 'error',
              message: '请减少查詢的数据条数'
            })
            this.$emit('close', false)
            return
          }

          const sourceArr = []

          for (const index in resData['hits']['hits']) {
            const _source = resData['hits']['hits'][index]['_source']
            _source['_id'] = resData['hits']['hits'][index]['_id']
            _source['_score'] = resData['hits']['hits'][index]['_score']
            _source['xwl_index'] = index
            sourceArr.push(_source)
          }

          this.isArray = false
          this.tableData = sourceArr

          for (const index in this.tableData) {
            for (const key in this.tableData[index]) {
              if (typeof this.tableData[index][key] === 'object' || typeof this.tableData[index][key] === 'array') {
                this.tableData[index][key] = JSON.stringify(this.tableData[index][key])
              }
            }
          }
        }
      }

      console.log('END')
    }
    this.tableHeader = Object.keys(this.tableData[0])
  },
  methods: {
    strToNum(str) {
      var convertNum = Number(str) // 将字符串强制转换为数字
      if (str === '') { // 排除空字符串
        return false
      } else {
        if (str.includes(' ')) { // 排除空格
          return false
        } else {
          if (isNaN(convertNum)) {
            return false
          } else {
            return convertNum
          }
        }
      }
    },
    openDetail(row, index, tmp) {
      this.index = row.xwl_index
      this.drawerShow = true
    },
    look(index) {
      this.index = index
      this.drawerShow = true
    },
    drawerHandleClose(done) {
      done()
    },
    close() {
      this.$emit('close', false)
    }

  }
}
</script>

<style scoped>

</style>
