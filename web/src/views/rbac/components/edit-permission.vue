<template>
  <dialog-panel title="权限修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="permissionForm" ref="permission_update" label-width="100px" size="small">
      <el-form-item label="权限名称" prop="desc" :rules="{required: true, message: '请填写权限名称'}">
        <el-input v-model="permissionForm.desc"/>
      </el-form-item>
      <el-form-item label="路由地址" prop="permission" :rules="{required: true, message: '请填写路由地址'}">
        <el-input v-model="permissionForm.permission" placeholder="请严格填写，示例：/api/user/list"/>
      </el-form-item>
      <el-form-item label="请求类型" prop="allow_act" :rules="{required: true, message: '请选择允许的请求类型'}">
        <el-radio-group v-model="permissionForm.allow_act" size="mini">
          <el-radio-button label="*"></el-radio-button>
          <el-radio-button label="GET"></el-radio-button>
          <el-radio-button label="POST"></el-radio-button>
          <el-radio-button label="PUT"></el-radio-button>
          <el-radio-button label="DELETE"></el-radio-button>
        </el-radio-group>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {permissionUpdate} from '@a/role'

  export default {
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        permissionForm: {
          id: 0,
          desc: '',
          permission: '',
          allow_act: '',
          parent_id: 0,
        }
      }
    },
    methods: {
      cancel() {
        this.$refs.permission_update.resetFields()
        this.visible = false
      },
      save() {
        this.$refs.permission_update.validate(v => {
          if (v) {
            this.loading = true
            permissionUpdate(this.permissionForm).then(res => {
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
