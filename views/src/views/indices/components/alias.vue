<template>
  <div>
    <el-form>
      <el-form-item>
        <el-button icon="el-icon-plus" @click="addAlias">新增别名</el-button>
        <el-button icon="el-icon-delete" @click="getAlias">重置</el-button>
      </el-form-item>
      <el-form-item
        v-for="(alias, index) in aliasList"
        :key="index"
        :label="'别名' + Number(index+1)"
      >
        <el-input v-model="aliasList[index].name" style="width:300px" />
        <el-button type="primary" icon="el-icon-check" @click="submitForm(index)">提交</el-button>
        <el-button icon="el-icon-delete" @click.prevent="removeAlias(index)">删除</el-button>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>
import { GetAliasAction } from '@/api/es-index'
export default {
  name: 'Alias',
  props: {
    indexName: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      aliasList: []
    }
  },
  mounted() {
    this.getAlias()
  },
  methods: {
    async getAlias() {
      this.aliasList = []
      const input = {}
      input['es_connect'] = this.$store.state.baseData.EsConnect
      input['index_name'] = this.indexName
      const { data, code, msg } = await GetAliasAction(input)
      console.log(data, code, msg)
      if (code == 0) {
        for (const k in data) {
          this.aliasList.push({ name: data[k].AliasName, types: 'old' })
        }
      } else {
        this.$message({
          message: msg,
          type: 'error'
        })
      }
    },
    submitForm() {

    },
    removeAlias(index) {
      if (this.aliasList[index].types == 'new') {
        this.aliasList.splice(index, 1)
      }
    },
    addAlias() {
      this.aliasList.push({ name: '', types: 'new' })
    }
  }
}
</script>

<style scoped>

</style>
