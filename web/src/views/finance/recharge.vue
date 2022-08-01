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
          <el-form-item label="">
            <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="24" style="margin-bottom: 15px;">
        <el-button v-permission="'recharge/create'" icon="el-icon-plus" size="small" type="primary" @click="add">添加流水</el-button>
        <el-button v-permission="'recharge/download'" icon="el-icon-download" size="small" :loading="loadings.downloadLoading" @click="download">导出</el-button>
      </el-col>
      <el-col :span="24">
        <el-table v-loading="loadings.pageLoading" :data="rechargeList.list" highlight-current-row stripe border show-summary
                  :summary-method="getSummaries" size="mini" class="first-summary-table" @sort-change="tableSort">
          <el-table-column prop="recharge_sn" label="流水编号" width="120" align="center" />
          <el-table-column prop="customer_name" label="客户名称" width="110" show-overflow-tooltip>
            <template slot-scope="scope">{{ scope.row.customer.customer_name }}</template>
          </el-table-column>
          <el-table-column prop="amount" label="打款金额" width="120" align="right">
            <template slot-scope="scope">{{ scope.row.amount|money }}</template>
          </el-table-column>
          <el-table-column prop="trans_date" label="交易日期" width="100" align="center" sortable/>
          <el-table-column prop="operator" label="操作人" width="110" />
          <el-table-column prop="remark" label="备注" min-width="220"/>
          <el-table-column label="账号名称" width="150" show-overflow-tooltip>
            <template slot-scope="scope">{{ scope.row.account ? scope.row.account.account_name : '' }}</template>
          </el-table-column>
          <el-table-column prop="created_at" label="添加时间" width="140"/>
          <el-table-column label="操作" width="120" align="center">
            <template slot-scope="scope">
              <el-button-group class="table-operate">
                <el-button v-permission="'recharge/destroy'" type="danger" plain :disabled="scope.row.state === 0"
                           @click.native.prevent="destroyRow(scope.row)">作废</el-button>
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
  </div>
</template>

<script>
import { rechargeList, rechargeListDownload, rechargeDestroy } from '@a/recharge'
import Page from '@c/Page'
import RechargeCreate from './components/recharge-add'
import { formatMoney } from '@/utils'

export default {
  name: 'Recharge',
  components: {
    Page, RechargeCreate
  },
  filters: {
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
      rechargeList: {
        total: 0,
        amount: 0,
        list: [],
        flow_state: {},
        types: {}
      },
      search: {
        recharge_sn: '',
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
      rechargeList(this.search).then(response => {
        this.rechargeList = response.data
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
      rechargeListDownload(this.search).then(res => {
        this.loadings.downloadLoading = false

        import('@/vendor/Export2Excel').then(excel => {
          const tHeader = []
          res.data.headers.map(item => tHeader.push(item.title))
          const data = excel.formatJson(res.data)
          if (data.length === 0) {
            this.$message.info('没有筛选到需导出的数据')
            return
          }
          excel.export_json_to_excel({ header: tHeader, data, filename: '客户打款记录' })
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
    tableSort({_, prop, order}){
      if (order !== null) {
        this.$set(this.search, 'order_type', order)
        this.$set(this.search, 'order_field', prop)
      } else {
        this.$set(this.search, 'order_type', '')
        this.$set(this.search, 'order_field', '')
      }
      this.getRechargeList()
    }
  }
}
</script>
