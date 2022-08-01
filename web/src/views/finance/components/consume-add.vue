<template>
  <dialog-panel
    title="添加消耗"
    confirm-text="添加"
    :visible="visible"
    width="900px"
    :confirm-loading="loading"
    @cancel="cancel"
    @confirm="confirm"
  >
    <el-form ref="form" :model="flowForm" :rules="rules" label-width="120px" size="mini">
      <el-form-item prop="customer_id" label="客户">
        <el-select
          v-model="flowForm.customer_id"
          remote
          filterable
          placeholder="请输入客户名称查询"
          clearable
          style="width: 100%;"
          :remote-method="remoteMethod"
          :loading="remoteLoading"
          @change="handleSelect"
        >
          <el-option v-for="item in customers" :label="item.customer_name" :value="item.id">
            <span style="float: left">{{ item.customer_name }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.full_name }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item prop="account_id" label="账号">
        <el-select v-model="flowForm.account_id" placeholder="请选择客户账号" style="width: 100%;" @change="handleSelectAccount">
          <el-option v-for="account in accounts" :label="account.account_name" :value="account.id" />
        </el-select>
      </el-form-item>
      <el-form-item prop="accounts" label="金额">
        <div v-if="userAccounts.length > 0" class="account-container">
          <div v-for="item in userAccounts" class="account-card">
            <div class="account-name">{{ item.account_name }}</div>
            <input v-model="flowForm.accounts[item.account_type]" type="text" class="number">
          </div>
        </div>
        <span v-else>请选择账号</span>
      </el-form-item>
      <el-form-item prop="trans_date" label="交易日期">
        <el-date-picker
          v-model="flowForm.trans_date"
          type="date"
          placeholder="开始时间"
          value-format="yyyy-MM-dd"
          style="width: 100%;"
        />
      </el-form-item>
      <el-form-item prop="consume_type" label="消耗类型">
        <el-radio-group v-model="flowForm.consume_type" @change="selectRate">
          <el-radio v-for="(name, val) in consumeTypes" :label="val">{{ name }}</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item prop="settlement_rate" label="返点比例(%)">
        <el-input-number v-model="flowForm.settlement_rate" :min="0" class="w200" :step="1" />
      </el-form-item>
      <el-form-item prop="remark" label="备注">
        <el-input v-model="flowForm.remark" type="textarea" />
      </el-form-item>
    </el-form>

    <el-button slot="operate" icon="el-icon-check" type="primary" :loading="loading" @click="continueCreate">添加并继续</el-button>
  </dialog-panel>
</template>

<script>
import DialogPanel from '@c/DialogPanel'
import { consumeCreate } from '@/api/marketing'
import { getAccounts } from '@a/account'
import { getDefaultCustomers, searchCustomer, flowBelongsAccount } from '@a/customer'
import { parseTime } from '@/utils'

export default {
  components: {
    DialogPanel
  },
  props: {
    consumeTypes: Object
  },
  data() {
    const checkAmount = (rule, value, callback) => {
      if (!value instanceof Object) {
        callback(new Error('请填写金额'))
      } else {
        let allEmpty = true
        for (const account in value) {
          const v = Number(value[account])
          if (!/^(-)?[0-9]+(\.?[0-9]+)?$/.test(v)) {
            return callback(new Error('金额只能填写数字或小数'))
          }
          if (v !== 0) {
            allEmpty = false
          }
        }
        if (allEmpty) {
          return callback(new Error('金额至少填写一个'))
        }
        callback()
      }
    }
    return {
      visible: false,
      loading: false,
      remoteLoading: false,
      customers: [],
      accounts: [],
      userAccounts: [],
      flowForm: {
        customer_id: '',
        account_id: '',
        consume_type: '',
        accounts: {},
        trans_date: parseTime(new Date(), '{y}-{m}-{d}'),
        settlement_rate: 0,
        remark: ''
      },
      rules: {
        accounts: { validator: checkAmount },
        customer_id: { required: true, message: '请选择客户' },
        account_id: { required: true, message: '请选择客户账号' },
        trans_date: { required: true, message: '请选择交易日期' },
        consume_type: { required: true, message: '请选择消耗类型' }
      }
    }
  },
  methods: {
    initCreate() {
      Promise.all([
        this.defaultCustomers()
      ]).then(res => {
        this.customers = res[0]
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
    confirm() {
      this.$refs.form.validate(v => {
        if (v) {
          this.loading = true
          for (const account in this.flowForm.accounts) {
            this.flowForm.accounts[account] = Number(this.flowForm.accounts[account])
          }
          consumeCreate(this.flowForm).then(res => {
            this.$message.success('添加成功')
            this.$emit('success')
            this.cancel()
          }).catch(() => {
            this.loading = false
          })
        }
      })
    },
    continueCreate() {
      this.$refs.form.validate(v => {
        if (v) {
          this.loading = true
          for (const account in this.flowForm.accounts) {
            this.flowForm.accounts[account] = Number(this.flowForm.accounts[account])
          }
          consumeCreate(this.flowForm).then(res => {
            this.$message.success('添加成功，请继续填写流水信息')
            this.$emit('success')
            this.loading = false
            this.$refs.form.resetFields()
          }).catch(() => {
            this.loading = false
          })
        }
      })
    },
    cancel() {
      this.loading = false
      this.$refs.form.resetFields()
      this.userAccounts = []
      this.visible = false
    },
    remoteMethod(query) {
      if (query.trim() !== '') {
        this.remoteLoading = true
        searchCustomer(query).then(res => {
          this.remoteLoading = false
          this.customers = res.data
        }).catch(() => {
          this.remoteLoading = false
        })
      } else {
        this.options = []
      }
    },
    handleSelect(customer_id) {
      if (Number(customer_id) === 0) {
        return
      }
      this.flowForm.account_id = ''
      flowBelongsAccount({ customer_id }).then(res => {
        this.accounts = res.data
      }).catch(() => {})
    },
    handleSelectAccount(account_id) {
      this.selectRate()
      this.userAccounts = []
      this.flowForm.accounts = {}
      getAccounts(account_id).then(res => {
        res.data.map(item => {
          if (item.account_id > 0) {
            this.userAccounts.push(item)
            this.$set(this.flowForm.accounts, item.account_type, 0)
          }
        })
      }).catch(() => {})
    },
    // 从后台返回带出返点比例
    selectRate() {
      let selected = false
      this.accounts.map(item => {
        if (this.flowForm.account_id === item.id) {
          const _key = 'CONSUME_' + this.flowForm.consume_type
          if (item.rates && item.rates.hasOwnProperty(_key)) {
            selected = true
            this.$set(this.flowForm, 'settlement_rate', item.rates[_key])
          }
        }
      })
      if (!selected) {
        this.$set(this.flowForm, 'settlement_rate', 0)
      }
    }
  }
}
</script>
