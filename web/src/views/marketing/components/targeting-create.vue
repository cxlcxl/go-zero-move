<template>
  <drawer-panel title="创建定向" confirm-text="创建" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="800px" direction="ltr">
    <el-form :model="targetingForm" ref="targetingForm" label-width="120px" size="mini" :rules="targetingRules">
      <el-form-item label="定向包名称" prop="targeting_name">
        <el-input v-model="targetingForm.targeting_name" placeholder="限 100 字内的中英文、数字以及破折号 [ - ]"/>
      </el-form-item>
      <el-form-item label="语言" prop="language">
        <el-radio-group v-model="targetingForm.language_check" @change="handleChangeDiy($event, 'language')">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">自定义</el-radio-button>
        </el-radio-group>
        <el-transfer v-model="targetingForm.language" class="dictionary-transfer" v-show="targetingForm.language_check === '1'"
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
      <el-form-item label="APP 行为" prop="app_behavior">
        <el-radio-group v-model="targetingForm.app_behavior" @change="handleChangeDiy($event, 'app_behaviors')">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="app_category_active_struct">一个月内活跃</el-radio-button>
          <el-radio-button label="app_category_installed_struct">已安装</el-radio-button>
          <el-radio-button label="not_app_category_install_struct">未安装</el-radio-button>
        </el-radio-group>
        <el-transfer v-model="targetingForm.app_behaviors" class="dictionary-transfer" v-show="targetingForm.app_behavior !== ''"
                     :props="{key: 'value',label: 'label' }" :data="dictionaries.app_category" filterable
                     :titles="['待选列表', '已选列表']"/>
      </el-form-item>
      <el-form-item label="APP 兴趣" prop="app_interest">
        <el-radio-group v-model="targetingForm.app_interest" @change="handleChangeDiy($event, 'app_interests')">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="unlimit_app_interest_struct">兴趣不限</el-radio-button>
          <el-radio-button label="normal_app_interest_struct">中等兴趣</el-radio-button>
          <el-radio-button label="high_app_interest_struct">高兴趣</el-radio-button>
        </el-radio-group>
        <el-cascader-panel v-model="targetingForm.app_interests" class="dictionary-transfer" :options="dictionaries.app_interest|dataToTree"
                           v-show="targetingForm.app_interest !== ''" :props="{ multiple: true }"/>
      </el-form-item>
      <el-form-item label="设备" prop="series_type">
        <el-radio-group v-model="targetingForm.series_type" @change="handleChangeDiy($event, 'series')">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">按品牌划分</el-radio-button>
        </el-radio-group>
        <el-transfer v-model="targetingForm.series" class="dictionary-transfer" v-show="targetingForm.series_type === '1'"
                     :props="{key: 'value',label: 'label' }" :data="dictionaries.series_type" filterable
                     :titles="['待选列表', '已选列表']"/>
      </el-form-item>
      <el-form-item label="联网方式" prop="network_type_struct">
        <el-checkbox-group v-model="targetingForm.network_type_struct" @change="handleChange($event, 'network_type_struct')">
          <el-checkbox-button label="">不限</el-checkbox-button>
          <el-checkbox-button v-for="item in dictionaries.network_type" :label="item.value">{{ item.label }}</el-checkbox-button>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="自定义人群" prop="audience">
        <el-radio-group v-model="targetingForm.audience">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">自定义</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="媒体类型" prop="media_app_category">
        <el-radio-group v-model="targetingForm.media_app_category">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">包含</el-radio-button>
        </el-radio-group>
        <el-cascader-panel v-model="targetingForm.app_category_of_media_struct" class="dictionary-transfer"
                           :options="dictionaries.media_app_category|dataToTree"
                           v-show="targetingForm.media_app_category === '1'" :props="{ multiple: true }"/>
      </el-form-item>
    </el-form>
  </drawer-panel>
</template>

<script>
import DrawerPanel from '@c/Drawer'
import {dictionaryQuery} from '@a/marketing'
import {arrayToTree} from '@/utils'

export default {
  name: 'TargetingCreate',
  components: {
    DrawerPanel
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
        app_behavior: '',
        app_behaviors: [],
        app_interest: '',
        app_interests: [],
        audience: '',
        audience_struct: [],
        not_audience_struct: [],
        series_type: '',
        series: [],
        media_app_category: '',
        app_category_of_media_struct: [],
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
        app_category: [],
        app_interest: [],
        media_app_category: [],
      }
    }
  },
  filters: {
    dataToTree(origin) {
      let d = []
      if (Array.isArray(origin)) {
        d = arrayToTree(origin, 0, 'pid', 'children')
      }
      return d
    }
  },
  methods: {
    initCreate(tab) {
      let dict_key = "age,series_type,network_type,gender,language,app_category,app_interest,media_app_category"
      dictionaryQuery({dict_key}).then(res => {
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
      if (Number(v) === 0) {

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
