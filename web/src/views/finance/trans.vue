<template>
  <div class="payment-flow">
    <el-row>
      <el-col :span="24" class="search-container">
        <el-form ref="_search" :model="search" inline size="mini">
          <el-form-item>
            <el-input v-model="search.trans_sn" clearable placeholder="流水单号" class="w130"/>
          </el-form-item>
          <el-form-item>
            <el-date-picker v-model="search.start_date" type="date" placeholder="开始时间" value-format="yyyy-MM-dd" class="w130"/>&nbsp;
            <el-date-picker v-model="search.end_date" type="date" placeholder="结束时间" value-format="yyyy-MM-dd" class="w130"/>
          </el-form-item>
          <el-form-item>
            <el-input v-model="search.customer_name" class="w150" clearable placeholder="客户名称"/>
          </el-form-item>
          <el-form-item>
            <el-select v-model="search.user_account" clearable placeholder="账户类型">
              <el-option v-for="(key,val) in transList.user_accounts" :label="key" :value="val"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-cascader v-model="search.profit_quarter" :options="quarterOptions" class="w110" clearable placeholder="返点归属"/>
          </el-form-item>
          <el-form-item>
            <el-select v-model="search.transfer_mode" clearable placeholder="转账方式" class="w100">
              <el-option v-for="(key,val) in transList.transfer_modes" :label="key" :value="val"/>
            </el-select>
          </el-form-item>
          <el-form-item label="">
            <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="24" style="margin-bottom: 15px;">
        <el-button icon="el-icon-plus" size="small" type="primary" @click="add" v-permission="'trans/create'">添加流水</el-button>
        <el-button icon="el-icon-download" size="small" @click="download" :loading="loadings.downloadLoading" v-permission="'trans/download'">导出</el-button>
      </el-col>
      <el-col :span="24">
        <el-table v-loading="loadings.pageLoading" :data="transList.list" highlight-current-row stripe border
                  show-summary :summary-method="getSummaries" size="mini" class="first-summary-table" @sort-change="tableSort">
          <el-table-column prop="trans_sn" label="流水编号" width="110" align="center">
            <template slot-scope="scope">
              <span v-if="scope.row.is_bound_recharge === 1" class="text-primary">{{scope.row.trans_sn}}</span>
              <span v-else>{{scope.row.trans_sn}}</span>
            </template>
          </el-table-column>
          <!--<el-table-column prop="from_account_name" label="创梦转出账号" width="110" show-overflow-tooltip/>-->
          <el-table-column prop="to_customer_name" label="转入客户名称" width="120" show-overflow-tooltip/>
          <el-table-column prop="to_account_name" label="转入账号" width="160" show-overflow-tooltip/>
          <el-table-column label="类型" width="120" align="center" show-overflow-tooltip>
            <template slot-scope="scope">{{transList.user_accounts[scope.row.user_account]}}</template>
          </el-table-column>
          <el-table-column prop="amount" label="金额" width="120" align="right">
            <template slot-scope="scope">{{scope.row.amount|money}}</template>
          </el-table-column>
          <el-table-column prop="trans_date" label="交易日期" width="100" align="center" sortable/>
          <el-table-column prop="operator" label="操作人" width="110"/>
          <el-table-column prop="remark" label="备注" min-width="220"/>
          <el-table-column prop="profit_quarter" label="返点归属" width="70" align="center"/>
          <el-table-column label="标签" width="100" align="center">
            <template slot-scope="scope">
              <el-tag size="mini" v-show="scope.row.is_mark === 1">返</el-tag>&nbsp;
              <el-tag size="mini" v-show="scope.row.transfer_mode > 0" :type="scope.row|repaidFinish">
                {{transList.transfer_modes[scope.row.transfer_mode]}}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="created_at" label="添加时间" width="140"/>
          <el-table-column label="操作" width="140" align="center">
            <template slot-scope="scope">
              <el-button-group class="table-operate">
                <el-button type="primary" plain @click.native.prevent="relationDetail(scope.row)">关联</el-button>
                <el-button type="primary" plain @click.native.prevent="editRow(scope.row)" v-permission="'trans-platform/update'">修改</el-button>
                <el-button type="danger" plain @click.native.prevent="destroyRow(scope.row)" v-permission="'trans/destroy'"
                           :disabled="scope.row.state === 0">作废</el-button>
              </el-button-group>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
      <el-col :span="24" style="margin-top: 15px;">
        <page :page="search.page" :total="transList.total" @current-change="handlePage" @size-change="handlePageSize" :limit="search.limit"/>
      </el-col>
    </el-row>

    <trans-create ref="create" @success="getTransList" :transfer-modes="transList.transfer_modes"/>
    <relation-flow ref="relation" flow-type="recharge"/>
    <rebate-quarter-update ref="update" @success="getTransList"/>
  </div>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {transList, transListDownload, transDestroy} from '@a/trans'
  import Page from '@c/Page'
  import TransCreate from './components/trans-add'
  import RebateQuarterUpdate from './components/rebate-quarter-edit'
  import RelationFlow from './components/relation-flow'
  import {flowStateClass} from "./variables"
  import {formatMoney} from "@/utils"
  import {quarters} from "@/views/report/report"

  export default {
    name: "Transfer",
    components: {
      Page, TransCreate, DialogPanel, RelationFlow, RebateQuarterUpdate
    },
    data() {
      return {
        loadings: {
          pageLoading: false,
          downloadLoading: false
        },
        rechargeAccounts: {},
        transList: {
          amount: 0,
          total: 0,
          list: [],
          user_accounts: {},
          transfer_modes: {},
          flow_state: {}
        },
        search: {
          trans_sn: '', customer_name: '', user_account: '', start_date: '', end_date: '', profit_quarter: [], qs: '',
          transfer_mode: '', order_type: '', order_field: '', page: 1, limit: 10
        },
      }
    },
    mounted() {
      this.getTransList()
    },
    computed: {
      quarterOptions() {
        const monthQuarters = []
        for (let i = 2021; i <= (new Date()).getFullYear(); i++) {
          monthQuarters.push({value: i, label: i, children: quarters})
        }
        return monthQuarters
      }
    },
    filters: {
      stateClassFilter(state) {
        return flowStateClass[state]
      },
      money(n) {
        return formatMoney(n)
      },
      repaidFinish(row) {
        if (row.transfer_mode === 1 && row.amount === row.repaid_amount) {
          return 'success'
        } else {
          return 'primary'
        }
      }
    },
    methods: {
      getTransList() {
        this.loadings.pageLoading = true
        if (this.search.profit_quarter.length === 2) {
          this.search.qs = this.search.profit_quarter[0] + '-' + this.search.profit_quarter[1]
        } else {
          this.search.qs = ''
        }
        transList(this.search).then(response => {
          this.transList = response.data
          this.loadings.pageLoading = false
        }).catch(() => {
          this.loadings.pageLoading = false
        })
      },
      getSummaries(param) {
        const {columns} = param
        const sums = []
        columns.forEach((column, index) => {
          if (index === 0) {
            sums[index] = '合计'
          } else if (['amount'].includes(column.property)) {
            sums[index] = formatMoney(this.transList.amount)
          }
        })
        return sums
      },
      add() {
        this.$refs.create.initCreate()
      },
      editRow(row) {
        this.$refs.update.initUpdate(row)
      },
      download() {
        this.loadings.downloadLoading = true
        transListDownload(this.search).then(res => {
          this.loadings.downloadLoading = false

          import('@/vendor/Export2Excel').then(excel => {
            let tHeader = []
            res.data.headers.map(item => tHeader.push(item.title))
            const data = excel.formatJson(res.data)
            if (data.length === 0) {
              this.$message.info("没有筛选到需导出的数据")
              return
            }
            excel.export_json_to_excel({header: tHeader, data, filename: "客户转账记录"})
            this.loadings.downloadLoading = false
          }).catch(err => {
            console.log(err)
            this.loadings.downloadLoading = false
          })
        }).catch(() => {
          this.loadings.downloadLoading = false
        })
      },
      doSearch() {
        this.search.page = 1
        this.getTransList()
      },
      handlePage(p) {
        this.search.page = p
        this.getTransList()
      },
      handlePageSize(s) {
        this.search.limit = s
        this.getTransList()
      },
      destroyRow(row) {
        this.$confirm('确认作废此流水么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.loadings.pageLoading = true
          transDestroy(row.id).then(() => {
            this.$message.success('作废成功')
            this.getTransList()
          }).catch(err => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
        })
      },
      relationDetail(row) {
        this.$refs.relation.getFlows(row.id)
      },
      tableSort({_, prop, order}){
        if (order !== null) {
          this.$set(this.search, 'order_type', order)
          this.$set(this.search, 'order_field', prop)
        } else {
          this.$set(this.search, 'order_type', '')
          this.$set(this.search, 'order_field', '')
        }
        this.getTransList()
      }
    }
  }
</script>
