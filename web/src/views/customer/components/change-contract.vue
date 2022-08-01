<template>
  <dialog-panel title="关联合同修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="changeForm" ref="changeForm" label-width="120px" size="small" :rules="contractRules">
      <el-form-item label="新合同" prop="contract_id">
        <template v-if="Array.isArray(contracts) && contracts.length > 0">
          <el-select v-model="changeForm.contract_id" placeholder="请选择合同" style="width: 100%;">
            <el-option v-for="item in contracts" :label="item.contract_name" :value="item.id"/>
          </el-select>
        </template>
        <span v-else>此账号所属客户暂无合同，请先上传</span>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {changeContract} from '@a/account'
  import {hasManyContract} from '@a/customer'

  export default {
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        changeForm: {
          id: 0,
          contract_id: '',
        },
        contractRules: {
          contract_id: {required: true, message: '请选择新合同'},
        },
        contracts: []
      }
    },
    methods: {
      initUpdate(row) {
        this.changeForm.id = row.id
        hasManyContract(row.customer_id).then(res => {
          this.contracts = res.data
          this.visible = true
        }).catch(() => {})
      },
      cancel() {
        this.$refs.changeForm.resetFields()
        this.contracts = []
        this.visible = false
      },
      save() {
        this.$refs.changeForm.validate(v => {
          if (v) {
            this.loading = true
            changeContract(this.changeForm).then(res => {
              this.$message.success("修改成功")
              this.$emit('success')
              this.loading = false
              this.cancel()
            }).catch(err => {
              this.loading = false
              console.log(err)
            })
          } else {
            return false
          }
        })
      }
    }
  }
</script>
