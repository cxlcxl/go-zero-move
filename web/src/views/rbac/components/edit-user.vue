<template>
  <dialog-panel title="用户修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="userForm" ref="userForm" label-width="120px" size="small" :rules="userRules">
      <el-form-item label="用户名称" prop="username">
        <el-input v-model="userForm.username" placeholder="请填写用户名"/>
      </el-form-item>
      <el-form-item label="邮箱地址" prop="email">
        <el-input v-model="userForm.email" placeholder="请填写邮箱，登陆使用"/>
      </el-form-item>
      <el-form-item label="手机号" prop="mobile">
        <el-input v-model="userForm.mobile" placeholder="请填写手机号码"/>
      </el-form-item>
      <el-form-item label="角色" prop="role_id">
        <el-select v-model="userForm.role_id" placeholder="请选择" style="width: 100%;" clearable>
          <el-option v-for="item in roles" :key="item.id" :label="item.role_name" :value="item.id"/>
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-switch v-model="userForm.state" :active-value="1" :inactive-value="0"/>
      </el-form-item>
      <el-form-item label="登录密码" prop="pass">
        <el-input v-model="userForm.pass" placeholder="字母开头，数字特殊字符 [@.&!#?,%$] 的 6 - 18 位"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {userUpdate} from '@a/user'
  import {validPass, validEmail} from "@/utils/validate"

  export default {
    components: {
      DialogPanel
    },
    props: {
      roles: {
        default: () => [],
        type: Array
      }
    },
    data() {
      var checkPass = (rule, value, callback) => {
        if (value === '') {
          callback()
        } else {
          if (!validPass(value)) {
            callback(new Error('密码格式不符合要求'));
          } else {
            callback()
          }
        }
      }
      var checkEmail = (rule, value, callback) => {
        if (!validEmail(value)) {
          callback(new Error('邮箱格式不正确'));
        } else {
          callback()
        }
      }
      return {
        visible: false,
        loading: false,
        remoteLoading: false,
        userForm: {
          id: 0,
          role_id: 0,
          username: '',
          email: '',
          mobile: '',
          state: 1,
          pass: '',
        },
        userRules: {
          username: {required: true, message: '请填写用户名称'},
          email: [{required: true, message: '请填写邮箱'}, {validator: checkEmail}],
          role_id: {required: true, message: '请选择角色'},
          pass: {validator: checkPass},
        },
      }
    },
    methods: {
      initUpdate(row) {
        this.userForm = row
        this.visible = true
      },
      cancel() {
        this.$refs.userForm.resetFields()
        this.visible = false
      },
      save() {
        this.$refs.userForm.validate(v => {
          if (v) {
            this.loading = true
            userUpdate(this.userForm).then(res => {
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
