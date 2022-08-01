<template>
  <dialog-panel title="汇率修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading">
    <el-form :model="rateForm" ref="rateForm" label-width="100px" size="small">
      <el-form-item label="月份">
        <el-input :value="rateForm.date_month" disabled/>
      </el-form-item>
      <el-form-item label="转换币种">
        <el-select v-model="rateForm.from_currency" style="width: 100%;" disabled>
          <el-option v-for="item in currencies" :label="item._name" :value="item._val"/>
        </el-select>
      </el-form-item>
      <el-form-item label="汇率" prop="rate" :rules="{required: true, message: '请填写汇率'}">
        <el-input-number v-model="rateForm.rate" :min="0" :step="0.1" style="width: 200px;" :precision="4"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {exchangeUpdate} from '@a/exchange'

  export default {
    components: {
      DialogPanel
    },
    props: {
      currencies: Array
    },
    data() {
      return {
        visible: false,
        loading: false,
        rateForm: {
          id: 0,
          date_month: '',
          rate: 1,
          from_currency: ''
        },
      }
    },
    methods: {
      initPanel(rate) {
        this.rateForm = rate
        this.visible = true
      },
      cancel() {
        this.$refs.rateForm.resetFields()
        this.visible = false
      },
      add() {
        this.$refs.rateForm.validate(v => {
          if (v) {
            this.loading = true
            exchangeUpdate(this.rateForm).then(res => {
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
