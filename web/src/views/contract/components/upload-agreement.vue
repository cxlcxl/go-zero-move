<template>
  <dialog-panel title="协议/附件上传" :visible="visible" @confirm="add" @cancel="cancel" confirm-text="确定"
                :confirm-loading="loading">
    <el-form ref="agreement" :model="agreementForm" :rules="rules" label-width="100px" class="contract-form">
      <el-form-item label="附件" prop="file_id">
        <upload ref="upload" accept=".pdf" @success="uploadSuccess"
                trigger-text="选取文件（PDF 扫描文件，且不超过 5M），必需上传"/>
      </el-form-item>
      <el-form-item label="协议名称" prop="agreement_name">
        <el-input v-model="agreementForm.agreement_name"/>
      </el-form-item>
      <el-form-item label="描述" prop="desc">
        <el-input v-model="agreementForm.desc"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import Upload from '@c/Upload'
  import {agreementCreate} from "@a/contract"

  export default {
    components: {
      DialogPanel, Upload
    },
    data() {
      return {
        visible: false,
        loading: false,
        agreementForm: {
          contract_id: 0,
          file_id: 0,
          agreement_name: '',
          desc: '',
        },
        rules: {
          agreement_name: {required: true, message: '请填写协议名称'},
        },
        valid_dates: []
      }
    },
    methods: {
      initPanel(contract_id) {
        this.$set(this.agreementForm, 'contract_id', contract_id)
        this.visible = true
      },
      cancel() {
        this.$refs.upload.fileList = []
        this.$refs.agreement.resetFields()
        this.valid_dates = []
        this.visible = false
      },
      uploadSuccess(data) {
        this.$set(this.agreementForm, 'file_id', data.id)
        this.doUpload()
      },
      add() {
        this.$refs.agreement.validate(v => {
          if (v) {
            if (this.agreementForm.file_id > 0) {
              this.doUpload()
            } else {
              this.$refs.upload.submit()
            }
          }
        })
      },
      doUpload() {
        this.loading = true
        agreementCreate(this.agreementForm).then(res => {
          this.$message.success('上传成功')
          this.loading = false
          this.cancel()
          this.$emit('success')
        }).catch(() => {
          this.loading = false
        })
      }
    }
  }
</script>
