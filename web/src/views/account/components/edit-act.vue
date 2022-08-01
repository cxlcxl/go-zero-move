<template>
  <dialog-panel title="账号修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="accountForm" ref="accountForm" label-width="120px" size="small" :rules="userRules">
      <el-form-item label="账号名称" prop="account_name">
        <el-input v-model="accountForm.account_name" placeholder="请填写账号名称"/>
      </el-form-item>
      <el-form-item label="平台账号ID" prop="platform_id">
        <el-input v-model="accountForm.platform_id" placeholder="请填写广告平台系统账号 ID"/>
      </el-form-item>
      <el-form-item label="账号类型" prop="account_type">
        <el-select v-model="accountForm.account_type" style="width: 100%;">
          <el-option v-for="(key, val) in accountTypes" :label="key" :value="Number(val)"/>
        </el-select>
      </el-form-item>
      <el-form-item label="开户服务商" prop="provider_account_id">
        <el-select v-model="accountForm.provider_account_id" style="width: 100%;" clearable>
          <el-option v-for="(key, val) in provider" :label="key" :value="Number(val)"/>
        </el-select>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {accountUpdate} from '@a/account'

  export default {
    components: {
      DialogPanel
    },
    props: {
      accountTypes: Object,
      provider: Object,
    },
    data() {
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        accountForm: {
          id: 0,
          account_name: '',
          platform_id: '',
        },
        userRules: {
          account_name: {required: true, message: '请填写账号名称'},
          platform_id: {required: true, message: '请填写广告平台系统账号 ID'},
          account_type: {required: true, message: '请选择账号类型'},
        },
        customers: []
      }
    },
    methods: {
      initUpdate(row) {
        this.accountForm = row
        this.visible = true
      },
      cancel() {
        this.$refs.accountForm.resetFields()
        this.customers = []
        this.visible = false
      },
      save() {
        this.$refs.accountForm.validate(v => {
          if (v) {
            this.loading = true
            accountUpdate(this.accountForm).then(res => {
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
