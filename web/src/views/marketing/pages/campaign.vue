<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.campaign_id" class="w120" clearable placeholder="计划 ID"/>
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.campaign_name" class="w230" clearable placeholder="计划名称"/>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24" v-show="search.app_id !== ''">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">创建计划</el-button>
    </el-col>
    <el-col :span="24" style="margin-top: 10px;">
      <el-table v-loading="campaignLoading" :data="campaign.campaigns" highlight-current-row stripe border size="mini">
        <el-table-column label="计划 ID" width="90" align="center" fixed="left">
          <template slot-scope="scope">
            <span :class="{'text-success': scope.row.is_callback === 1}">{{ scope.row.campaign_id }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="campaign_name" label="计划名称" min-width="160" show-overflow-tooltip fixed="left"/>
        <el-table-column label="计划类型" width="90" align="center">
          <template slot-scope="scope">{{ resources.campaign_type[scope.row.campaign_type] }}</template>
        </el-table-column>
        <el-table-column label="计划状态" width="140">
          <template slot-scope="scope">{{ resources.show_status[scope.row.show_status] }}</template>
        </el-table-column>
        <el-table-column label="操作状态" width="90" align="center">
          <template slot-scope="scope">{{ resources.opt_status[scope.row.opt_status] }}</template>
        </el-table-column>
        <el-table-column label="日预算状态" width="130">
          <template slot-scope="scope">{{ resources.campaign_daily[scope.row.campaign_daily_budget_status] }}
          </template>
        </el-table-column>
        <el-table-column label="账户余额状态" width="100">
          <template slot-scope="scope">{{ resources.balance_status[scope.row.user_balance_status] }}</template>
        </el-table-column>
        <el-table-column label="推广产品" width="90">
          <template slot-scope="scope">{{ resources.product_type[scope.row.product_type] }}</template>
        </el-table-column>
        <el-table-column prop="today_daily_budget" label="当日限额" width="80" align="right"/>
        <el-table-column prop="tomorrow_daily_budget" label="次日限额" width="80" align="right"/>
        <el-table-column label="投放网络" width="110">
          <template slot-scope="scope">{{ resources.flow_resource[scope.row.flow_resource] }}</template>
        </el-table-column>
        <el-table-column prop="sync_flow_resource_searchad" label="同步网络" width="100"/>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center">
          <template slot-scope="scope">{{ scope.row.created_at|timeFormat }}</template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="120">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
              <el-button type="primary" plain @click.native.prevent="addAdgroup(scope.row)">建任务</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>

    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="campaign.total" @current-change="handlePage" @size-change="handlePageSize"/>

      <create-campaign ref="create" @success="fetchData" :product-type="resources.product_type" :app-id="search.app_id"
                       :campaign-type="resources.campaign_type" :sync-flow="resources.sync_flow"/>
    </el-col>
  </el-row>
</template>

<script>
import { campaignList, campaignResources } from '@/api/campaign'
import Page from '@c/Page'
import CreateCampaign from '../components/add-campaign'
import { parseTime } from '@/utils'

export default {
  components: { Page, CreateCampaign },
  data() {
    return {
      campaignLoading: false,
      search: {
        app_id: '',
        campaign_id: '',
        campaign_name: '',
        page: 1,
        page_size: 10
      },
      campaign: {
        total: 0,
        campaigns: []
      },
      resources: {
        campaign_type: {},
        daily_budget_opts: {},
        show_status: {},
        opt_status: {},
        product_type: {},
        sync_flow: {},
        campaign_daily: {},
        balance_status: {},
        flow_resource: {}
      }
    }
  },
  filters: {
    timeFormat(timestamp) {
      return parseTime(timestamp)
    }
  },
  methods: {
    initData(app_id) {
      this.$set(this.search, 'app_id', app_id)
      this.fetchData()
    },
    fetchData() {
      this.campaignLoading = true
      Promise.all([
        this.getResources()
      ]).then(() => {
        campaignList(this.search).then(res => {
          this.campaignLoading = false
          this.campaign = res.data
        }).catch(err => {
          this.campaignLoading = false
        })
      }).catch(() => {
        this.campaignLoading = false
      })
    },
    getResources() {
      return new Promise((resolve, reject) => {
        if (Object.keys(this.resources.campaign_type).length > 0) {
          resolve()
          return
        }
        campaignResources().then(res => {
          this.resources = res.data
          resolve()
        }).catch(() => {
          reject()
        })
      })
    },
    doSearch() {
      this.search.page = 1
      this.fetchData()
    },
    add() {
      if (this.search.app_id === '') {
        this.$message.error("请先选择应用")
        return
      }
      this.$refs.create.initCreate()
    },
    editRow() {

    },
    addAdgroup(campaign) {
      this.$router.push({ name: 'AdgroupCreate', query: { campaign_id: campaign.campaign_id } })
    },
    handlePage(p) {
      this.search.page = p
      this.fetchData()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.fetchData()
    }
  }
}
</script>

<style scoped>

</style>
