<template>
  <dialog-panel title="添加流水" confirm-text="添加" :visible="visible" @cancel="cancel" width="800px"
                @confirm="confirm" :confirm-loading="loading">
    <el-form ref="form" :model="flowForm" :rules="rules" label-width="100px" size="mini">
      <el-form-item prop="from_account_id" label="创梦账号">
        <el-select v-model="flowForm.from_account_id" placeholder="请先选择创梦客户数据" style="width: 100%;" @change="handleFromSelect">
          <el-option v-for="item in from_customers" :label="item.account_name" :value="item.account_id"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="to_customer_id" label="客户名称">
        <el-select v-model="flowForm.to_customer_id" style="width: 100%;" remote filterable :loading="remoteLoading" clearable
                   placeholder="请输入客户名称查询" @change="handleSelect($event, 'to')" :remote-method="remoteMethod">
          <el-option v-for="item in to_customers" :label="item.customer_name" :value="item.id">
            <span style="float: left">{{ item.customer_name }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.full_name }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item prop="to_account_id" label="客户账号">
        <el-select v-model="flowForm.to_account_id" placeholder="请选择客户账号" style="width: 100%;">
          <el-option v-for="item in to_accounts" :label="item.account_name" :value="item.id"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="user_account" label="账户类型">
        <el-select v-model="flowForm.user_account" placeholder="请选择账户类型" @change="handleAccount" style="width: 100%;">
          <el-option v-for="item in accounts" :label="item._name" :value="item._val"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="amount" label="金额">
        <el-input-number v-model="flowForm.amount" class="w180" :step="1000"/>
        <div style="margin-left: 10px; display: inline-block;" v-if="flowForm.user_account === 'URGE'">
          <el-select v-model="profit.year" placeholder="归属年份" class="w100">
            <el-option v-for="year in years" :label="year" :value="year"/>
          </el-select>&nbsp;
          <el-select v-model="profit.quarter" placeholder="归属季度" class="w100">
            <el-option label="Q1" :value="1"/>
            <el-option label="Q2" :value="2"/>
            <el-option label="Q3" :value="3"/>
            <el-option label="Q4" :value="4"/>
          </el-select>
          <span class="form-item-tip">返利请填写归属季度</span>
        </div>
      </el-form-item>
      <el-form-item prop="transfer_mode" label="转账类型">
        <el-radio-group v-model="flowForm.transfer_mode" @change="handleAccount">
          <el-radio v-for="(name, val) in TransferModes" :label="Number(val)">{{name}}</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item prop="recharge_id" label="关联充值"
                    v-if="flowForm.transfer_mode === 0">
        <el-select v-model="flowForm.recharge_id" placeholder="请选择关联的客户充值数据" multiple style="width: 100%;">
          <el-option v-for="item in recharges" :label="item|rechargeFilter" :value="item.id"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="is_mark" label="计算消耗返">
        <el-switch v-model="flowForm.is_mark" :active-value="1" :inactive-value="0"/>
      </el-form-item>
      <el-form-item prop="trans_date" label="交易日期">
        <el-date-picker v-model="flowForm.trans_date" type="date" placeholder="开始时间"
                        value-format="yyyy-MM-dd" style="width: 100%;"/>
      </el-form-item>
      <el-form-item prop="remark" label="备注">
        <el-input type="textarea" v-model="flowForm.remark"/>
      </el-form-item>
    </el-form>

    <el-button slot="operate" icon="el-icon-check" type="primary" @click="continueCreate" :loading="loading">添加并继续</el-button>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {transCreate} from "@a/trans"
  import {getConfigs} from "@a/global"
  import {rechargeForTrans} from "@a/recharge"
  import {getDefaultCustomers, searchCustomer, hasManyAccount, getCmCustomers} from "@a/customer"
  import {parseTime} from '@/utils'

  export default {
    components: {DialogPanel},
    props: {
      TransferModes: {
        type: Object,
        required: true
      }
    },
    data() {
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        from_customers: [],
        from_accounts: [],
        to_customers: [],
        to_accounts: [],
        accounts: [],
        recharges: [],
        profit: {year: '', quarter: ''},
        flowForm: {
          from_customer_id: '',
          from_account_id: '',
          to_customer_id: '',
          to_account_id: '',
          user_account: '',
          profit_quarter: '',
          is_mark: 0,
          transfer_mode: 0,
          amount: 0,
          recharge_id: [],
          trans_date: parseTime(new Date(), "{y}-{m}-{d}"),
          remark: '',
        },
        rules: {
          from_customer_id: {required: true, message: '请选择客户'},
          to_customer_id: {required: true, message: '请选择客户'},
          from_account_id: {required: true, message: '请选择账号'},
          to_account_id: {required: true, message: '请选择账号'},
          trans_date: {required: true, message: '请选择交易日期'},
          amount: {required: true, message: '请填写金额'},
          user_account: {required: true, message: '请选择账户类型'},
          transfer_mode: {required: true, message: '请选择转账类型'},
        }
      }
    },
    computed: {
      years() {
        let y = [], d = new Date()
        for (let i = 2021;i <= d.getFullYear();i++) {
          y.push(i)
        }
        return y
      }
    },
    filters: {
      rechargeFilter(recharge) {
        return recharge.trans_date + ' / ' + recharge.amount + ' / 未拨：' + (recharge.amount - recharge.paid_amount).toFixed(2)
      }
    },
    methods: {
      initCreate() {
        console.log("现金与授信类型 必需关联充值")
        console.log("计算消耗返的 必需关联充值")
        Promise.all([
          this.cmCustomer(),
          this.defaultCustomer(),
          this.userAccounts()
        ]).then(res => {
          this.recharges = []
          this.from_customers = res[0]
          if (Array.isArray(this.from_customers)) {
            this.$set(this.flowForm, 'from_account_id', this.from_customers[0].account_id)
            this.handleFromSelect(this.from_customers[0].account_id)
          }

          this.to_customers = res[1]
          this.accounts = res[2]
          this.visible = true
        }).catch(() => {})
      },
      cmCustomer() {
        return new Promise((resolve, reject) => {
          getCmCustomers().then(res => {
            resolve(res.data)
          }).catch(() => {
            reject()
          })
        })
      },
      defaultCustomer() {
        return new Promise((resolve, reject) => {
          getDefaultCustomers().then(res => {
            resolve(res.data)
          }).catch(() => {
            reject()
          })
        })
      },
      userAccounts() {
        return new Promise((resolve, reject) => {
          getConfigs('user_account').then(res => {
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
            if (this.flowForm.transfer_mode > 0 || this.flowForm.recharge_id.length === 0) {
              this.flowForm.recharge_id = []
            }
            if (this.profit.year && this.profit.quarter) {
              this.flowForm.profit_quarter = this.profit.year + '-' + this.profit.quarter
            }
            transCreate(this.flowForm).then(res => {
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
            if (this.flowForm.transfer_mode > 0 || this.flowForm.recharge_id.length === 0) {
              this.flowForm.recharge_id = []
            }
            if (this.profit.year && this.profit.quarter) {
              this.flowForm.profit_quarter = this.profit.year + '-' + this.profit.quarter
            }
            transCreate(this.flowForm).then(res => {
              this.$message.success('添加成功，请继续填写流水信息')
              this.$emit('success')
              this.loading = false
              this.profit = {year: '', quarter: ''}
              this.$refs.form.resetFields()
              this.$set(this.flowForm, "accounts", [])
            }).catch(() => {
              this.loading = false
            })
          }
        })
      },
      cancel() {
        this.loading = false
        this.$refs.form.resetFields()
        this.profit = {year: '', quarter: ''}
        this.visible = false
      },
      remoteMethod(query) {
        if (query.trim() !== '') {
          this.remoteLoading = true;
          searchCustomer(query).then(res => {
            this.remoteLoading = false
            this.to_customers = res.data
          }).catch(() => {
            this.remoteLoading = false
          })
        } else {
          this.options = []
        }
      },
      handleSelect(customer_id, t) {
        this.$set(this.flowForm, 'recharge_id', '')
        if (Number(customer_id) === 0) {
          return
        }
        // 获取用户绑定的客户充值
        rechargeForTrans(customer_id).then(res => {
          this.recharges = res.data
        }).catch(() => {})
        this.$set(this.flowForm, 'to_account_id', '')
        hasManyAccount(customer_id).then(res => {
          if (t === 'from') {
            this.from_accounts = res.data
          } else {
            this.to_accounts = res.data
          }
        }).catch(() => {
        })
      },
      handleFromSelect(account_id) {
        this.from_customers.map(item => {
          if (item.account_id === account_id) {
            this.$set(this.flowForm, "from_customer_id", item.customer_id)
          }
        })
      },
      handleAccount() {
        if ([2, 3, 4].includes(this.flowForm.transfer_mode)) {
          this.$set(this.flowForm, 'is_mark', 0)
        } else if (['CASH', 'CREDIT'].includes(this.flowForm.user_account)) {
          this.$set(this.flowForm, 'is_mark', 1)
        }
      }
    }
  }
</script>
