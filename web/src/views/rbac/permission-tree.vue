<template>
  <div class="permission-page">
    <permission-create ref="_create" @success="getPermissions"/>
    <permission-update ref="_update" @success="getPermissions"/>

    <test-org-chart @handle-click="handleNode" :tree-data="list"/>
  </div>
</template>

<script>
  import TestOrgChart from '@c/OrgChart'
  import {permissionList, permissionDestroy} from '@a/role'
  import PermissionCreate from './components/add-permission'
  import PermissionUpdate from './components/edit-permission'

  export default {
    name: "PermissionTree",
    components: {
      TestOrgChart, PermissionCreate, PermissionUpdate
    },
    data() {
      return {
        dataLoading: false,
        list: {
          id: 0,
          desc: '根节点（不参与权限设置）',
          allow_act: '*',
          permission: '*',
          children: []
        },
      }
    },
    mounted() {
      this.getPermissions()
    },
    methods: {
      getPermissions() {
        this.dataLoading = true
        permissionList().then(res => {
          this.list.children = res.data
          this.dataLoading = false
        }).catch(err => {
          this.dataLoading = false
        })
      },
      handleNode(node, t) {
        if (t === 'create') {
          this.$refs._create.initCreate(node)
        } else if (t === 'update') {
          this.$refs._update.permissionForm = {
            id: node.id, desc: node.desc, permission: node.permission, allow_act: node.allow_act, parent_id: node.parent_id
          }
          this.$refs._update.visible = true
        } else if (t === 'delete') {
          this.destroyRow(node)
        }
      },
      destroyRow(row) {
        this.$confirm('此操作会同步删除已分配的权限, 是否继续?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.dataLoading = true
          permissionDestroy(row.id).then(() => {
            this.$message.success('删除成功')
            this.getPermissions()
          }).catch(err => {
            this.dataLoading = false
          })
        }).catch(() => {})
      }
    }
  }
</script>
