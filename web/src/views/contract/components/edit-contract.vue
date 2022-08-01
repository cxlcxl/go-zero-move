<template>
  <el-row v-loading="loading">
    <el-col :span="24">
      <el-form ref="contract" :model="contractForm" :rules="contractRules" label-width="100px" class="contract-form">
        <el-row :gutter="15">
          <el-col :span="13">
            <el-form-item prop="contract_sn" label="合同编号">
              <el-input v-model="contractForm.contract_sn"/>
            </el-form-item>
            <el-form-item prop="contract_name" label="合同名称">
              <el-input v-model="contractForm.contract_name"/>
            </el-form-item>
            <el-form-item prop="start_date" label="有效期">
              <el-date-picker v-model="valid_dates" type="daterange" @change="handleDates"
                              start-placeholder="Start" end-placeholder="End" value-format="yyyy-MM-dd"/>
            </el-form-item>
            <el-form-item prop="currency" label="币种">
              <el-select v-model="contractForm.currency" placeholder="请选择币种" style="width: 100%;">
                <el-option v-for="item in currencies" :label="item._name" :value="item._val"/>
              </el-select>
            </el-form-item>
            <el-form-item prop="customer_id" label="客户">
              <el-select v-model="contractForm.customer_id" remote filterable placeholder="可输入客户名称查询"
                         @change="handleCustomer" :remote-method="remoteMethod" :loading="remoteLoading" style="width: 100%;">
                <el-option v-for="item in customers" :label="item.customer_name" :value="item.id">
                  <span style="float: left">{{ item.customer_name }}</span>
                  <span style="float: right; color: #8492a6; font-size: 13px">{{ item.full_name }}</span>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item prop="account_ids" label="账号">
              <el-select v-model="contractForm.account_ids" multiple placeholder="请选择客户账号" style="width: 100%;">
                <el-option v-for="item in customer_accounts" :label="item.account_name" :value="item.id"/>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="11">
            <el-form-item prop="file_id" label="合同文件">
              <upload ref="upload" accept=".pdf" @success="handleSuccess" @change="handleChange"
                      trigger-text="选取文件（PDF 扫描文件，且不超过 5M）必需上传"/>
            </el-form-item>
            <el-form-item prop="party_a" label="甲方">
              <el-input v-model="contractForm.party_a"/>
            </el-form-item>
            <el-form-item prop="party_b" label="乙方">
              <el-input v-model="contractForm.party_b"/>
            </el-form-item>
            <el-form-item prop="sign_date" label="签订日期">
              <el-date-picker v-model="contractForm.sign_date" value-format="yyyy-MM-dd" type="date"/>
            </el-form-item>
            <el-form-item prop="desc" label="合同描述">
              <el-input v-model="contractForm.desc" type="textarea"/>
            </el-form-item>
          </el-col>

          <el-col :span="24" style="border-top: 1px solid #e2e2e2; padding-top: 15px;">
            <el-form-item prop="terms">
              <div slot="label">
                合同条款
                <el-tooltip effect="dark" content="查看阶梯比例说明" placement="right">
                  <i class="el-icon-question text-primary" @click="showTermsFAQ" style="cursor: pointer;"/>
                </el-tooltip>
              </div>
              <div class="contract-terms">
                <ul>
                  <li v-for="item in termsList" v-show="item.state === 1 || contractForm.terms.includes(item._key)">
                    <div class="term-items">
                      <el-checkbox v-model="contractForm.terms" :label="item._key"
                                   @change="handleTerms($event, item._key)">{{item._name}}<span v-show="item.state === 0">（已停用）</span>
                      </el-checkbox>
                      <el-button icon="el-icon-plus" plain type="primary" size="mini" v-show="contractForm.terms.includes(item._key)"
                                 style="width: 100%; margin-left: 10px;" @click="addTerms(item._key)"/>
                    </div>
                    <div class="term-values" v-show="contractForm.terms.includes(item._key)">
                      <table cellspacing="0" cellpadding="0">
                        <tr>
                          <th width="260">阶梯值<span v-show="contractForm.currency">（{{contractForm.currency}}）</span>
                          </th>
                          <th width="150">比例（%）</th>
                          <th>计算方式</th>
                          <th width="40"></th>
                        </tr>
                        <tr v-for="(terms, index) in contractForm.terms_values[item._key]">
                          <td>
                            <el-input-number v-model="terms.amount" :precision="2" :min="0" :step="10000" :max="9999999999"
                                             size="mini" class="value-txt" :disabled="index === 0"/>
                          </td>
                          <td>
                            <el-input-number v-model="terms.terms_val" :step="1" :min="0" :max="100" size="mini" class="value-txt"/>
                          </td>
                          <td>
                            <template v-if="index === 0">
                              <el-select v-model="terms.calculate_type" placeholder="请选择计算方式" size="mini" disabled>
                                <el-option v-for="(name, id) in calcTypes" :label="name" :value="Number(id)" v-show="id === 0"/>
                              </el-select>
                            </template>
                            <template v-else>
                              <el-select v-model="terms.calculate_type" placeholder="请选择计算方式" size="mini" @change="changeCalc($event, item._key)">
                                <el-option v-for="(name, id) in calcTypes" :label="name" :value="Number(id)" v-show="id > 0"/>
                              </el-select>
                            </template>
                          </td>
                          <td><i class="el-icon-close text-error" v-show="index > 0" @click="removeTerms(item._key, index)"/></td>
                        </tr>
                      </table>
                    </div>
                  </li>
                </ul>
              </div>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="el-icon-check" @click="save">保存</el-button>
              <el-button icon="el-icon-close" @click="cancel">取消</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>

      <terms-help ref="help"/>
    </el-col>
  </el-row>
</template>

<script>
  import Upload from '@c/Upload'
  import {contractUpdate, contractInfo} from '@a/contract'
  import TermsHelp from "./terms-help"
  import {searchCustomer, getDefaultCustomers, hasManyAccount} from "@a/customer"
  import {variable, getConfigs} from '@a/global'
  import {termsList} from "@a/contract-terms"

  export default {
    components: {
      Upload, TermsHelp
    },
    data() {
      const termsValid = (rule, value, callback) => {
        let terms = this.contractForm.terms
        for (let i = 0; i < terms.length; i++) {
          // 条款项
          let terms_values = this.contractForm.terms_values[terms[i]]
          if (!Array.isArray(terms_values)) {
            return callback(new Error('合同条款出错，请刷新页面重试'))
          }
          if (terms_values.length === 0) {
            return callback(new Error('合同条款出错，请刷新页面重试'))
          }
          terms_values.forEach((terms_value, idx) => {
            if (idx > 0) {
              if (terms_value.amount === 0) {
                return callback(new Error('请填写阶梯比例的阶梯值'))
              }
              if (terms_value.calculate_type === '' || terms_value.calculate_type === 0) {
                return callback(new Error('请填写阶梯比例的计算方式'))
              }
            }
          })
        }
        callback()
      }
      return {
        loading: false,
        calcTypes: {},
        isReUpload: false,
        remoteLoading: false,
        contractForm: {
          id: 0,
          file_id: 0,
          customer_id: 0,
          contract_sn: '',
          account_ids: [],
          contract_name: '',
          currency: '',
          start_date: '',
          end_date: '',
          sign_date: '',
          party_a: '',
          party_b: '',
          desc: '',
          terms: [],
          terms_values: {}
        },
        valid_dates: [],
        currencies: [],
        termsList: [],
        customers: [],
        customer_accounts: [],
        contractRules: {
          contract_sn: {required: true, message: '请填写合同编号'},
          contract_name: {required: true, message: '请填写合同名称'},
          start_date: {required: true, message: '请选择合同有效期'},
          currency: {required: true, message: '请选择币种'},
          customer_id: {required: true, message: '请填写所属客户'},
          account_ids: {required: true, message: '请选择客户账号'},
          terms: {validator: termsValid}
        }
      }
    },
    mounted() {
      this.initUpdate()
    },
    methods: {
      defaultCustomers() {
        return new Promise((resolve, reject) => {
          getDefaultCustomers().then(res => {
            this.customers = res.data
            resolve()
          }).catch(() => {
            reject()
          })
        })
      },
      getTermsList() {
        return new Promise((resolve, reject) => {
          if (this.termsList.length > 0) {
            resolve()
          } else {
            termsList().then(res => {
              this.termsList = res.data
              resolve()
            }).catch(() => {
              reject()
            })
          }
        })
      },
      getCurrencies() {
        return new Promise((resolve, reject) => {
          if (this.currencies.length > 0) {
            resolve()
          } else {
            getConfigs('currency').then(res => {
              this.currencies = res.data
              resolve()
            }).catch(() => {
              reject()
            })
          }
        })
      },
      getVariable() {
        return new Promise((resolve, reject) => {
          variable({key: 'calculate_type'}).then(res => {
            this.calcTypes = res.data
            resolve()
          }).catch(() => {
            reject()
          })
        })
      },
      initUpdate() {
        let id = Number(this.$route.query.id)
        if (id === 0) {
          this.$message.error("参数错误")
          return
        }
        this.loading = true
        Promise.all([
          this.defaultCustomers(),
          this.getCurrencies(),
          this.getTermsList(),
          this.getVariable()
        ]).then(res => {
          contractInfo(id).then(res => {
            this.contractForm = res.data
            if (this.contractForm.sign_date === '0001-01-01') {
              this.contractForm.sign_date = ''
            }
            if (!this.inDefaultCustomers(this.contractForm.customer_id)) {
              this.customers.push({
                id: this.contractForm.customer_id,
                customer_name: this.contractForm.customer_name,
                full_name: this.contractForm.full_name
              })
            }
            this.customer_accounts = this.contractForm.accounts
            this.valid_dates = [this.contractForm.start_date, this.contractForm.end_date]
            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        }).catch(err => {
          this.loading = false
        })
      },
      cancel() {
        this.$store.dispatch('tagsView/delView', this.$route).then(() => {
          this.$router.replace({name: 'ContractList'})
        })
      },
      save() {
        this.$refs.contract.validate(v => {
          if (v) {
            if (!this.isReUpload) {
              this.doUpdate()
            } else {
              // 通过附件上传成功调度表单数据提交
              this.$refs.upload.submit()
            }
          }
        })
      },
      handleDates(d) {
        if (d === null) {
          this.$set(this.contractForm, 'start_date', '')
          this.$set(this.contractForm, 'end_date', '')
        } else {
          this.$set(this.contractForm, 'start_date', d[0])
          this.$set(this.contractForm, 'end_date', d[1])
        }
      },
      handleSuccess(data) {
        this.$set(this.contractForm, 'file_id', data.id)
        this.doUpdate()
      },
      doUpdate() {
        contractUpdate(this.contractForm).then(res => {
          this.$message.success('修改成功，请刷新页面或点击查询查看')
          this.cancel()
        }).catch(() => {})
      },
      handleChange(file, list) {
        if (Array.isArray(list)) {
          this.isReUpload = list.length > 0
        } else {
          this.isReUpload = false
        }
      },
      addTerms(terms_key) {
        let calc_type = ''
        let tmp = this.contractForm.terms_values[terms_key]
        if (Array.isArray(tmp) && tmp.length > 1) {
          for (let i = 1; i < tmp.length; i++) {
            if (Number(tmp[i].calculate_type) > 0) {
              calc_type = Number(tmp[i].calculate_type)
              break
            }
          }
        }
        this.contractForm.terms_values[terms_key].push({terms_val: 0, amount: 0, calculate_type: calc_type})
      },
      changeCalc(v, terms_key) {
        for (let i = 1; i < this.contractForm.terms_values[terms_key].length; i++) {
          this.$set(this.contractForm.terms_values[terms_key][i], 'calculate_type', v)
        }
      },
      removeTerms(terms_key, index) {
        this.contractForm.terms_values[terms_key].splice(index, 1)
      },
      // 切换时填充默认属性，否则 v-model 不能识别
      handleTerms(v, terms_key) {
        if (v) {
          this.$set(this.contractForm.terms_values, terms_key, [])
          this.contractForm.terms_values[terms_key].push({terms_val: 0, amount: 0, calculate_type: 0})
        } else {
          delete this.contractForm.terms_values[terms_key]
        }
      },
      showTermsFAQ() {
        this.$refs.help.visible = true
      },
      handleCustomer(customer_id) {
        this.contractForm.account_ids = []
        hasManyAccount(customer_id).then(res => {
          this.customer_accounts = res.data
          if (Array.isArray(this.customer_accounts)) {
            this.customer_accounts.map(act => this.contractForm.account_ids.push(act.id))
          }
        }).catch(() => {})
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
      },
      inDefaultCustomers(customer_id) {
        let _in = false
        this.customers.map(item => {
          if (item.id === customer_id) {
            _in = true
          }
        })

        return _in
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import "terms";
</style>