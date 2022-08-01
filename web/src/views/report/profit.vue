<template>
  <div class="profit-container">
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
          <el-date-picker v-model="search.month_range" :picker-options="pickerOptions" class="report-range w200" :clearable="false"
                          value-format="yyyy-MM" type="monthrange" start-placeholder="开始月份" end-placeholder="结束月份"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" @click="getData">查询</el-button>
        </el-form-item>
        <div style="float: right; display: inline-block;">
          <el-button-group>
            <el-button icon="el-icon-download" size="mini" @click="downloadInfo">下载数据</el-button>
          </el-button-group>
        </div>
      </el-form>
    </div>
    <div class="data-tips">
      <p>注：从其他报表取数，最长的数据更新周期为每天凌晨 1 次</p>
    </div>

    <advance-table :data="profitList" :data-loading="loading" :headers="headers"/>
  </div>
</template>

<script>
import {profit} from '@a/report'
import AdvanceTable from '@c/AdvanceTable'
import {profitHeaders, quarters} from "@/views/report/report"

export default {
  name: "ReportProfit",
  components: {
    AdvanceTable
  },
  data() {
    return {
      profitList: [],
      endOptions: [],
      loading: false,
      search: {
        date_type: 1,
        start_quarter: [2021, 1],
        end_quarter: [],
        month_range: [],
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
    headers() {
      return profitHeaders
    },
    quarterOptions() {
      const monthQuarters = []
      for (let i = 2021; i <= (new Date()).getFullYear(); i++) {
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
  },
  methods: {
    getData() {
      this.loading = true
      profit(this.search).then(res => {
        this.profitList = res.data
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    downloadInfo() {
      import('@/vendor/Export2Excel').then(excel => {
        let exportHeaders = []
        let tHeaders = []
        profitHeaders.map(headerInfo => {
          if (headerInfo.hasOwnProperty('children')) {
            headerInfo.children.map(item => {
              exportHeaders.push({code: item.prop})
              tHeaders.push(headerInfo.title + '-' + item.title)
            })
          } else {
            exportHeaders.push({code: headerInfo.prop})
            tHeaders.push(headerInfo.title)
          }
        })
        const data = excel.formatJson({list: this.profitList, headers: exportHeaders})
        if (data.length === 0) {
          this.$message.info("没有筛选到需导出的数据")
          return
        }
        excel.export_json_to_excel({header: tHeaders, data, filename: "仪表盘-利润数据"})
      }).catch(err => {})
    },
    setEndQuarters(v) {
      const monthQuarters = []
      for (let i = v[0]; i <= (new Date()).getFullYear(); i++) {
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
  }
}
</script>
