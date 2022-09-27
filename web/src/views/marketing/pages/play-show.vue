<template>
  <dialog-panel title="素材查看" :visible="visible" @cancel="cancel" :confirm-loading="loading" width="1000px">
    <p style="margin: 0 0 10px 0; color: #1890ff;font-weight: 600;">
      <span v-show="file.app_id !== ''" style="display: inline-block;margin-right: 8px;">应用: {{file.app_id}}</span>
      {{file.width + '*' + file.height + ' / ' + file.asset_name}}
    </p>
    <div class="file-box">
      <img :src="file.file_url" v-if="file.asset_type === 'CREATIVE_ASSET_PICTURE'" class="preview__file"/>
      <video class="preview__file" :src="file.file_url" v-else controls="true"/>
    </div>
  </dialog-panel>
</template>

<script>
import DialogPanel from '@c/DialogPanel'
export default {
  components: {
    DialogPanel
  },
  data() {
    return {
      visible: false,
      loading: false,
      file: {
        app_id: "",
        asset_id: 0,
        asset_name: "",
        asset_type: "CREATIVE_ASSET_PICTURE",
        file_format: "image/png",
        file_size: 0,
        file_url: "",
        height: 0,
        video_play_duration: 0,
        width: 0
      },
    }
  },
  methods: {
    initPlay(file) {
      this.file = file
      setTimeout(() => {
        this.visible = true
      }, 200)
    },
    cancel() {
      this.visible = false
    }
  }
}
</script>

<style scoped lang="scss">
.file-box {
  text-align: center;
  .preview__file {
    display: block;
    margin: 0 auto;
    max-width: 900px;
    max-height: 500px;
  }
}
</style>
