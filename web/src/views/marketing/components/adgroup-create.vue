<template>
  <el-row :gutter="15" v-loading="loadings.pageLoading">
    <el-form :model="adgroupForm" ref="adgroupForm" label-width="120px" size="mini" :rules="adgroupRules">
      <el-col :span="12">
        <el-form-item label="所属计划" prop="campaign_id" :rules="{required: true, message: '计划必需选择'}">
          <el-input :value="campaign.campaign_name" disabled/>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="推广产品" prop="product_id" :rules="{required: true, message: '计划必需选择'}">
          <el-input :value="campaign.app.app_name" disabled/>
        </el-form-item>
      </el-col>
      <el-col :span="12">

      </el-col>
      <el-col :span="24" style="border-top: 1px solid #e3e3e3;">
        <p style="margin: 10px 0 8px 0;">
          <el-button size="mini" @click="addTab(editableTabsValue)" icon="el-icon-plus" type="primary">添加任务</el-button>
        </p>

        <el-tabs v-model="editableTabsValue" type="card" closable @tab-remove="removeTab">
          <el-tab-pane v-for="(__adgroup, idx) in adgroupForm.adgroups"
                       :label="__adgroup.adgroup_name" :name="__adgroup.tab_name" :key="__adgroup.tab_name">
            <el-form-item label="定向" prop="targeting_package_id"
                          :prop="'adgroups.' + idx + '.targeting_package_id'" :rules="{required: true, message: '请选择定向'}">
              <el-select v-model="__adgroup.targeting_package_id" placeholder="选择已有定向" style="width: 380px;">
                <el-option v-for="targeting in targetings" :label="targeting.targeting_name" :value="targeting.targeting_id"/>
              </el-select>
              <el-button type="primary" style="margin-left: 5px;" @click="updateTargeting(__adgroup.tab_name, idx)" icon="el-icon-edit">修改</el-button>
              <el-button type="primary" style="margin-left: 5px;" @click="createTargeting(__adgroup.tab_name, idx)" icon="el-icon-plus">新建</el-button>
              <el-button icon="el-icon-document-copy" type="primary" @click="copyTab(__adgroup.tab_name, idx)" style="float: right;">拷贝此任务</el-button>
            </el-form-item>

            <el-form-item label="版位" prop="creative_size_id"
                          :prop="'adgroups.' + idx + '.creative_size_id'" :rules="{required: true, message: '请选择版位'}">
              <el-radio-group v-model="__adgroup.category" @change="getCreativeList($event, __adgroup.tab_name)">
                <el-radio-button v-for="(c_label, c_k) in resources.creative_category" :label="c_k">{{ c_label }}</el-radio-button>
              </el-radio-group>
              <el-row :gutter="15" v-loading="loadings.creativeLoading" v-show="creativeList[__adgroup.tab_name].length > 0">
                <el-col :span="17">
                  <div class="creative-container">
                    <div v-for="creative in creativeList[__adgroup.tab_name]"
                         :class="{'creative-list': true, 'creative-active': creative.creative_size_id === __adgroup.creative_size_id}"
                         @click="handleSelectCreative(creative.creative_size_id, creative.samples, creative.support_price_type)">
                      <span class="creative-name">{{creative.creative_size_name_dsp}}</span>
                      <span class="creative-desc">{{creative.creative_size_description}}</span>
                    </div>
                  </div>
                </el-col>
                <el-col :span="7" style="background: #99a9bf;">
                  <div class="creative-banner" v-show="creativeSamples[__adgroup.tab_name].length > 0">
                    <el-carousel height="305px">
                      <el-carousel-item v-for="(sample, sampleIdx) in creativeSamples[__adgroup.tab_name]" :key="sampleIdx">
                        <div class="banner-panel">
                          <p>{{sample.preview_title}}</p>
                          <img :src="sample.creative_size_ample"/>
                        </div>
                      </el-carousel-item>
                    </el-carousel>
                  </div>
                </el-col>
              </el-row>
            </el-form-item>

            <el-form-item label="投放日期"
                          :prop="'adgroups.' + idx + '.adgroup_begin_date'" :rules="{required: true, message: '请选择投放开始日期'}">
              <el-date-picker v-model="__adgroup.adgroup_begin_date" type="date" placeholder="选择开始日期"/> -
              <el-date-picker v-model="__adgroup.adgroup_end_date" type="date" placeholder="结束日期可不选"/>
            </el-form-item>

            <el-form-item label="出价" prop="price">
              <div v-if="Number(__adgroup.creative_size_id) === 0" style="color: #909399;">未选择版位与产品信息</div>
              <div v-else>
                <el-radio-group v-model="__adgroup.price_type">
                  <el-radio-button :label="key" v-for="(label,key) in resources.pricing"
                                   v-if="pricing[__adgroup.tab_name].includes(label)">{{ label }}</el-radio-button>
                </el-radio-group>
                <p style="margin: 5px 0 0 0;">
                  <el-input-number class="w180" v-model="__adgroup.price"/>
                </p>
              </div>
            </el-form-item>

            <el-form-item label="任务名称"
                          :prop="'adgroups.' + idx + '.adgroup_name'" :rules="{required: true, message: '请填写任务名称'}">
              <el-input v-model="__adgroup.adgroup_name" placeholder="限 100 字内的中英文、数字以及破折号 [ - ]"/>
            </el-form-item>

            <el-divider />
          </el-tab-pane>
        </el-tabs>

        <el-divider />
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
import {targetingList, creativeList} from '@a/marketing'
import {creativeCategory, pricing} from '@a/marketing-resource'
import {campaignInfo} from '@a/campaign'

export default {
  // name: 'AdgroupCreate',
  components: {TargetingCreate},
  data() {
    return {
      loadings: {
        pageLoading: false,
        remoteAppLoading: false,
        creativeLoading: false,
      },
      targetings: [],
      campaign: {app: {}},
      creativeList: {adgroup0: []},
      creativeSamples: {adgroup0: []},
      pricing: {adgroup0: []},
      editableTabsValue: 'adgroup0',
      adgroupForm: {
        campaign_id: '',
        product_id: '',
        adgroups: [
          {
            tab_name: 'adgroup0',
            adgroup_name: '任务 0',
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
        creative_category: {},
        pricing: {},
      }
    }
  },
  mounted() {
    this.getResources()
  },
  methods: {
    getPricing() {
      return new Promise((resolve, reject) => {
        pricing().then(res => {
          this.resources.pricing = res.data
          resolve()
        }).catch(() => {
          this.$message.error('付费方式读取失败')
          reject()
        })
      })
    },
    getCreativeCategory() {
      return new Promise((resolve, reject) => {
        creativeCategory().then(res => {
          this.resources.creative_category = res.data
          resolve()
        }).catch(() => {
          this.$message.error('定向包列表读取失败')
          reject()
        })
      })
    },
    getCampaignInfo() {
      return new Promise((resolve, reject) => {
        campaignInfo(this.adgroupForm.campaign_id).then(res => {
          this.campaign = res.data
          if (!this.campaign.app.hasOwnProperty('product_id') || this.campaign.app.product_id === '') {
            this.$message.error('产品信息未同步，请到应用列表拉取关联账户应用数据')
          } else {
            this.$set(this.adgroupForm, 'product_id', this.campaign.app.product_id)
          }
          resolve()
        }).catch(() => {
          this.$message.error('计划信息读取失败')
          reject()
        })
      })
    },
    getTargetingList() {
      return new Promise((resolve, reject) => {
        targetingList({campaign_id: this.adgroupForm.campaign_id}).then(res => {
          this.targetings = res.data
          resolve()
        }).catch(() => {
          this.$message.error('定向包列表读取失败')
          reject()
        })
      })
    },
    getCreativeList(category, tabName) {
      this.loadings.creativeLoading = true
      creativeList({category, account_id: this.campaign.account_id}).then(res => {
        this.creativeList[tabName] = res.data
        this.loadings.creativeLoading = false
      }).catch(() => {
        this.creativeList[tabName] = []
        this.loadings.creativeLoading = false
      })
    },
    getResources() {
      this.loadings.pageLoading = true
      if (this.$route.query.hasOwnProperty('campaign_id') && this.$route.query.campaign_id !== '') {
        this.$set(this.adgroupForm, 'campaign_id', this.$route.query.campaign_id)

        Promise.all([
          this.getPricing(),
          this.getCreativeCategory(),
          this.getCampaignInfo(),
          this.getTargetingList(),
        ]).then(() => {
          this.loadings.pageLoading = false
        }).catch(() => {
          this.loadings.pageLoading = false
        })
      } else {
        this.$message.error("计划参数错误，请重新进入此页面")
      }
    },
    createTargeting(tab, idx) {
      this.$refs.targeting_create.initCreate(tab, this.adgroupForm.campaign_id, {}, idx)
    },
    updateTargeting(tab, idx) {
      let targetingId = Number(this.adgroupForm.adgroups[idx].targeting_package_id)
      if (targetingId === 0) {
        this.$message.error('请先选择定向包')
        return
      }
      let tmpTargeting = {}
      for (let i = 0; i < this.targetings.length; i++) {
        if (this.targetings[i].targeting_id === targetingId) {
          tmpTargeting = Object.assign({}, this.targetings[i])
          tmpTargeting.targeting_name = ''
          break
        }
      }
      this.$refs.targeting_create.initCreate(tab, this.adgroupForm.campaign_id, tmpTargeting, idx)
    },
    targetingCallback(targeting, adgroupIdx) {
      this.targetings.push(targeting)
      this.$set(this.adgroupForm.adgroups[adgroupIdx], 'targeting_package_id', targeting.targeting_id)
    },
    handleSelectCreative(creativeSizeId, samples, supportPriceType) {
      this.$set(this.creativeSamples, this.editableTabsValue, (Array.isArray(samples) ? samples : []))
      this.$set(this.pricing, this.editableTabsValue, supportPriceType.split(","))
      this.adgroupForm.adgroups.forEach((tab, index) => {
        if (tab.tab_name === this.editableTabsValue) {
          this.$set(this.adgroupForm.adgroups[index], 'creative_size_id', creativeSizeId)
        }
      })
    },
    addTab(targetName) {
      let adgroupName = 'adgroup' + this.tabIndex
      this.adgroupForm.adgroups.push({
        tab_name: adgroupName,
        adgroup_name: '任务 ' + this.tabIndex,
        category: '',
        targeting_package_id: '',
        creative_size_id: '',
        adgroup_begin_date: '',
        adgroup_end_date: '',
        price_type: '',
        price: 0,
      })
      this.editableTabsValue = adgroupName
      this.tabIndex++
    },
    copyTab(targetName, idx) {
      let adgroupName = 'adgroup' + this.tabIndex
      let copyAdgroup = Object.assign({}, this.adgroupForm.adgroups[idx])
      copyAdgroup.tab_name = adgroupName
      copyAdgroup.adgroup_name = '任务 ' + this.tabIndex
      this.adgroupForm.adgroups.push(copyAdgroup)
      this.editableTabsValue = adgroupName
      this.tabIndex++
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
@import "../styles/adgroup-create.scss";
</style>
