<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.account_id" class="w230" clearable placeholder="AccountID"/>
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.account_name" class="w200" clearable placeholder="账户名"/>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.account_type" class="w120" clearable>
            <el-option label="账号类型" :value="0"/>
            <el-option v-for="(key, val) in account_types" :label="key" :value="Number(val)"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加账号&客户</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="accountList.list" highlight-current-row stripe border size="mini"
                style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="70" align="center"/>
        <el-table-column prop="account_name" label="账号名称" width="150" show-overflow-tooltip/>
        <el-table-column label="账号类型" width="100" align="center" show-overflow-tooltip>
          <template slot-scope="scope">{{ account_types[scope.row.account_type] }}</template>
        </el-table-column>
        <el-table-column prop="email" label="邮箱" width="180"/>
        <el-table-column prop="account_id" label="AccountID" width="160"/>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center"/>
        <el-table-column label="状态" width="80" align="center">
          <template slot-scope="scope">
            <span :class="(scope.row.state === 0 ? 'text-error' : '')">{{scope.row.state|stateFilter(accountList.state)}}</span>
          </template>
        </el-table-column>
        <el-table-column label="ClientID" min-width="300" show-overflow-tooltip>
          <template slot-scope="scope">
            <p v-for="item in scope.row.account_client_ids">
              {{item.client_id}} / {{item.secret}} / {{item.comment}}
            </p>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="100">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="accountList.total" @current-change="handlePage" @size-change="handlePageSize"/>

      <account-create ref="accountCreate" @success="getAccountList" :account-types="account_types"/>
      <account-update ref="accountUpdate" @success="getAccountList" :account-types="account_types"/>
    </el-col>
  </el-row>
</template>

<script>
  import {accountList, accountDestroy} from '@a/account'
  import AccountCreate from './components/add-act'
  import AccountUpdate from './components/edit-act'
  import Page from '@c/Page'

  export default {
    // name: 'CustomerAccountList',
    components: {
      AccountCreate,
      AccountUpdate,
      Page,
    },
    data() {
      return {
        loadings: {
          pageLoading: false
        },
        account_types: {},
        accountList: {
          total: 0,
          list: [],
          state: {}
        },
        roles: [],
        search: {
          account_id: '',
          account_name: '',
          account_type: 0,
          state: 1,
          page: 1,
          page_size: 10
        }
      }
    },
    mounted() {
      this.getAccountList()
    },
    filters: {
      stateFilter(s, state) {
        return state[s]
      }
    },
    methods: {
      getAccountList() {
        this.loadings.pageLoading = true
        Promise.all([

        ]).then(response => {
          accountList(this.search).then(res => {
            this.accountList.list = res.data
            this.accountList.total = res.total
            this.account_types = res.account_types
            this.loadings.pageLoading = false
          }).catch(() => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
          this.loadings.pageLoading = false
        })
      },
      add() {
        this.$refs.accountCreate.initCreate()
      },
      accountInfo(row) {
        this.$refs.account.initAccountInfo(row.id)
      },
      editRow(row) {
        this.$refs.accountUpdate.initUpdate({
          id: row.id,
          account_name: row.account_name,
          platform_id: row.platform_id,
          account_type: row.account_type,
          provider_account_id: row.provider_account_id,
        })
      },
      changeContract(row) {
        this.$refs.contract.initUpdate({
          id: row.id,
          customer_id: row.customer_id,
          contract_id: row.contract_id,
        })
      },
      destroyRow(row) {
        this.$confirm('确认作废此账号么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.loadings.pageLoading = true
          accountDestroy(row.id).then(() => {
            this.$message.success('作废成功')
            this.getAccountList()
          }).catch(err => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
        })
      },
      doSearch() {
        this.search.page = 1
        this.getAccountList()
      },
      handlePage(p) {
        this.search.page = p
        this.getAccountList()
      },
      handlePageSize(p) {
        this.search.limit = p
        this.getAccountList()
      }
    }
  }
</script>
