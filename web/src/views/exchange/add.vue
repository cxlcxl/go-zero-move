<template>
  <dialog-panel title="汇率添加" confirm-text="添加" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading">
    <el-form :model="rateForm" ref="rateForm" label-width="100px" size="small">
      <el-form-item label="月份" prop="date_month" :rules="{required: true, message: '请选择月份'}">
        <el-date-picker v-model="rateForm.date_month" type="month" value-format="yyyy-MM" placeholder="选择月份"/>
      </el-form-item>
      <el-form-item label="转换币种" prop="from_currency" :rules="{required: true, message: '请选择转换币种'}">
        <el-select v-model="rateForm.from_currency" placeholder="请选择转换币种" style="width: 100%;">
          <el-option v-for="item in currencies" :label="item._name" :value="item._val" v-show="item._val !== 'USD'"/>
        </el-select>
        <i class="el-icon-info" style="color: #909399; font-size: 12px;"> 系统默认美元计算，以选择的币种转化为美元</i>
      </el-form-item>
      <el-form-item label="汇率" prop="rate" :rules="{required: true, message: '请填写汇率'}">
        <el-input-number v-model="rateForm.rate" :min="0" :step="0.1" style="width: 200px;" :precision="4"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {exchangeCreate} from '@a/exchange'

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
          date_month: '',
          rate: 1,
          from_currency: ''
        },
      }
    },
    methods: {
      cancel() {
        this.$refs.rateForm.resetFields()
        this.visible = false
      },
      add() {
        this.$refs.rateForm.validate(v => {
          if (v) {
            this.loading = true
            exchangeCreate(this.rateForm).then(res => {
              this.$message.success("创建成功")
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
