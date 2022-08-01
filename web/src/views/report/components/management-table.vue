<template>
  <i-table
    ref="itable"
    v-loading="IncomeLoading"
    :data="IncomeData"
    :columns="IColumns"
    :summary-method="handleSummary"
    size="small"
    show-summary
    border
    highlight-row
  />
</template>

<script>
import { formatMoney } from '@/utils'

export default {
  name: 'ManagementTable',
  props: {
    IncomeData: {
      type: Array,
      required: true
    },
    IColumns: {
      type: Array,
      required: true
    },
    IncomeLoading: {
      type: Boolean,
      required: true
    },
    excludeSummary: {
      type: Array,
      default: () => {
        return []
      }
    }
  },
  methods: {
    handleSummary({ columns, data }) {
      const sums = {}
      columns.forEach((column, index) => {
        const key = column.key
        if (index === 0) {
          sums[key] = { key, value: '合计' }
          return
        }
        if (this.excludeSummary.includes(key)) {
          sums[key] = { key, value: '' }
          return
        }
        const values = data.map(item => Number(item[key]))
        if (!values.every(value => isNaN(value))) {
          const v = values.reduce((prev, curr) => {
            const value = Number(curr)
            if (!isNaN(value)) {
              return prev + curr
            } else {
              return prev
            }
          }, 0)
          sums[key] = { key, value: formatMoney(v) }
        } else {
          sums[key] = { key, value: '0' }
        }
      })

      return sums
    }
  }
}
</script>
