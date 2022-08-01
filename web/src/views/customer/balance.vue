<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <div class="search-form">
        <el-form ref="search" :inline="true" size="mini">
          <el-form-item>
            <el-input v-model="search.customer_name" placeholder="客户名称" class="w230" clearable/>
          </el-form-item>
          <el-form-item>
            <el-radio-group v-model="search.show_account" @change="getBalanceList">
              <el-radio-button :label="0">客户维度</el-radio-button>
              <el-radio-button :label="1">账号维度</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" @click="doSearch">查询</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-col>

    <el-col :span="24">
      <el-table v-loading="loading" :data="balanceList.list" :span-method="objectSpanMethod" highlight-current-row size="mini" border>
        <el-table-column prop="customer_name" label="客户名" width="200" align="center"/>
        <el-table-column prop="full_name" label="客户主体" width="300"/>
        <el-table-column prop="account_name" label="账号名称" width="230" show-overflow-tooltip v-if="search.show_account === 1"/>
        <el-table-column prop="account_type" label="账户类型" width="160"/>
        <el-table-column label="余额" width="130" align="right">
          <template slot-scope="scope">
            <span v-if="scope.row.amount < 0" class="text-error">{{scope.row.amount|money}}</span>
            <span v-else class="text-primary">{{scope.row.amount|money}}</span>
          </template>
        </el-table-column>
      </el-table>
    </el-col>

    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :limit="search.limit" :total="balanceList.total" @current-change="handlePage" @size-change="handlePageSize"/>
    </el-col>
  </el-row>
</template>

<script>
  import Page from '@c/Page'
  import {customerBalance} from '@a/customer'
  import {formatMoney} from "@/utils"

  export default {
    name: "CustomerBalance",
    components: {Page},
    data() {
      return {
        loading: false,
        search: {
          customer_name: '',
          show_account: 1,
          page: 1,
          limit: 8,
        },
        spanRows: [],
        spanAccount: [],
        balanceList: {
          total: 0,
          list: []
        },
      };
    },
    filters: {
      money(n) {
        return formatMoney(n)
      }
    },
    mounted() {
      this.getBalanceList()
    },
    methods: {
      getBalanceList() {
        this.loading = true
        customerBalance(this.search).then(res => {
          this.balanceList = res.data
          this.getRowSpan(res.data.list)
          this.loading = false
        }).catch(() => {
          this.loading = false
        })
      },
      handlePage(p) {
        this.search.page = p
        this.getBalanceList()
      },
      handlePageSize(p) {
        this.search.limit = p
        this.getBalanceList()
      },
      doSearch() {
        this.search.page = 1
        this.getBalanceList()
      },
      getRowSpan(data) {
        this.spanRows = []
        this.spanAccount = []
        let idx = 0
        let spanIdx = 0
        if (Array.isArray(data)) {
          for (let i = 0; i < data.length; i++) {
            if (i === 0) {
              this.spanRows.push(1)
              this.spanAccount.push(1)
            } else {
              if (data[i].customer_id === data[i-1].customer_id) {
                this.spanRows[idx] += 1
                this.spanRows.push(0)
              } else {
                this.spanRows.push(1)
                idx = i
              }
              if (data[i].account_id === data[i-1].account_id) {
                this.spanAccount[spanIdx] += 1
                this.spanAccount.push(0)
              } else {
                this.spanAccount.push(1)
                spanIdx = i
              }
            }
          }
        }
      },
      objectSpanMethod({row, column, rowIndex, columnIndex}) {
        if (columnIndex <= 1) {
          const _row = this.spanRows[rowIndex]
          const _col = _row > 0 ? 1 : 0
          return {
            rowspan: _row,
            colspan: _col
          };
        }
      }
    }
  }
</script>

<style scoped lang="scss">
  .account-type {
    display: inline-block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
</style>