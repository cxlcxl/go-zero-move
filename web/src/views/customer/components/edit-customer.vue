<template>
  <dialog-panel title="客户信息修改" :visible="visible" @confirm="confirm" @cancel="cancel" confirm-text="保存"
                :confirmLoading="loading">
    <el-form ref="form" :model="customer" label-width="110px" size="small" :rules="rules">
      <el-form-item prop="customer_name" label="客户名称">
        <el-input v-model="customer.customer_name"/>
      </el-form-item>
      <el-form-item prop="customer_type" label="客户类型">
        <el-select v-model="customer.customer_type" style="width: 100%;" disabled>
          <!--value 转化 Number 类型解决回写显示 value 而不显示 label 的问题-->
          <el-option v-for="(val, key) in customerTypes" :label="val" :value="Number(key)"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="full_name" label="客户全称">
        <el-input v-model="customer.full_name"/>
      </el-form-item>
      <el-form-item prop="weight" label="权重">
        <el-input v-model.number="customer.weight"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {customerUpdate} from "@a/customer"

  export default {
    name: "EditCustomer",
    components: {
      DialogPanel
    },
    props: {
      customerTypes: Object
    },
    data() {
      return {
        visible: false,
        loading: false,
        customer: {
          id: 0,
          customer_name: '',
          full_name: '',
          customer_type: 0,
          weight: 99,
        },
        rules: {
          customer_name: {required: true, message: '请填写客户名称'},
          customer_type: {required: true, message: '请选择客户类型'},
        }
      }
    },
    methods: {
      initUpdate(row) {
        Object.assign(this.customer, row)
        this.visible = true
      },
      confirm() {
        this.$refs.form.validate(v => {
          if (v) {
            this.loading = true
            customerUpdate(this.customer).then(res => {
              this.loading = false
              this.$message.success('修改成功')
              this.$emit('success')
              this.cancel()
            }).catch(err => {
              this.loading = false
            })
          }
        })
      },
      cancel() {
        this.visible = false
      }
    }
  }
</script>
