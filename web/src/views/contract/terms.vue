<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search._name" class="w180" clearable placeholder="条款名称"/>
        </el-form-item>
        <el-form-item>
          <el-input v-model="search._key" class="w180" clearable placeholder="条款代码"/>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" placeholder="条款状态" class="w110">
            <el-option :key="1" label="正常" :value="1"/>
            <el-option :key="0" label="停用" :value="0"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="getTermsList">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24" v-permission="'terms/create'">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加合同条款</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="termsList" highlight-current-row stripe border size="mini"
                style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="90" align="center"/>
        <el-table-column prop="_name" label="条款名称" width="220"/>
        <el-table-column prop="_key" label="条款代码" width="180"/>
        <el-table-column label="状态" width="90" align="center">
          <template slot-scope="scope">
            {{ scope.row.state === 1 ? '正常' : '停用' }}
          </template>
        </el-table-column>
        <el-table-column prop="desc" label="描述"/>
        <el-table-column prop="created_at" label="创建时间" width="160"/>
        <el-table-column align="center" label="操作" fixed="right" width="130">
          <template slot-scope="scope">
            <el-button-group class="table-operate" v-permission="'terms/update'">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <terms-create ref="termsCreate" @success="getTermsList"/>
      <terms-update ref="termsUpdate" @success="getTermsList"/>
    </el-col>
  </el-row>
</template>

<script>
  import {termsList, termsClose} from '@a/contract-terms'
  import TermsCreate from './components/add-terms'
  import TermsUpdate from './components/edit-terms'

  export default {
    name: 'ContractTerms',
    components: {
      TermsCreate,
      TermsUpdate
    },
    data() {
      return {
        loadings: {
          pageLoading: false
        },
        termsList: [],
        search: {_key: '', _name: '', state: 1}
      }
    },
    computed: {},
    mounted() {
      this.getTermsList()
    },
    methods: {
      getTermsList() {
        this.loadings.pageLoading = true

        termsList(this.search).then(response => {
          this.termsList = response.data
          this.loadings.pageLoading = false
        }).catch(() => {
          this.loadings.pageLoading = false
        })
      },
      add() {
        this.$refs.termsCreate.visible = true
      },
      editRow(row) {
        this.$refs.termsUpdate.termsForm = {
          id: row.id,
          _name: row._name,
          _key: row._key,
          state: row.state,
          desc: row.desc,
        }
        this.$refs.termsUpdate.visible = true
      },
      closeTerms(row) {
        this.$confirm('确定停用此条款么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.loadings.pageLoading = true
          termsClose(row.id).then(() => {
            this.$message.success('停用成功')
            this.getTermsList()
          }).catch(err => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
        })
      }
    }
  }
</script>
