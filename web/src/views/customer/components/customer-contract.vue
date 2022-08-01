<template>
  <dialog-panel title="合同查看" :visible="visible" width="80%" @cancel="handleCancel">
    <el-table :data="contractList" border stripe size="mini" height="360px" v-loading="loading" style="margin-bottom: 20px;">
      <el-table-column prop="id" label="ID" width="120" align="center"/>
      <el-table-column prop="contract_sn" label="合同编号" width="150"/>
      <el-table-column prop="contract_name" label="合同名称" width="350"/>
      <el-table-column label="有效期" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.start_date + ' / ' + scope.row.end_date }}
        </template>
      </el-table-column>
      <el-table-column prop="desc" label="描述"/>
      <el-table-column prop="created_at" label="上传时间" width="150"/>
      <el-table-column align="center" label="操作" width="110" v-permission="'file'">
        <template slot-scope="scope">
          <el-button-group class="table-operate">
            <el-button type="primary" plain @click.native.prevent="viewContract(scope.row)">预览</el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>

    <pdf ref="pdf"/>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import Pdf from '@c/Pdf'
  import {hasManyContract} from "@a/customer"

  export default {
    name: "CustomerContract",
    components: {
      DialogPanel, Pdf
    },
    data() {
      return {
        visible: false,
        loading: false,
        contractList: []
      }
    },
    methods: {
      initView(row) {
        hasManyContract(row.id).then(res => {
          this.contractList = res.data
          this.visible = true
        }).catch(() => {})
      },
      handleCancel() {
        this.visible = false
      },
      viewContract(row) {
        this.$refs.pdf.previewPDF(row.file_id)
      },
    }
  }
</script>
