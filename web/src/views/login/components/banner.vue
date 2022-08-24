<template>
  <el-carousel :style="{height: '100%'}">
    <el-carousel-item v-for="item in images" :style="{height: '100%'}">
      <img :src="item" :style="{height: '100%', width: '100%'}"/>
    </el-carousel-item>
  </el-carousel>
</template>

<script>
  export default {
    data() {
      return {
        images: [
          '/images/1.jpg',
          '/images/1.jpg',
          '/images/1.jpg',
          '/images/1.jpg',
        ]
      }
    }
  }
</script>

<style>
  .el-carousel__container {
    height: 100% !important;
  }
</style>