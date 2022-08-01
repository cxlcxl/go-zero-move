<template>
  <dialog-panel title="添加流水" confirm-text="添加" :visible="visible" @cancel="cancel" width="900px"
                @confirm="confirm" :confirm-loading="loading">
    <el-form ref="form" :model="flowForm" :rules="rules" label-width="120px" size="mini">
      <el-form-item prop="from_account_id" label="平台账号">
        <el-select v-model="flowForm.from_account_id" placeholder="请先选择广告平台客户数据" style="width: 100%;" @change="handleFromSelect">
          <el-option v-for="item in from_customers" :label="item.account_name" :value="item.account_id"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="to_account_id" label="创梦账号">
        <el-select v-model="flowForm.to_account_id" placeholder="请先选择创梦客户数据" style="width: 100%;"
                   @change="handleSelectAccount">
          <el-option v-for="item in to_customers" :label="item.account_name" :value="item.account_id"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="accounts" label="金额">
        <div class="account-container" v-if="accounts.length > 0">
          <div class="account-card" v-for="item in accounts">
            <div class="account-name">{{item.account_name}}</div>
            <input type="text" class="number" v-model="flowForm.accounts[item.account_type]" @change="calcProfitQuarter"/>
          </div>
        </div>
        <span v-else>请选择账号</span>
      </el-form-item>
      <el-form-item prop="trans_date" label="交易日期">
        <el-date-picker v-model="flowForm.trans_date" type="date" placeholder="开始时间" @change="calcProfitQuarter"
                        value-format="yyyy-MM-dd" style="width: 100%;"/>
      </el-form-item>
      <el-form-item prop="profit_quarter" label="返利金季度归属">
        <el-select v-model="profit.year" placeholder="归属年份" style="width: 100px; margin-right: 10px;">
          <el-option v-for="year in years" :label="year" :value="year"/>
        </el-select>
        <el-select v-model="profit.quarter" placeholder="归属季度">
          <el-option label="Q1" :value="1"/>
          <el-option label="Q2" :value="2"/>
          <el-option label="Q3" :value="3"/>
          <el-option label="Q4" :value="4"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="remark" label="备注">
        <el-input type="textarea" v-model="flowForm.remark"/>
      </el-form-item>
    </el-form>

    <el-button slot="operate" icon="el-icon-check" type="primary" @click="continueCreate" :loading="loading">添加并继续
    </el-button>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {transPlatformCreate} from "@a/trans"
  import {hasManyAccount, getPlatformCustomers, getCmCustomers} from "@a/customer"
  import {getAccounts} from "@a/account"
  import {parseTime} from "@/utils"

  export default {
    components: {DialogPanel},
    data() {
      const checkAmount = (rule, value, callback) => {
        if (!value instanceof Object) {
          callback(new Error('请填写金额'))
        } else {
          let allEmpty = true
          for (let account in value) {
            let v = Number(value[account])
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
        from_customers: [],
        from_accounts: [],
        to_customers: [],
        to_accounts: [],
        accounts: [],
        profit: {year: '', quarter: ''},
        flowForm: {
          from_customer_id: '',
          from_account_id: '',
          to_customer_id: '',
          to_account_id: '',
          profit_quarter: '',
          accounts: {},
          trans_date: parseTime(new Date(), "{y}-{m}-{d}"),
          remark: '',
        },
        rules: {
          from_customer_id: {required: true, message: '请选择客户'},
          to_customer_id: {required: true, message: '请选择客户'},
          from_account_id: {required: true, message: '请选择账号'},
          to_account_id: {required: true, message: '请选择账号'},
          trans_date: {required: true, message: '请选择交易日期'},
          accounts: {validator: checkAmount},
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
    methods: {
      initCreate() {
        Promise.all([
          this.platformCustomer(),
          this.cmCustomer()
        ]).then(res => {
          this.from_customers = res[0]
          this.to_customers = res[1]
          if (Array.isArray(this.from_customers)) {
            this.$set(this.flowForm, 'from_account_id', this.from_customers[0].account_id)
            this.handleFromSelect(this.from_customers[0].account_id)
          }
          if (Array.isArray(this.to_customers)) {
            this.$set(this.flowForm, 'to_account_id', this.to_customers[0].account_id)
            this.handleSelectAccount(this.to_customers[0].account_id)
          }

          this.visible = true
        }).catch(() => {
        })
      },
      platformCustomer() {
        return new Promise((resolve, reject) => {
          getPlatformCustomers().then(res => {
            resolve(res.data)
          }).catch(() => {
            reject()
          })
        })
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
      confirm() {
        this.$refs.form.validate(v => {
          if (v) {
            this.loading = true
            for (let account in this.flowForm.accounts) {
              this.flowForm.accounts[account] = Number(this.flowForm.accounts[account])
            }
            if (this.profit.year && this.profit.quarter) {
              this.flowForm.profit_quarter = this.profit.year + '-' + this.profit.quarter
            }
            transPlatformCreate(this.flowForm).then(res => {
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
            for (let account in this.flowForm.accounts) {
              this.flowForm.accounts[account] = Number(this.flowForm.accounts[account])
            }
            transPlatformCreate(this.flowForm).then(res => {
              this.$message.success('添加成功，请继续填写流水信息')
              this.$emit('success')
              this.loading = false
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
        this.visible = false
      },
      handleSelect(customer_id, t) {
        hasManyAccount(customer_id).then(res => {
          if (t === 'from') {
            this.from_accounts = res.data
          } else {
            this.to_accounts = res.data
          }
        }).catch(() => {
        })
      },
      handleSelectAccount(account_id) {
        this.accounts = []
        this.flowForm.accounts = {}
        this.to_customers.map(item => {
          if (item.account_id === account_id) {
            this.$set(this.flowForm, "to_customer_id", item.customer_id)
          }
        })
        getAccounts(account_id).then(res => {
          res.data.map(item => {
            if (item.account_id > 0) {
              this.accounts.push(item)
              this.$set(this.flowForm.accounts, item.account_type, 0)
            }
          })
        }).catch(() => {})
      },
      handleFromSelect(account_id) {
        this.from_customers.map(item => {
          if (item.account_id === account_id) {
            this.$set(this.flowForm, "from_customer_id", item.customer_id)
          }
        })
      },
      calcProfitQuarter() {
        if (this.flowForm.accounts.hasOwnProperty('URGE')) {
          if (this.flowForm.accounts.URGE !== 0 && this.flowForm.trans_date !== '' && !this.profit.quarter) {
            let m = {1: 1, 2: 1, 3: 1, 4: 2, 5: 2, 6: 2, 7: 3, 8: 3, 9: 3, 10: 4, 11: 4, 12: 4}
            let d = this.flowForm.trans_date.split("-")
            let quarter = m[Number(d[1])]
            if (quarter === 1) {
              this.$set(this.profit, 'quarter', 4)
              this.$set(this.profit, 'year', Number(d[0]) - 1)
            } else {
              this.$set(this.profit, 'year', Number(d[0]))
              this.$set(this.profit, 'quarter', quarter - 1)
            }
          }
        }
      }
    }
  }
</script>
