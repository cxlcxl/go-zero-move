<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.promotion_name" class="w230" clearable placeholder="计划名称"/>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">创建计划</el-button>
    </el-col>
    <el-col :span="24" style="margin-top: 15px">

    </el-col>
  </el-row>
</template>

<script>
export default {
  //name: 'PromotionList',
  data() {
    return {
      loadings: {
        pageLoading: false
      },
      userList: {
        total: 0,
        list: []
      },
      roles: [],
      search: {
        promotion_name: '',
        page: 1,
        page_size: 10
      }
    }
  },
  methods: {
    doSearch() {},
    add() {

    }
  }
}
</script>

<style scoped>

</style>
