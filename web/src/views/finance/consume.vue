<template>
  <div class="payment-flow">
    <el-row>
      <el-col :span="24" class="search-container">
        <el-form ref="_search" :model="search" inline size="mini">
          <el-form-item>
            <el-input v-model="search.consume_sn" clearable placeholder="流水单号"/>
          </el-form-item>
          <el-form-item>
            <el-date-picker v-model="search.start_date" type="date" placeholder="开始时间" value-format="yyyy-MM-dd" class="w130"/>&nbsp;
            <el-date-picker v-model="search.end_date" type="date" placeholder="结束时间" value-format="yyyy-MM-dd" class="w130"/>
          </el-form-item>
          <el-form-item>
            <el-input v-model="search.customer_name" class="w200" clearable placeholder="客户名称"/>
          </el-form-item>
          <el-form-item>
            <el-select v-model="search.user_account" clearable placeholder="请选择账户类型">
              <el-option v-for="(key,val) in consumeList.user_accounts" :label="key" :value="val"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="">
            <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="24" style="margin-bottom: 15px;">
        <el-button icon="el-icon-plus" size="small" type="primary" @click="add" v-permission="'consume/create'">添加流水</el-button>
        <el-button icon="el-icon-download" size="small" @click="download" :loading="loadings.downloadLoading" v-permission="'consume/download'">导出</el-button>
      </el-col>
      <el-col :span="24">
        <el-table v-loading="loadings.pageLoading" :data="consumeList.list" highlight-current-row stripe border
                  show-summary :summary-method="getSummaries" size="mini" class="first-summary-table" @sort-change="tableSort">
          <el-table-column prop="consume_sn" label="流水编号" width="120" align="center"/>
          <el-table-column label="客户名称" width="120">
            <template slot-scope="scope">{{scope.row.customer.customer_name}}</template>
          </el-table-column>
          <el-table-column prop="customer_name" label="账号" width="160">
            <template slot-scope="scope">{{scope.row.account.account_name}}</template>
          </el-table-column>
          <el-table-column label="账户类型" width="120" align="center">
            <template slot-scope="scope">{{consumeList.user_accounts[scope.row.user_account]}}</template>
          </el-table-column>
          <el-table-column prop="amount" label="金额" width="120" align="right">
            <template slot-scope="scope">{{scope.row.amount|money}}</template>
          </el-table-column>
          <el-table-column prop="settlement_rate" label="返点比例(%)" width="100" align="right"/>
          <el-table-column prop="trans_date" label="交易日期" width="110" align="center" sortable/>
          <el-table-column prop="operator" label="操作人" width="110"/>
          <el-table-column prop="remark" label="备注" min-width="220"/>
          <el-table-column label="消耗类型" width="80" align="center">
            <template slot-scope="scope">{{ consume_types[scope.row.consume_type]}}</template>
          </el-table-column>
          <el-table-column label="状态" width="80" align="center">
            <template slot-scope="scope">
              <span :class="scope.row.state|stateClassFilter">{{consumeList.flow_state[scope.row.state]}}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="70" align="center" v-permission="'consume/destroy'">
            <template slot-scope="scope">
              <el-button-group class="table-operate">
                <el-button type="danger" plain @click.native.prevent="destroyRow(scope.row)"
                           :disabled="scope.row.state === 0">作废</el-button>
              </el-button-group>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
      <el-col :span="24" style="margin-top: 15px;">
        <page :page="search.page" :total="consumeList.total" @current-change="handlePage" @size-change="handlePageSize"/>
      </el-col>
    </el-row>

    <consume-create ref="create" @success="getConsumeList" :consume-types="consume_types"/>
  </div>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {consumeList, consumeListDownload, consumeDestroy} from '@/api/marketing'
  import {getConfigMap} from '@a/global'
  import Page from '@c/Page'
  import ConsumeCreate from './components/consume-add'
  import {flowStateClass} from "./variables"
  import {formatMoney} from "@/utils"

  export default {
    name: "Consume",
    components: {
      Page, ConsumeCreate, DialogPanel
    },
    data() {
      return {
        loadings: {
          pageLoading: false,
          downloadLoading: false
        },
        consumeAccounts: {},
        consumeList: {
          total: 0,
          amount: 0,
          list: [],
          user_accounts: {},
          flow_state: {},
        },
        search: {consume_sn: '', customer_name: '', user_account: '', start_date: '', end_date: '', order_type: '', order_field: '', page: 1, limit: 10},
        consume_types: {}
      }
    },
    mounted() {
      this.getConsumeList()
    },
    filters: {
      stateClassFilter(state) {
        return flowStateClass[state]
      },
      money(n) {
        return formatMoney(n)
      }
    },
    methods: {
      getConsumeList() {
        this.loadings.pageLoading = true
        Promise.all([
          this.getConsumeTypes()
        ]).then(res => {
          consumeList(this.search).then(response => {
            this.consumeList = response.data
            this.loadings.pageLoading = false
          }).catch(() => {
            this.loadings.pageLoading = false
          })
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
            sums[index] = formatMoney(this.consumeList.amount)
          }
        })
        return sums
      },
      add() {
        this.$refs.create.initCreate()
      },
      download() {
        this.loadings.downloadLoading = true
        consumeListDownload(this.search).then(res => {
          this.loadings.downloadLoading = false

          import('@/vendor/Export2Excel').then(excel => {
            let tHeader = []
            res.data.headers.map(item => tHeader.push(item.title))
            const data = excel.formatJson(res.data)
            if (data.length === 0) {
              this.$message.info("没有筛选到需导出的数据")
              return
            }
            excel.export_json_to_excel({header: tHeader, data, filename: "消耗记录"})
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
        this.getConsumeList()
      },
      handlePage(p) {
        this.search.page = p
        this.getConsumeList()
      },
      handlePageSize(s) {
        this.search.limit = s
        this.getConsumeList()
      },
      destroyRow(row) {
        this.$confirm('确认作废此流水么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.loadings.pageLoading = true
          consumeDestroy(row.id).then(() => {
            this.$message.success('作废成功')
            this.getConsumeList()
          }).catch(err => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
        })
      },
      getConsumeTypes() {
        return new Promise((resolve, reject) => {
          if (Object.keys(this.consume_types).length > 0) {
            return resolve()
          }
          getConfigMap('consume_type').then(res => {
            this.consume_types = res.data
            resolve()
          }).catch(() => {
            reject()
          })
        })
      },
      tableSort({_, prop, order}){
        if (order !== null) {
          this.$set(this.search, 'order_type', order)
          this.$set(this.search, 'order_field', prop)
        } else {
          this.$set(this.search, 'order_type', '')
          this.$set(this.search, 'order_field', '')
        }
        this.getConsumeList()
      }
    }
  }
</script>
