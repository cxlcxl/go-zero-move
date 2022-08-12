<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.app_name" class="w150" clearable placeholder="应用名"/>
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.app_id" class="w120" clearable placeholder="APP ID"/>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.app_type" class="w130">
            <el-option label="全部应用类型" :value="0"/>
            <el-option v-for="(key, val) in appList.app_type" :label="key" :value="Number(val)"/>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.channel" class="w100">
            <el-option label="全部渠道" :value="0"/>
            <el-option v-for="(key, val) in appList.app_channel" :label="key" :value="Number(val)"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加应用</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loading" :data="appList.data" highlight-current-row stripe border size="mini" style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="80" align="center"/>
        <el-table-column prop="app_name" label="应用名称" min-width="120" show-overflow-tooltip/>
        <el-table-column prop="app_id" label="AppID" width="140"/>
        <el-table-column prop="account_name" label="关联账户" min-width="150" show-overflow-tooltip/>
        <el-table-column label="渠道" width="120" align="center">
          <template slot-scope="scope">{{appList.app_channel[scope.row.channel]}}</template>
        </el-table-column>
        <el-table-column label="应用包名" prop="pkg_name" min-width="150" show-overflow-tooltip/>
        <el-table-column label="应用类型" width="130">
          <template slot-scope="scope">{{ appList.app_type[scope.row.app_type] }}</template>
        </el-table-column>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center">
          <template slot-scope="scope">{{scope.row.created_at|timeFormat}}</template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="130">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row.id)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="appList.total" @current-change="handlePage" @size-change="handlePageSize"/>
    </el-col>

    <create-application ref="appCreate" :app-type="appList.app_type" :app-channel="appList.app_channel" @success="getApplicationList"/>
    <update-application ref="appUpdate" :app-type="appList.app_type" :app-channel="appList.app_channel" @success="getApplicationList"/>
  </el-row>
</template>

<script>
import { appList } from '@a/app'
import { parseTime } from '@/utils'
import Page from '@c/Page'
import CreateApplication from './components/add-app'
import UpdateApplication from './components/edit-app'

export default {
  name: 'AppList',
  components: {Page, CreateApplication, UpdateApplication},
  data() {
    return {
      loading: false,
      search: {
        app_name: '',
        app_id: '',
        app_type: 0,
        channel: 0,
        page: 1,
        page_size: 10
      },
      appList: {
        total: 0,
        app_type: {},
        app_channel: {},
        data: []
      }
    }
  },
  mounted() {
    this.getApplicationList()
  },
  filters: {
    timeFormat(timestamp) {
      return parseTime(timestamp)
    }
  },
  methods: {
    add() {
      this.$refs.appCreate.initCreate()
    },
    doSearch() {
      this.search.page = 1
      this.getApplicationList()
    },
    getApplicationList() {
      this.loading = true
      appList(this.search).then(res => {
        this.appList.data = res.data
        this.appList.total = res.total
        this.appList.app_type = res.app_type
        this.appList.app_channel = res.app_channel
        this.loading = false
      }).catch(err => {
        this.loading = false
      })
    },
    handlePage(p) {
      this.search.page = p
      this.getApplicationList()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.getApplicationList()
    },
    editRow(id) {
      this.$refs.appUpdate.initUpdate(id)
    },
  }
}
</script>

<style scoped>

</style>
