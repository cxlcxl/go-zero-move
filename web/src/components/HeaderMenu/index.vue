<template>
  <div id="header-menu">
    <!--    <el-menu :default-active="activeIndex" @select="handleSelect" mode="horizontal">
      <el-menu-item index="1">基础信息</el-menu-item>
      <el-menu-item index="2">处理中心</el-menu-item>
    </el-menu>-->
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeIndex: '1'
    }
  },
  created() {

  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style lang="scss" scoped>
#header-menu {
  margin-left: 10px;
  display: inline-block;
  .el-menu {
    border-right: none !important;
    .el-menu-item {
      height: 50px !important;
      line-height: 50px !important;
    }
  }
}
</style>
