<template>
  <dialog-panel title="创建定向" confirm-text="创建" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="1000px">
    <el-form :model="targetingForm" ref="targetingForm" label-width="120px" size="mini" :rules="targetingRules">
      <el-form-item label="定向包名称" prop="targeting_name">
        <el-input v-model="targetingForm.targeting_name" placeholder="限 100 字内的中英文、数字以及破折号 [ - ]"/>
      </el-form-item>
      <el-form-item label="语言" prop="language">
        <el-radio-group v-model="targetingForm.language_check" @change="handleChangeDiy($event, 'language')">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button :label="1">自定义</el-radio-button>
        </el-radio-group>
        <el-transfer v-model="targetingForm.language" class="dictionary-transfer" v-show="targetingForm.language_check === 1"
                     :props="{key: 'value',label: 'label' }" :data="dictionaries.language" filterable
                     :titles="['待选列表', '已选列表']"/>
      </el-form-item>
      <el-form-item label="性别" prop="gender_struct">
        <el-radio-group v-model="targetingForm.gender_struct">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button v-for="item in dictionaries.gender" :label="item.value">{{ item.label }}</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="年龄" prop="age_struct">
        <el-checkbox-group v-model="targetingForm.age_struct" @change="handleChange($event, 'age_struct')">
          <el-checkbox-button label="">不限</el-checkbox-button>
          <el-checkbox-button v-for="item in dictionaries.age" :label="item.value">{{ item.label }}</el-checkbox-button>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="APP 安装" prop="installed_apps">
        <el-radio-group v-model="targetingForm.installed_apps">
          <el-radio-button label="1">已安装</el-radio-button>
          <el-radio-button label="0">未安装</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="设备" prop="series_type">
        <el-radio-group v-model="targetingForm.series_type">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">按品牌划分</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="联网方式" prop="network_type_struct">
        <el-checkbox-group v-model="targetingForm.network_type_struct" @change="handleChange($event, 'network_type_struct')">
          <el-checkbox-button label="">不限</el-checkbox-button>
          <el-checkbox-button v-for="item in dictionaries.network_type" :label="item.value">{{ item.label }}</el-checkbox-button>
        </el-checkbox-group>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from '@c/DialogPanel'
import {dictionaryQuery} from '@a/marketing'

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
        series_type: '',
        language_check: '',
        language: [],
        installed_apps: '1',
      },
      targetingRules: {
        targeting_name: {required: true, message: '请填写定向包名称'},
      },
      creativeTab: '',
      dictionaries: {
        age: [],
        series_type: [],
        network_type: [],
        gender: [],
        language: [],
      }
    }
  },
  methods: {
    initCreate(tab) {
      dictionaryQuery({dict_key: "age,series_type,network_type,gender,language"}).then(res => {
        this.dictionaries = res.data
        this.creativeTab = tab
        this.visible = true
      }).catch(() => {
        this.$message.error('字典参数请求失败')
      })
    },
    cancel() {
      this.$refs.targetingForm.resetFields()
      this.visible = false
    },
    add() {
      console.log(this.targetingForm)
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
    },
    handleChangeDiy(v, attr) {
      if (Number(v) === 1) {

      }
    }
  }
}
</script>

<style lang="scss">
.dictionary-transfer {
  margin-top: 5px;

  .el-transfer__buttons {
    padding: 0 10px;
  }
}
</style>
