<template>
  <dialog-panel title="添加流水" confirm-text="添加" :visible="visible" @cancel="cancel" width="900px"
                @confirm="confirm" :confirm-loading="loading">
    <el-form ref="form" :model="flowForm" :rules="rules" label-width="100px" size="mini">
      <el-form-item prop="recharge_type" label="类型">
        <el-radio-group v-model="flowForm.recharge_type">
          <el-radio-button v-for="(key,val) in rechargeTypes" :label="Number(val)">{{key}}</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item prop="amount" label="金额">
        <el-input-number v-model="flowForm.amount" :precision="2" class="w200" :step="1000"/>
      </el-form-item>
      <el-form-item prop="settlement_rate" label="返点比例(%)">
        <el-input-number v-model="flowForm.settlement_rate" :min="0" class="w200" :step="1"/>
      </el-form-item>
      <el-form-item prop="trans_date" label="交易日期">
        <el-date-picker v-model="flowForm.trans_date" type="date" placeholder="开始时间"
                        value-format="yyyy-MM-dd" style="width: 100%;"/>
      </el-form-item>
      <el-form-item prop="file_id" label="合同文件">
        <upload ref="upload" accept=".pdf,.xlsx,.xls" @success="uploadSuccess" :post-data="{module: 'recharge_platform'}"
                trigger-text="选择需上传的附件" :required="false"/>
      </el-form-item>
      <el-form-item prop="remark" label="备注">
        <el-input type="textarea" v-model="flowForm.remark"/>
      </el-form-item>
    </el-form>

    <el-button slot="operate" icon="el-icon-check" type="primary" @click="continueCreate" :loading="loading">添加并继续
    </el-button>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {rechargePlatformCreate} from "@a/recharge"
  import {parseTime} from '@/utils'
  import Upload from "@c/Upload"

  export default {
    components: {DialogPanel, Upload},
    props: {
      rechargeTypes: Object
    },
    data() {
      const checkNumber = (rule, value, callback) => {
        if (value === 0) {
          callback(new Error('不可填写 0'))
        } else {
          callback()
        }
      }
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        isContinue: false,
        flowForm: {
          amount: 0,
          recharge_type: 1,
          trans_date: parseTime(new Date(), "{y}-{m}-{d}"),
          settlement_rate: 0,
          file_id: 0,
          remark: '',
        },
        rules: {
          amount: [
            {required: true, message: '请填写金额'},
            {validator: checkNumber}
          ],
          /*settlement_rate: [
            {required: true, message: '请填写计算比例'},
            {validator: checkNumber}
          ],*/
          customer_id: {required: true, message: '请选择客户'},
          trans_date: {required: true, message: '请选择交易日期'},
          recharge_type: {required: true, message: '请选择类型'},
        }
      }
    },
    methods: {
      initCreate() {
        this.visible = true
      },
      uploadSuccess(data) {
        // 当文件上传成功后提交表单数据
        this.$set(this.flowForm, 'file_id', data.id)
        this.loading = true
        if (this.isContinue) {
          rechargePlatformCreate(this.flowForm).then(res => {
            this.$message.success('添加成功，请继续填写流水信息')
            this.$emit('success')
            this.loading = false
            this.$refs.form.resetFields()
            this.$set(this.flowForm, "accounts", [])
          }).catch(() => {
            this.loading = false
          })
        } else {
          rechargePlatformCreate(this.flowForm).then(res => {
            this.$message.success('添加成功')
            this.$emit('success')
            this.loading = false
            this.cancel()
          }).catch(() => {
            this.loading = false
          })
        }
      },
      uploadFile() {
        this.$refs.form.validate(v => {
          if (v) {
            if (this.flowForm.file_id > 0) {
              this.doCreate()
            } else {
              // 通过附件上传成功调度表单数据提交
              this.$refs.upload.submit()
            }
          }
        })
      },
      confirm() {
        this.isContinue = false
        this.uploadFile()
      },
      continueCreate() {
        this.isContinue = true
        this.uploadFile()
      },
      cancel() {
        this.loading = false
        this.$refs.form.resetFields()
        this.visible = false
      }
    }
  }
</script>
