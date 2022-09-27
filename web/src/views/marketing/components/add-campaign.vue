<template>
  <dialog-panel title="创建计划" confirm-text="创建" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="700px">
    <el-form :model="campaignForm" ref="campaignForm" label-width="120px" size="small" :rules="campaignRules">
      <el-form-item label="计划名称" prop="campaign_name">
        <el-input v-model="campaignForm.campaign_name" placeholder="不能使用“^”,“|”，换行符；最大长度不得超过100；计划名称不得重复"/>
      </el-form-item>
      <el-form-item label="推广产品" prop="product_type">
        <el-radio-group v-model="campaignForm.product_type">
          <el-radio-button v-for="(v,k) in ProductType" :label="k">{{v}}</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="日预算" prop="daily_budget">
        <el-input-number v-model="campaignForm.daily_budget" placeholder="请填写日预算" class="w180" :min="8" :max="999999999"/>
        <span class="form-item-tip">不限制填写：999999999</span>
      </el-form-item>
      <el-form-item label="计划类型" prop="role_id">
        <el-radio-group v-model="campaignForm.campaign_type">
          <el-radio-button v-for="(v,k) in CampaignType" :label="k">{{v}}</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="同步投放网络" prop="sync_flow_resource_searchad">
        <el-radio-group v-model="campaignForm.sync_flow_resource_searchad">
          <el-radio-button v-for="(v,k) in SyncFlow" :label="k">{{v}}</el-radio-button>
        </el-radio-group>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {campaignCreate} from '@a/campaign'

  export default {
    components: {
      DialogPanel
    },
    props: {
      CampaignType: {
        required: true,
        type: Object
      },
      ProductType: {
        required: true,
        type: Object
      },
      SyncFlow: {
        required: true,
        type: Object
      },
      AppId: {
        required: true,
        type: String
      }
    },
    data() {
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        campaignForm: {
          app_id: '',
          campaign_name: '',
          daily_budget: 10,
          product_type: '',
          campaign_type: '',
          sync_flow_resource_searchad: '',
        },
        campaignRules: {
          campaign_name: {required: true, message: '请填写计划名称'},
          daily_budget: {required: true, message: '请填写日预算'},
          campaign_type: {required: true, message: '请选择计划类型'},
          product_type: {required: true, message: '请选择推广产品'},
          sync_flow_resource_searchad: {required: true, message: '请选择同步投放网络'},
        }
      }
    },
    methods: {
      initCreate() {
        this.visible = true
        this.$set(this.campaignForm, 'app_id', this.AppId)
        console.log(this.AppId)
      },
      cancel() {
        this.$refs.campaignForm.resetFields()
        this.visible = false
      },
      add() {
        this.$refs.campaignForm.validate(v => {
          if (v) {
            this.loading = true
            campaignCreate(this.campaignForm).then(res => {
              this.$message.success("创建成功")
              this.$emit('success')
              this.loading = false
              this.cancel()
            }).catch(err => {
              this.loading = false
              console.log(err)
            })
          } else {
            return false
          }
        })
      }
    }
  }
</script>
