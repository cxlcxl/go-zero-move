<template>
  <div v-loading="balanceLoading" class="dashboard-balance">
    <el-card v-for="item in list.accounts">
      <p class="balance-account">{{ item.account_name }}</p>
      <table class="balance-list">
        <tr v-for="balance in list.balances[item.id]">
          <td><span class="balance-account-type">{{ balance.account_type }}：</span></td>
          <td>
            <span class="balance-amount text-error">
              <vue-count-number :start-val="0" :end-val="balance.amount" :duration="2000" :decimals="2" prefix="$"/>
            </span>
          </td>
        </tr>
      </table>
    </el-card>
  </div>
</template>

<script>
import { balance } from '@a/global'
import VueCountNumber from '_vue-count-to@1.0.13@vue-count-to'

export default {
  name: 'Balance',
  components: {
    VueCountNumber
  },
  data() {
    return {
      balanceLoading: false,
      list: {
        balances: {},
        accounts: []
      }
    }
  },
  mounted() {
    this.getBalance()
  },
  methods: {
    getBalance() {
      this.balanceLoading = true
      balance().then(res => {
        this.list = res.data
        this.balanceLoading = false
      }).catch(err => {
        this.balanceLoading = false
      })
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
