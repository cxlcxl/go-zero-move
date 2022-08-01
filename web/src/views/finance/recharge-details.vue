<template>
  <div class="payment-flow">
    <el-row>
      <el-col :span="24" class="search-container">
        <el-form ref="_search" :model="search" inline size="mini">
          <el-form-item>
            <el-input v-model="search.recharge_sn" clearable placeholder="流水单号" />
          </el-form-item>
          <el-form-item>
            <el-date-picker v-model="search.start_date" type="date" placeholder="开始时间" value-format="yyyy-MM-dd" class="w130" />&nbsp;
            <el-date-picker v-model="search.end_date" type="date" placeholder="结束时间" value-format="yyyy-MM-dd" class="w130" />
          </el-form-item>
          <el-form-item>
            <el-input v-model="search.customer_name" class="w200" clearable placeholder="客户名称" />
          </el-form-item>
          <el-form-item>
            <el-select v-model="search.state" placeholder="状态" clearable class="w130">
              <el-option v-for="(name, key) in rechargeList.flow_state" :label="name" :value="key" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-select v-model="search.recharge_type" placeholder="充值类型" clearable class="w130">
              <el-option v-for="(name, key) in rechargeList.types" :label="name" :value="key" />
            </el-select>
          </el-form-item>
          <el-form-item label="">
            <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="24" style="margin-bottom: 15px;">
        <el-button v-permission="'recharge/create'" icon="el-icon-plus" size="small" type="primary" @click="add">添加流水</el-button>
        <el-button v-permission="'recharge/detail-download'" icon="el-icon-download" size="small" :loading="loadings.downloadLoading" @click="download">导出</el-button>
      </el-col>
      <el-col :span="24">
        <el-table v-loading="loadings.pageLoading" :data="rechargeList.list" highlight-current-row stripe border show-summary :span-method="objectSpanMethod"
                  :summary-method="getSummaries" size="mini" class="first-summary-table" @sort-change="tableSort">
          <el-table-column prop="recharge_sn" label="所属打款单编号" width="120" align="center" />
          <el-table-column prop="payment_amount" label="所属打款单金额" width="120" align="right">
            <template slot-scope="scope">{{ scope.row.payment_amount|money }}</template>
          </el-table-column>
          <el-table-column prop="customer_name" label="客户名称" width="120" show-overflow-tooltip/>
          <el-table-column label="充值类型" width="110" align="center">
            <template slot-scope="scope">{{ rechargeList.types[scope.row.recharge_type] }}</template>
          </el-table-column>
          <el-table-column prop="amount" label="金额" width="120" align="right">
            <template slot-scope="scope">{{ scope.row.amount|money }}</template>
          </el-table-column>
          <el-table-column prop="trans_date" label="交易日期" width="100" align="center" sortable/>
          <el-table-column prop="remark" label="备注" min-width="220"/>
          <el-table-column label="状态" width="120" align="center">
            <template slot-scope="scope">
              <span :class="scope.row.state|stateClassFilter">
                {{ rechargeList.flow_state[scope.row.state] }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="created_at" label="添加时间" width="140"/>
          <el-table-column label="操作" width="100" align="center">
            <template slot-scope="scope">
              <el-button-group class="table-operate">
                <el-button type="primary" plain @click.native.prevent="relationDetail(scope.row)">关联详情</el-button>
              </el-button-group>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
      <el-col :span="24" style="margin-top: 15px;">
        <page :page="search.page" :total="rechargeList.total" @current-change="handlePage" @size-change="handlePageSize"/>
      </el-col>
    </el-row>

    <recharge-create ref="create" :recharge-types="rechargeList.types" @success="getRechargeList" />
    <relation-flow ref="relation" flow-type="transfer"/>
  </div>
</template>

<script>
import DialogPanel from '@c/DialogPanel'
import { rechargeDetailList, rechargeDetailDownload, rechargeDestroy } from '@a/recharge'
import Page from '@c/Page'
import RechargeCreate from './components/recharge-add'
import { StateRechargeTheme } from './variables'
import { formatMoney } from '@/utils'
import RelationFlow from './components/relation-flow'

export default {
  name: 'RechargeDetail',
  components: {
    Page, RechargeCreate, DialogPanel, RelationFlow
  },
  filters: {
    stateClassFilter(state) {
      return StateRechargeTheme[state]
    },
    money(n) {
      return formatMoney(n)
    }
  },
  data() {
    return {
      loadings: {
        pageLoading: false,
        downloadLoading: false
      },
      spanRows: [],
      rechargeList: {
        total: 0,
        amount: 0,
        list: [],
        flow_state: {},
        types: {}
      },
      search: {
        recharge_sn: '',
        recharge_type: '',
        customer_name: '',
        state: '',
        start_date: '',
        end_date: '',
        order_field: '',
        order_type: '',
        page: 1,
        limit: 10
      }
    }
  },
  mounted() {
    this.getRechargeList()
  },
  methods: {
    getRechargeList() {
      this.loadings.pageLoading = true
      rechargeDetailList(this.search).then(res => {
        this.rechargeList = res.data
        this.getRowSpan(res.data.list)
        this.loadings.pageLoading = false
      }).catch(() => {
        this.loadings.pageLoading = false
      })
    },
    getSummaries(param) {
      const { columns } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
        } else if (['amount'].includes(column.property)) {
          sums[index] = formatMoney(this.rechargeList.amount)
        }
      })

      return sums
    },
    add() {
      this.$refs.create.initCreate()
    },
    download() {
      this.loadings.downloadLoading = true
      rechargeDetailDownload(this.search).then(res => {
        this.loadings.downloadLoading = false

        import('@/vendor/Export2Excel').then(excel => {
          const tHeader = []
          res.data.headers.map(item => tHeader.push(item.title))
          const data = excel.formatJson(res.data)
          if (data.length === 0) {
            this.$message.info('没有筛选到需导出的数据')
            return
          }
          excel.export_json_to_excel({ header: tHeader, data, filename: '客户打款明细' })
          this.loadings.downloadLoading = false
        }).catch(err => {
          console.log(err)
          this.loadings.downloadLoading = false
        })
      }).catch(() => {
        this.loadings.downloadLoading = false
      })
    },
    doSearch() {
      this.search.page = 1
      this.getRechargeList()
    },
    handlePage(p) {
      this.search.page = p
      this.getRechargeList()
    },
    handlePageSize(s) {
      this.search.limit = s
      this.getRechargeList()
    },
    destroyRow(row) {
      this.$confirm('确认作废此流水么?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }).then(() => {
        this.loadings.pageLoading = true
        rechargeDestroy(row.id).then(() => {
          this.$message.success('作废成功')
          this.getRechargeList()
        }).catch(err => {
          this.loadings.pageLoading = false
        })
      }).catch(() => {
      })
    },
    relationDetail(row) {
      this.$refs.relation.getFlows(row.id)
    },
    tableSort({_, prop, order}){
      if (order !== null) {
        this.$set(this.search, 'order_type', order)
        this.$set(this.search, 'order_field', prop)
      } else {
        this.$set(this.search, 'order_type', '')
        this.$set(this.search, 'order_field', '')
      }
      this.getRechargeList()
    },
    objectSpanMethod({ row, column, rowIndex, columnIndex }) {
      if (columnIndex <= 1) {
        const _row = this.spanRows[rowIndex]
        const _col = _row > 0 ? 1 : 0
        return {
          rowspan: _row,
          colspan: _col
        };
      }
    },
    getRowSpan(data) {
      this.spanRows = []
      let idx = 0
      if (Array.isArray(data)) {
        for (let i = 0; i < data.length; i++) {
          if (i === 0) {
            this.spanRows.push(1)
          } else {
            if (data[i].recharge_sn === data[i-1].recharge_sn) {
              this.spanRows[idx] += 1
              this.spanRows.push(0)
            } else {
              this.spanRows.push(1)
              idx = i
            }
          }
        }
      }
    },
  }
}
</script>
