<template>
  <dialog-panel title="应用拉取" confirm-text="确认拉取" :visible="visible" @cancel="cancel" @confirm="pull" :confirm-loading="loading">
    <el-form :model="appForm" ref="appForm" label-width="120px" size="small" :rules="userRules">
      <el-form-item prop="account_id" label="账户">
        <el-select v-model="appForm.account_id" remote filterable placeholder="可输入名称查询" @change="handleChange"
                   :remote-method="remoteMethod" :loading="remoteLoading" style="width: 100%;">
          <el-option v-for="item in accounts" :label="item.account_name" :value="Number(item.id)"/>
        </el-select>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {appPull} from "@a/app"
  import {searchAccounts, defaultAccounts} from "@a/account"

  export default {
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        appForm: {
          account_id: '',
          advertiser_id: '',
        },
        accounts: [],
        userRules: {
          account_id: {required: true, message: '请选择账户'},
        },
      }
    },
    methods: {
      initPull() {
        defaultAccounts().then(res => {
          this.visible = true
          if (Array.isArray(res.data)) {
            this.accounts = res.data
          } else {
            this.accounts = []
          }
        }).catch(err => {
          this.$message.error("请求错误")
        })
      },
      cancel() {
        this.$refs.appForm.resetFields()
        this.visible = false
      },
      pull() {
        this.$refs.appForm.validate(v => {
          if (v) {
            this.loading = true
            appPull(this.appForm).then(res => {
              this.$message.success("拉取成功")
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
      handleChange(v) {
        for (let i = 0; i < this.accounts.length; i++) {
          if (Number(v) === this.accounts[i].id) {
            this.$set(this.appForm, 'advertiser_id', this.accounts[i].advertiser_id)
          }
        }
      },
      remoteMethod(query) {
        if (query.trim() !== '') {
          this.remoteLoading = true;
          searchAccounts(query).then(res => {
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
