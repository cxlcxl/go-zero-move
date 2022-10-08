<template>
  <el-row :gutter="15" v-loading="creativeLoading" v-show="CreativeSizeList.length > 0">
    <el-col :span="17">
      <div class="creative-container">
        <div v-for="creative in CreativeSizeList"
             :class="{'creative-list': true, 'creative-active': creative.creative_size_id === creative_size_id}"
             @click="handleChangeCreative(creative)">
          <span class="creative-name">{{creative.creative_size_name_dsp}}</span>
          <span class="creative-desc">{{creative.creative_size_description}}</span>
        </div>
      </div>
    </el-col>
    <el-col :span="7" style="background: #99a9bf;">
      <div class="creative-banner" v-show="creative_samples.length > 0">
        <el-carousel height="255px">
          <el-carousel-item v-for="(sample, sampleIdx) in creative_samples" :key="sampleIdx">
            <div class="banner-panel">
              <p>{{sample.preview_title}}</p>
              <img :src="sample.creative_size_ample"/>
            </div>
          </el-carousel-item>
        </el-carousel>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import { positionPlacement } from '@/api/marketing-position'

export default {
  props: {
    CreativeSizeList: {
      type: Array,
      required: true
    },
    CreativeSizeId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      creativeLoading: false,
      creative_samples: [],
      creative_size_id: ''
    }
  },
  mounted() {
    this.initComponent()
  },
  methods: {
    initComponent() {
      this.creative_size_id = this.CreativeSizeId
      if (this.creative_size_id > 0) {
        for (let i = 0; i < this.CreativeSizeList.length; i++) {
          if (this.creative_size_id === this.CreativeSizeList[i].creative_size_id) {
            this.creative_samples = this.CreativeSizeList[i].samples
            break
          }
        }
      }
    },
    handleChangeCreative(creative) {
      this.creative_size_id = creative.creative_size_id
      if (this.CreativeSizeId === creative.creative_size_id) {
        return
      }
      // 版位缩略图轮播
      this.creative_samples = Array.isArray(creative.samples) ? creative.samples : []
      // 创意格式获取
      positionPlacement({creative_size_id: creative.creative_size_id}).then(res => {
        // 回调
        this.$emit('handle-change', this.creative_size_id, creative.support_price_type.split(","), res.data)
      }).catch(() => {
        this.$emit('handle-change', this.creative_size_id, creative.support_price_type.split(","), {})
      })
    }
  }
}
</script>

<style scoped lang="scss">
@import "../styles/creative-size.scss";
</style>

