<template>
  <el-row v-loading="summaryLoading" :gutter="15">
    <el-col :span="12">
      <summary-panel :val="summary.monthRecharge" :prefix="summary.currency_type">本月累计充值</summary-panel>
    </el-col>
    <el-col :span="12">
      <summary-panel :val="summary.yearRecharge" :prefix="summary.currency_type">本年累计充值</summary-panel>
    </el-col>
  </el-row>
</template>

<script>
import SummaryPanel from './summary-panel'
import { rechargeSummary } from '@a/global'

export default {
  name: 'Summary',
  components: {
    SummaryPanel
  },
  data() {
    return {
      summaryLoading: false,
      summary: {
        currency_type: '',
        yearRecharge: 0,
        monthRecharge: 0
      }
    }
  },
  mounted() {
    this.getSummaryData()
  },
  methods: {
    getSummaryData() {
      rechargeSummary().then(res => {
        this.summary = res.data
      }).catch(err => {})
    }
  }
}
</script>
