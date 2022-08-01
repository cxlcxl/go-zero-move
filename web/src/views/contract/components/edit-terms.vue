<template>
  <dialog-panel title="条款信息修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="termsForm" ref="termsForm" label-width="100px" size="small">
      <el-form-item label="条款名称" prop="_name" :rules="{required: true, message: '请填写条款名称'}">
        <el-input v-model="termsForm._name"/>
      </el-form-item>
      <el-form-item label="条款代码" prop="_key" :rules="{required: true, message: '请填写条款代码'}">
        <el-input v-model="termsForm._key" placeholder="仅支持字符开头的大小写字母数字下划线组合 [50 位以内]"/>
        <i class="el-icon-info" style="font-size: 12px; color: #909399;"> 涉及流水计算的代码：流水模块大写字母 + '_' + 数据类型大写字母</i>
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-switch v-model="termsForm.state" :active-value="1" :inactive-value="0"/>
      </el-form-item>
      <el-form-item label="条款描述" prop="desc">
        <el-input v-model="termsForm.desc"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {termsUpdate} from '@a/contract-terms'

  export default {
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        termsForm: {
          id: 0,
          _name: '',
          _key: '',
          desc: '',
          state: 1,
        }
      }
    },
    methods: {
      cancel() {
        this.$refs.termsForm.resetFields()
        this.visible = false
      },
      save() {
        this.$refs.termsForm.validate(v => {
          if (v) {
            this.loading = true
            termsUpdate(this.termsForm).then(res => {
              this.$message.success("修改成功")
              this.$emit('success')
              this.loading = false
              this.cancel()
            }).catch(err => {
              this.loading = false
              console.log(err)
            })
          } else {
            return false
          }
        })
      }
    }
  }
</script>
