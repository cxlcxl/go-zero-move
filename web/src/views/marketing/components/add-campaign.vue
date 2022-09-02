<template>
  <dialog-panel title="创建计划" confirm-text="创建" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="700px">
    <el-form :model="campaignForm" ref="campaignForm" label-width="120px" size="small" :rules="campaignRules">
      <el-form-item label="计划名称" prop="campaign_name">
        <el-input v-model="campaignForm.campaign_name" placeholder="不能使用“^”,“|”，换行符；最大长度不得超过100；计划名称不得重复"/>
      </el-form-item>
      <el-form-item label="选择账户" prop="account_id">
        <el-select v-model="campaignForm.account_id" remote filterable placeholder="可输入名称查询"
                   :remote-method="remoteMethod" :loading="remoteLoading" style="width: 100%;">
          <el-option v-for="item in accounts" :label="item.account_name" :value="Number(item.id)" v-if="Number(item.is_auth) === 1"/>
        </el-select>
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
  import {campaignCreate} from '@a/marketing'
  import { defaultAccounts, searchAccounts } from '@/api/account'

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
      }
    },
    data() {
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        campaignForm: {
          account_id: '',
          campaign_name: '',
          daily_budget: 10,
          product_type: '',
          campaign_type: '',
          sync_flow_resource_searchad: '',
        },
        campaignRules: {
          campaign_name: {required: true, message: '请填写计划名称'},
          daily_budget: {required: true, message: '请填写日预算'},
          account_id: {required: true, message: '请选择账户'},
          campaign_type: {required: true, message: '请选择计划类型'},
          product_type: {required: true, message: '请选择推广产品'},
          sync_flow_resource_searchad: {required: true, message: '请选择同步投放网络'},
        },
        accounts: []
      }
    },
    methods: {
      initCreate() {
        defaultAccounts().then(res => {
          this.visible = true
          if (Array.isArray(res.data)) {
            this.accounts = res.data
          } else {
            this.accounts = []
          }
        }).catch(err => {
          this.$message.error("请求错误")
        })
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
      },
      remoteMethod(query) {
        if (query.trim() !== '') {
          this.remoteLoading = true;
          searchAccounts(query).then(res => {
            this.remoteLoading = false
            if (Array.isArray(res.data)) {
              this.accounts = res.data
            } else {
              this.accounts = []
            }
          }).catch(() => {
            this.remoteLoading = false
          })
        } else {
          this.options = [];
        }
      }
    }
  }
</script>
