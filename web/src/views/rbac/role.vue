<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.role_name" class="w150" clearable placeholder="角色名称"/>
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" placeholder="用户状态" class="w110">
            <el-option :key="1" label="正常" :value="1"/>
            <el-option :key="0" label="停用" :value="0"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="getRoleList">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24" style="margin-bottom: 15px;">
      <el-button icon="el-icon-plus" size="mini" type="primary" @click="addRole">新增角色</el-button>
    </el-col>
    <el-col :span="24">
      <el-table ref="user" v-loading="loadings.pageLoading" :data="roles.list" highlight-current-row stripe border
                size="mini">
        <el-table-column prop="id" label="ID" width="110"/>
        <el-table-column prop="role_name" label="角色名称"/>
        <el-table-column align="center" label="操作" fixed="right" width="130">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>

    <role-create ref="roleCreate" :permissions="permissions" @success="getRoleList"/>
    <role-update ref="roleUpdate" :permissions="permissions" @success="getRoleList"/>
  </el-row>
</template>

<script>
  import {roleList, permissionList, roleDestroy} from '@a/role'
  import RoleCreate from './components/add-role'
  import RoleUpdate from './components/edit-role'

  export default {
    name: 'RoleList',
    components: {
      RoleCreate, RoleUpdate
    },
    data() {
      return {
        loadings: {
          changeLoading: false,
          pageLoading: false
        },
        roles: {
          list: []
        },
        search: {
          role_name: '',
          state: 1
        },
        permissions: []
      }
    },
    mounted() {
      this.getRoleList()
    },
    methods: {
      getRoleList() {
        this.loadings.pageLoading = true
        roleList(this.search).then(res => {
          this.roles.list = res.data
          this.loadings.pageLoading = false
        }).catch(() => {
          this.loadings.pageLoading = false
        })
      },
      editRow(row) {
        if (this.permissions.length === 0) {
          this.getPermissions()
        }

        this.$refs.roleUpdate.initUpdate(row)
      },
      destroyRow(row) {
        this.$confirm('此操作会同步删除已分配的权限, 是否继续?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.loadings.pageLoading = true
          roleDestroy(row.id).then(() => {
            this.$message.success('删除成功')
            this.getRoleList()
          }).catch(err => {
            this.loadings.pageLoading = false
          })
        }).catch(() => {
        })
      },
      getPermissions() {
        permissionList().then(res => {
          this.permissions = res.data
        }).catch(err => {
        })
      },
      addRole() {
        if (this.permissions.length === 0) {
          this.getPermissions()
        }
        this.$refs.roleCreate.visible = true
      }
    }
  }
</script>
