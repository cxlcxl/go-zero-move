<template>
  <el-dialog :title="title" :visible.sync="visible" :width="width" append-to-body modal-append-to-body
             :close-on-click-modal="false" :close-on-press-escape="false" v-el-drag-dialog :show-close="false">
    <slot/>

    <div slot="footer" class="dialog-footer">
      <el-button @click="handleCancel" :loading="confirmLoading" icon="el-icon-close">取消</el-button>
      <el-button type="primary" @click="handleConfirm" icon="el-icon-check" v-show="confirmText"
                 :loading="confirmLoading">{{confirmText}}</el-button>
      <slot name="operate"/>
    </div>
  </el-dialog>
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
        default: "600px",
        type: String
      },
      confirmLoading: {
        default: false,
        type: Boolean
      },
      visible: {
        default: false,
        type: Boolean
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

<style>
  .el-dialog__body {
    padding: 15px 20px 10px;
  }
</style>