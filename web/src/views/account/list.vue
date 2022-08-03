<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.account_name" class="w200" clearable placeholder="账户名"/>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.account_type" class="w130">
            <el-option label="全部账户类型" :value="0"/>
            <el-option v-for="(key, val) in account_types" :label="key" :value="Number(val)"/>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" class="w100">
            <el-option v-for="(key, val) in accountList.state" :label="key" :value="Number(val)"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加账户</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="accountList.list" highlight-current-row stripe border size="mini"
                style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="70" align="center"/>
        <el-table-column prop="account_name" label="账号名称" width="150" show-overflow-tooltip/>
        <el-table-column label="账号类型" width="100" align="center" show-overflow-tooltip>
          <template slot-scope="scope">{{ account_types[scope.row.account_type] }}</template>
        </el-table-column>
        <el-table-column prop="advertiser_id" label="广告主ID" width="160"/>
        <el-table-column prop="developer_id" label="开发者ID" width="160"/>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center"/>
        <el-table-column label="状态" width="80" align="center">
          <template slot-scope="scope">
            <span :class="(scope.row.state === 0 ? 'text-error' : '')">{{scope.row.state|stateFilter(accountList.state)}}</span>
          </template>
        </el-table-column>
        <el-table-column label="ClientID" show-overflow-tooltip>
          <template slot-scope="scope">{{scope.row.client_id}} / {{scope.row.secret}}</template>
        </el-table-column>
        <el-table-column label="认证状态" width="100" align="center">
          <template slot-scope="scope">
            <span v-if="scope.row.is_auth === 1" class="text-success">已认证</span>
            <span v-else>待认证</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="130">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row.id)">编辑</el-button>
              <el-button type="primary" plain @click.native.prevent="doAuth(scope.row.id, scope.row.is_auth)">认证</el-button>
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
  import {accountList, accountAuth} from '@a/account'
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
            this.accountList.state = res.state
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
      editRow(id) {
        this.$refs.accountUpdate.initUpdate(id)
      },
      doAuth(id, is_auth) {
        let msg = is_auth === 1 ? '此账户已做授权认证，确定重新认证此账户授权信息吗?' : '确认认证此账号么?'
        this.$confirm(msg, '确认信息', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'success'
        }).then(() => {
          this.loadings.pageLoading = true
          accountAuth(id).then(res => {
            this.loadings.pageLoading = false
            window.open(res.data)
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
