<template>
  <drawer title="角色创建" confirm-text="添加" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading">
    <el-form :model="roleForm" ref="roleForm" label-width="100px" size="small">
      <el-form-item label="角色名称" prop="role_name" :rules="{required: true, message: '请填写角色名称'}">
        <el-input v-model="roleForm.role_name"/>
      </el-form-item>
      <el-form-item label="选择权限" prop="permissions">
        <el-tree :data="permissions" show-checkbox node-key="id" :props="defaultProps"
                 @check-change="handleCheck" ref="permission_tree"/>
      </el-form-item>
    </el-form>
  </drawer>
</template>

<script>
  import Drawer from '@c/Drawer'
  import {roleCreate} from '@a/role'
  import {removeArrayItem} from '@/utils'

  export default {
    components: {
      Drawer
    },
    props: {
      permissions: Array,
      roleType: Object
    },
    data() {
      return {
        visible: false,
        loading: false,
        roleForm: {
          role_name: '',
          sys: 0,
          permissions: []
        },
        defaultProps: {
          children: 'children',
          label: 'desc'
        }
      }
    },
    methods: {
      cancel() {
        this.$refs.roleForm.resetFields()
        this.$set(this.roleForm, 'permissions', [])
        this.$refs.permission_tree.setCheckedKeys([])
        this.visible = false
      },
      add() {
        this.$refs.roleForm.validate(v => {
          if (v) {
            this.loading = true
            roleCreate(this.roleForm).then(res => {
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
      },
      // 	共三个参数，依次为：传递给 data 属性的数组中该节点所对应的对象、节点本身是否被选中、节点的子树中是否有被选中的节点
      handleCheck(data, selfIsChecked, hasBrotherChecked) {
        if (selfIsChecked && !this.roleForm.permissions.includes(data.permission)) {
          this.roleForm.permissions.push(data.permission)
        } else {
          this.roleForm.permissions = removeArrayItem(this.roleForm.permissions, data.permission)
        }
      }
    }
  }
</script>
