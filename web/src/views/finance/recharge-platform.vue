<template>
  <div class="payment-flow">
    <el-row>
      <el-col :span="24" class="search-container">
        <el-form ref="_search" :model="search" inline size="mini">
          <el-form-item>
            <el-input v-model="search.recharge_sn" clearable placeholder="流水单号"/>
          </el-form-item>
          <el-form-item>
            <el-date-picker v-model="search.start_date" type="date" placeholder="开始时间" value-format="yyyy-MM-dd" class="w130"/>&nbsp;
            <el-date-picker v-model="search.end_date" type="date" placeholder="结束时间" value-format="yyyy-MM-dd" class="w130"/>
          </el-form-item>
          <el-form-item>
            <el-radio-group v-model="search.recharge_type">
              <el-radio-button label="">-</el-radio-button>
              <el-radio-button v-for="(key, val) in rechargeList.recharge_types" :label="val">{{key}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="">
            <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="24" style="margin-bottom: 15px;">
        <el-button icon="el-icon-plus" size="small" type="primary" @click="add" v-permission="'recharge-platform/create'">添加流水</el-button>
        <el-button icon="el-icon-download" size="small" @click="download" :loading="loadings.downloadLoading" v-permission="'recharge-platform/download'">导出</el-button>
      </el-col>
      <el-col :span="24">
        <el-table v-loading="loadings.pageLoading" :data="rechargeList.list" highlight-current-row stripe border
                  show-summary :summary-method="getSummaries" size="mini" class="first-summary-table" @sort-change="tableSort">
          <el-table-column prop="recharge_sn" label="流水编号" width="110" align="center"/>
          <el-table-column label="类型" width="100" align="center">
            <template slot-scope="scope">{{rechargeList.recharge_types[scope.row.recharge_type]}}</template>
          </el-table-column>
          <el-table-column prop="amount" label="金额" width="120" align="right">
            <template slot-scope="scope">{{scope.row.amount|money}}</template>
          </el-table-column>
          <el-table-column prop="settlement_rate" label="返点比例(%)" width="90" align="right"/>
          <el-table-column prop="trans_date" label="交易日期" width="110" align="center" sortable/>
          <el-table-column prop="operator" label="操作人" width="110"/>
          <el-table-column prop="remark" label="备注" min-width="220"/>
          <el-table-column label="状态" width="80" align="center">
            <template slot-scope="scope">
              <span :class="scope.row.state|stateClassFilter">{{rechargeList.flow_state[scope.row.state]}}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="100" align="center">
            <template slot-scope="scope">
              <el-button-group class="table-operate">
                <el-button type="danger" plain @click.native.prevent="destroyRow(scope.row)" v-permission="'recharge-platform/destroy'"
                           :disabled="scope.row.state === 0">作废</el-button>
                <el-button type="primary" plain @click.native.prevent="viewFile(scope.row.file_id)"
                           :disabled="scope.row.file_id === 0">附件</el-button>
              </el-button-group>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
      <el-col :span="24" style="margin-top: 15px;">
        <page :page="search.page" :total="rechargeList.total" @current-change="handlePage" @size-change="handlePageSize"/>
      </el-col>
    </el-row>

    <recharge-platform-create ref="create" @success="getRechargeList" :recharge-types="rechargeList.recharge_types"/>
  </div>
</template>

<script>
  import {rechargePlatformList, rechargePlatformListDownload, rechargePlatformDestroy} from '@a/recharge'
  import {fileView} from '@a/global'
  import Page from '@c/Page'
  import RechargePlatformCreate from './components/recharge-platform-add'
  import {flowStateClass} from "./variables"
  import {formatMoney} from "@/utils"

  export default {
    name: "RechargePlatform",
    components: {
      Page, RechargePlatformCreate
    },
    data() {
      return {
        visible: false,
        loadings: {
          pageLoading: false,
          downloadLoading: false
        },
        rechargeAccounts: {},
        rechargeList: {
          total: 0,
          amount: 0,
          list: [],
          recharge_types: {}
        },
        search: {recharge_sn: '', recharge_type: '', start_date: '', end_date: '', order_type: '', order_field: '', page: 1, limit: 10},
      }
    },
    mounted() {
      this.getRechargeList()
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
      getRechargeList() {
        this.loadings.pageLoading = true
        rechargePlatformList(this.search).then(response => {
          this.rechargeList = response.data
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
            sums[index] = formatMoney(this.rechargeList.amount)
          }
        })
        return sums
      },
      add() {
        this.$refs.create.initCreate()
      },
      doSearch() {
        this.search.page = 1
        this.getRechargeList()
      },
      handlePage(p) {
        this.search.page = p
        this.getRechargeList()
      },
      handlePageSize(s) {
        this.search.limit = s
        this.getRechargeList()
      },
      download() {
        this.loadings.downloadLoading = true
        rechargePlatformListDownload(this.search).then(res => {
          this.loadings.downloadLoading = false

          import('@/vendor/Export2Excel').then(excel => {
            let tHeader = []
            res.data.headers.map(item => tHeader.push(item.title))
            const data = excel.formatJson(res.data)
            if (data.length === 0) {
              this.$message.info("没有筛选到需导出的数据")
              return
            }
            excel.export_json_to_excel({header: tHeader, data, filename: "广告平台充值记录"})
            this.loadings.downloadLoading = false
          }).catch(err => {
            console.log(err)
            this.loadings.downloadLoading = false
          })
        }).catch(() => {
          this.loadings.downloadLoading = false
        })
      },
      viewFile(file_id) {
        window.open(fileView + file_id)
      },
      destroyRow(row) {
        this.$confirm('确认作废此流水么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.loadings.pageLoading = true
          rechargePlatformDestroy(row.id).then(() => {
            this.$message.success('作废成功')
            this.getRechargeList()
          }).catch(err => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
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
        this.getRechargeList()
      }
    }
  }
</script>
