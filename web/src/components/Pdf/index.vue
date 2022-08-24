<template>
  <el-dialog fullscreen :visible.sync="visible" modal-append-to-body append-to-body :show-close="false"
             class="pdf-preview-container">
    <div v-loading="isLoading" class="pdf-pages">
      <p class="page-controller">
        <span v-show="pageCount > 0">{{page + ' / ' + pageCount}}</span>
        <el-button type="danger" size="mini" icon="el-icon-close" plain
                   style="margin-left: 20px;" @click="closePanel">关闭</el-button>
      </p>
      <div class="pdf-box" :style="{width: '65%', margin: '0 auto'}">
        <preview-pdf ref="pdf" :src="pdfUrl" :page="page" @num-pages="pageCount = $event" @link-clicked="page = $event"/>
      </div>
      <i class="el-icon-arrow-left" @click="changePage('prev')"/>
      <i class="el-icon-arrow-right" @click="changePage('next')"/>
    </div>
  </el-dialog>
</template>

<script>
  import {getFile} from "@a/global"
  import PreviewPdf from 'vue-pdf'
  import axios from 'axios'
  import {getToken} from "@/utils/auth"

  export default {
    components: {
      PreviewPdf
    },
    data() {
      return {
        visible: false,
        isLoading: false,
        pdfUrl: "",
        page: 1,
        pageCount: 0,
        rotate: 0,
      }
    },
    methods: {
      previewPDF(file_id) {
        this.isLoading = true
        this.page = 1
        axios({
          method: 'get',
          url: getFile,
          responseType: 'blob',
          params: {file_id},
          headers: {'Authorization': 'Bearer ' + getToken()}
        }).then(res => {
          // 成功 - 返回文件流
          let blob = new Blob([res.data], {
            type: "application/pdf",
          })
          this.pdfUrl = URL.createObjectURL(blob)
          this.isLoading = false
          this.visible = true
        }).catch(error => {
          this.$message.error('请求异常，请确认访问的文件是否存在，或有无权限')
          this.isLoading = false
        })
      },
      changePage(t) {
        if (this.page === 1 && t === 'prev') {
          return true
        }
        if (this.pageCount !== 0 && this.page === this.pageCount && t === 'next') {
          return true
        }
        if (t === 'prev') {
          this.page --
        } else {
          this.page ++
        }
      },
      closePanel() {
        this.pdfUrl = ''
        this.visible = false
      }
    }
  }
</script>

<style lang="scss">
  .pdf-preview-container {
    .el-dialog__body,.el-dialog__header {
      padding: 0 !important;
    }
    .el-dialog__body,.pdf-pages {
      height: 100%;
    }
    .pdf-pages {
      background: #404657 !important;
      padding-bottom: 40px;
      overflow-y: auto;

      &::-webkit-scrollbar {
        /*滚动条整体样式*/
        width: 7px; /*高宽分别对应横竖滚动条的尺寸*/
        height: 0;
      }

      &::-webkit-scrollbar-thumb {
        /*滚动条里面小方块*/
        border-radius: 10px;
        background-color: rgb(136, 143, 156);
      }

      .el-icon-arrow-right,.el-icon-arrow-left {
        color: #fff;
        font-size: 30px;
        padding: 15px;
        -webkit-border-radius: 50%;
        -moz-border-radius: 50%;
        border-radius: 50%;
        cursor: pointer;
        position: fixed;
        top: 50%;
        margin-top: -23px;

        &:hover {
          background: #51596e;
        }
      }
      .el-icon-arrow-left {
        left: 5px;
      }
      .el-icon-arrow-right {
        right: 12px;
      }

      .page-controller {
        background: #404657;
        width: 100%;
        height: 50px;
        line-height: 50px;
        text-align: center;
        color: #ffffff;
        position: sticky;
        top: 0;
        z-index: 999;
      }
    }
  }
</style>