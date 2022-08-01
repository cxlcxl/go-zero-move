<template>
  <dialog-panel title="账号列表" :visible="visible" width="80%" @cancel="handleCancel">
    <el-table :data="accountList" border stripe size="mini" height="260px" v-loading="loading" style="margin-bottom: 20px;">
      <el-table-column prop="account_name" label="名称" width="220" show-overflow-tooltip/>
      <el-table-column label="账户信息">
        <template slot-scope="scope">
          <div v-show="Array.isArray(scope.row.account_numbers)">
            <p v-for="account in scope.row.account_numbers" class="account-list">
              {{account.account_name}}：<span>{{account.amount}}</span>
            </p>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {accountInfo} from "@a/account"

  export default {
    name: "CustomerUser",
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        accountList: []
      }
    },
    methods: {
      initView(row) {
        accountInfo(row.id).then(res => {
          this.accountList = res.data
          this.visible = true
        }).catch(() => {})
      },
      handleCancel() {
        this.visible = false
      },
    }
  }
</script>

<style scoped lang="scss">
  .account-list {
    height: 20px;
    line-height: 20px;
    font-weight: 600;
    display: inline-block;
    color: #606266;

    span {
      display: inline-block;
      width: 100px;
      color: #f4516c;
      padding-right: 10px;
    }
  }
</style>