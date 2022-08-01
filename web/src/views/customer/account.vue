<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.account_name" class="w230" clearable placeholder="账号名称/华为Ads-ID"/>
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.customer_name" class="w200" clearable placeholder="所属客户"/>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.account_type" class="w120" clearable>
            <el-option v-for="(key, val) in accountList.account_types" :label="key" :value="Number(val)"/>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" class="w120">
            <el-option v-for="(key, val) in accountList.state" :label="key" :value="Number(val)"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add" v-permission="'account/create'">添加账号&客户</el-button>
      <el-button icon="el-icon-refresh" size="mini" @click="refreshAccountNumbers" v-permission="'account/refresh-number'" type="warning"
                 style="float: right;">同步新增账户配置</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="accountList.list" highlight-current-row stripe border size="mini"
                style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="90" align="center"/>
        <el-table-column prop="account_name" label="账号名称" width="190" show-overflow-tooltip/>
        <el-table-column label="所属客户">
          <template slot-scope="scope">
            {{ scope.row.customer.customer_name }}
          </template>
        </el-table-column>
        <el-table-column label="账号类型" width="100" align="center" show-overflow-tooltip>
          <template slot-scope="scope">{{ accountList.account_types[scope.row.account_type] }}</template>
        </el-table-column>
        <el-table-column label="开户服务商" width="140" show-overflow-tooltip>
          <template slot-scope="scope">{{ providerActs[scope.row.provider_account_id] }}</template>
        </el-table-column>
        <el-table-column label="关联的合同" width="180" show-overflow-tooltip>
          <template slot-scope="scope">
            {{ scope.row.contract ? scope.row.contract.contract_name : '' }}
          </template>
        </el-table-column>
        <el-table-column prop="platform_id" label="广告平台客户 ID" width="180"/>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center"/>
        <el-table-column label="状态" width="80" align="center">
          <template slot-scope="scope">
            <span :class="(scope.row.state === 0 ? 'text-error' : '')">{{scope.row.state|stateFilter(accountList.state)}}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="150">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)" v-permission="'account/update'">编辑</el-button>
              <!--<el-button type="primary" plain @click.native.prevent="accountInfo(scope.row)">账户</el-button>-->
              <el-button type="primary" plain @click.native.prevent="changeContract(scope.row)" v-permission="'account/change-contract'">改合同</el-button>
              <el-button type="danger" plain @click.native.prevent="destroyRow(scope.row)" v-permission="'account/destroy'">作废</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="accountList.total" @current-change="handlePage" @size-change="handlePageSize"/>

      <account-create ref="accountCreate" @success="getAccountList" :account-types="accountList.account_types" :provider="providerActs"/>
      <account-update ref="accountUpdate" @success="getAccountList" :account-types="accountList.account_types" :provider="providerActs"/>
      <change-contract ref="contract" @success="getAccountList"/>
      <account-numbers ref="account"/>
    </el-col>
  </el-row>
</template>

<script>
  import {accountList, refreshNumbers, accountDestroy, providerAccounts} from '@a/account'
  import AccountCreate from './components/add-act'
  import AccountUpdate from './components/edit-act'
  import AccountNumbers from './components/account-numbers'
  import ChangeContract from './components/change-contract'
  import Page from '@c/Page'

  export default {
    name: 'CustomerAccountList',
    components: {
      AccountCreate,
      AccountUpdate,
      ChangeContract,
      Page,
      AccountNumbers
    },
    data() {
      return {
        loadings: {
          pageLoading: false
        },
        providerActs: {},
        accountList: {
          account_types: {},
          total: 0,
          list: [],
          state: {}
        },
        roles: [],
        search: {
          account_name: '',
          customer_name: '',
          account_type: '',
          account_list: 1,
          state: 1,
          page: 1,
          limit: 10
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
          this.provider()
        ]).then(res => {
          accountList(this.search).then(response => {
            this.accountList = response.data
            this.loadings.pageLoading = false
          }).catch(() => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
          this.loadings.pageLoading = false
        })
      },
      provider() {
        return new Promise((resolve, reject) => {
          providerAccounts().then(res => {
            this.providerActs = {}
            res.data.map(d => {
              this.providerActs[d.id] = d.account_name
            })
            resolve(res.data)
          }).catch(() => {
            reject()
          })
        })
      },
      add() {
        this.$refs.accountCreate.initCreate()
      },
      refreshAccountNumbers() {
        this.loadings.pageLoading = true
        refreshNumbers().then(() => {
          this.loadings.pageLoading = false
          this.$message.success("同步完成")
        }).catch(() => {
          this.loadings.pageLoading = false
        })
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
