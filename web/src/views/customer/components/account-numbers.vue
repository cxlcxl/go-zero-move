<template>
  <el-dialog title="账户信息" :visible.sync="visible" width="500px" append-to-body modal-append-to-body
             :close-on-click-modal="false" :close-on-press-escape="false" v-el-drag-dialog :show-close="false">

    <ul class="account-list">
      <li v-for="account in accounts">
        <div class="title">{{account.account_name}}</div>
        <div class="content">
          <template v-if="account.account_id > 0">{{account.amount}}</template>
          <template v-else>
            尚未开通，<a class="open-account" @click="openAccount(account)">立即开通</a>
          </template>
        </div>
      </li>
    </ul>

    <div slot="footer" class="dialog-footer">
      <el-button @click="handleCancel" :loading="confirmLoading" icon="el-icon-close">关闭</el-button>
    </div>
  </el-dialog>
</template>

<script>
  import {getAccounts, openAccountNumber} from '@a/account'

  export default {
    name: "UserAccount",
    components: {

    },
    data() {
      return {
        visible: false,
        confirmLoading: false,
        accounts: [],
        account_id: 0
      }
    },
    methods: {
      initAccountInfo(account_id) {
        this.account_id = account_id
        getAccounts(this.account_id).then(res => {
          this.accounts = res.data
          this.visible = true
        }).catch(() => {})
      },
      handleCancel() {
        this.visible = false
      },
      openAccount(account) {
        this.$confirm('确定开启此账户么?', '操作提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'success'
        }).then(() => {
          openAccountNumber({account_id: this.account_id, account_type: account.account_type}).then(res => {
            this.initAccountInfo(this.account_id)
          }).catch(() => {})
        }).catch(() => {})
      }
    }
  }
</script>

<style scoped lang="scss">
.account-list {
  li {
    height: 30px;
    line-height: 30px;
    font-size: 14px;
    display: flex;
    margin-bottom: 5px;

    .title {
      display: inline-block;
      width: 120px;
      text-align: right;
      font-weight: 600;
      color: #606266;
      margin-right: 10px;
    }
    .content {
      flex: 1;
      /*color: #f4516c;*/
      border-bottom: 1px dashed #dfe6ec;

      .open-account {
        color: #1890ff;
        text-decoration: underline;
      }
    }
  }
}
</style>