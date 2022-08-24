<template>
  <el-drawer :title="title" :visible.sync="visible" :size="width" append-to-body :direction="direction"
             :wrapperClosable="false" :show-close="false" :close-on-press-escape="false">
    <div style="padding: 0 15px 0 0;">
      <slot />
    </div>

    <div style="text-align: center;">
      <el-button type="primary" @click="handleConfirm" v-show="confirmText" icon="el-icon-check"
                 :loading="confirmLoading">{{confirmText}}</el-button>
      <el-button @click="handleCancel" :loading="confirmLoading" icon="el-icon-close">取消</el-button>
    </div>
  </el-drawer>
</template>

<script>
  export default {
    props: {
      confirmText: {
        default: "",
        type: String
      },
      title: {
        default: "",
        type: String
      },
      width: {
        default: "500px",
        type: String
      },
      confirmLoading: {
        default: false,
        type: Boolean
      },
      visible: {
        default: false,
        type: Boolean
      },
      direction: {
        default: "rtl",
        type: String
      }
    },
    methods: {
      handleCancel() {
        this.$emit('cancel')
      },
      handleConfirm() {
        this.$emit('confirm')
      }
    }
  }
</script>

<style lang="scss">
  .el-drawer__body {
    overflow: auto;
    padding-bottom: 15px;
  }
</style>
