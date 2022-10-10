<template>
  <div class="creative-select-container" v-show="selected">
    <p class="creative-title text-primary">创意素材选择</p>
    <p class="creative-sub-type mt10">
      <el-radio-group v-model="element_search.sub_type" @change="handleChangeSubType">
        <el-radio-button :label="label" v-for="(item, label) in placement_sub_types.placements">
          <span>{{placement_sub_types.sub_types[label]}}</span>
          <span v-show="creatives.hasOwnProperty(label) && creatives[label].number > 0" style="margin-left: 10px;">{{creatives[label].number}}</span>
        </el-radio-button>
      </el-radio-group>
    </p>
    <div class="creative-opt">
      <el-button icon="el-icon-plus" plain @click="addCreative" :disabled="creativesForm.length >= 5">添加创意</el-button>

      <span :class="{'creative-item': true, 'creative-item-active': creative.creative_index === creative_index-1}"
            v-for="(creative, idx) in creativesForm" v-if="creative.sub_type === element_search.sub_type">
        创意 {{creative.creative_index}} <i class="el-icon-delete text-error" @click="removeCreative(idx)"/>
      </span>
    </div>
    <p class="creative-size mt10" v-show="creative_size.length > 0">
      <el-radio-group v-model="element_search.creative_size" @change="handleChangeSize">
        <el-radio-button :label="v" v-for="(v, label) in creative_size">{{v}}</el-radio-button>
      </el-radio-group>
    </p>
    <el-form class="creative-element" label-width="120px" size="mini" ref="elementForm" :model="elementForm">
      <el-form-item v-for="__e in elements" :label="__e.element_title" :prop="__e.element_name">
        <div v-if="['image', 'icon', 'video'].includes(__e.element_name)">
          <div class="element-upload" @click="elementAsset(__e)" v-if="!assets[__e.element_name].hasOwnProperty('asset_id')">
            <i class="el-icon-plus"/>
            <p class="element-tip">{{__e.width + 'px * ' + __e.height + 'px'}}</p>
            <p class="element-tip">{{__e.file_format}}</p>
            <p class="element-tip">不超过 {{__e.file_size_kb_limit|sizeLimitFilter}}</p>
          </div>
          <div class="element-selected" v-else>
            <p class="asset-remove" @click="reselect(__e.element_name)">重新选择</p>
            <img class="asset" :src="assets[__e.element_name].file_url" v-if="['image', 'icon'].includes(__e.element_name)"/>
            <video class="asset" :type="assets[__e.element_name].file_format" :src="assets[__e.element_name].file_url" v-else/>
          </div>

          <position-element-asset :ref="'elementAsset__' + __e.element_id" :element-type="__e.element_name" @confirm-select="selectedAsset"/>
        </div>
        <el-input :placeholder="__e.element_caption" v-model="elementForm[__e.element_name]" v-else/>
      </el-form-item>
      <el-form-item prop="creative_name" label="创意名称" v-show="elements.length > 0" :rules="{required: true, message: '请填写创意名称'}">
        <el-input placeholder="请填写创意名称" v-model="elementForm.creative_name"/>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import PositionElementAsset from './position-element-asset'
import { positionElements } from '@/api/marketing-position'

export default {
  components: { PositionElementAsset },
  data() {
    return {
      selected: false,
      selectedSize: false,
      app_id: "",
      placement_sub_types: {
        placements: {},
        sub_types: {}
      },
      element_search: {
        creative_size_id: '',
        sub_type: '',
        creative_size: ''
      },
      creative_size: [],
      elements: [],
      creatives: {},
      headers: {},
      element_image: {},
      fileList: [],
      uploadUrl: "",
      elementForm: {
        creative_name: '',
      },
      creative_index: 1,
      creativesForm: [], // {creative_name: '', sub_type: '', creative_size: '', creative_index: 1}
      assets: {}
    }
  },
  filters: {
    acceptsFilter(v) {
      return v
    },
    sizeLimitFilter(v) {
      if (v > 1024) {
        return (v/1024).toFixed(2) + 'MB'
      } else {
        return v + 'KB'
      }
    }
  },
  methods: {
    setPositionSubTypes(placements, creativeSizeId, appId) {
      this.app_id = appId
      this.element_search.creative_size_id = creativeSizeId
      this.creatives = {}
      this.creative_size = []
      this.elements = []
      this.placement_sub_types = placements
      let filled = false
      for (const placement in placements.placements) {
        this.creatives[placement] = {number: 0}
        if (!filled) {
          this.$set(this.element_search, 'sub_type', placement)
          this.$set(this.element_search, 'creative_size', '')
          filled = true
        }
      }
      this.selected = true
    },
    handleChangeSubType(sub_type) {
      if (this.creativesForm.length > 0) {
        for (let i = 0; i < this.creativesForm.length; i++) {
          if (this.creativesForm[i].sub_type === sub_type) {
            this.$set(this.element_search, 'creative_size', this.creativesForm[i].creative_size)
            this.fetchElements()
            return
          }
        }
      }
      this.creative_size = []
      this.elements = []
      this.$set(this.element_search, 'creative_size', '')
    },
    handleChangeSize() {
      if (!this.selectedSize) {
        this.fetchElements()
      } else {
        this.$confirm('将清除已填入的创意, 是否继续?', '确认修改尺寸?', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.fetchElements()
        }).catch(() => {})
      }
      this.selectedSize = true
    },
    addCreative() {
      if (this.element_search.sub_type === '') {
        this.$message.error('请先选择版位形式')
        return
      }

      this.creative_size = this.placement_sub_types.placements[this.element_search.sub_type]
      if (this.creative_size.length === 1) {
        this.element_search.creative_size = this.creative_size[0]
        this.selectedSize = true
        this.fetchElements()
      } else {
        this.selectedSize = false
      }

      this.creativesForm.push({
        creative_name: '',
        sub_type: this.element_search.sub_type,
        creative_index: this.creative_index,
        creative_size: this.element_search.creative_size,
      })
      this.creative_index++
    },
    removeCreative(idx) {
      this.creativesForm.splice(idx, 1)
      if (this.creativesForm.length === 0) {
        this.creative_index = 1
      } else {
        let creative = this.creativesForm[this.creativesForm.length - 1]
        this.creative_index = creative.creative_index
        this.$set(this.element_search, 'creative_size', creative.creative_size)
        this.$set(this.element_search, 'sub_type', creative.sub_type)
      }
    },
    fetchElements() {
      positionElements(this.element_search).then(res => {
        this.elements = res.data
        for (let e of this.elements) {
          this.elementForm[e.element_name] = ''
          this.assets[e.element_name] = {}
        }
      }).catch(() => {})
    },
    elementAsset(element) {
      this.$refs['elementAsset__' + element.element_id][0].initElementSelect(element, this.app_id)
    },
    selectedAsset(asset, element_name) {
      this.$set(this.elementForm, element_name, asset.asset_id)
      this.$set(this.assets, element_name, asset)
      console.log(this.assets, this.elementForm)
    },
    reselect(element_name) {
      this.$set(this.elementForm, element_name, 0)
      this.$set(this.assets, element_name, {})
    }
  }
}
</script>

<style scoped lang="scss">
.creative-select-container {
  padding-left: 20px;

  .creative-title {
    height: 30px;
    line-height: 30px;
    padding-left: 10px;
    margin: 0;
    font-weight: 600;
    font-size: 14px;
    border-left: 3px solid #1890ff;
  }
  .creative-opt {
    margin-top: 10px;

    .creative-item {
      font-size: 12px;
      height: 32px;
      line-height: 32px;
      display: inline-block;
      padding: 0 10px;
      margin: 0 0 0 5px;
      border: 1px solid #e3e3e3;
      border-radius: 3px;
      color: #606266;
      cursor: pointer;

      .el-icon-delete {
        margin-left: 5px;
      }

      &:hover {
        color: #1890ff;
        border-color: #1890ff;
      }
    }

    .creative-item-active {
      color: #1890ff;
      border-color: #1890ff;
    }
  }
  .mt10 {
    margin-top: 10px;
  }
  .creative-element {
    margin-top: 30px;

    .element-selected {
      width: 120px;

      .asset-remove {
        font-size: 12px;
        color: #606266;
        text-align: center;
        height: 25px;line-height: 25px;
        border: 1px solid #e3e3e3;
        margin-bottom: 3px;
        cursor: pointer;

        &:hover {
          color: #ff4949;
          border: 1px solid #ff4949;
        }
      }
      .asset {
        max-width: 150px;
        max-height: 120px;
      }
    }
    .element-upload {
      width: 120px;
      height: 120px;
      border: 1px dashed #e3e3e3;
      cursor: pointer;
      border-radius: 4px;

      .el-icon-plus {
        display: block;
        width: 30px;
        height: 30px;
        text-align: center;
        line-height: 30px;
        margin: 10px auto 5px;
        background: #a2a7ad;
        color: #ffffff;
        border-radius: 50%;
        font-weight: 600;
      }
      .element-tip {
        color: #86898f;
        height: 20px;
        margin: 0;
        padding: 0;
        line-height: 20px;
        font-size: 12px;
        text-align: center;
      }
    }
  }
}
</style>
