<template>
  <el-table :data="dataList.list" highlight-current-row stripe border
            show-summary :summary-method="getSummaries" size="mini" class="first-summary-table">
    <el-table-column v-for="column in dataList.columns" :label="column.title" :prop="column.field"
                     :sortable="column.sortable" :align="column.align" :width="column.width > 0 ? column.width : 'auto'">
      <template slot-scope="scope">
        {{scope.row[column.field] | money(summaryColumns, column.field)}}
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
  import {formatMoney} from '@/utils'

  export default {
    name: "ReportDataTable",
    props: {
      dataList: Object,
      summaryColumns: Array
    },
    filters: {
      money(n, columns, field) {
        if (columns.includes(field)) {
          return formatMoney(n)
        }
        return n
      }
    },
    methods: {
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
