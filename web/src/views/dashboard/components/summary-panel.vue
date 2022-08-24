<template>
  <el-card>
    <div class="panel-container">
      <section>
        <slot />
      </section>

      <div class="number text-error">
        <vue-count-number
          :start-val="0"
          :end-val="val"
          :duration="2000"
          class="card-panel-num"
          :decimals="2"
          :prefix="prefix"
        />
      </div>
    </div>
  </el-card>
</template>

<script>
import VueCountNumber from '_vue-count-to@1.0.13@vue-count-to'

export default {
  name: 'SummaryPanel',
  components: {
    VueCountNumber
  },
  props: {
    val: Number,
    prefix: {
      default: '$',
      type: String
    }
  }
}
</script>

<style scoped lang="scss">
.panel-container {
  display: flex;
  font-size: 16px;
  color: #606266;

  .number {
    flex: 1;
    text-align: right;
    font-weight: 600;

    .card-panel-num {
      font-size: 16px;
    }
  }
}
</style>
