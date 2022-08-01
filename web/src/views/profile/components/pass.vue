<template>
  <el-form :model="form" ref="form" label-width="120px" :rules="passRules">
    <el-form-item label="旧密码" prop="old_pass">
      <el-input v-model.trim="form.old_pass" type="password"/>
    </el-form-item>
    <el-form-item label="新密码" prop="pass">
      <el-input v-model.trim="form.pass" type="password"/>
    </el-form-item>
    <el-form-item label="确认新密码" prop="confirmation_pass">
      <el-input v-model.trim="form.confirmation_pass" type="password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit" :loading="loading">修改</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
  import { validPass } from '@/utils/validate'
  import { resetPass } from '@a/user'

  export default {
    name: 'Pass',
    data() {
      var validatePass = (rule, value, callback) => {
        if (!validPass(value)) {
          callback(new Error('密码限字母开头 6-18 位的字符'))
        } else {
          callback()
        }
      }
      var confirmPass = (rule, value, callback) => {
        if (this.form.pass !== value) {
          callback(new Error('两次输入密码不一致'))
        } else {
          callback()
        }
      }
      return {
        loading: false,
        form: {
          old_pass: '',
          pass: '',
          confirmation_pass: '',
        },
        passRules: {
          old_pass: [{required: true, message: '请填写旧密码'}, { validator: validatePass, trigger: 'blur' }],
          pass: [{required: true, message: '请填写新密码'}, { validator: validatePass, trigger: 'blur' }],
          confirmation_pass: [{required: true, message: '请确认新密码'}, { validator: confirmPass, trigger: 'blur' }],
        }
      }
    },
    methods: {
      submit() {
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.loading = true
            resetPass(this.form).then(() => {
              this.loading = false
              this.$message.success('修改成功，下次请使用新密码登陆')
              this.$refs.form.resetFields()
            }).catch(() => {
              this.loading = false
            })
          }
        })
      }
    }
  }
</script>
