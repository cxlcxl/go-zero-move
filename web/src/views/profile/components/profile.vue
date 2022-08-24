<template>
  <el-form :model="user" ref="user" label-width="120px" :rules="formRules">
    <el-form-item label="用户名" prop="username">
      <el-input v-model.trim="user.username"/>
    </el-form-item>
    <el-form-item label="邮箱" prop="email">
      <el-input v-model.trim="user.email"/>
    </el-form-item>
    <el-form-item label="手机号" prop="mobile">
      <el-input v-model.trim="user.mobile"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit" :loading="loading">修改</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
  import {getInfo, infoUpdate} from '@a/user'

  export default {
    name: 'UserInfo',
    data() {
      return {
        loading: false,
        user: {
          username: '',
          email: '',
          mobile: '',
        },
        formRules: {
          username: {required: true, message: '请填写用户名'},
          email: [{required: true, message: '请填写邮箱号码'}, {type: 'email', message: '请填写正确的邮箱号'}],
          mobile: {required: true, message: '请填写手机号码'},
        }
      }
    },
    mounted() {
      this.getUserInfo()
    },
    methods: {
      getUserInfo() {
        getInfo().then(res => {
          this.$set(this.user, 'username', res.data.username)
          this.$set(this.user, 'email', res.data.email)
          this.$set(this.user, 'mobile', res.data.mobile)
        }).catch(() => {})
      },
      submit() {
        this.$refs.user.validate((valid) => {
          if (valid) {
            this.loading = true
            infoUpdate(this.user).then(() => {
              this.loading = false
              this.$message.success('修改成功')
            }).catch(() => {
              this.loading = false
            })
          }
        })
      }
    }
  }
</script>
