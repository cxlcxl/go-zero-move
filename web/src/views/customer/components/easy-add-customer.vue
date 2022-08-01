<template>
  <dialog-panel title="添加客户信息" :visible="visible" @confirm="confirm" @cancel="cancel" confirm-text="确认"
                :confirmLoading="loading" append-to-body>
    <el-form ref="form" :model="customer" label-width="110px" size="small" :rules="rules">
      <el-form-item prop="customer_name" label="客户名称">
        <el-input v-model="customer.customer_name"/>
      </el-form-item>
      <el-form-item prop="customer_type" label="客户类型">
        <el-select v-model="customer.customer_type" style="width: 100%;" placeholder="类型一经添加不可修改">
          <!--value 转化 Number 类型解决回写显示 value 而不显示 label 的问题-->
          <el-option v-for="(val, key) in customerTypes" :label="val" :value="Number(key)" v-if="Number(key) > 2"/>
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
  import {getCustomerTypes} from '@a/customer'

  export default {
    name: "EasyAddCustomer",
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        customerTypes: {},
        customer: {
          customer_name: '',
          full_name: '',
          customer_type: 3,
          state: 1,
          weight: 99,
        },
        rules: {
          customer_name: {required: true, message: '请填写客户名称'},
          customer_type: {required: true, message: '请选择客户类型'},
        }
      }
    },
    methods: {
      easyCreate() {
        getCustomerTypes().then(res => {
          this.customerTypes = res.data
          this.visible = true
        }).catch(() => {})
      },
      confirm() {
        this.$refs.form.validate(v => {
          if (v) {
            this.$emit('confirm', this.customer)
            this.visible = false
          }
        })
      },
      cancel() {
        this.$refs.form.resetFields()
        this.visible = false
      }
    }
  }
</script>
