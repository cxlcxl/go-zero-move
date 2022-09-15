<template>
  <el-row :gutter="15" v-loading="loadings.pageLoading">
    <el-form :model="adgroupForm" ref="adgroupForm" label-width="120px" size="mini" :rules="adgroupRules">
      <el-col :span="24">
        <el-form-item label="选择计划" prop="campaign_id">
          <el-select v-model="adgroupForm.campaign_id" remote filterable placeholder="可输入名称查询"
                     :remote-method="remoteCampaign" :loading="loadings.remoteCampaignLoading" style="width: 100%;">
            <el-option v-for="item in campaigns" :label="item.campaign_name" :value="Number(item.campaign_id)"/>
          </el-select>
        </el-form-item>
        <el-form-item label="选择应用" prop="campaign_id">
          <el-select v-model="adgroupForm.campaign_id" remote filterable placeholder="可输入名称查询"
                     :remote-method="remoteApplication" :loading="loadings.remoteAppLoading" style="width: 100%;">
            <el-option v-for="item in campaigns" :label="item.campaign_name" :value="Number(item.campaign_id)"/>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="24" style="border-top: 1px solid #e3e3e3;">
        <p style="margin: 10px 0 8px 0;">
          <el-button size="mini" @click="addTab(editableTabsValue)" icon="el-icon-plus" type="primary">添加任务</el-button>
        </p>

        <el-tabs v-model="editableTabsValue" type="card" closable @tab-remove="removeTab">
          <el-tab-pane v-for="(__adgroup, idx) in adgroupForm.adgroups"
                       :label="__adgroup.adgroup_name" :name="__adgroup.tab_name" :key="__adgroup.tab_name">
            <el-form-item label="定向" prop="targeting_package_id">
              <el-select v-model="adgroupForm.targeting_package_id" placeholder="选择已有定向" style="width: 380px;">
                <el-option v-for="targeting in targetingList" :label="targeting.targeting_name" :value="targeting.targeting_id"/>
              </el-select>
              <el-button type="primary" style="margin-left: 5px;" @click="createTargeting(__adgroup.tab_name)" icon="el-icon-edit"/>
              <el-button type="primary" style="margin-left: 5px;" @click="createTargeting(__adgroup.tab_name)" icon="el-icon-plus"/>
              <el-button style="margin-left: 5px;" @click="pullTargeting" icon="el-icon-refresh"/>
              <el-button icon="el-icon-document-copy" type="primary" @click="copyTab(__adgroup.tab_name, idx)" style="float: right;">拷贝此任务</el-button>
            </el-form-item>
            <el-form-item label="版位" prop="creative_size_id">
              <el-radio-group v-model="__adgroup.category" @change="remoteCreative">
                <el-radio-button label="CREATIVE_SIZE_CATEGORY_THIRD_PARTY">三方媒体资源</el-radio-button>
                <el-radio-button label="CREATIVE_SIZE_CATEGORY_SELF_OWNED">自有媒体资源</el-radio-button>
                <el-radio-button label="CREATIVE_SIZE_CATEGORY_OTHER">其他首选资</el-radio-button>
              </el-radio-group>
              <el-row :gutter="10" v-loading="loadings.creativeLoading">
                <el-col :span="17">
                  <div class="creative-container">
                    <div v-for="creative in creativeList"
                         :class="{'creative-list': true, 'creative-active': creative.creative_size_id === __adgroup.creative_size_id}"
                         @click="handleSelectCreative(creative.creative_size_id)">
                      <span class="creative-name">{{creative.name}}</span>
                      <span class="creative-desc">{{creative.desc}}</span>
                    </div>
                  </div>
                </el-col>
                <el-col :span="7">
                  <div class="creative-banner"></div>
                </el-col>
              </el-row>
            </el-form-item>
            <el-form-item label="投放日期" prop="adgroup_begin_date">
              <el-date-picker v-model="__adgroup.adgroup_begin_date" type="date" placeholder="选择开始日期"/> -
              <el-date-picker v-model="__adgroup.adgroup_end_date" type="date" placeholder="结束日期可不选"/>
            </el-form-item>
            <el-form-item label="出价" prop="price">
              <el-radio-group v-model="__adgroup.price_type">
                <el-radio-button label="CREATIVE_SIZE_CATEGORY_THIRD_PARTY">CPC</el-radio-button>
                <el-radio-button label="CREATIVE_SIZE_CATEGORY_SELF_OWNED">CPM</el-radio-button>
              </el-radio-group>
              <p style="margin: 5px 0 0 0;">
                <el-input-number class="w180" v-model="__adgroup.price"/>
              </p>
            </el-form-item>
            <el-form-item label="任务名称" prop="adgroup_name">
              <el-input v-model="__adgroup.adgroup_name" placeholder="限 100 字内的中英文、数字以及破折号 [ - ]"/>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>

        <el-form-item>
          <el-button type="primary" icon="el-icon-check" @click="confirmAdgroup">提交</el-button>
          <el-button icon="el-icon-close">取消创建</el-button>
        </el-form-item>
      </el-col>
    </el-form>

    <targeting-create ref="targeting_create" @handle-success="targetingCallback"/>
  </el-row>
</template>

<script>
import TargetingCreate from './targeting-create'

export default {
  // name: 'AdgroupCreate',
  components: {TargetingCreate},
  data() {
    return {
      loadings: {
        pageLoading: false,
        remoteAppLoading: false,
        remoteCampaignLoading: false,
        creativeLoading: false,
      },
      targetingList: [],
      campaigns: [],
      creativeList: [
        {creative_size_id: '111', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 1'},
        {creative_size_id: '222', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 2'},
        {creative_size_id: '333', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 3'},
        {creative_size_id: '444', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 4'},
        {creative_size_id: '555', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 5'},
        {creative_size_id: '666', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 6'},
        {creative_size_id: '777', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 7'},
        {creative_size_id: '888', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 8'},
        {creative_size_id: '999', name: 'Global High-quality traffic Banner_VIP', desc: 'Global High-quality traffic Banner_VIP 9'},
      ],
      editableTabsValue: 'adgroup-1',
      adgroupForm: {
        campaign_id: '',
        adgroups: [
          {
            tab_name: 'adgroup-1',
            adgroup_name: 'adgroup-1',
            category: '', // 版位类型
            targeting_package_id: '', // 定向 ID
            creative_size_id: '', // 版位 ID
            adgroup_begin_date: '',
            adgroup_end_date: '',
            price_type: '',
            price: 0,
          }
        ]
      },
      adgroupRules: {

      },
      tabIndex: 1,
      resources: {

      }
    }
  },
  mounted() {
    this.getResources()
  },
  methods: {
    getResources() {
      this.loadings.pageLoading = true
      setTimeout(() => {
        this.loadings.pageLoading = false
      }, 1000)
    },
    createTargeting(tab) {
      this.$refs.targeting_create.initCreate(tab, "69856985")
    },
    pullTargeting() {
      this.$message.info("定向数据拉取中，请稍后...")
    },
    targetingCallback(targeting_id) {
      console.log(targeting_id)
    },
    remoteCampaign() {

    },
    remoteCreative(q) {
      this.loadings.creativeLoading = true
      setTimeout(() => {
        this.loadings.creativeLoading = false
      }, 1000)
    },
    handleSelectCreative(creativeSizeId) {
      this.adgroupForm.adgroups.forEach((tab, index) => {
        if (tab.tab_name === this.editableTabsValue) {
          this.$set(this.adgroupForm.adgroups[index], 'creative_size_id', creativeSizeId)
        }
      })
    },
    remoteApplication() {
      if (query.trim() !== '') {
        this.remoteLoading = true;
        // searchAccounts(query).then(res => {
        //   this.remoteLoading = false
        //   if (Array.isArray(res.data)) {
        //     this.accounts = res.data
        //   } else {
        //     this.accounts = []
        //   }
        // }).catch(() => {
        //   this.remoteLoading = false
        // })
      } else {
        this.campaigns = [];
      }
    },
    addTab(targetName) {
      this.tabIndex++
      let adgroupName = 'adgroup-' + this.tabIndex
      this.adgroupForm.adgroups.push({
        tab_name: adgroupName,
        adgroup_name: adgroupName,
        category: '',
        targeting_package_id: '',
        creative_size_id: '',
        adgroup_begin_date: '',
        adgroup_end_date: '',
        price_type: '',
        price: 0,
      })
      this.editableTabsValue = adgroupName;
    },
    copyTab(targetName, idx) {
      this.tabIndex++
      let adgroupName = 'adgroup-' + this.tabIndex
      let copyAdgroup = Object.assign({}, this.adgroupForm.adgroups[idx])
      copyAdgroup.tab_name = adgroupName
      copyAdgroup.adgroup_name = adgroupName
      this.adgroupForm.adgroups.push(copyAdgroup)
      this.editableTabsValue = adgroupName
    },
    removeTab(targetName) {
      if (this.adgroupForm.adgroups.length === 1) {
        this.$message.error("不能删除最后一个任务")
        return
      }
      this.$confirm('已填写的信息会丢失, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }).then(() => {
        let tabs = this.adgroupForm.adgroups;
        let activeName = this.editableTabsValue;
        if (activeName === targetName) {
          tabs.forEach((tab, index) => {
            if (tab.tab_name === targetName) {
              let nextTab = index === 0 ? tabs[index + 1] : tabs[index - 1];
              if (nextTab) {
                activeName = nextTab.tab_name;
              }
            }
          });
        }
        this.editableTabsValue = activeName;
        this.adgroupForm.adgroups = tabs.filter(tab => tab.tab_name !== targetName);
      }).catch(() => {})
    },
    confirmAdgroup() {
      this.$refs.adgroupForm.validate(v => {
        if (v) {

        } else {
          return false
        }
      })
    }
  }
}
</script>

<style scoped lang="scss">
.creative-container {
  height: 230px;
  margin-top: 5px;
  border: 1px solid #e3e3e3;
  border-radius: 3px;
  overflow-y: scroll;

  .creative-active {
    color: #1890ff;
    font-weight: 600;
  }

  .creative-list {
    cursor: pointer;
    display: flex;
    line-height: 20px;
    border-bottom: 1px solid #e3e3e3;

    &:hover {
      background: #f5f5f5;
    }
  }

  .creative-name,.creative-desc {
    display: inline-block;
    padding: 8px;
    font-size: 12px;
  }
  .creative-name {
    width: 300px;
    border-right: 1px solid #e3e3e3;
  }
  .creative-desc {
    flex: 1;
  }

  &::-webkit-scrollbar {
    /*滚动条整体样式*/
    width: 6px; /*高宽分别对应横竖滚动条的尺寸*/
    height: 0;
  }

  &::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    border-radius: 10px;
    background-color: rgb(131, 138, 150);
  }

  &::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    box-shadow: inset 0 0 5px rgba(100, 98, 98, 0.2);
    background: #ededed;
    border-radius: 2px;
  }
}
</style>
