<template>
  <el-card>
    <advance-table :data="profitList" :data-loading="loading" :headers="tHeaders"/>
  </el-card>
</template>

<script>
import {dashboardProfit} from '@a/report'
import {profitHeaders} from "@/views/report/report"
import AdvanceTable from '@c/AdvanceTable'

export default {
  name: 'DashboardProfit',
  components: {
    AdvanceTable
  },
  data() {
    return {
      loading: false,
      profitList: []
    }
  },
  computed: {
    tHeaders() {
      return profitHeaders
    }
  },
  mounted() {
    this.initDashboardProfit()
  },
  methods: {
    initDashboardProfit() {
      this.loading = true
      dashboardProfit().then(res => {
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
    }
  }
}
</script>
