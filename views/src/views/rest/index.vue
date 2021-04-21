<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button
        class="filter-item"

        @click="addTab()"
      >
        添加查询窗口
      </el-button>
    </div>
    <el-card class="box-card">

      <el-tabs v-model="editableTabsValue" type="card" closable @tab-remove="removeTab">
        <el-tab-pane
          v-for="(item, index) in editableTabs"
          :key="item.uniqueId"
          :label="item.title"
          :name="Number(item.uniqueId)"
        >
          <tools
            :title="item.title"
            :query-data="queryData"
            :input="item.input"
            :unique-id="item.uniqueId"
            :sql-str="item.sqlStr"
            @saveData="saveData"
          />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>
<script>
import { esPathKeyWords } from '@/utils/base-data'
import { ListAction } from '@/api/es-map'

export default {
  components: {
    'Tools': () => import('@/views/rest/components/tools')
  },
  data() {
    return {
      editableTabsValue: 1,
      editableTabs: [],
      queryData: [],
      test: 1
    }
  },
  mounted() {
    this.$nextTick(() => {
      var introJS = require('intro.js')

      introJS().setOptions({
        prevLabel: '上一步',
        nextLabel: '下一步',
        skipLabel: '跳过',
        doneLabel: '完成'
      }).onbeforechange((e) => {
        console.log(e.getAttribute('guid'))
        console.log(e)
      }).oncomplete(() => {
        // 点击结束按钮后执行的事件
      }).onexit(() => {
        // 点击跳过按钮后执行的事件
      }).start()
      // introJS().start() // 退出引导调用 exit() 即可
    })
    this.mergeEsPathKeyWords()
    const editableTabs = sessionStorage.getItem('editableTabs')
    const editableTabsValue = sessionStorage.getItem('editableTabsValue')

    if (editableTabsValue != null && editableTabs != '' && editableTabs != 'null') {
      this.editableTabsValue = Number(editableTabsValue)
    } else {
      this.editableTabsValue = 1
    }

    if (editableTabs != null && editableTabs != '' && editableTabs != 'null') {
      this.editableTabs = JSON.parse(editableTabs)
    } else {
      this.editableTabs.push({
        title: '新窗口1',
        uniqueId: 1,
        input: {
          body: '{}',
          method: 'GET',
          path: ''
        },
        sqlStr: 'select * from '
      })
    }
  },
  destroyed() {
    const editableTabs = JSON.stringify(this.editableTabs)
    sessionStorage.setItem('editableTabs', editableTabs)
    sessionStorage.setItem('editableTabsValue', this.editableTabsValue)
  },
  methods: {
    saveData(uniqueId, input, sqlStr, title) {
      for (const editableTabIndex in this.editableTabs) {
        if (this.editableTabs[editableTabIndex].uniqueId == uniqueId) {
          this.editableTabs[editableTabIndex].input = input
          this.editableTabs[editableTabIndex].sqlStr = sqlStr
          this.editableTabs[editableTabIndex].title = title
        }
      }
    },
    async mergeEsPathKeyWords() {
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnectID
      const res = await ListAction(input)

      if (res.code == 0) {
        const list = res.data
        const indices = Object.keys(list)
        this.queryData = []
        const queryData = esPathKeyWords

        for (const esPathKeyWord of queryData) {
          if (esPathKeyWord.value.indexOf('{indices}') !== -1) {
            for (const indice of indices) {
              const mappings = Object.keys(list[indice]['mappings'])
              if (mappings.length == 0) {
                mappings[0] = '{type}'
              }
              const obj = {
                'value': esPathKeyWord.value.replace(/{indices}/g, indice).replace(/{type}/g, mappings[0]),
                'data': esPathKeyWord.data.replace(/{indices}/g, indice).replace(/{type}/g, mappings[0])
              }
              queryData.push(obj)
            }
          }
        }
        this.queryData = queryData
      }
    },
    addTab() {
      this.$prompt('请输入新窗口标题', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      }).then(({ value }) => {
        if (value == '' || value == null) {
          const count = Number(this.editableTabs.length + 1)
          value = '新窗口' + count
        }
        var timestamp = new Date().getTime()
        this.editableTabs.push({
          title: value,
          uniqueId: timestamp,
          input: {
            body: '{}',
            method: 'GET',
            path: ''
          },
          sqlStr: ''
        })
        this.editableTabsValue = timestamp
      }).catch(err => {
        console.log('err', err)
      })
    },
    removeTab(targetName) {
      const tabs = this.editableTabs
      let activeName = this.editableTabsValue
      if (activeName === targetName) {
        tabs.forEach((tab, index) => {
          if (tab.uniqueId === targetName) {
            const nextTab = tabs[index + 1] || tabs[index - 1]
            if (nextTab) {
              activeName = nextTab.uniqueId
            }
          }
        })
      }

      this.editableTabsValue = activeName
      this.editableTabs = tabs.filter(tab => tab.uniqueId !== targetName)
    }
  }
}
</script>
<!--

<template>
  <div class="app-container">
    <el-card class="box-card">
      <tools></tools>
    </el-card>
  </div>
</template>

<script>

export default {

  components: {
    'Tools': () => import('@/views/rest/components/tools')
  },
  data() {
    return {

    }
  },
  computed: {

  },
  created() {

  },
  destroyed() {

  },
  methods: {

  }
}
</script>

-->
