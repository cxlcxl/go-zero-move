<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.customer_name" class="w200" clearable placeholder="客户名称"/>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.customer_type" placeholder="客户类型" clearable>
            <el-option label="请选择客户类型" key="" value=""/>
            <el-option v-for="(val,key) in customerList.customer_types" :label="val" :key="key" :value="key"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--<el-col :span="24" v-permission="'customer/create'">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加客户</el-button>
    </el-col>-->
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="customerList.list" highlight-current-row stripe border
                size="mini">
        <el-table-column prop="id" label="ID" width="80" align="center"/>
        <el-table-column prop="customer_name" label="客户名称" show-overflow-tooltip/>
        <el-table-column prop="full_name" label="客户全称" show-overflow-tooltip/>
        <el-table-column label="客户类型" width="100" align="center">
          <template slot-scope="scope">
            {{ customerList.customer_types[scope.row.customer_type] }}
          </template>
        </el-table-column>
        <el-table-column prop="weight" label="权重" width="70" align="center"/>
        <el-table-column prop="created_at" label="创建时间" width="160"/>
        <el-table-column align="center" label="操作" width="140">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)" v-permission="'customer/update'">编辑</el-button>
              <el-button type="primary" plain @click.native.prevent="accountInfo(scope.row)">账号</el-button>
              <el-button type="primary" plain @click.native.prevent="contracts(scope.row)">合同</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" @current-change="handlePage" @size-change="handlePageSize" :total="customerList.total" :limit="search.limit"/>

      <customer-edit ref="customerUpdate" @success="getCustomerList" :customer-types="customerList.customer_types"/>
      <!--<customer-add ref="customerCreate" @success="getCustomerList" :customer-types="customerList.customer_types"/>-->
      <customer-contract ref="contract"/>
      <customer-account ref="account"/>
    </el-col>
  </el-row>
</template>

<script>
  import {customerList} from '@a/customer'
  import Page from '@c/Page'
  import CustomerEdit from './components/edit-customer'
  import CustomerAdd from './components/add-customer'
  import CustomerContract from './components/customer-contract'
  import CustomerAccount from './components/customer-account'

  export default {
    name: 'CustomerList',
    components: {
      CustomerEdit, /*CustomerAdd,*/ CustomerContract, CustomerAccount, Page
    },
    data() {
      return {
        visible: false,
        loadings: {
          pageLoading: false
        },
        customerList: {
          total: 0,
          list: [],
          customer_types: {},
        },
        search: {customer_name: '', customer_type: '', page: 1, limit: 10},
      }
    },
    mounted() {
      this.getCustomerList()
    },
    methods: {
      getCustomerList() {
        this.loadings.pageLoading = true

        customerList(this.search).then(response => {
          this.customerList = response.data
          this.loadings.pageLoading = false
        }).catch(() => {
          this.loadings.pageLoading = false
        })
      },
      add() {
        this.$refs.customerCreate.visible = true
      },
      editRow(row) {
        this.$refs.customerUpdate.initUpdate(row)
      },
      contracts(row) {
        this.$refs.contract.initView(row)
      },
      accountInfo(row) {
        this.$refs.account.initView(row)
      },
      uploadAgreement(row) {
        this.$refs._upload.initPanel(row.id)
      },
      showAgreements(agreements) {
        if (agreements === null || !agreements) {
          this.$message.info('此合同未上传协议或附件')
        } else {
          this.$refs.agreement.initAgreements(agreements)
        }
      },
      doSearch() {
        this.search.page = 1
        this.getCustomerList()
      },
      handlePage(p) {
        this.search.page = p
        this.getCustomerList()
      },
      handlePageSize(s) {
        this.search.limit = s
        this.getCustomerList()
      }
    }
  }
</script>
