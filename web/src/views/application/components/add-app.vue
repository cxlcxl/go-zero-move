<template>
  <dialog-panel title="添加应用" confirm-text="添加" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading">
    <el-form :model="appForm" ref="appForm" label-width="120px" size="small" :rules="userRules">
      <el-form-item label="应用名称" prop="app_name">
        <el-input v-model="appForm.app_name" placeholder="请填写应用名称"/>
      </el-form-item>
      <el-form-item label="APP ID" prop="app_id">
        <el-input v-model="appForm.app_id" placeholder="请填写APP ID"/>
      </el-form-item>
      <el-form-item prop="account_id" label="关联账户">
        <el-select v-model="appForm.account_id" remote filterable placeholder="可输入名称查询" @change="handleChange"
                   :remote-method="remoteMethod" :loading="remoteLoading" style="width: 100%;">
          <el-option v-for="item in accounts" :label="item.account_name" :value="Number(item.id)"/>
        </el-select>
      </el-form-item>
      <el-form-item label="应用包名" prop="pkg_name">
        <el-input v-model="appForm.pkg_name" placeholder="请填写应用包名"/>
      </el-form-item>
      <el-form-item label="渠道" prop="channel">
        <el-select v-model="appForm.channel" style="width: 100%;">
          <el-option v-for="(key, val) in appChannel" :label="key" :value="Number(val)"/>
        </el-select>
      </el-form-item>
      <el-form-item label="应用类型" prop="app_type">
        <el-select v-model="appForm.app_type" style="width: 100%;">
          <el-option v-for="(key, val) in appType" :label="key" :value="val"/>
        </el-select>
      </el-form-item>
      <el-form-item label="标签" prop="tags">
        <el-input v-model="appForm.tags" placeholder="多个请用英文 ',' 号隔开"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {appCreate} from "@a/app"
  import {searchAccounts, defaultAccounts} from "@a/account"

  export default {
    components: {
      DialogPanel
    },
    props: {
      appType: Object,
      appChannel: Object,
    },
    data() {
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        appForm: {
          app_name: '',
          app_id: '',
          account_id: '',
          advertiser_id: '',
          pkg_name: '',
          tags: '',
          channel: '',
          app_type: '',
        },
        accounts: [],
        userRules: {
          app_name: {required: true, message: '请填写应用名称'},
          app_id: {required: true, message: '请填写 APP ID'},
          channel: {required: true, message: '请选择渠道'},
          pkg_name: {required: true, message: '请填写应用包名'},
          account_id: {required: true, message: '请关联账户'},
        },
      }
    },
    methods: {
      initCreate() {
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
      add() {
        this.$refs.appForm.validate(v => {
          if (v) {
            this.loading = true
            appCreate(this.appForm).then(res => {
              this.$message.success("创建成功")
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
