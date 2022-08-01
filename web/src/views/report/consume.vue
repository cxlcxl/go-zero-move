<template>
  <div v-loading="loading" class="report-container">
    <div class="search-form">
      <el-form ref="search" :inline="true" size="mini">
        <el-form-item style="margin-left: 5px;">
          <el-checkbox v-model="search.has_customer" border label="客户筛选" :false-label="0" :true-label="1" style="margin-right: 5px;" />
          <div v-show="search.has_customer" style="display: inline-block;">
            <el-select
              v-model="search.customer_ids"
              filterable
              multiple
              collapse-tags
              placeholder="请选择客户，不选等于全选"
              class="w230"
              clearable
            >
              <el-option v-for="(name, id) in customers" :label="name" :value="Number(id)" />
            </el-select>
            <el-button plain type="primary" style="margin-left: 3px;" @click="allCustomer">全选</el-button>
          </div>
        </el-form-item>
        <el-form-item>
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
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" @click="getReportData">查询</el-button>
          <el-button icon="el-icon-download" @click="output">导出</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-tabs v-model="activeTopic" type="border-card" size="mini" @tab-click="handleChangeTopic">
      <el-tab-pane v-for="(title, topic) in consumeList.topics" :label="title" :name="topic" />
      <line-chart v-if="renderChart" ref="line" :chart-data="lineChartData" :x-axis="search.month_range" />
    </el-tabs>

    <data-table ref="data" style="margin-top: 10px;" :data-list="consumeList" :summary-columns="summaryColumns" />
  </div>
</template>

<script>
import LineChart from './charts/line'
import DataTable from './data-tables'
import { consumeReport } from '@a/report'
import { getAllCustomer } from '@a/customer'
import { parseTime } from '@/utils'
import { formatLineChart } from './report'

export default {
  name: 'ReportConsume',
  components: {
    LineChart, DataTable
  },
  data() {
    return {
      loading: false,
      renderChart: true,
      activeTopic: 'CASH',
      summaryColumns: [],
      search: {
        has_customer: 0,
        customer_ids: [],
        month_range: []
      },
      lineChartData: [],
      customers: [],
      consumeList: {
        list: [],
        columns: [],
        topics: {}
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
  created() {
    this.setDefaultDateRange()
  },
  mounted() {
    this.getReportData()
  },
  methods: {
    getReportData() {
      this.loading = true
      Promise.all([
        this.getCustomers()
      ]).then(res => {
        consumeReport(this.search).then(res => {
          if (res.data.topics) {
            for (const topic in res.data.topics) {
              this.summaryColumns.push(topic)
            }
          }
          // 部分金额返回的字符串类型影响 el-table 排序
          if (Array.isArray(res.data.list)) {
            for (let i = 0; i < res.data.list.length; i++) {
              res.data.list[i].amount = Number(res.data.list[i].amount)
              this.summaryColumns.map(key => {
                res.data.list[i][key] = Number(res.data.list[i][key])
              })
            }
          }
          this.consumeList = res.data
          this.reRenderChart(this.activeTopic)
          this.loading = false
        }).catch(() => {
          this.loading = false
        })
      }).catch(() => {
        this.loading = false
      })
    },
    // 重新渲染图表
    reRenderChart(topic) {
      const chartData = formatLineChart(topic, this.consumeList.list, this.search.month_range, this.search.has_customer)
      this.renderChart = false
      this.$nextTick(() => {
        this.lineChartData = chartData
        this.renderChart = true
      })
    },
    handleChangeTopic(tab) {
      this.reRenderChart(tab.name)
    },
    setDefaultDateRange() {
      const start = new Date()
      this.$set(this.search, 'month_range', [
        parseTime(start.setMonth(start.getMonth() - 5), '{y}-{m}'),
        parseTime(new Date(), '{y}-{m}')
      ])
    },
    allCustomer() {
      for (const id in this.customers) {
        if (!this.search.customer_ids.includes(Number(id))) {
          this.search.customer_ids.push(Number(id))
        }
      }
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
    output() {
      this.loading = true
      import('@/vendor/Export2Excel').then(excel => {
        const tHeader = []; const tmpHeaders = []
        this.consumeList.columns.map(item => {
          tHeader.push(item.title)
          tmpHeaders.push({ code: item.field })
        })
        const data = excel.formatJson({ headers: tmpHeaders, list: this.consumeList.list })
        if (data.length === 0) {
          this.$message.info('没有筛选到需导出的数据')
          this.loading = false
          return
        }
        excel.export_json_to_excel({ header: tHeader, data, filename: '消耗报表数据' })
        this.loading = false
      }).catch(err => {
        this.loading = false
      })
    }
  }
}
</script>

<style lang="scss">
  @import "style";
</style>
