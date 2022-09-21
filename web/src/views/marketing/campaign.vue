<template>
  <div class="campaign-container">
    <div class="app-list">
      <app-list ref="appList" :app-id="app_id" @handle-change="selectApp"/>
    </div>
    <div class="campaign-list">
      <el-tabs type="border-card" v-model="tabPanel" @tab-click="handleClick">
        <el-tab-pane label="计划列表" name="campaign" key="campaign"/>
        <el-tab-pane label="任务列表" name="adgroup" key="adgroup"/>
        <el-tab-pane label="素材库" name="resource" key="resource"/>

        <campaign-panel ref="campaign" v-show="tabPanel === 'campaign'"/>
        <resource-panel ref="resource" v-show="tabPanel === 'resource'"/>
        <adgroup-panel ref="adgroup" v-show="tabPanel === 'adgroup'"/>
      </el-tabs>
    </div>
  </div>
</template>

<script>
import AppList from './pages/app-list'
import CampaignPanel from './pages/campaign'
import ResourcePanel from './pages/resource'
import AdgroupPanel from './pages/adgroup'

export default {
  //name: 'CampaignList',
  components: { AppList, CampaignPanel, ResourcePanel, AdgroupPanel },
  data() {
    return {
      app_id: '',
      loadings: {
        pageLoading: false
      },
      tabPanel: 'campaign',
    }
  },
  mounted() {
    this.$refs[this.tabPanel].initData(this.app_id)
  },
  methods: {
    selectApp(app_id) {
      this.app_id = app_id
      this.$refs[this.tabPanel].initData(app_id)
    },
    handleClick(tab, event) {
      this.$refs[tab.name].initData(this.app_id)
    }
  }
}
</script>

<style scoped lang="scss">
.campaign-container {
  display: flex;

  .app-list {
    border: 1px solid #DCDFE6;
  }

  .campaign-list {
    width: 0;
    flex: 1;
    padding-left: 10px;
  }
}
</style>
