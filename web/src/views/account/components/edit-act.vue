<template>
  <dialog-panel title="账户信息修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="accountForm" ref="accountForm" label-width="120px" size="small" :rules="userRules">
      <el-form-item label="账号名称" prop="account_name">
        <el-input v-model="accountForm.account_name" placeholder="请填写账号名称"/>
      </el-form-item>
      <el-form-item label="广告主 ID" prop="advertiser_id">
        <el-input v-model="accountForm.advertiser_id" placeholder="请填写广告主 ID"/>
      </el-form-item>
      <el-form-item label="账号类型" prop="account_type">
        <el-select v-model="accountForm.account_type" style="width: 100%;">
          <el-option v-for="(key, val) in accountTypes" :label="key" :value="Number(val)"/>
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-switch v-model="accountForm.state" :active-value="1" :inactive-value="0"/>
      </el-form-item>
      <el-form-item label="开发者 ID" prop="developer_id">
        <el-input v-model="accountForm.developer_id" placeholder="请填写开发者 ID"/>
      </el-form-item>
      <el-form-item label="ClientID" prop="client_id">
        <el-input v-model="accountForm.client_id" :disabled="lockSecret"/>
      </el-form-item>
      <el-form-item label="Secret" prop="secret">
        <el-input v-model="accountForm.secret" :disabled="lockSecret"/>
      </el-form-item>
      <el-form-item prop="parent_id" label="上级账户">
        <el-select v-model="accountForm.parent_id" remote filterable placeholder="可输入名称查询"
          :remote-method="remoteMethod" :loading="remoteLoading" style="width: 100%;">
          <el-option v-for="item in accounts" :label="item.account_name" :value="Number(item.id)"
            v-show="Number(item.id) !== Number(accountForm.id)"/>
        </el-select>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import { accountInfo, accountUpdate, parentAccounts } from '@a/account'

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
        lockSecret: false,
        remoteLoading: false,
        accounts: [],
        accountForm: {
          id: 0,
          parent_id: '',
          account_name: '',
          advertiser_id: '',
          client_id: '',
          is_auth: 0,
          developer_id: '',
          secret: '',
          account_type: 1,
          state: 1,
        },
        userRules: {
          account_name: {required: true, message: '请填写账户名称'},
          advertiser_id: {required: true, message: '请填写广告主 ID'},
          account_type: {required: true, message: '请选择账户类型'},
          state: {required: true, message: '请选择'},
        },
      }
    },
    methods: {
      initUpdate(id) {
        accountInfo(id).then(res => {
          this.accountForm = res.data
          this.getDefaultAccounts(this.accountForm.parent_id)
          this.visible = true
        }).catch(() => {
          this.$message.error("用户信息请求错误")
        })
      },
      getDefaultAccounts(parent_id) {
        parentAccounts({parent_id}).then(res => {
          if (Array.isArray(res.data)) {
            this.accounts = res.data
          } else {
            this.accounts = []
          }
          this.accounts.push({id: 0, account_name: '无上级账户'})
        }).catch(err => {
          this.$message.error("账户信息请求错误")
        })
      },
      cancel() {
        this.$refs.accountForm.resetFields()
        this.visible = false
      },
      save() {
        this.$refs.accountForm.validate(v => {
          if (v) {
            this.loading = true
            this.$set(this.accountForm, 'parent_id', Number(this.accountForm.parent_id))
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
      },
      remoteMethod(query) {
        if (query.trim() !== '') {
          this.remoteLoading = true;
          parentAccounts({account_name: query}).then(res => {
            this.remoteLoading = false
            if (Array.isArray(res.data)) {
              this.accounts = res.data
            } else {
              this.accounts = []
            }
          }).catch(() => {
            this.remoteLoading = false
          })
        } else {
          this.options = [];
        }
      },
    }
  }
</script>
