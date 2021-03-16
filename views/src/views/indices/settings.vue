<template>
  <div>
    <el-dialog :visible.sync="open" title="索引配置" width="70%" @close="closeDialog">
      <el-card class="box-card">
        <el-form label-width="300px" label-position="left">
          <el-form-item label="索引名称">
            <el-input v-model="indexName" placeholder="索引名称" />
          </el-form-item>

          <el-form-item label="分片数">
            <el-input v-model="form.number_of_shards" type="number" style="width: 300px" />
          </el-form-item>
          <el-form-item label="副本数">
            <el-input v-model="form.number_of_replicas" type="number" style="width: 300px" />
          </el-form-item>
          <el-form-item label="索引的刷新时间间隔">
            <el-input v-model="form.refresh_interval" placeholder="索引的刷新时间间隔" />
          </el-form-item>
          <el-form-item label="translog同步到磁盘的时间间隔">
            <el-input v-model="form['index.translog.sync_interval']" />
          </el-form-item>
          <el-form-item label="tanslog同步到本地的方式">
            <el-input v-model="form['index.translog.durability']" />
          </el-form-item>
          <el-form-item label="满足translog同步的容量">
            <el-input v-model="form['index.translog.flush_threshold_size']" />
          </el-form-item>
          <el-form-item label="调高合并的最大线程">
            <el-input v-model="form['index.merge.scheduler.max_thread_count']" />
          </el-form-item>
          <el-form-item label="最大分段大小">
            <el-input v-model="form['index.merge.policy.max_merged_segment']" />
          </el-form-item>
          <el-form-item label="每层所允许的分段数">
            <el-input v-model="form['index.merge.policy.segments_per_tier']" type="number" style="width: 300px" />
          </el-form-item>
          <el-form-item label="是否应在索引打开前检查分片是否损坏">
            <el-select v-model="form['index.shard.check_on_startup']" filterable>
              <el-option label="是" :value="true" />
              <el-option label="否" :value="false" />
            </el-select>
          </el-form-item>

          <el-form-item label="自定义路由值可以转发的目的分片数">
            <el-input v-model="form['index.routing_partition_size']" type="number" style="width: 300px" />
          </el-form-item>

          <el-form-item label="默认使用LZ4压缩方式存储数据">
            <el-input v-model="form['index.codec']" />
          </el-form-item>

          <el-form-item label="基于可用节点的数量自动分配副本数量">
            <el-select v-model="form['index.auto_expand_replicas']" filterable>
              <el-option label="是" :value="true" />
              <el-option label="否" :value="false" />
            </el-select>
          </el-form-item>

          <el-form-item label="用于索引搜索的 from+size 的最大值">
            <el-input v-model="form['index.max_result_window']" type="number" style="width: 300px" />
          </el-form-item>

          <el-form-item label="在搜索此索引中 rescore 的 window_size 的最大值">
            <el-input v-model="form['index.max_rescore_window']" type="number" style="width: 300px" />
          </el-form-item>

          <el-form-item label="允许写入和元数据更改">
            <el-select v-model="form['index.blocks.read_only']" filterable>
              <el-option label="是" :value="true" />
              <el-option label="否" :value="false" />
            </el-select>
          </el-form-item>
          <el-form-item label="禁用对索引的读取操作">
            <el-select v-model="form['index.blocks.read']" filterable>
              <el-option label="是" :value="true" />
              <el-option label="否" :value="false" />
            </el-select>
          </el-form-item>
          <el-form-item label="禁用对索引的写入操作">
            <el-select v-model="form['index.blocks.write']" filterable>
              <el-option label="是" :value="true" />
              <el-option label="否" :value="false" />
            </el-select>
          </el-form-item>
          <el-form-item label="禁用索引元数据的读取和写入">
            <el-select v-model="form['index.blocks.metadata']" filterable>
              <el-option label="是" :value="true" />
              <el-option label="否" :value="false" />
            </el-select>
          </el-form-item>
          <el-form-item label="索引的每个分片上可用的最大刷新侦听器数">
            <el-input v-model="form['index.max_refresh_listeners']" type="number" style="width: 300px" />
          </el-form-item>
        </el-form>
        <div style="text-align:right;">
          <el-button type="danger" icon="el-icon-close" @click="closeDialog">取消</el-button>
          <el-button type="primary" icon="el-icon-check" @click="confirmRole">确认</el-button>
        </div>
      </el-card>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Add',
  components: {

  },
  props: {
    indexName: {
      type: String,
      default: ''
    },
    open: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      isClose: false,
      form: {
        'index.shard.check_on_startup': false,
        'index.routing_partition_size': 1,
        'index.codec': 'LZ4',
        'index.auto_expand_replicas': false,
        'index.max_result_window': 10000,
        'index.max_rescore_window': 10000,
        'index.blocks.read_only': false,
        'index.blocks.read': false,
        'index.blocks.write': false,
        'index.blocks.metadata': true,
        'index.max_refresh_listeners': 10,
        'number_of_shards': 0,
        'number_of_replicas': 0,
        'refresh_interval': '1s',
        'index.translog.sync_interval': '5s',
        'index.translog.durability': 'async',
        'index.translog.flush_threshold_size': '5g',
        'index.merge.scheduler.max_thread_count': 20,
        'index.merge.policy.max_merged_segment': '5gb',
        'index.merge.policy.segments_per_tier': '10'
      },
      dialogVisible: true
    }
  },

  created() {
    console.log(111)
    this.searchData()
  },
  methods: {
    confirmRole() {
      console.log('新增')
    },
    searchData() {
      console.log(this.indexName, 'index')
    },
    closeDialog() {
      this.$emit('close')
    }
  }
}
</script>

<style scoped>

</style>
