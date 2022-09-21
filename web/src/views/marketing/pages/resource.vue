<template>
  <div class="resource-container">
    <el-card class="opt" shadow="hover" v-show="search.app_id !== ''">
      <i class="el-icon-upload"/>
    </el-card>
    <el-card shadow="hover" class="resource">
      <el-popconfirm icon="el-icon-info" class="remove" icon-color="red" title="确认删除此素材吗?" @onConfirm="removeResource">
        <i class="el-icon-delete" slot="reference"> 删除</i>
      </el-popconfirm>

    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      search: {
        app_id: ''
      }
    }
  },
  methods: {
    initData(app_id) {
      this.$set(this.search, 'app_id', app_id)
    },
    removeResource() {
      this.$message.info("删除素材")
    }
  }
}
</script>

<style lang="scss">
.resource-container {
  .el-card {
    cursor: pointer;
    display: inline-block;
    width: 120px;
    height: 150px;
    margin-right: 10px;
    margin-bottom: 10px;

    .el-card__body {
      padding: 0;
    }
  }
  .opt {
    font-size: 60px;
    font-weight: 800;
    color: #909399;
    text-align: center;
    line-height: 150px;
  }
  .resource {
    position: relative;

    .remove {
      color: #ffffff;
      height: 25px;
      line-height: 25px;
      background: rgba(255, 0, 0, 0.5);
      position: absolute;
      width: 100%;
      text-align: center;
      bottom:  0;
      display: none;
      font-size: 12px;
    }
    &:hover {
      .remove {
        display: block;
      }
    }
  }
}
</style>
