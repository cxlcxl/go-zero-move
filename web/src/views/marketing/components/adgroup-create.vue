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
              <el-radio-group v-model="__adgroup.category" @change="getPositionList($event, __adgroup.tab_name)">
                <el-radio-button v-for="(c_label, c_k) in resources.creative_category" :label="c_k">{{ c_label }}</el-radio-button>
              </el-radio-group>

                <adgroup-creative-size :creative-size-list="positionList[__adgroup.tab_name]" :ref="'creativeSize' + __adgroup.tab_name"
                                       :creative-size-id="__adgroup.creative_size_id" v-show="__adgroup.category !== ''"
                                       @handle-change="handleSelectPosition"/>
            </el-form-item>

            <el-form-item label="投放日期"
                          :prop="'adgroups.' + idx + '.adgroup_begin_date'" :rules="{required: true, message: '请选择投放开始日期'}">
              <el-date-picker v-model="__adgroup.adgroup_begin_date" type="date" placeholder="选择开始日期" value-format="yyyy-MM-dd"/> -
              <el-date-picker v-model="__adgroup.adgroup_end_date" type="date" placeholder="结束日期可不选" value-format="yyyy-MM-dd"/>
            </el-form-item>

            <el-form-item label="出价" :prop="'adgroups.' + idx + '.price'" :rules="{required: true, message: '请填写出价金额'}">
              <div v-if="Number(__adgroup.creative_size_id) === 0" style="color: #909399;">未选择版位与产品信息</div>
              <div v-else v-loading="loadings.priceLoading">
                <el-radio-group v-model="__adgroup.price_type" @change="handleChangePriceType($event, __adgroup.creative_size_id, idx)">
                  <el-radio-button :label="key" v-for="(label,key) in resources.pricing"
                                   v-if="pricing[__adgroup.tab_name].includes(label)">{{ label }}</el-radio-button>
                </el-radio-group>
                <div style="margin: 5px 0 0 0;" v-show="__adgroup.price_type !== ''">
                  <div v-if="__adgroup.price_type !== 'PRICING_OCPC'">
                    <el-input-number class="w180" v-model="__adgroup.price" :min="minPrice" :precision="2" :controls="false"/>
                    <span class="text-error" style="margin-left: 10px; font-size: 12px;">请填写不低于 {{minPrice}} 的金额</span>
                  </div>
                  <div v-else class="price-ocpc">
                    <p class="ocpc-item">
                      <span class="ocpc-item-label">转化名称</span>
                      <el-select v-model="__adgroup.tracking_id" placeholder="请选择">
                        <el-option v-for="track in trackings" :key="track.tracking_id" :label="track.effect_name" :value="track.tracking_id"/>
                      </el-select>
                      <el-button icon="el-icon-refresh" style="margin-left: 5px;" plain @click="refreshTracking"/>
                    </p>
                    <p class="ocpc-item">
                      <span class="ocpc-item-label">期望转化成本</span>
                      <el-input-number v-model="__adgroup.conversion_cost" class="w130" :min="0.01" :max="10" :precision="2" :controls="false"/>
                    </p>
                    <p class="ocpc-item">
                      <span class="ocpc-item-label">学习期点击出价</span>
                      <el-input-number v-model="__adgroup.price" class="w130" :min="0.01" :max="10" :precision="2" :controls="false"/>
                    </p>
                  </div>
                </div>
              </div>
            </el-form-item>

            <el-form-item label="任务名称"
                          :prop="'adgroups.' + idx + '.adgroup_name'" :rules="{required: true, message: '请填写任务名称'}">
              <el-input v-model="__adgroup.adgroup_name" placeholder="限 100 字内的中英文、数字以及破折号 [ - ]"/>
            </el-form-item>

            <el-divider />
            <creative-select :ref="'creativeSelect__'+__adgroup.tab_name"/>
          </el-tab-pane>
        </el-tabs>

        <el-divider />
        <el-form-item>
          <el-button type="primary" icon="el-icon-check" @click="confirmAdgroup">提交</el-button>
          <el-button icon="el-icon-close" @click="cancelCreate">取消创建</el-button>
        </el-form-item>
      </el-col>
    </el-form>

    <targeting-create ref="targeting_create" @handle-success="targetingCallback"/>
  </el-row>
</template>

<script>
import TargetingCreate from './targeting-create'
import { targetingList } from '@a/marketing'
import { positionList, positionPrice, positionCategory } from '@a/marketing-position'
import { pricing, trackingList, trackingRefresh } from '@a/marketing-resource'
import { campaignInfo } from '@a/campaign'
import AdgroupCreativeSize from './adgroup-creative-size'
import CreativeSelect from './creative-select'

export default {
  // name: 'AdgroupCreate',
  components: {TargetingCreate, AdgroupCreativeSize, CreativeSelect},
  data() {
    return {
      loadings: {
        pageLoading: false,
        remoteAppLoading: false,
        positionLoading: false,
        priceLoading: false,
      },
      minPrice: 0,
      targetings: [],
      trackings: [],
      campaign: {app: {}},
      positionList: {adgroup0: []},
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
            conversion_cost: 0,
            tracking_id: 0,
            creative: {}
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
    getPositionCategory() {
      return new Promise((resolve, reject) => {
        positionCategory().then(res => {
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
          this.getTrackingList(this.campaign.app_id)
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
    getTrackingList(app_id) {
      trackingList({app_id}).then(res => {
        this.trackings = res.data
      }).catch(() => {
        this.$message.error('转化跟踪列表读取失败')
      })
    },
    getPositionList(category, tabName) {
      this.loadings.positionLoading = true
      positionList({category, account_id: this.campaign.account_id, product_type: this.campaign.product_type}).then(res => {
        this.positionList[tabName] = res.data
        this.loadings.positionLoading = false
      }).catch(() => {
        this.positionList[tabName] = []
        this.loadings.positionLoading = false
      })
    },
    getResources() {
      this.loadings.pageLoading = true
      if (this.$route.query.hasOwnProperty('campaign_id') && this.$route.query.campaign_id !== '') {
        this.$set(this.adgroupForm, 'campaign_id', this.$route.query.campaign_id)

        Promise.all([
          this.getPricing(),
          this.getPositionCategory(),
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
    handleSelectPosition(creativeSizeId, supportPriceType, placements) {
      this.$set(this.pricing, this.editableTabsValue, supportPriceType)
      this.adgroupForm.adgroups.forEach((tab, idx) => {
        if (tab.tab_name === this.editableTabsValue) {
          this.$refs['creativeSelect__' + this.editableTabsValue][idx].setPositionSubTypes(placements, creativeSizeId)
          this.$set(this.adgroupForm.adgroups[idx], 'creative_size_id', creativeSizeId)
          // 出价联动
          this.$set(this.adgroupForm.adgroups[idx], 'price_type', '')
          this.$set(this.adgroupForm.adgroups[idx], 'price', 0)
        }
      })
    },
    addTab(targetName) {
      let adgroupName = 'adgroup' + this.tabIndex
      this.adgroupForm.adgroups.push({
        tab_name: adgroupName, adgroup_name: '任务 ' + this.tabIndex, category: '', targeting_package_id: '',
        creative_size_id: '', adgroup_begin_date: '', adgroup_end_date: '', price_type: '', price: 0,
        conversion_cost: 0, tracking_id: 0, creative: {}
      })
      this.editableTabsValue = adgroupName

      this.setTabInfo(adgroupName, '')
      this.tabIndex++
    },
    copyTab(targetName, idx) {
      // 本身任务数据
      let adgroupName = 'adgroup' + this.tabIndex
      let copyAdgroup = Object.assign({}, this.adgroupForm.adgroups[idx])
      copyAdgroup.tab_name = adgroupName
      copyAdgroup.adgroup_name = '任务 ' + this.tabIndex
      this.adgroupForm.adgroups.push(copyAdgroup)
      this.editableTabsValue = adgroupName

      this.setTabInfo(adgroupName, targetName)
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

        this.setTabInfo(targetName, '')
      }).catch(() => {})
    },
    setTabInfo(forSetTab, resourceTab) {
      // 版位信息
      this.positionList[forSetTab] = resourceTab === '' ? [] : this.positionList[resourceTab]
      // 出价
      this.pricing[forSetTab] = resourceTab === '' ? [] : this.pricing[resourceTab]
    },
    handleChangePriceType(v, creative_size_id, idx) {
      this.loadings.priceLoading = true
      this.minPrice = 0
      this.$set(this.adgroupForm.adgroups[idx], 'price', 0)
      positionPrice({creative_size_id, price_type: v}).then(res => {
        this.minPrice = res.data
        this.loadings.priceLoading = false
      }).catch(() => {
        this.loadings.priceLoading = false
      })
    },
    refreshTracking() {
      this.loadings.priceLoading = true
      trackingRefresh({app_id: this.campaign.app_id}).then(res => {
        this.trackings = res.data
        this.loadings.priceLoading = false
      }).catch(() => {
        this.$message.error('转化跟踪拉取失败')
        this.loadings.priceLoading = false
      })
    },
    confirmAdgroup() {
      this.$refs.adgroupForm.validate(v => {
        if (v) {
          console.log(this.adgroupForm)
        } else {
          return false
        }
      })
    },
    cancelCreate() {
      this.$confirm('已填写的数据将会丢失, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }).then(() => {
        this.$store.dispatch('tagsView/delView', this.$route).then(() => {
          this.$router.replace({name: "CampaignList"})
        })
      }).catch(() => {})
    }
  }
}
</script>

<style scoped lang="scss">
.price-ocpc {
  background: #f6f6f6;
  padding: 10px;

  .ocpc-item {
    height: 35px;
    line-height: 35px;
    font-size: 12px;

    .ocpc-item-label {
      display: inline-block;
      width: 130px;
      text-align: right;
      padding-right: 10px;
    }
  }
}
</style>
