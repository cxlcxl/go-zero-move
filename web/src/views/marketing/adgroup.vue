<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.adgroup_id" class="w120" clearable placeholder="任务 ID"/>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="adgroup.adgroups" highlight-current-row stripe border size="mini">
        <el-table-column prop="adgroup_id" label="任务 ID" width="120" align="center" fixed="left"></el-table-column>
        <el-table-column prop="campaign_name" label="任务名称" min-width="160" show-overflow-tooltip fixed="left"/>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center">
          <template slot-scope="scope">{{scope.row.created_at|timeFormat}}</template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="80">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="adgroup.total" @current-change="handlePage" @size-change="handlePageSize"/>
    </el-col>
  </el-row>
</template>

<script>
import { parseTime } from '@/utils'
import Page from '@c/Page'

export default {
  //name: 'AdgroupList',
  components: {Page},
  data() {
    return {
      loadings: {
        pageLoading: false
      },
      adgroup: {
        total: 0,
        adgroups: []
      },
      search: {
        adgroup_id: '',
        page: 1,
        page_size: 10
      },
      resources: {}
    }
  },
  filters: {
    timeFormat(t) {
      return parseTime(t)
    }
  },
  mounted() {
    this.getAdgroupList()
  },
  methods: {
    getAdgroupList() {
      //this.loadings.pageLoading = true
      Promise.all([
        this.getResources()
      ]).then(() => {
        /* adgroupList(this.search).then(res => {
          this.loadings.pageLoading = false
          this.campaign = res.data
        }).catch(err => {
          this.loadings.pageLoading = false
        }) */
      }).catch(() => {
        this.loadings.pageLoading = false
      })
    },
    getResources() {
      return new Promise((resolve, reject) => {
        resolve()
        /* if (Object.keys(this.resources.campaign_type).length > 0) {
          resolve()
          return
        }
        adgroupResources().then(res => {
          this.resources = res.data
          resolve()
        }).catch(() => {
          reject()
        }) */
      })
    },
    doSearch() {
      this.search.page = 1
      this.getAdgroupList()
    },
    editRow() {

    },
    handlePage(p) {
      this.search.page = p
      this.getAdgroupList()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.getAdgroupList()
    }
  }
}
</script>

<style scoped>

</style>
