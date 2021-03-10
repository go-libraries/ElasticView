<template>
  <div class="filter-item">
    <el-select v-model="gameNo" placeholder="请选择游戏" clearable filterable @change="change">
      <el-option
        v-for="(v,k,index) in gameMap"
        :key="index"
        :label="v.concat('-',k)"
        :value="Number(k)"
      />
    </el-select>
  </div>
</template>

<script>
import { getGameList } from '@/api/game'
export default {
  name: 'Index',
  props: [
    'game'
  ],
  data() {
    return {
      gameNo: this.game,
      gameMap: []
    }
  },
  mounted() {
    this.getGameList()
  },
  methods: {
    change() {
      this.$emit('changeGame', this.gameNo)
      this.$emit('searchData', 1)
    },
    async getGameList() {
      var input = { 'action': 'all' }
      const res = await getGameList(input)
      this.gameMap = res.game
    }
  }
}
</script>

<style scoped>

</style>
