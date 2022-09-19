<template>
  <drawer-panel title="创建定向" confirm-text="创建" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="800px" direction="ltr">
    <el-form :model="targetingForm" ref="targetingForm" label-width="120px" size="mini" :rules="targetingRules" v-loading="loading">
      <el-form-item label="定向包名称" prop="targeting_name">
        <el-input v-model="targetingForm.targeting_name" placeholder="限 100 字内的中英文、数字以及破折号 [ - ]"/>
      </el-form-item>
      <el-form-item label="地域" prop="location">
        <el-radio-group v-model="targetingForm.location">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">自定义地理位置</el-radio-button>
        </el-radio-group>
        <div v-show="targetingForm.location === '1'" style="margin-top: 5px;">
          <el-radio-group v-model="targetingForm.location_type">
            <el-radio label="current">当前在这里的人</el-radio>
            <el-radio label="residence">常住在这里的人</el-radio>
          </el-radio-group>
          <div class="dictionary-location">
            <el-button plain type="primary" icon="el-icon-plus" @click="selectLocation">选择国家/地区</el-button>
            <targeting-location ref="location" :continents="targetingLocations.continents"
                                :countries="targetingLocations.countries" @handle-select="locationSelector"/>

            <targeting-location-view ref="locationView" :countries="targetingLocations.countries"
                                     @handle-change="locationChange"/>
          </div>
        </div>
      </el-form-item>
      <el-form-item label="运营商" prop="carrier">
        <el-radio-group v-model="targetingForm.carrier">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">自定义</el-radio-button>
        </el-radio-group>
        <el-cascader-panel v-model="targetingForm.carriers" class="dictionary-transfer" :options="dictionaries.carrier" filterable clearable
                           v-show="targetingForm.carrier === '1'" :props="{ multiple: true }"/>
      </el-form-item>
      <el-form-item label="语言" prop="language">
        <el-radio-group v-model="targetingForm.language_check">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">自定义</el-radio-button>
        </el-radio-group>
        <el-transfer v-model="targetingForm.language" class="dictionary-transfer" v-show="targetingForm.language_check === '1'"
                     :props="{key: 'value',label: 'label' }" :data="dictionaries.language"
                     :titles="['待选列表', '已选列表']"/>
      </el-form-item>
      <el-form-item label="性别" prop="gender">
        <el-radio-group v-model="targetingForm.gender">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button v-for="item in dictionaries.gender" :label="item.value">{{ item.label }}</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="年龄" prop="age">
        <el-checkbox-group v-model="targetingForm.age" @change="handleChange($event, 'age')">
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
      <el-form-item label="APP 行为" prop="app_category">
        <el-radio-group v-model="targetingForm.app_category">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">一个月内活跃</el-radio-button>
          <el-radio-button label="2">已安装</el-radio-button>
          <el-radio-button label="3">未安装</el-radio-button>
        </el-radio-group>
        <el-transfer v-model="targetingForm.app_categories" class="dictionary-transfer" v-show="targetingForm.app_category !== ''"
                     :props="{key: 'value',label: 'label' }" :data="dictionaries.app_category"
                     :titles="['待选列表', '已选列表']"/>
      </el-form-item>
      <el-form-item label="APP 兴趣" prop="app_interest">
        <el-radio-group v-model="targetingForm.app_interest">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">兴趣不限</el-radio-button>
          <el-radio-button label="2">中等兴趣</el-radio-button>
          <el-radio-button label="3">高兴趣</el-radio-button>
        </el-radio-group>
        <el-cascader-panel v-model="targetingForm.app_interests" class="dictionary-transfer" :options="dictionaries.app_interest"
                           v-show="targetingForm.app_interest !== ''" :props="{ multiple: true }"/>
      </el-form-item>
      <el-form-item label="设备" prop="series_type">
        <el-radio-group v-model="targetingForm.series_type">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">按品牌划分</el-radio-button>
        </el-radio-group>
        <el-transfer v-model="targetingForm.series" class="dictionary-transfer" v-show="targetingForm.series_type === '1'"
                     :props="{key: 'value',label: 'label' }" :data="dictionaries.series_type"
                     :titles="['待选列表', '已选列表']"/>
      </el-form-item>
      <el-form-item label="联网方式" prop="network_type">
        <el-checkbox-group v-model="targetingForm.network_type" @change="handleChange($event, 'network_type')">
          <el-checkbox-button label="">不限</el-checkbox-button>
          <el-checkbox-button v-for="item in dictionaries.network_type" :label="item.value">{{ item.label }}</el-checkbox-button>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="自定义人群" prop="audience">
        <el-radio-group v-model="targetingForm.audience">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">自定义</el-radio-button>
        </el-radio-group>

        <el-row :gutter="5" class="dictionary-audience" v-show="targetingForm.audience === '1'">
          <el-col :span="24">
            <p>包含人群</p>
            <el-table :data="dictionaries.pre_define_audience" size="mini" max-height="244px" border ref="audience"
                      @selection-change="audienceSelectionChange">
              <el-table-column prop="id" type="selection" width="50" align="center" :selectable="audienceSelectable"/>
              <el-table-column prop="label" label="人群包名称"/>
              <el-table-column label="覆盖量" width="90" align="right">
                <template slot-scope="scope">{{scope.row.value|audienceNumFilter}}</template>
              </el-table-column>
            </el-table>
          </el-col>

          <el-col :span="24">
            <p>排除人群</p>
            <el-table :data="dictionaries.not_pre_define_audience" size="mini" max-height="244px" border ref="not_audience"
                      @selection-change="notAudienceSelectionChange">
              <el-table-column prop="id" type="selection" width="50" align="center" :selectable="notAudienceSelectable"/>
              <el-table-column prop="label" label="人群包名称"/>
              <el-table-column label="覆盖量" width="90" align="right">
                <template slot-scope="scope">{{scope.row.value|audienceNumFilter}}</template>
              </el-table-column>
            </el-table>
          </el-col>
        </el-row>
      </el-form-item>
      <el-form-item label="媒体类型" prop="media_app_category">
        <el-radio-group v-model="targetingForm.media_app_category">
          <el-radio-button label="">不限</el-radio-button>
          <el-radio-button label="1">包含</el-radio-button>
        </el-radio-group>
        <el-cascader-panel v-model="targetingForm.app_category_of_media" class="dictionary-transfer"
                           :options="dictionaries.media_app_category"
                           v-show="targetingForm.media_app_category === '1'" :props="{ multiple: true }"/>
      </el-form-item>
    </el-form>
  </drawer-panel>
</template>

<script>
import DrawerPanel from '@c/Drawer'
import TargetingLocation from './targeting-location'
import TargetingLocationView from './targeting-location-view'
import {dictionaryQuery, targetingLocation, targetingCreate} from '@a/marketing'

export default {
  name: 'TargetingCreate',
  components: {
    DrawerPanel, TargetingLocation, TargetingLocationView
  },
  data() {
    return {
      visible: false,
      loading: false,
      remoteLoading: false,
      adgroupIndex: 0,
      targetingForm: {
        targeting_id: 0,
        campaign_id: '',
        targeting_type: '',
        targeting_name: '',
        gender: '',
        age: [''],
        network_type: [''],
        location: '',
        location_type: 'current',
        include_location: [],
        exclude_location: [],
        carrier: '',
        carriers: [],
        app_category: '',
        app_categories: [],
        app_interest: '',
        app_interests: [],
        audience: '',
        audiences: [],
        not_audience: [],
        series_type: '',
        series: [],
        media_app_category: '',
        app_category_of_media: [],
        language_check: '',
        language: [],
        installed_apps: '1',
      },
      defaultDictKeys: 'age,series_type,network_type,gender,language,app_category,app_interest,media_app_category,carrier,pre_define_audience,not_pre_define_audience',
      targetingRules: {
        targeting_name: {required: true, message: '请填写定向包名称'},
      },
      targetingLocations: {
        countries: [],
        continents: []
      },
      creativeTab: '',
      dictionaries: {
        age: [],
        series_type: [],
        network_type: [],
        carrier: [],
        gender: [],
        language: [],
        app_category: [],
        app_interest: [],
        media_app_category: [],
        pre_define_audience: [],
        not_pre_define_audience: [],
      }
    }
  },
  filters: {
    audienceNumFilter(d) {
      if (Number(d) > 0) {
        return (Number(d)/10000).toFixed(0) + '万'
      }
      return d
    }
  },
  methods: {
    getLocationInfo() {
      return new Promise((resolve, reject) => {
        targetingLocation().then(res => {
          this.targetingLocations = res.data
          resolve()
        }).catch(() => {
          this.$message.error('地域信息读取失败')
          reject()
        })
      })
    },
    initCreate(tab, campaignId, source, idx) {
      this.adgroupIndex = idx
      Promise.all([
        this.getLocationInfo(),
      ]).then(() => {
        dictionaryQuery({dict_key: this.defaultDictKeys}).then(res => {
          this.dictionaries = res.data
          this.creativeTab = tab
          this.visible = true

          if (source.hasOwnProperty('targeting_name')) {
            this.loading = true
            this.$nextTick(() => {
              this.$refs.locationView.watchChange({
                _include: source.include_location,
                _exclude: source.exclude_location
              })
            })
            this.targetingForm = source
            this.loading = false
            // 自定义人群设置 toggleRowSelection
            if (Array.isArray(this.targetingForm.audiences) && this.targetingForm.audiences.length > 0) {
              for (let i = 0; i < this.dictionaries.pre_define_audience.length; i++) {
                if (this.targetingForm.audiences.includes(this.dictionaries.pre_define_audience[i].id)) {
                  this.$nextTick(() => {
                    this.$refs.audience.toggleRowSelection(this.dictionaries.pre_define_audience[i], true)
                  })
                }
              }
            }
            if (Array.isArray(this.targetingForm.not_audience) && this.targetingForm.not_audience.length > 0) {
              for (let i = 0; i < this.dictionaries.not_pre_define_audience.length; i++) {
                if (this.targetingForm.not_audience.includes(this.dictionaries.not_pre_define_audience[i].id)) {
                  this.$nextTick(() => {
                    this.$refs.not_audience.toggleRowSelection(this.dictionaries.not_pre_define_audience[i], true)
                  })
                }
              }
            }
          } else {
            this.resetForm()
          }
          this.$set(this.targetingForm, 'campaign_id', campaignId)
        }).catch(() => {
          this.$message.error('字典参数请求失败')
        })
      }).catch(() => {})
    },
    resetForm() {
      this.targetingForm = {
        targeting_id: 0,
        campaign_id: '',
        targeting_type: '',
        targeting_name: '',
        gender: '',
        age: [''],
        network_type: [''],
        location: '',
        location_type: 'current',
        include_location: [],
        exclude_location: [],
        carrier: '',
        carriers: [],
        app_category: '',
        app_categories: [],
        app_interest: '',
        app_interests: [],
        audience: '',
        audiences: [],
        not_audience: [],
        series_type: '',
        series: [],
        media_app_category: '',
        app_category_of_media: [],
        language_check: '',
        language: [],
        installed_apps: '1'
      }
    },
    cancel() {
      this.resetForm()
      this.visible = false
    },
    add() {
      this.$refs.targetingForm.validate(v => {
        if (v) {
          this.loading = true
          targetingCreate(this.targetingForm).then(res => {
            this.loading = false
            let assignForm = Object.assign({}, this.targetingForm)
            assignForm.targeting_id = res.data.targeting_id
            this.$emit('handle-success', assignForm, this.adgroupIndex)
            this.cancel()
          }).catch(() => {
            this.loading = false
          })
          this.$emit('success', this.creativeTab)
        } else {
          return false
        }
      })
    },
    handleChange(v, attr) {
      if (v.length > 1) {
        if (v[v.length-1] === '') {
          this.$set(this.targetingForm, attr, [''])
        } else {
          let vs = []
          v.forEach((_v, k) => {
            if (_v !== '') {
              vs.push(_v)
            }
          })
          this.$set(this.targetingForm, attr, vs)
        }
      }
    },
    notAudienceSelectionChange(v) {
      let val = []
      v.forEach((item, k) => {val.push(item.id)})
      this.$set(this.targetingForm, 'not_audience', val)
    },
    audienceSelectionChange(v) {
      let val = []
      v.forEach((item, k) => {val.push(item.id)})
      this.$set(this.targetingForm, 'audiences', val)
    },
    audienceSelectable(row, index) {
      return this.targetingForm.not_audience.includes(row.id) === false
    },
    notAudienceSelectable(row, index) {
      return this.targetingForm.audiences.includes(row.id) === false
    },
    selectLocation() {
      this.$refs.location.initCreate(this.targetingForm.include_location, this.targetingForm.exclude_location)
    },
    locationSelector(v) {
      this.$set(this.targetingForm, 'include_location', v._include)
      this.$set(this.targetingForm, 'exclude_location', v._exclude)
      this.$refs.locationView.watchChange(v)
    },
    locationChange(v) {
      this.$set(this.targetingForm, 'include_location', v._include)
      this.$set(this.targetingForm, 'exclude_location', v._exclude)
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

  .el-transfer-panel {
    width: 280px;

    .el-transfer-panel__header {
      height: 30px;
      line-height: 30px;

      .el-checkbox {
        line-height: 30px;

        .el-checkbox__label {
          font-size: 14px;
        }
      }
    }
    .el-transfer-panel__body {
      .el-transfer-panel__item {
        height: 25px;
        line-height: 25px;

        .el-checkbox__label {
          font-size: 12px;
        }
      }
    }
  }

  .el-transfer__buttons {
    .el-button--mini {
      display: block;
      padding: 5px 10px;
      margin: 0 0 5px 0;
    }
  }
}
.dictionary-select {
  display: block;
  margin-top: 5px;
}
.dictionary-audience {
  margin-top: 5px;

  .el-table--mini th, .el-table--mini td {
    padding: 2px 0;
  }
}
.dictionary-location {
  margin-top: 5px;
}
</style>
