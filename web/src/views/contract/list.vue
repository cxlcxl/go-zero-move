<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.contract_sn" class="w200" clearable placeholder="合同编号"/>
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.contract_name" class="w200" clearable placeholder="合同名称"/>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24" v-permission="'contract/create'">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加合同</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="contractList.list" highlight-current-row stripe border
                size="mini"
                style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="80" align="center"/>
        <el-table-column prop="contract_sn" label="合同编号" width="180" show-overflow-tooltip/>
        <el-table-column prop="contract_name" label="合同名称" width="240" show-overflow-tooltip>
          <template slot-scope="scope">
            <el-link @click="showContractDetail(scope.row.id)" type="primary">{{scope.row.contract_name}}</el-link>
          </template>
        </el-table-column>
        <el-table-column label="有效期" width="180" align="center">
          <template slot-scope="scope">
            {{ scope.row.start_date + ' / ' + scope.row.end_date }}
          </template>
        </el-table-column>
        <el-table-column prop="currency" label="币种" width="70" align="center"/>
        <el-table-column prop="desc" label="描述" show-overflow-tooltip/>
        <el-table-column prop="created_at" label="创建时间" width="140"/>
        <el-table-column label="协议/附件" width="80" align="center">
          <template slot-scope="scope">
            <el-link @click="showAgreements(scope.row.agreements)" type="primary" size="mini">
              {{scope.row.agreements | agreementLengthFilter}}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="170">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)" v-permission="'contract/update'">编辑</el-button>
              <el-button type="primary" plain @click.native.prevent="uploadAgreement(scope.row)" v-permission="'agreement/create'">上传协议</el-button>
              <el-button type="primary" plain @click.native.prevent="viewContract(scope.row)" v-permission="'file'">预览</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="contractList.total" @current-change="handlePage" @size-change="handlePageSize"/>

      <contract-agreement ref="agreement" @success="getContractList"/>
      <upload-agreement ref="_upload" @success="getContractList"/>

      <pdf ref="pdf"/>
    </el-col>
  </el-row>
</template>

<script>
  import {contractList} from '@a/contract'
  import {termsList} from '@a/contract-terms'
  import {getConfigs} from '@a/global'
  import ContractAgreement from './components/contract-agreement'
  import UploadAgreement from './components/upload-agreement'
  import Pdf from '@c/Pdf'
  import Page from '@c/Page'

  export default {
    name: 'ContractList',
    components: {
      ContractAgreement,
      UploadAgreement,
      Pdf, Page
    },
    data() {
      return {
        visible: false,
        loadings: {
          pageLoading: false
        },
        contractList: {
          total: 0,
          list: [],
          calculate_types: {}
        },
        termsList: [],
        currencies: [],
        search: {contract_sn: '', contract_name: '', /*customer_name: '', */page: 1, limit: 10},
      }
    },
    mounted() {
      this.getContractList()
    },
    filters: {
      agreementLengthFilter(agreements) {
        if (Array.isArray(agreements)) {
          return agreements.length
        } else {
          return 0
        }
      }
    },
    methods: {
      getContractList() {
        this.loadings.pageLoading = true
        Promise.all([
          this.getTermsList(),
          this.getCurrencies()
        ]).then(res => {
          contractList(this.search).then(response => {
            this.contractList = response.data
            this.loadings.pageLoading = false
          }).catch(() => {
            this.loadings.pageLoading = false
          })
        }).catch(err => {
          this.loadings.pageLoading = false
        })
      },
      getTermsList() {
        return new Promise((resolve, reject) => {
          if (this.termsList.length > 0) {
            resolve()
          } else {
            termsList({state: 1}).then(res => {
              this.termsList = res.data
              resolve()
            }).catch(() => {
              reject()
            })
          }
        })
      },
      getCurrencies() {
        return new Promise((resolve, reject) => {
          if (this.currencies.length > 0) {
            resolve()
          } else {
            getConfigs('currency').then(res => {
              this.currencies = res.data
              resolve()
            }).catch(() => {
              reject()
            })
          }
        })
      },
      add() {
        this.$router.push({name: 'ContractCreate'})
      },
      editRow(row) {
        this.$router.push({name: 'ContractUpdate', query: {id: row.id}})
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
      viewContract(row) {
        if (row.file_id === 0) {
          return this.$message.info("未上传附件")
        }
        this.$refs.pdf.previewPDF(row.file_id)
      },
      doSearch() {
        this.search.page = 1
        this.getContractList()
      },
      handlePage(p) {
        this.search.page = p
        this.getContractList()
      },
      handlePageSize(p) {
        this.search.limit = p
        this.getContractList()
      },
      showContractDetail(id) {
        this.$router.push({name: 'ContractDetail', query: {id}})
      }
    }
  }
</script>
