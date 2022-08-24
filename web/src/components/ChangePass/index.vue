<template>
  <el-dialog title="密码修改" :visible.sync="visible" append-to-body width="380px" v-el-drag-dialog
             :close-on-click-modal="false">
    <el-form ref="reset_pass" :model="form" label-width="100px" :rules="rules">
      <el-form-item label="原密码" prop="old_pass">
        <el-input v-model="form.old_pass" type="password" placeholder="请输入原密码" />
      </el-form-item>
      <el-form-item label="新密码" prop="pass">
        <el-input v-model="form.pass" :type="passType" placeholder="请输入新密码">
          <el-button slot="append" @click="showPwd">
            <svg-icon :icon-class="passType === 'password' ? 'eye' : 'eye-open'" />
          </el-button>
        </el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="confirmation_pass">
        <el-input v-model="form.confirmation_pass" type="password" placeholder="请确认新密码" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button :loading="loading" @click="close">取消</el-button>
      <el-button type="primary" :loading="loading" @click="handleSubmit">修改</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { resetPass } from '@a/user'
import { validPass } from '@/utils/validate'

export default {
  name: 'ChangePass',
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
      visible: false,
      loading: false,
      passType: 'password',
      form: {
        pass: '',
        confirmation_pass: '',
        old_pass: ''
      },
      rules: {
        pass: [{ required: true, message: '新密码必填' }, { validator: validatePass, trigger: 'blur' }],
        confirmation_pass: [{ required: true, message: '确认密码必填' }, { validator: confirmPass, trigger: 'blur' }],
        old_pass: [{ required: true, message: '原密码必填' }, { validator: validatePass, trigger: 'blur' }]
      }
    }
  },
  methods: {
    initPanel() {
      this.visible = true
    },
    handleSubmit() {
      this.$refs.reset_pass.validate((valid) => {
        if (valid) {
          this.loading = true
          resetPass(this.form).then(() => {
            this.loading = false
            this.$message.success('修改成功，下次请使用新密码登陆')
            this.close()
          }).catch(() => {
            this.loading = false
          })
        }
      })
    },
    close() {
      this.$refs.reset_pass.resetFields()
      this.visible = false
    },
    showPwd() {
      this.passType = this.passType === 'password' ? 'text' : 'password'
    }
  }
}
</script>
