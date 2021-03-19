<template>
  <div class="app-container">
    <back-to-top />
    <el-card class="box-card">
      <el-tabs v-model="activeName">
        <el-tab-pane label="查看索引信息" name="CatIndices">
          <cat-indices v-if="activeName == 'CatIndices'" />
        </el-tab-pane>
        <el-tab-pane label="显示别名,过滤器,路由信息" name="CatAliases">
          <cat-aliases v-if="activeName == 'CatAliases'" />
        </el-tab-pane>
        <el-tab-pane label="显示每个节点分片数量、占用空间" name="CatAllocation">
          <cat-allocation v-if="activeName == 'CatAllocation'" />
        </el-tab-pane>
        <el-tab-pane label="显示索引文档的数量" name="CatCount">
          <cat-count v-if="activeName == 'CatCount'" />
        </el-tab-pane>
        <el-tab-pane label="查看集群健康状况" name="CatHealth">
          <cat-health v-if="activeName == 'CatHealth'" />
        </el-tab-pane>
        <el-tab-pane label="显示索引分片信息" name="CatShards">
          <cat-shards v-if="activeName == 'CatShards'" />
        </el-tab-pane>
        <el-tab-pane label="显示任务信息" name="CatTask">
          <cat-task v-if="activeName == 'CatTask'" />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
export default {
  components: {
    'BackToTop': () => import('@/components/BackToTop/index'),
    'CatAliases': () => import('@/components/Cat/CatAliases'),
    'CatAllocation': () => import('@/components/Cat/CatAllocation'),
    'CatCount': () => import('@/components/Cat/CatCount'),
    'CatHealth': () => import('@/components/Cat/CatHealth'),
    'CatIndices': () => import('@/components/Cat/CatIndices'),
    'CatShards': () => import('@/components/Cat/CatShards'),
    'CatTask': () => import('@/components/Cat/CatTask')

  },
  data() {
    return {
      activeName: 'CatIndices'
    }
  },
  computed: {},
  created() {

  },
  methods: {}
}
</script>

<style lang="scss" scoped>
  .app-container {
    .roles-table {
      margin-top: 30px;
    }

    .permission-tree {
      margin-bottom: 30px;
    }
  }
</style>
