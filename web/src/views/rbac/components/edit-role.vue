<template>
  <drawer title="角色修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="roleForm" ref="roleForm" label-width="100px" size="small">
      <el-form-item label="角色名称" prop="role_name" :rules="{required: true, message: '请填写角色名称'}">
        <el-input v-model="roleForm.role_name"/>
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-switch v-model="roleForm.state" :active-value="1" :inactive-value="0"/>
      </el-form-item>
      <!--<el-form-item label="角色类型" prop="role_type" :rules="{required: true, message: '请选择角色类型'}">
        <el-radio v-for="(name, val) in roleType" v-model="roleForm.role_type" :label="Number(val)">{{name}}</el-radio>
      </el-form-item>
      <el-form-item label="角色描述" prop="desc">
        <el-input v-model="roleForm.desc"/>
      </el-form-item>
      <el-form-item label="选择权限" prop="permissions" :rules="{required: true, message: '请选择权限'}">
        <el-tree :data="permissions" show-checkbox node-key="permission" :props="defaultProps"
                 @check-change="handleCheck" ref="tree" :default-checked-keys="roleForm.permissions"/>
      </el-form-item>-->
    </el-form>
  </drawer>
</template>

<script>
  import Drawer from '@c/Drawer'
  import {roleUpdate, rolePermissions} from '@a/role'
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
          id: 0,
          role_name: '',
          state: 1,
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
      // Tips：el-tree 默认选中有bug，需设置 setCheckedNodes
      initUpdate(row) {
        this.$set(this.roleForm, 'id', row.id)
        this.$set(this.roleForm, 'role_name', row.role_name)
        this.$set(this.roleForm, 'state', row.state)
        this.visible = true
        /*Promise.all([
          this.getRolePermissions(row.id)
        ]).then(res => {
          this.$set(this.roleForm, 'permissions', res[0])


          this.$refs.tree.setCheckedNodes(res[0])
        }).catch(err => {
        })*/
      },
      cancel() {
        this.$refs.roleForm.resetFields()
        this.visible = false
      },
      save() {
        // this.$set(this.roleForm, 'permissions', this.$refs.tree.getCheckedKeys())
        this.$refs.roleForm.validate(v => {
          if (v) {
            this.loading = true
            roleUpdate(this.roleForm).then(res => {
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
      },
      // 	共三个参数，依次为：传递给 data 属性的数组中该节点所对应的对象、节点本身是否被选中、节点的子树中是否有被选中的节点
      handleCheck(data, selfIsChecked, hasBrotherChecked) {
        if (selfIsChecked && !this.roleForm.permissions.includes(data.permission)) {
          this.roleForm.permissions.push(data.permission)
        } else {
          this.roleForm.permissions = removeArrayItem(this.roleForm.permissions, data.permission)
        }
      },
      getRolePermissions(id) {
        return new Promise((resolve, reject) => {
          rolePermissions(id).then(res => {
            resolve(res.data)
          }).catch(() => {
            reject()
          })
        })
      }
    }
  }
</script>
