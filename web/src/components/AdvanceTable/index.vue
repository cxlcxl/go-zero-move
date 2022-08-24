<template>
  <!--因 element-ui 递归组件生成的表格会出现顺序错乱问题，仅支持三层（含三层）以内数据表头-->
  <el-table :data="data" v-loading="dataLoading" border size="mini" highlight-current-row stripe show-summary
            :summary-method="getSummaries">
    <el-table-column :prop="c1.prop" :min-width="c1.minWidth" :align="c1.align" v-for="c1 in headers">
      <template slot="header" slot-scope="scope">
        {{ c1.title }}
        <el-tooltip v-if="c1.desc !== ''" :content="c1.desc" placement="top"><i class="el-icon-info"/>
        </el-tooltip>
      </template>
      <template slot-scope="scope"><span :class="{'text-error': c1.highlight}">{{ scope.row[c1.prop]|isMoney(c1.isFormat) }}</span></template>

      <el-table-column v-if="c1.hasOwnProperty('children') && Array.isArray(c1.children) && c1.children.length > 0"
                       v-for="c2 in c1.children" :prop="c2.prop" :min-width="c2.minWidth" :align="c2.align">
        <template slot="header" slot-scope="c_scope">
          {{ c2.title }}
          <el-tooltip v-if="c2.desc !== ''" :content="c2.desc" placement="top"><i class="el-icon-info"/></el-tooltip>
        </template>
        <template slot-scope="scope"><span :class="{'text-error': c2.highlight}">{{ scope.row[c2.prop]|isMoney(c2.isFormat) }}</span></template>

        <el-table-column v-if="c2.hasOwnProperty('children') && Array.isArray(c2.children) && c2.children.length > 0"
                         v-for="c3 in c2.children" :prop="c3.prop" :min-width="c3.minWidth" :align="c3.align">
          <template slot="header" slot-scope="c_scope">
            {{ c3.title }}
            <el-tooltip v-if="c3.desc !== ''" :content="c3.desc" placement="top"><i class="el-icon-info"/></el-tooltip>
          </template>
          <template slot-scope="scope"><span :class="{'text-error': c3.highlight}">{{ scope.row[c3.prop]|isMoney(c3.isFormat) }}</span></template>
        </el-table-column>
      </el-table-column>
    </el-table-column>
  </el-table>
</template>

<script>
import {formatMoney} from "@/utils"

export default {
  name: "AdvanceTable",
  filters: {
    isMoney(n, isFormat) {
      return isFormat ? formatMoney(n) : n
    }
  },
  props: {
    data: {
      type: Array,
      required: true
    },
    dataLoading: {
      type: Boolean,
      required: true
    },
    headers: {
      type: Array,
      required: true
    },
    summaryText: {
      type: String,
      default: '汇总'
    },
    excludeSummaryColumns: {
      // 不汇总的列，非数字类型会自动判断
      type: Array,
      default: () => []
    }
  },
  methods: {
    getSummaries(param) {
      const {columns, data} = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = this.summaryText
        } else {
          let isSum = true
          const values = data.map(item => {
            if (typeof item[column.property] !== 'number') {
              isSum = false
            }
            return Number(item[column.property])
          })
          if (this.excludeSummaryColumns.includes(column.property) || !isSum) {
            sums[index] = 0
          } else {
            let tmp = 0
            tmp = values.reduce((prev, curr) => (prev + curr), 0).toFixed(2)
            sums[index] = formatMoney(tmp)
          }
        }
      })
      return sums
    },
  }
}
</script>
