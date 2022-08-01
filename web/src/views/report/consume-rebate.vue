<template>
  <div v-loading="loading" class="report-container">
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
          <el-button type="primary" icon="el-icon-search" @click="getReportData">查询</el-button>
          <el-button v-permission="'report/rebate-refresh'" icon="el-icon-refresh" @click="refreshRebate">计算消耗返点</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="data-tips">
      <p>注：返点报表自动刷新周期：每天 03:05</p>
      <p>平台返点数据不涉及客户，不支持按客户查询；真实返点以季月单位，不支持按月查询</p>
    </div>
    <el-table :data="consumeList" highlight-current-row stripe border
              show-summary :summary-method="getSummaries" size="mini" class="first-summary-table">
      <el-table-column label="时间" prop="quarter" align="center" width="100"/>
      <el-table-column label="客户名称" align="center" width="120" v-if="search_customer">
        <template slot-scope="scope">{{customers[scope.row.customer_id]}}</template>
      </el-table-column>
      <el-table-column label="消耗" prop="consume_amount" align="right" min-width="110">
        <template slot-scope="scope">{{scope.row.consume_amount | money}}</template>
      </el-table-column>
      <el-table-column label="平台返点(试算)" prop="rebate_platform_sum" align="right" min-width="110">
        <template slot-scope="scope">{{scope.row.rebate_platform_sum | money}}</template>
      </el-table-column>
      <el-table-column label="客户返点(试算)" prop="rebate_customer_sum" align="right" min-width="110">
        <template slot-scope="scope">{{scope.row.rebate_customer_sum | money}}</template>
      </el-table-column>
      <el-table-column label="返点盈余(试算)" prop="rebate_sum" align="right" min-width="110">
        <template slot-scope="scope">{{scope.row.rebate_sum | money}}</template>
      </el-table-column>
      <el-table-column label="平台返点(真实)" prop="rebate_platform_real" align="right" min-width="110">
        <template slot-scope="scope">{{scope.row.rebate_platform_real | money}}</template>
      </el-table-column>
      <el-table-column label="客户返点(真实)" prop="rebate_customer_real" align="right" min-width="110">
        <template slot-scope="scope">{{scope.row.rebate_customer_real | money}}</template>
      </el-table-column>
      <el-table-column label="返点盈余(真实)" prop="rebate_real" align="right" min-width="110">
        <template slot-scope="scope">{{scope.row.rebate_real | money}}</template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import BarChart from './charts/bar'
import DataTable from './data-tables'
import { consumeRebate, refreshConsumeRebate } from '@a/report'
import { quarters } from './report'
import { getAllCustomer } from '@a/customer'
import {formatMoney} from "@/utils"

export default {
  name: "ReportConsumeRebate",
  components: {
    BarChart, DataTable
  },
  data() {
    return {
      loading: false,
      search_customer: false,
      summaryColumns: [
        'consume_amount', 'rebate_platform_sum', 'rebate_customer_sum', 'rebate_sum', 'rebate_platform_real',
        'rebate_customer_real', 'rebate_real'
      ],
      now: new Date(),
      search: {
        date_type: 1,
        start_quarter: [2021, 1],
        end_quarter: [],
        month_range: [],
        has_customer: 0,
        customer_ids: []
      },
      endOptions: [],
      consumeList: [],
      customers: {},
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
  filters: {
    money(n) {
      return formatMoney(n)
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
    this.getReportData()
    this.setEndQuarters(this.search.start_quarter)
  },
  methods: {
    getReportData() {
      this.loading = true
      Promise.all([
        this.getCustomers()
      ]).then(() => {
        consumeRebate(this.search).then(res => {
          this.consumeList = res.data
          if (this.search.has_customer === 1) {
            this.search_customer = true
          } else {
            this.search_customer = false
          }
          this.loading = false
        }).catch(() => {
          this.loading = false
        })
      }).catch(() => {
        this.loading = false
      })
    },
    refreshRebate() {
      this.$confirm('从开始时间刷新到本年的最后一个月，只能单年重新计算一次，确认重新计算此设定的返点数据么?', '操作提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.loading = true
        refreshConsumeRebate(this.search).then(() => {
          this.$message.success('返点数据重跑成功')
          this.getReportData()
        }).catch(err => {
          this.loading = false
        })
      }).catch(() => {})
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
    allCustomer() {
      for (const id in this.customers) {
        if (!this.search.customer_ids.includes(Number(id))) {
          this.search.customer_ids.push(Number(id))
        }
      }
    },
    getSummaries(param) {
      const {columns, data} = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
        } else if (this.summaryColumns.includes(column.property)) {
          if (Array.isArray(data)) {
            const values = data.map(item => Number(item[column.property]))
            let tmp = 0
            tmp = values.reduce((prev, curr) => {
              return prev + curr
            }, 0).toFixed(2)
            sums[index] = formatMoney(tmp)
          }
        }
      })

      return sums
    }
  }
}
</script>

<style lang="scss">
  @import "style";
</style>
