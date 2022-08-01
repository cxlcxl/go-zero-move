<template>
  <div v-loading="balanceLoading" class="dashboard-balance">
    <el-card>
      <table class="balance-list">
        <template v-if="update">
          <tr v-for="(item, idx) in newBalances">
            <td><el-input size="mini" placeholder="账户名称" v-model="item.account_type_name"/></td>
            <td><el-input-number size="mini" placeholder="账户余额" v-model="item.amount"/></td>
            <td><i class="el-icon-delete text-error" @click="removeLine(idx)" style="cursor: pointer;"/></td>
          </tr>
          <tr>
            <td colspan="3" align="center">
              <el-button-group class="table-operate">
                <el-button icon="el-icon-plus" @click="addBalance">增加</el-button>
                <el-button icon="el-icon-close" @click="cancelUpdate">取消</el-button>
                <el-button type="primary" icon="el-icon-check" @click="saveBalance">保存</el-button>
              </el-button-group>
            </td>
          </tr>
        </template>
        <template v-else>
          <tr v-for="balance in list">
            <td><span class="balance-account-type">{{ balance.account_type_name }}：</span></td>
            <td>
            <span class="balance-amount text-error">
              <vue-count-number :start-val="0" :end-val="balance.amount" :duration="2000" :decimals="2" prefix="$"/>
            </span>
            </td>
          </tr>
        </template>
      </table>
    </el-card>
  </div>
</template>

<script>
import { platformBalance, savePlatformBalance } from '@a/global'
import VueCountNumber from '_vue-count-to@1.0.13@vue-count-to'

export default {
  name: 'PlatformBalance',
  components: {
    VueCountNumber
  },
  data() {
    return {
      update: false,
      balanceLoading: false,
      list: [],
      newBalances: []
    }
  },
  mounted() {
    this.getBalance()
  },
  methods: {
    doUpdate() {
      this.addBalance()
      this.update = true
    },
    getBalance() {
      platformBalance().then(res => {
        this.list = res.data
      }).catch(err => {})
    },
    addBalance() {
      this.newBalances.push({account_type_name: '', amount: 0})
    },
    cancelUpdate() {
      this.newBalances = []
      this.update = false
    },
    saveBalance() {
      if (this.newBalances.length === 0) {
        this.$message.error("请填写账户余额")
        return
      }
      let pass = true
      this.newBalances.map(item => {
        if (!/^(-)?[0-9]+(\.?[0-9]+)?$/.test(item.amount) || item.account_type_name === '') {
          pass = false
        }
      })
      if (!pass) {
        this.$message.error("数据填写有误，请检查")
        return
      }
      savePlatformBalance(this.newBalances).then(res => {
        this.cancelUpdate()
        this.getBalance()
      }).catch(() => {})
    },
    removeLine(idx) {
      this.newBalances.splice(idx, 1)
    }
  }
}
</script>

<style scoped lang="scss">
.dashboard-balance {
  display: flex;

  .el-card {
    flex: 1;
    min-width: 200px;

    .balance-account {
      height: 30px;
      line-height: 30px;
      font-size: 16px;
      font-weight: 600;
      color: #6e7075;
    }

    &:not(:last-child) {
      margin-right: 15px;
    }

    .balance-list {
      margin-top: 3px;

      tr {
        height: 25px;
        line-height: 25px;

        .balance-account-type {
          display: inline-block;
          color: #86888e;
        }

        .balance-amount {
          font-weight: 600;
        }
      }
    }
  }
}
</style>
