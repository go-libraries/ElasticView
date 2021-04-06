<template>
  <div>
    <el-dialog :visible.sync="open" title="新增仓库" @close="closeDialog">
      <el-card class="box-card">
        <el-form label-width="500px" label-position="left">
          <el-form-item label="仓库名">
            <el-input v-model="form.name" placeholder="仓库名"/>
          </el-form-item>
          <el-form-item label="类型（type）">
            <el-input v-model="form.type" placeholder="类型（type）"/>
          </el-form-item>
          <el-form-item label="位置（location）">
            <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 4}" v-model="form.location" placeholder="位置（location）"/>
          </el-form-item>
          <el-form-item label="是否压缩 (compress)">
            <el-select v-model="form.compress" clearable filterable>
              <el-option label="是" value="true"/>
              <el-option label="否" value="false"/>
            </el-select>
          </el-form-item>
          <el-form-item label="节点恢复速率">
            <el-input v-model="form.max_restore_bytes_per_sec" placeholder="节点恢复速率"/>
          </el-form-item>
          <el-form-item label="每个节点快照速率">
            <el-input v-model="form.max_snapshot_bytes_per_sec" placeholder="每个节点快照速率"/>
          </el-form-item>
        </el-form>
        <div style="text-align:right;">
          <el-button type="danger" icon="el-icon-close" @click="closeDialog">取消</el-button>
          <el-button type="primary" icon="el-icon-check" @click="confirm">确认</el-button>
        </div>
      </el-card>
    </el-dialog>
  </div>
</template>

<script>
    import {SnapshotCreateRepositoryAction} from '@/api/es-backup'

    export default {
        name: 'Add',
        components: {},
        props: {
            open: {
                type: Boolean,
                default: false
            },
            snapshotData: {
                type: Object,
                default: null
            }
        },
        data() {
            return {
                isOpen: false,
                form: {
                    name: '',
                    type: 'fs',
                    location: '',
                    compress: "false",
                    max_restore_bytes_per_sec: '40mb',
                    max_snapshot_bytes_per_sec: '40mb'
                }
            }
        },
        computed: {},

        created() {
            if (this.snapshotData != null) {
                this.form = Object.assign({}, this.snapshotData)
            }
        },
        methods: {
            closeDialog() {
                this.$emit('close', false)
            },
            async confirm() {
                let input = this.form
                input['es_connect'] = this.$store.state.baseData.EsConnectID
                const {code, data, msg} = await SnapshotCreateRepositoryAction(input)
                if (code == 0) {
                    this.$emit('close', true)
                    this.$message({
                        type: 'success',
                        message: msg
                    })
                    return
                } else {
                    this.$message({
                        type: 'error',
                        message: msg
                    })
                    return
                }
            }
        }
    }
</script>

<style scoped>

</style>
