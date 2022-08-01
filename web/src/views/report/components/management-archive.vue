<template>
  <dialog-panel title="收支数据归档" :visible="visible" :confirm-loading="loading" confirm-text="确定" @confirm="confirm" @cancel="cancel">
    <el-form ref="form" :model="archiveForm" :rules="archiveRules">
      <el-form-item prop="quarters">
        <el-checkbox-group v-model="archiveForm.quarters">
          <el-checkbox v-for="item in archiveQs" :label="item" />
        </el-checkbox-group>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from '@c/DialogPanel'
import { archiveQuarters, saveArchive } from '@a/report'

export default {
  name: 'ManagementArchive',
  components: {
    DialogPanel
  },
  data() {
    return {
      visible: false,
      loading: false,
      archiveQs: [],
      archiveForm: {
        quarters: []
      },
      archiveRules: {
        quarters: { required: true, message: '请选择归档季度' }
      }
    }
  },
  methods: {
    getArchiveQs() {
      this.loading = true
      archiveQuarters().then(res => {
        this.archiveQs = res.data
        this.loading = false
        this.visible = true
      }).catch(() => {
        this.loading = false
      })
    },
    confirm() {
      this.$refs.form.validate(v => {
        if (v) {
          if (!this.checkQuarters()) {
            this.$message.error('需按顺序归档数据')
            return
          }
          this.loading = true
          saveArchive(this.archiveForm).then(res => {
            this.$message.success('归档成功')
            this.$emit('success')
            this.cancel()
          }).catch(() => {
            this.loading = false
          })
        }
      })
    },
    checkQuarters() {
      let pass = true
      let selectIdx = 0
      for (let i = this.archiveQs.length - 1; i >= 0; i--) {
        if (this.archiveForm.quarters.includes(this.archiveQs[i])) {
          selectIdx = i
        } else {
          if (selectIdx > i) {
            pass = false
            break
          }
        }
      }
      return pass
    },
    cancel() {
      this.$refs.form.resetFields()
      this.visible = false
    }
  }
}
</script>
