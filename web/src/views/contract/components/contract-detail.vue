<template>
  <el-row v-loading="loading">
    <el-col :span="24">
      <el-form ref="contract" :model="contractForm" label-width="100px" class="detail-form">
        <el-row :gutter="15">
          <el-col :span="13">
            <el-form-item label="合同编号">{{contractForm.contract_sn}}</el-form-item>
            <el-form-item label="合同名称">{{contractForm.contract_name}}</el-form-item>
            <el-form-item label="有效期">{{contractForm.start_date + ' 至 ' + contractForm.end_date}}</el-form-item>
            <el-form-item label="币种">{{contractForm.currency}}</el-form-item>
            <el-form-item label="客户">{{contractForm.customer_name}}</el-form-item>
            <el-form-item label="账号">
              <p v-if="Array.isArray(contractForm.accounts) && contractForm.accounts.length > 0">
                <el-tag size="mini" v-for="item in contractForm.accounts" style="margin-right: 5px;">{{item.account_name}}</el-tag>
              </p>
            </el-form-item>
          </el-col>
          <el-col :span="11">
            <el-form-item label="甲方">{{contractForm.party_a}}</el-form-item>
            <el-form-item label="乙方">{{contractForm.party_b}}</el-form-item>
            <el-form-item label="签订日期">{{contractForm.sign_date}}</el-form-item>
            <el-form-item label="合同描述">{{contractForm.desc}}</el-form-item>
          </el-col>

          <el-col :span="24" style="border-top: 1px solid #e2e2e2; padding-top: 15px;">
            <el-form-item prop="terms" label="合同条款">
              <div class="contract-detail-terms">
                <ul>
                  <li v-for="(item, terms_key) in contractForm.terms_values">
                    <div class="term-items">{{terms_key|termsKeyMap(termsList)}}</div>
                    <div class="term-values">
                      <table cellspacing="0" cellpadding="0">
                        <tr>
                          <th>阶梯值（{{contractForm.currency}}）</th>
                          <th width="150">比例（%）</th>
                          <th width="180">计算方式</th>
                        </tr>
                        <tr v-for="(terms, index) in item">
                          <td>
                            <span v-if="index === 0">{{terms.amount|money}}</span>
                            <span v-else>{{item[index-1].amount|money}} ~ {{terms.amount|money}}</span>
                          </td>
                          <td>{{terms.terms_val}}</td>
                          <td>{{calcTypes[terms.calculate_type]}}</td>
                        </tr>
                      </table>
                    </div>
                  </li>
                </ul>
              </div>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>

      <p class="text-center">
        <el-button icon="el-icon-close" type="primary" @click="closePage">关闭</el-button>
      </p>
    </el-col>
  </el-row>
</template>

<script>
  import {contractInfo} from '@a/contract'
  import {variable} from '@a/global'
  import {termsList} from "@a/contract-terms"
  import {formatMoney} from "@/utils"

  export default {
    name: 'ContractDetail',
    data() {
      return {
        loading: false,
        valid_dates: [],
        termsList: [],
        calcTypes: {},
        contractForm: {
          id: 0,
          file_id: 0,
          contract_sn: '',
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
      }
    },
    filters: {
      termsKeyMap(key, list) {
        let __name = ''
        if (Array.isArray(list)) {
          for (let i = 0; i < list.length; i++) {
            if (key === list[i]._key) {
              __name = list[i]._name
              if (list[i].state === 0) {
                __name += '（已停用）'
              }
              break
            }
          }
        }
        return __name
      },
      money(n) {
        return formatMoney(n)
      }
    },
    mounted() {
      this.contractDetail()
    },
    methods: {
      contractDetail() {
        let id = Number(this.$route.query.id)
        if (id === 0) {
          this.$message.error("参数错误")
          return
        }
        this.loading = true
        Promise.all([
          this.getTermsList(),
          this.getVariable()
        ]).then(res => {
          contractInfo(id).then(res => {
            this.contractForm = res.data
            if (this.contractForm.sign_date === '0001-01-01') {
              this.contractForm.sign_date = ''
            }

            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        }).catch(err => {
          this.loading = false
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
      closePage() {
        this.$store.dispatch('tagsView/delView', this.$route).then(() => {
          this.$router.replace({name: 'ContractList'})
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
  .contract-detail-terms {
    li {
      border: 1px solid #DCDFE6;
      display: flex;
      margin-bottom: 5px;
      -webkit-border-radius: 5px;
      -moz-border-radius: 5px;
      border-radius: 5px;
      padding-left: 10px;

      .term-items {
        width: 220px;
      }

      .term-values {
        flex: 1;
        padding-left: 20px;
        font-size: 12px;

        table {
          width: 100%;

          tr {
            th {
              background: #e0e4ee;
            }
            th, td {
              text-align: center;
              border-left: 1px solid #DCDFE6;
            }
            td {
              line-height: 25px !important;
            }
            &:not(:last-child) td, &>th {
              border-bottom: 1px solid #DCDFE6;
            }
          }
        }

        .el-icon-close {
          cursor: pointer;
        }
      }
    }
  }
</style>
