<template>
  <el-drawer :visible.sync="visible" title="阶梯比例说明" append-to-body :modal="false">
    <table class="terms-help-table" cellspacing="0" cellpadding="0">
      <tr>
        <th width="100">阶梯值</th>
        <th width="60">比例</th>
        <th>说明</th>
      </tr>
      <tr>
        <td>0</td>
        <td>20</td>
        <td>第一层为基础比例，阶梯值不可填</td>
      </tr>
      <tr>
        <td>50000</td>
        <td>0</td>
        <td>相当于：0 ~ 5w 比例为 0%</td>
      </tr>
      <tr>
        <td>100000</td>
        <td>5</td>
        <td>相当于：50001 ~ 10w 比例为 5%</td>
      </tr>
      <tr>
        <td>150000</td>
        <td>10</td>
        <td>相当于：100001 ~ 15w 比例为 10%</td>
      </tr>
      <tr>
        <td>9999999999</td>
        <td>15</td>
        <td>相当于：150001 ~ 无穷 比例为 15%（理论上 9999999999 不会达到）</td>
      </tr>
    </table>
    <div style="text-align: center;">
      <el-button @click="closeHelp" icon="el-icon-close">关闭</el-button>
    </div>
  </el-drawer>
</template>

<script>
  export default {
    name: "TermsHelp",
    data() {
      return {
        visible: false
      }
    },
    methods: {
      closeHelp() {
        this.visible = false
      }
    }
  }
</script>

<style lang="scss">
  .terms-help-table {
    color: #909399;
    font-size: 14px;
    border: 1px solid #eee;
    margin: 0 10px 15px 10px;

    tr td,th {
      border-left: 1px solid #eee;
      border-top: 1px solid #eee;
      padding: 3px 5px;
    }
    th {
      background: #fafafa;
      height: 25px;
      line-height: 25px;
    }
  }
</style>