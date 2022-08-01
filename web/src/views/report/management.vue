<template>
  <div class="profit-items-container">
    <div class="search-form">
      <el-form ref="search" :inline="true" size="mini">
        <el-form-item>
          <el-radio-group v-model="search.date_type" size="mini">
            <el-radio-button :label="1">季</el-radio-button>
            <el-radio-button :label="2">月</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="search.date_type === 1">
          <el-cascader v-model="search.start_quarter" :options="quarterOptions" class="w100" @change="setEndQuarters" />&nbsp;
          <el-cascader v-model="search.end_quarter" :options="endOptions" class="w100" clearable />
        </el-form-item>
        <el-form-item v-else>
          <el-date-picker
            v-model="search.month_range"
            :picker-options="pickerOptions"
            class="report-range w200"
            :clearable="false"
            value-format="yyyy-MM"
            type="monthrange"
            start-placeholder="开始月份"
            end-placeholder="结束月份"
          />
        </el-form-item>
        <el-form-item style="margin-left: 5px;">
          <el-checkbox v-model="search.has_customer" border label="客户筛选" :false-label="0" :true-label="1" style="margin-right: 5px;" />
          <div v-show="search.has_customer" style="display: inline-block;">
            <el-select v-model="search.customer_ids" filterable multiple collapse-tags placeholder="请选择客户，不选等于全选" class="w230" clearable>
              <el-option v-for="(name, id) in customers" :label="name" :value="Number(id)" />
            </el-select>
            <el-button plain type="primary" style="margin-left: 3px;" @click="allCustomer">全选</el-button>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" @click="getData">查询</el-button>
        </el-form-item>
        <div style="float: right; display: inline-block;">
          <el-button-group>
            <el-button icon="el-icon-download" size="mini" @click="exportData">一键导出</el-button>
          </el-button-group>
        </div>
      </el-form>

      <el-tabs type="border-card">
        <el-tab-pane label="收付明细(每天 01:35 刷新)">
          <p style="margin-bottom: 10px; text-align: right;">
            <el-button-group>
              <el-button v-permission="'report/profit-refresh'" icon="el-icon-refresh" size="mini" @click="refreshReport">手动计算明细</el-button>
              <el-button v-permission="'report/archive'" icon="el-icon-setting" size="mini" @click="archiveData">归档数据</el-button>
            </el-button-group>
          </p>
          <income-table ref="profit" :income-data="profitData.list" :i-columns="columns.profit" :income-loading="pageLoading" :exclude-summary="['is_archive']" />
        </el-tab-pane>
        <el-tab-pane label="充值收支(实时)">
          <income-table ref="recharge" :income-data="rechargeData.list" :i-columns="columns.recharge" :income-loading="pageLoading" />
        </el-tab-pane>
        <el-tab-pane label="转账收支(每 2 小时刷新)">
          <p style="margin-bottom: 10px; text-align: right;">
            <el-button-group>
              <el-button v-permission="'report/transfer-refresh'" icon="el-icon-refresh" size="mini" @click="refreshTransfer">更新转账统计</el-button>
            </el-button-group>
          </p>
          <income-table ref="transfer" :income-data="transferData.list" :i-columns="columns.transfer" :income-loading="pageLoading" />
        </el-tab-pane>
      </el-tabs>
    </div>

    <management-archive ref="archive" @success="flowProfit" />
  </div>
</template>

<script>
import { incomeProfit, incomeRecharge, incomeTransfer, refreshManagementReport, refreshTransferReport } from '@a/report'
import { quarters, profitColumns, rechargeColumns, transferColumns, managementExportHeaders, managementExportHeaderCodes } from './report'
import { getAllCustomer } from '@a/customer'
import IncomeTable from './components/management-table'
import ManagementArchive from './components/management-archive'

export default {
  name: "ReportManagement",
  components: { IncomeTable, ManagementArchive },
  data() {
    return {
      pageLoading: false,
      endOptions: [],
      customers: [],
      now: new Date(),
      search: {
        date_type: 1,
        start_quarter: [2021, 1],
        month_range: [],
        has_customer: 0,
        end_quarter: [],
        customer_ids: []
      },
      profitData: { list: [], account_types: {}},
      rechargeData: { list: [], recharge_types: {}, platform_recharge_types: {}},
      transferData: { list: [], accounts: {}},
      columns: {
        profit: [],
        recharge: [],
        transfer: []
      },
      pickerOptions: {
        shortcuts: [{
          text: '本月',
          onClick(picker) {
            picker.$emit('pick', [new Date(), new Date()])
          }
        }, {
          text: '今年至今',
          onClick(picker) {
            const end = new Date()
            const start = new Date(new Date().getFullYear(), 0)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '近六月',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setMonth(start.getMonth() - 5)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '上一年',
          onClick(picker) {
            const baseDate = new Date().getFullYear() - 1
            picker.$emit('pick', [baseDate.toString() + '-01', baseDate.toString() + '-12'])
          }
        }]
      }
    }
  },
  computed: {
    quarterOptions() {
      const monthQuarters = []
      for (let i = 2021; i <= this.now.getFullYear(); i++) {
        monthQuarters.push({
          value: i,
          label: i,
          children: quarters
        })
      }
      return monthQuarters
    }
  },
  mounted() {
    this.getData()
    this.setEndQuarters(this.search.start_quarter)
  },
  methods: {
    getData() {
      this.pageLoading = true
      Promise.all([
        this.flowProfit(),
        this.flowRecharge(),
        this.flowTransfer(),
        this.getCustomers()
      ]).then(() => {
        this.pageLoading = false
      }).catch(() => {
        this.pageLoading = false
      })
    },
    getCustomers() {
      return new Promise((resolve, reject) => {
        if (this.customers.length > 0) {
          return resolve()
        }
        getAllCustomer().then(res => {
          this.customers = res.data
          return resolve()
        }).catch(() => {
          return reject()
        })
      })
    },
    flowProfit() {
      return new Promise((resolve, reject) => {
        incomeProfit(this.search).then(res => {
          this.profitData = res.data
          this.$set(this.columns, 'profit', profitColumns(this.search.date_type, this.search.has_customer))
          resolve()
        }).catch(() => {
          reject()
        })
      })
    },
    flowRecharge() {
      return new Promise((resolve, reject) => {
        incomeRecharge(this.search).then(res => {
          this.rechargeData = res.data
          this.$set(this.columns, 'recharge', rechargeColumns(res.data.platform_recharge_types, res.data.recharge_types))
          resolve()
        }).catch(() => {
          reject()
        })
      })
    },
    flowTransfer() {
      return new Promise((resolve, reject) => {
        incomeTransfer(this.search).then(res => {
          this.transferData = res.data
          this.$set(this.columns, 'transfer', transferColumns(res.data.accounts))
          resolve()
        }).catch(() => {
          reject()
        })
      })
    },
    setEndQuarters(v) {
      const monthQuarters = []
      for (let i = v[0]; i <= this.now.getFullYear(); i++) {
        let q = []
        if (v[0] === i) {
          for (let k = v[1] - 1; k < 4; k++) {
            q.push(quarters[k])
          }
        } else {
          q = quarters
        }
        monthQuarters.push({ value: i, label: i, children: q })
      }
      this.endOptions = monthQuarters
    },
    archiveData() {
      this.$refs.archive.getArchiveQs()
    },
    refreshReport() {
      this.$confirm('确认刷新收付明细么?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.pageLoading = true
        refreshManagementReport().then(() => {
          this.$message.success('刷新成功')
          this.flowProfit()
          this.pageLoading = false
        }).catch(() => {
          this.pageLoading = false
        })
      }).catch(() => {})
    },
    refreshTransfer() {
      this.$confirm('确认刷新转账收支明细么?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.pageLoading = true
        refreshTransferReport().then(() => {
          this.$message.success('刷新成功')
          this.flowTransfer()
          this.pageLoading = false
        }).catch(() => {
          this.pageLoading = false
        })
      }).catch(() => {})
    },
    exportData() {
      this.loading = true
      const tHeaders = Object.assign({}, managementExportHeaders)
      const tmpHeaders = Object.assign({}, managementExportHeaderCodes)
      const keyName = { payment: '收付明细', recharge: '充值收支', transfer: '转账收支' }
      import('@/vendor/Export2Excel').then(excel => {
        const data = { payment: [], transfer: [] }
        for (const k in tHeaders) {
          let d = []
          if (k === 'payment') {
            d = this.profitData.list
            if (this.search.has_customer === 1) {
              // unshift 会改变原数组数据
              tHeaders[k] = ['客户'].concat(tHeaders[k])
              tmpHeaders[k] = [{ code: 'customer_name' }].concat(tmpHeaders[k])
            }
            tHeaders[k] = ['时间'].concat(tHeaders[k])
            tmpHeaders[k] = [{ code: this.search.date_type === 1 ? 'date_quarter' : 'date_month' }].concat(tmpHeaders[k])
          } else if (k === 'recharge') {
            d = this.rechargeData.list
          } else {
            d = this.transferData.list
          }
          data[k] = excel.formatJson({ headers: tmpHeaders[k], list: d }, 0)
        }
        excel.export_mutilate_sheet({ headers: tHeaders, data, filename: '经营报表', keyName })
        this.loading = false
      }).catch(err => {
        this.loading = false
      })
    },
    allCustomer() {
      for (const id in this.customers) {
        if (!this.search.customer_ids.includes(Number(id))) {
          this.search.customer_ids.push(Number(id))
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  .gird-title{
    margin-bottom: 8px;
    color: #606266;
    font-weight: 600;
    padding-left: 10px;
    box-sizing: border-box;
    border-left: 2px solid #409eff;
    height: 26px;
    line-height: 26px;

    .export {
      float: right;
    }
  }
</style>
