<template>
  <dialog-panel title="素材选择" confirm-text="确认选择" :visible="visible" @cancel="cancel" @confirm="confirm"
                :confirm-loading="loading" width="800px">
    <div class="asset-box">
      <ul class="asset-list">
        <li :class="{'asset-item': true, 'asset-item-active': item.asset_id === asset.asset_id}"
            v-for="item in assetList" @click="handleSelect(item)">
          <p class="asset-title">{{item.asset_name}}</p>
          <img :src="item.file_url" v-if="['image', 'icon'].includes(ElementType)" @click="showFile(item)" class="asset-file"/>
          <video :type="item.file_format" :src="item.file_url" v-else @click="showFile(item)" class="asset-file"/>
        </li>
      </ul>
    </div>

    <file-show ref="fileShow"/>
  </dialog-panel>
</template>

<script>
import FileShow from '../pages/play-show'
import DialogPanel from '@c/DialogPanel'
import { elementAsset } from '@/api/marketing-resource'

export default {
  props: {
    ElementType: {
      type: String,
      required: true
    },
    ElementId: {
      type: String,
      default: ''
    }
  },
  components: {
    DialogPanel, FileShow
  },
  data() {
    return {
      loading: false,
      visible: false,
      assetSearch: {
        asset_name: '',
        app_id: '',
        file_format: '',
        file_type: '',
        file_size_kb_limit: 0,
        width: 0,
        height: 0,
      },
      assetList: [],
      asset: {
        asset_id: 0,
      }
    }
  },
  methods: {
    initElementSelect(element, app_id) {
      this.assetSearch = {
        app_id: app_id,
        file_format: element.file_format,
        file_size_kb_limit: element.file_size_kb_limit,
        width: element.width,
        height: element.height,
        file_type: element.element_name
      }
      elementAsset(this.assetSearch).then(res => {
        this.assetList = res.data
        this.visible = true
      }).catch(() => {})
    },
    cancel() {
      this.visible = false
    },
    confirm() {
      if (this.asset.asset_id === 0) {
        this.$message.error('请选择素材')
        return
      }
      this.$emit('confirm-select', this.asset, this.ElementType)
      this.assetList = []
      this.visible = false
    },
    handleSelect(asset) {
      this.asset = asset
    },
    showFile(f) {
      this.$refs.fileShow.initPlay(f)
    },
  }
}
</script>

<style scoped lang="scss">
.asset-box {
  max-height: 375px;
  overflow-y: scroll;
  &::-webkit-scrollbar {
    /*滚动条整体样式*/
    width: 3px;
    background-color: #aaa; /* or add it to the track */
  }
  &::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    background: #d5d3d3;
    border-radius: 10px;
  }
  &::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    background: #ffffff;
  }

  .asset-list {
    margin: 0;
    padding: 0;

    .asset-item {
      display: inline-block;
      width: 120px;
      height: 120px;
      border: 1px solid #e3e3e3;
      margin-right: 6px;
      margin-bottom: 5px;
      border-radius: 4px;
      cursor: pointer;
      text-align: center;

      .asset-title {
        height: 20px;
        line-height: 20px;
        padding-left: 5px;
        font-size: 12px;
        color: #909399;
        border-bottom: 1px solid #e3e3e3;
      }
      .asset-file {
        max-width: 118px;
        max-height: 98px;
      }
    }

    .asset-item-active {
      border-color: #1890ff;

      .asset-title {
        color: #1890ff;
      }
    }
  }
}
</style>
