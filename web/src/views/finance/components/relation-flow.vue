<template>
  <dialog-panel title="关联的流水" :visible="visible" @cancel="cancel" width="900px":confirm-loading="loading">
    <el-table v-loading="loading" :data="flows" highlight-current-row stripe border size="mini"
              show-summary :summary-method="flowSummaries">
      <el-table-column prop="id" label="ID" width="70" align="center"/>
      <el-table-column prop="flow_sn" label="流水编号" width="120" align="center"/>
      <el-table-column prop="trans_date" label="日期" width="100" align="center"/>
      <el-table-column prop="amount" label="流水金额" width="110" align="right">
        <template slot-scope="scope">{{scope.row.amount|money}}</template>
      </el-table-column>
      <el-table-column prop="bind_amount" label="此单关联金额" width="110" align="right">
        <template slot-scope="scope">{{scope.row.bind_amount|money}}</template>
      </el-table-column>
      <el-table-column prop="remark" label="备注"/>
    </el-table>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {boundFlows} from "@a/global"
  import {formatMoney} from "@/utils"

  export default {
    components: {DialogPanel},
    props: {
      flowType: {
        type: String,
        required: true
      }
    },
    data() {
      return {
        visible: false,
        loading: false,
        flows: []
      }
    },
    filters: {
      money(n) {
        return formatMoney(n)
      }
    },
    methods: {
      getFlows(id) {
        boundFlows(id, this.flowType).then(res => {
          this.flows = res.data
          this.visible = true
        }).catch(() => {})
      },
      cancel() {
        this.flows = []
        this.visible = false
      },
      flowSummaries(param) {
        const {columns, data} = param
        const sums = []
        columns.forEach((column, index) => {
          if (index === 0) {
            sums[index] = '总关联'
          } else if (['bind_amount'].includes(column.property)) {
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
