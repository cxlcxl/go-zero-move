<template>
  <dialog-panel title="添加账号" confirm-text="添加" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading">
    <el-form :model="accountForm" ref="accountForm" size="small" :rules="userRules">
      <el-form-item label="账号名称" prop="account_name" label-width="120px">
        <el-input v-model="accountForm.account_name" placeholder="请填写账号名称"/>
      </el-form-item>
      <el-form-item label="账号类型" prop="account_type" label-width="120px">
        <el-select v-model="accountForm.account_type" style="width: 100%;">
          <el-option v-for="(key, val) in accountTypes" :label="key" :value="Number(val)"/>
        </el-select>
      </el-form-item>
      <el-form-item label="平台账号ID" prop="platform_id" label-width="120px">
        <el-input v-model="accountForm.platform_id" placeholder="请填写广告平台系统账号 ID"/>
      </el-form-item>
      <el-form-item label="开户服务商" prop="provider_account_id" label-width="120px">
        <el-select v-model="accountForm.provider_account_id" style="width: 100%;">
          <el-option v-for="(key, val) in provider" :label="key" :value="Number(val)"/>
        </el-select>
      </el-form-item>
      <el-tabs type="border-card" v-model="accountForm.add_type">
        <el-tab-pane label="新客户信息填写" name="new">
          <el-form-item prop="customer_name"
                        :rules="{required: accountForm.add_type === 'new', message: '请选择开户服务商'}">
            <el-input v-model="accountForm.customer_name" placeholder="客户名称（简称）"/>
          </el-form-item>
          <el-form-item prop="full_name">
            <el-input v-model="accountForm.full_name" placeholder="客户全称"/>
          </el-form-item>
          <el-form-item prop="weight">
            <el-input v-model.number="accountForm.weight" placeholder="权重（数字越小添加流水客户下拉面板越靠前）"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane label="已有客户直接绑定" name="select">
          <el-form-item prop="customer_id"
                        :rules="{required: accountForm.add_type === 'select', message: '请选择客户'}">
            <el-select v-model="accountForm.customer_id" remote filterable placeholder="请输入客户名称查询" clearable
                       :remote-method="remoteMethod" :loading="remoteLoading" style="width: 100%; flex: 1;">
              <el-option v-for="item in customers" :label="item.customer_name" :value="item.id">
                <span style="float: left">{{ item.customer_name }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{ item.full_name }}</span>
              </el-option>
            </el-select>
          </el-form-item>
        </el-tab-pane>
      </el-tabs>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {getDefaultCustomers, searchCustomer} from "@a/customer"
  import {accountCreate} from "@a/account"

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
          add_type: 'new',
          account_name: '',
          customer_id: '',
          account_type: '',
          platform_id: '',
          provider_account_id: 2,
          customer_name: '',
          weight: 99,
          full_name: '',
        },
        providerActs: [],
        userRules: {
          account_name: {required: true, message: '请填写账号名称'},
          account_type: {required: true, message: '请选择账号类型'},
          platform_id: {required: true, message: '请填写广告平台系统账号 ID'},
          provider_account_id: {required: true, message: '请选择开户服务商'},
        },
        customers: [],
      }
    },
    methods: {
      initCreate() {
        Promise.all([
          this.defaultCustomers(),
        ]).then(res => {
          this.customers = res[0]
          this.providerActs = res[1]
          this.visible = true
        }).catch(() => {})
      },
      defaultCustomers() {
        return new Promise((resolve, reject) => {
          getDefaultCustomers().then(res => {
            resolve(res.data)
          }).catch(() => {
            reject()
          })
        })
      },
      cancel() {
        this.$refs.accountForm.resetFields()
        this.customers = []
        this.visible = false
      },
      add() {
        this.$refs.accountForm.validate(v => {
          if (v) {
            this.loading = true
            this.accountForm.customer_id = Number(this.accountForm.customer_id)
            accountCreate(this.accountForm).then(res => {
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
      remoteMethod(query) {
        if (query.trim() !== '') {
          this.remoteLoading = true;
          searchCustomer(query).then(res => {
            this.remoteLoading = false
            this.customers = res.data
          }).catch(() => {
            this.remoteLoading = false
          })
        } else {
          this.options = [];
        }
      }
    }
  }
</script>
