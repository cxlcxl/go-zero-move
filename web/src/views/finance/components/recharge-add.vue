<template>
  <dialog-panel title="添加流水" confirm-text="添加" :visible="visible" width="800px" :confirm-loading="loading"
                @cancel="cancel" @confirm="confirm">
    <el-form ref="form" :model="flowForm" :rules="rules" label-width="110px" size="mini">
      <el-form-item prop="customer_id" label="客户">
        <el-select v-model="flowForm.customer_id" remote filterable placeholder="请输入客户名称查询" clearable
                   style="width: 100%;" :remote-method="remoteMethod" :loading="remoteLoading" @change="handleSelect">
          <el-option v-for="item in customers" :label="item.customer_name" :value="item.id">
            <span style="float: left">{{ item.customer_name }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.full_name }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item prop="amounts" label="明细金额">
        <div class="account-container">
          <div class="account-card" v-for="(name, val) in rechargeTypes">
            <div class="account-name">{{ name }}</div>
            <input type="text" class="number" v-model="flowForm.amounts[Number(val)]" @change="handleCountAmount"/>
          </div>
        </div>
      </el-form-item>
      <el-form-item prop="amount" label="银行流水金额">
        <el-input-number v-model="flowForm.amount" :precision="2" class="w200" :step="1000"/>
        <el-tag size="mini" style="margin-left: 10px;">流水金额默认自动计算，出现误差时需手动调整</el-tag>
      </el-form-item>
      <el-form-item prop="trans_id" label="关联垫付" v-if="hasRepayment"
                    :rules="{required: true, message: '请选择关联的流水'}">
        <el-select v-model="flowForm.trans_id" placeholder="选择关联垫付流水" style="width: 100%;" multiple>
          <el-option v-for="item in transfers" :label="item|transferFilter" :value="item.id"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="account_id" label="账号">
        <el-select v-model="flowForm.account_id" placeholder="选择账号" style="width: 100%;">
          <el-option v-for="account in accounts" :label="account.account_name" :value="account.id"/>
        </el-select>
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
      <el-form-item prop="remark" label="备注">
        <el-input v-model="flowForm.remark" type="textarea"/>
      </el-form-item>
    </el-form>

    <el-button slot="operate" icon="el-icon-check" type="primary" :loading="loading" @click="continueCreate">添加并继续
    </el-button>
  </dialog-panel>
</template>

<script>
import DialogPanel from '@c/DialogPanel'
import {rechargeCreate} from '@a/recharge'
import {getAdvanceTransfers} from '@a/trans'
import {getDefaultCustomers, searchCustomer, flowBelongsAccount} from '@a/customer'
import {parseTime} from '@/utils'

export default {
  components: {DialogPanel},
  props: {
    rechargeTypes: Object
  },
  data() {
    const checkNumber = (rule, value, callback) => {
      if (value === 0) {
        callback(new Error('不可填写 0'))
      } else {
        callback()
      }
    }
    const checkAmount = (rule, value, callback) => {
      if (!value instanceof Object) {
        callback(new Error('请填写金额'))
      } else {
        let allEmpty = true
        for (let t in value) {
          let v = Number(value[t])
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
      hasRepayment: false,
      customers: [],
      accounts: [],
      transfers: [],
      flowForm: {
        customer_id: '',
        trans_id: [],
        account_id: 0,
        accounts: [],
        amount: 0,
        amounts: {},
        trans_date: parseTime(new Date(), '{y}-{m}-{d}'),
        remark: ''
      },
      rules: {
        amount: [
          {required: true, message: '请填写金额'},
          {validator: checkNumber}
        ],
        amounts: [{validator: checkAmount}, {required: true, message: '请填写金额'}],
        customer_id: {required: true, message: '请选择客户'},
        trans_date: {required: true, message: '请选择交易日期'},
      }
    }
  },
  filters: {
    transferFilter(trans) {
      return trans.trans_date + ' / ' + trans.amount + ' / 待还：' + (trans.amount - trans.paid_amount).toFixed(2)
    }
  },
  methods: {
    initCreate() {
      Promise.all([
        this.defaultCustomers()
      ]).then(res => {
        this.customers = res[0]
        for (let v in this.rechargeTypes) {
          this.$set(this.flowForm.amounts, Number(v), 0)
        }
        this.visible = true
      }).catch(() => {
      })
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
          for (let t in this.flowForm.amounts) {
            this.flowForm.amounts[t] = Number(this.flowForm.amounts[t])
          }
          rechargeCreate(this.flowForm).then(res => {
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
          for (let t in this.flowForm.amounts) {
            this.flowForm.amounts[t] = Number(this.flowForm.amounts[t])
          }
          rechargeCreate(this.flowForm).then(res => {
            this.$message.success('添加成功，请继续填写流水信息')
            this.$emit('success')
            this.loading = false
            this.$refs.form.resetFields()
            this.$set(this.flowForm, 'accounts', [])
          }).catch(() => {
            this.loading = false
          })
        }
      })
    },
    cancel() {
      this.loading = false
      this.$refs.form.resetFields()
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
      this.advanceTransfers(customer_id)
      this.$set(this.flowForm, 'account_id', 0)
      if (Number(customer_id) === 0) {
        return
      }
      flowBelongsAccount({customer_id}).then(res => {
        this.accounts = res.data
      }).catch(() => {
      })
    },
    advanceTransfers(customer_id) {
      this.transfers = []
      if (Number(customer_id) === 0 || !this.hasRepayment) {
        return
      }
      getAdvanceTransfers(customer_id).then(res => {
        this.transfers = res.data
      }).catch(() => {})
    },
    handleCountAmount() {
      let amount = 0
      let hasRepayment = false
      for (let v in this.flowForm.amounts) {
        let tmpAmount = Number(this.flowForm.amounts[v])
        amount += tmpAmount
        if (Number(v) === 2 && tmpAmount !== 0) {
          hasRepayment = true
        }
      }
      this.$set(this.flowForm, 'amount', Number(amount.toFixed(2)))
      // 关联拉取垫付流水
      this.hasRepayment = hasRepayment
      this.advanceTransfers(this.flowForm.customer_id)
    }
  }
}
</script>
