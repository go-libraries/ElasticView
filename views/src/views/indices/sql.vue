<template>
  <div class="big">
    <div class="box1">
      <el-input class="tmp" />
    </div>
    <div class="box2" />
    <div class="box3" />
  </div>
</template>

<style>
  * {
    margin: 0;
    padding: 0;
  }
</style>

<style scoped>
  .tmp{
    width: 300px;

  }
  .big {
    width: 100px;
    height: 100px;
    background: yellowgreen;
  }

  .big div {
    height: 100px;
  }

  .big img {
    width: 710px;
    height: 400px;
  }

  .box1 {
    width: 960px;
    background: tomato;
  }

  .box2 {
    width: 1920px;
    background-color: darkred;
  }
  .box3 {
    width: 1920px;
    background-color: #e1b59f;
  }
</style>
