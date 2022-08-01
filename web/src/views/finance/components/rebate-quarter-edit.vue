<template>
  <dialog-panel title="转账流水修改" confirm-text="保存" :visible="visible" @cancel="cancel" width="600px"
                @confirm="confirm" :confirm-loading="loading">
    <el-form ref="form" :model="flowForm" label-width="120px" size="mini">
      <el-form-item label="金额">
        <el-input :value="flowForm.amount" disabled/>
      </el-form-item>
      <el-form-item label="返利金季度归属" v-if="isRebate">
        <el-select v-model="flowForm.year" placeholder="归属年份" style="width: 100px; margin-right: 10px;">
          <el-option v-for="year in years" :label="year" :value="year"/>
        </el-select>
        <el-select v-model="flowForm.quarter" placeholder="归属季度">
          <el-option label="Q1" :value="1"/>
          <el-option label="Q2" :value="2"/>
          <el-option label="Q3" :value="3"/>
          <el-option label="Q4" :value="4"/>
        </el-select>
      </el-form-item>
      <el-form-item prop="remark" label="追加备注">
        <el-input type="textarea" v-model="flowForm.remark" placeholder="请在此填写追加的备注（追加在原备注信息后）"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {transPlatformUpdate} from "@a/trans"

  export default {
    components: {DialogPanel},
    data() {
      return {
        visible: false,
        loading: false,
        isRebate: false,
        profit: {},
        flowForm: {
          id: '',
          amount: 0,
          year: '',
          quarter: '',
          remark: '',
        }
      }
    },
    computed: {
      years() {
        let y = [], d = new Date()
        for (let i = 2021; i <= d.getFullYear(); i++) {
          y.push(i)
        }
        return y
      }
    },
    methods: {
      initUpdate(row) {
        this.isRebate = row.user_account === 'URGE'
        this.$set(this.flowForm, 'id', row.id)
        this.$set(this.flowForm, 'amount', row.amount)
        if (row.profit_quarter) {
          let d = row.profit_quarter.split('-')
          this.$set(this.flowForm, 'year', Number(d[0]))
          this.$set(this.flowForm, 'quarter', Number(d[1]))
        } else {
          this.$set(this.flowForm, 'year', '')
          this.$set(this.flowForm, 'quarter', '')
        }

        this.visible = true
      },
      confirm() {
        this.$refs.form.validate(v => {
          if (v) {
            this.loading = true
            if (this.flowForm.year === '' || this.flowForm.quarter === '') {
              this.flowForm.year =  0
              this.flowForm.quarter = 0
            }
            transPlatformUpdate(this.flowForm).then(res => {
              this.$message.success('修改成功')
              this.$emit('success')
              this.cancel()
            }).catch(() => {
              this.loading = false
            })
          }
        })
      },
      cancel() {
        this.loading = false
        this.$refs.form.resetFields()
        this.visible = false
      }
    }
  }
</script>
