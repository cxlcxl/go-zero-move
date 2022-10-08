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
    <p class="mt10">
      <el-button icon="el-icon-plus" plain @click="addCreative">添加创意</el-button>
    </p>
    <p class="creative-size mt10" v-show="creative_size.length > 0">
      <el-radio-group v-model="element_search.creative_size" @change="handleChangeSize">
        <el-radio-button :label="v" v-for="(v, label) in creative_size">{{v}}</el-radio-button>
      </el-radio-group>
    </p>
    <el-form class="creative-element" label-width="120px" size="mini">
      <el-form-item v-for="__e in elements" :label="__e.element_title">
        <el-input :placeholder="__e.element_caption"/>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { positionElements } from '@/api/marketing-position'

export default {
  data() {
    return {
      selected: false,
      selectedSize: false,
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
      creatives: {}
    }
  },
  methods: {
    setPositionSubTypes(placements, creativeSizeId) {
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
    handleChangeSubType() {
      this.creative_size = []
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
    },
    fetchElements() {
      positionElements(this.element_search).then(res => {
        this.elements = res.data
      }).catch(() => {})
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
  .mt10 {
    margin-top: 10px;
  }
  .creative-sub-type {
  }
  .creative-element {
    margin-top: 30px;
  }
}
</style>
