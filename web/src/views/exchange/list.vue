<template>
  <el-row>
    <el-col :span="24" style="margin-bottom: 15px;" v-permission="'exchange/create'">
      <el-button icon="el-icon-refresh" size="mini" @click="getExchangeList">刷新列表</el-button>
      <el-button icon="el-icon-plus" size="mini" type="primary" @click="addExchange">录入汇率</el-button>
    </el-col>

    <el-col :span="24">
      <el-table :data="list" v-loading="pageLoading" border size="mini" highlight-current-row stripe>
        <el-table-column prop="date_month" label="月份" width="100" align="center"/>
        <el-table-column prop="rate" label="汇率" width="120" align="right"/>
        <el-table-column label="币种转换">
          <template slot-scope="scope">{{scope.row.from_currency}} 转 {{scope.row.to_currency}}</template>
        </el-table-column>
        <el-table-column prop="operator" label="操作人" width="160"/>
        <el-table-column prop="created_at" label="添加时间" width="160" align="center"/>
        <el-table-column label="操作" width="100" align="center">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="updateRow(scope.row)" v-permission="'exchange/update'">修改</el-button>
              <el-button type="danger" plain @click.native.prevent="destroyRow(scope.row)" v-permission="'exchange/destroy'">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>

    <exchange-add ref="create" :currencies="currencies" @success="getExchangeList"/>
    <exchange-edit ref="update" :currencies="currencies" @success="getExchangeList"/>
  </el-row>
</template>

<script>
  import {exchangeList, exchangeDestroy} from '@a/exchange'
  import ExchangeAdd from './add'
  import ExchangeEdit from './edit'
  import {getConfigs} from '@a/global'

  export default {
    name: 'Exchange',
    components: {
      ExchangeAdd, ExchangeEdit
    },
    data() {
      return {
        pageLoading: false,
        list: [],
        currencies: [],
      }
    },
    mounted() {
      this.getExchangeList()
    },
    methods: {
      getExchangeList() {
        this.pageLoading = true
        Promise.all([
          this.getCurrencies()
        ]).then(res => {
          exchangeList().then(response => {
            this.list = response.data
            this.pageLoading = false
          }).catch(() => {
            this.pageLoading = false
          })
        }).catch(err => {
          this.pageLoading = false
        })
      },
      updateRow(row) {
        this.$refs.update.initPanel(row)
      },
      destroyRow(row) {
        this.$confirm('确认删除此汇率么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.pageLoading = true
          exchangeDestroy(row.id).then(() => {
            this.$message.success('删除成功')
            this.getExchangeList()
          }).catch(err => {
            this.pageLoading = false
          })
        }).catch(() => {
        })
      },
      addExchange() {
        this.$refs.create.visible = true
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
      }
    }
  }
</script>
