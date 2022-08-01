<template>
  <dialog-panel :visible="visible" title="协议/附件详情" confirm-text="确定" @cancel="close" @confirm="close" width="80%">
    <div style="padding: 0 0 20px 15px;">
      <el-table :data="agreementList" border stripe size="mini" height="250px" v-loading="agreementLoading">
        <el-table-column prop="id" label="ID" width="120" align="center"></el-table-column>
        <el-table-column prop="agreement_name" label="协议/附件名称" width="350"></el-table-column>
        <el-table-column prop="desc" label="描述"></el-table-column>
        <el-table-column prop="created_at" label="上传时间" width="150"></el-table-column>
        <el-table-column align="center" label="操作" width="160">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="viewAgreement(scope.row)">查看协议</el-button>
              <el-button type="danger" plain @click.native.prevent="removeAgreement(scope.$index, scope.row)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <pdf ref="pdf"/>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {agreementDestroy} from "@a/contract"
  import Pdf from '@c/Pdf'

  export default {
    components: {
      DialogPanel, Pdf
    },
    data() {
      return {
        visible: false,
        agreementLoading: false,
        agreementList: []
      }
    },
    methods: {
      initAgreements(agreementList) {
        this.agreementList = agreementList
        this.visible = true
      },
      viewAgreement(row) {
        this.$refs.pdf.previewPDF(row.file_id)
      },
      removeAgreement(idx, row) {
        this.$confirm('确定删除协议 ['+ row.agreement_name +'] 么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.agreementLoading = true
          agreementDestroy(row.id).then(res => {
            this.$message.success('删除成功')
            this.$emit('success')
            this.agreementList.splice(idx, 1)
            this.agreementLoading = false
          }).catch(() => {
            this.agreementLoading = false
          })
        }).catch(() => {})
      },
      close() {
        this.agreementList = []
        this.visible = false
      }
    }
  }
</script>
