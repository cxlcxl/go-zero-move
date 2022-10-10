<template>
  <el-upload ref="upload" :action="UploadUrl" style="text-align: left;" :limit="limit" :on-exceed="fileOverflow"
             :file-list="fileList" :auto-upload="false" :headers="headers" :data="PostData" :accept="accept"
             :on-success="uploadSuccess" :on-error="uploadError" :before-upload="beforeUpload" :on-change="onChange"
             :on-remove="removeFile" multiple>
    <el-button icon="el-icon-paperclip" slot="trigger" size="mini">{{triggerText}}</el-button>

    <slot name="opt"/>
    <div slot="tip" class="el-upload__tip">{{tip}}</div>
  </el-upload>
</template>

<script>
  import {getToken} from "@/utils/auth"

  export default {
    props: {
      UploadUrl: {
        type: String,
        required: true
      },
      PostData: {
        type: Object,
        required: true
      },
      accept: {
        type: String,
        default: '.png,.jpg,.jpeg,.pdf,.docx,.xlsx'
      },
      limit: {
        type: Number,
        default: 1
      },
      triggerText: {
        type: String,
        default: '选取文件'
      },
      tip: {
        type: String,
        default: ''
      },
      required: {
        type: Boolean,
        default: true
      }
    },
    data() {
      return {
        headers: {
          Authorization: 'Bearer ' + getToken()
        },
        fileList: [],
        selectedFiles: []
      }
    },
    methods: {
      fileOverflow() {
        this.$message.error("文件的最大上传数量为：" + this.limit)
      },
      uploadSuccess(res, file, list) {
        if (res.code !== 200) {
          this.$message.error("上传失败：" + res.message)
        } else {
          this.$emit('success', res.data)
        }
      },
      uploadError(err, file, list) {
        this.$message.error(JSON.parse(err.message).message)
      },
      beforeUpload(file) {
        return this.$emit('before', file)
      },
      submit() {
        if (this.selectedFiles.length === 0) {
          if (this.required) {
            this.$message.error('请先选择需要上传的文件')
            return false
          } else {
            this.$emit('success', {id: 0, path: ''})
          }
        } else {
          this.$refs.upload.submit()
        }
      },
      onChange(file, fileList) {
        let idx = this.inFiles(file.uid)
        if (idx === -1) {
          this.selectedFiles.push(file)
        }
        this.$emit('change', file, fileList)
      },
      removeFile(file, fileList) {
        let idx = this.inFiles(file.uid)
        if (idx !== -1) {
          this.selectedFiles.splice(idx, 1)
        }
        this.$emit('change', file, fileList)
      },
      inFiles(uid) {
        if (this.selectedFiles.length === 0) {
          return -1
        }
        for (let i = 0;i < this.selectedFiles.length; i ++) {
          if (uid === this.selectedFiles[i].uid) {
            return i
          }
        }
        return -1
      }
    }
  }
</script>
