<template>
  <dialog-panel title="素材上传 - 需要命名请直接对素材进行名字修改" :visible="visible" @cancel="cancel" confirm-text="确认上传" @confirm="upload"
                :confirm-loading="loading" width="900px">
    <el-form :model="assetForm" :rules="rules" ref="assetForm" label-width="110px" size="mini">
      <el-form-item label="应用 ID" prop="app_id">
        <el-input v-model="assetForm.app_id" disabled/>
      </el-form-item>
      <el-form-item label="素材类型" prop="asset_type">
        <el-radio-group v-model="assetForm.asset_type">
          <el-radio-button label="CREATIVE_ASSET_PICTURE">图片素材</el-radio-button>
          <el-radio-button label="CREATIVE_ASSET_VIDEO">视频素材</el-radio-button>
        </el-radio-group>
        <div style="font-size: 12px;">
          <p style="cursor: pointer;">
            <i :class="{'text-primary': true, 'el-icon-plus': !dimensionShow, 'el-icon-minus': dimensionShow}"
               @click="dimensionShow = !dimensionShow"> 点我查看允许尺寸</i>
          </p>
          <div v-show="dimensionShow">
            <el-tag v-for="dimension in assetDimensions[assetForm.asset_type]" size="mini" style="margin-right: 2px;">{{dimension}}</el-tag>
          </div>
        </div>
      </el-form-item>
      <el-form-item label="选择素材" prop="asset_type">
        <el-upload :action="uploadUrl" :headers="headers" multiple :limit="10" ref="assetUpload"
                   :on-exceed="overLimit" name="file"
                   :auto-upload="false" :file-list="fileList"
                   :data="assetForm" :on-change="handleChange" :on-remove="handleChange"
                   :accept="accepts[assetForm.asset_type]"
                   :before-upload="prepareUpload"
                   :on-error="uploadErr"
                   :on-success="uploadSuccess">
          <el-button slot="trigger" type="primary" plain icon="el-icon-plus">选取文件</el-button>
        </el-upload>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from '@c/DialogPanel'
import { getToken } from '@/utils/auth'
import {assetUpload, assetToken, assetDimension} from '@/api/marketing-resource'

export default {
  components: {
    DialogPanel
  },
  data() {
    return {
      visible: false,
      loading: false,
      dimensionShow: false,
      assetForm: {
        app_id: '',
        asset_type: 'CREATIVE_ASSET_PICTURE',
        file_token: '',
        duration: 0,
        width: 0,
        height: 0,
      },
      accepts: {
        CREATIVE_ASSET_PICTURE: ".jpg,.jpeg,.png,.gif",
        CREATIVE_ASSET_VIDEO: ".mp4",
      },
      headers: {
        "Authorization": 'Bearer ' + getToken()
      },
      rules: {
        app_id: {required: true, message: '应用必需先选择'},
        asset_type: {required: true, message: '请选择素材类型'}
      },
      assetDimensions: {
        CREATIVE_ASSET_PICTURE: [],
        CREATIVE_ASSET_VIDEO: []
      },
      fileList: [],
      assetNumbers: {
        total: 0,
        success: 0,
        failed: 0
      }
    }
  },
  computed: {
    uploadUrl() {
      return assetUpload
    }
  },
  methods: {
    initUpload(app_id) {
      assetDimension().then(res => {
        this.assetDimensions = res.data
        this.assetForm.app_id = app_id
        this.visible = true
      }).catch(() => {})
    },
    prepareUpload(file) {
      let _this = this
      return new Promise((resolve, reject) => {
        //普通的判断可以用return false
        // 获取文件尺寸，判断尺寸在不在规定范围之内
        let reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onload = function(theFile) {
          if (_this.assetForm.asset_type === 'CREATIVE_ASSET_PICTURE') {
            let image = new Image()
            image.src = theFile.target.result
            image.onload = function() {
              let dimension = `${this.width}*${this.height}`
              if (!_this.assetDimensions.CREATIVE_ASSET_PICTURE.includes(dimension)) {
                _this.$notify.error({ title: '素材上传提示', message: `${file.name} 尺寸不对，请重新上传`, duration: 10000 })
                _this.loading = false
                return reject()
              }
              _this.$set(_this.assetForm, 'width', this.width)
              _this.$set(_this.assetForm, 'height', this.height)
            }
          } else {
            const videoUrl = URL.createObjectURL(file)
            const videoObj = document.createElement('video')
            videoObj.preload = 'metadata'
            videoObj.src = videoUrl
            videoObj.onloadedmetadata = function (evt) {
              let dimension = `${videoObj.videoWidth}*${videoObj.videoHeight}`
              if (!_this.assetDimensions.CREATIVE_ASSET_VIDEO.includes(dimension)) {
                _this.$notify.error({ title: '素材上传提示', message: `${file.name} 尺寸不对，请重新上传`, duration: 10000 })
                _this.loading = false
                return reject()
              }
              _this.$set(_this.assetForm, 'width', this.width)
              _this.$set(_this.assetForm, 'height', this.height)
              _this.$set(_this.assetForm, 'duration', Math.round(videoObj.duration))
            }
          }
        }

        assetToken({app_id: this.assetForm.app_id}).then(res => {
          _this.$set(_this.assetForm, 'file_token', res.data)
          resolve()
        }).catch(() => {
          _this.loading = false
          reject()
        })
      })
    },
    handleChange(file, list) {
      this.assetNumbers.total = list.length
    },
    upload() {
      if (this.assetNumbers.total === 0) {
        this.$message.error("请先选择素材")
        return
      }
      this.loading = true
      this.$refs.assetUpload.submit()
    },
    uploadErr(err, file, list) {
      this.assetNumbers.failed ++
      this.$notify.error({ title: '上传错误提示', message: `${file.name} 上传失败：` + err, duration: 10000 })
      this.checkUploadComplete()
    },
    uploadSuccess(res, file, list) {
      this.assetNumbers.success ++
      this.checkUploadComplete()
    },
    checkUploadComplete() {
      if (this.assetNumbers.failed + this.assetNumbers.success >= this.assetNumbers.total) {
        this.loading = false
        if (this.assetNumbers.success > 0) {
          this.$emit('upload-success')
        }
        if (this.assetNumbers.failed === 0) {
          // this.visible = false
        }
      }
    },
    cancel() {
      this.visible = false
    },
    overLimit() {
      this.$message.error("最多一次选择 10 个素材")
    }
  }
}
</script>

