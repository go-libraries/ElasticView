<template>
  <div>
    <el-dialog :visible.sync="open" :title="title" width="95%" @close="closeDialog">
      <el-card class="box-card">
        <div class="filter-container">
          <el-tag class="filter-item">类型名</el-tag>
          <el-input
            v-model="typeName"
            placeholder="类型名"
            :readonly="typeName!=''"
            style="width: 500px"
            class="filter-item"
            clearable
          />
          <el-tag class="filter-item">dynamic</el-tag>

          <el-select v-model="dynamic" class="filter-item" clearable filterable>
            <el-option label="默认" value="" />
            <el-option label="true" value="true" />
            <el-option label="false" value="false" />
            <el-option label="strict" value="strict" />
          </el-select>
        </div>
        <mapping-table :mappings="mappingList" />
      </el-card>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Mapping',
  components: {
    'MappingTable': () => import('@/views/indices/components/mappingTable')
  },
  props: {
    indexName: {
      type: String,
      default: ''
    },
    mappings: {
      type: Object,
      default: {}
    },
    open: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: '新增映射结构'
    }
  },

  data() {
    return {
      connectLoading: false,
      mappingList: [],
      typeName: '',
      dynamic: ''
    }
  },
  mounted() {
    for (const k in this.mappings) {
      this.typeName = k
    }
    if (this.typeName != '') {
      if (this.mappings[this.typeName].hasOwnProperty('dynamic')) {
        this.dynamic = this.mappings[this.typeName]['dynamic']
      }

      this.mappingList = this.mappings[this.typeName]
    }
  },
  methods: {
    closeDialog() {
      this.$emit('close')
    }
  }
}
</script>

<style scoped>

</style>
