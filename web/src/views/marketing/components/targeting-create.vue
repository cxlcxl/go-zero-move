<template>
  <dialog-panel title="创建定向" confirm-text="创建" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="700px">
    <el-form :model="targetingForm" ref="targetingForm" label-width="120px" size="mini" :rules="targetingRules">
      <el-form-item label="定向包名称" prop="targeting_name">
        <el-input v-model="targetingForm.targeting_name" placeholder="限 100 字内的中英文、数字以及破折号 [ - ]"/>
      </el-form-item>
      <el-form-item label="性别" prop="gender_struct">
        <el-radio-group v-model="targetingForm.gender_struct">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="0">男</el-radio-button>
          <el-radio-button label="1">女</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="年龄" prop="age_struct">
        <el-checkbox-group v-model="targetingForm.age_struct" @change="handleChange($event, 'age_struct')">
          <el-checkbox-button label="">不限</el-checkbox-button>
          <el-checkbox-button label="1">18~23岁</el-checkbox-button>
          <el-checkbox-button label="2">24~34岁</el-checkbox-button>
          <el-checkbox-button label="3">35~44岁</el-checkbox-button>
          <el-checkbox-button label="4">45~54岁</el-checkbox-button>
          <el-checkbox-button label="5">55岁及以上</el-checkbox-button>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="APP 安装" prop="installed_apps">
        <el-radio-group v-model="targetingForm.installed_apps">
          <el-radio-button label="1">已安装</el-radio-button>
          <el-radio-button label="0">未安装</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="设备" prop="series_type">

      </el-form-item>
      <el-form-item label="联网方式" prop="network_type_struct">
        <el-checkbox-group v-model="targetingForm.network_type_struct" @change="handleChange($event, 'network_type_struct')">
          <el-checkbox-button label="">不限</el-checkbox-button>
          <el-checkbox-button label="WIFI">WIFI</el-checkbox-button>
          <el-checkbox-button label="2G">2G</el-checkbox-button>
          <el-checkbox-button label="3G">3G</el-checkbox-button>
          <el-checkbox-button label="4G">4G</el-checkbox-button>
          <el-checkbox-button label="5G">5G</el-checkbox-button>
        </el-checkbox-group>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from '@c/DialogPanel'

export default {
  name: 'TargetingCreate',
  components: {
    DialogPanel
  },
  props: {
    AccountId: {
      required: true,
      type: Number
    }
  },
  data() {
    return {
      visible: false,
      loading: false,
      remoteLoading: false,
      targetingForm: {
        account_id: 0,
        targeting_type: '',
        targeting_name: '',
        gender_struct: '',
        age_struct: [''],
        network_type_struct: [''],
        series_type: [''],
        installed_apps: '1',
      },
      targetingRules: {
        targeting_name: {required: true, message: '请填写定向包名称'},
      },
      creativeTab: ''
    }
  },
  methods: {
    initCreate(tab) {
      this.creativeTab = tab
      this.visible = true
    },
    cancel() {
      this.$refs.targetingForm.resetFields()
      this.visible = false
    },
    add() {
      this.$refs.targetingForm.validate(v => {
        if (v) {
          this.$emit('success', this.creativeTab)
        } else {
          return false
        }
      })
    },
    remoteMethod(query) {
      if (query.trim() !== '') {
        this.remoteLoading = true;

      } else {

      }
    },
    handleChange(v, attr) {
      if (v.length > 1) {
        if (v[v.length-1] === '') {
          this.$set(this.targetingForm, attr, [''])
        } else {
          this.targetingForm[attr] = v.filter(item => item !== '')
        }
      }
    }
  }
}
</script>
