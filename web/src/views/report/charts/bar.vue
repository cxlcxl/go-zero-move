<template>
  <div :class="className" :style="{height: height, width: width}"></div>
</template>

<script>
  import echarts from 'echarts'

  require('echarts/theme/macarons') // echarts theme
  import resize from './mixins/resize'

  export default {
    mixins: [resize],
    name: "Bar", // 柱状图
    props: {
      chartData: {
        type: Array,
        required: true
      },
      xAxis: {
        type: Array,
        required: true
      },
      className: {
        type: String,
        default: 'chart'
      },
      width: {
        type: String,
        default: '100%'
      },
      height: {
        type: String,
        default: '400px'
      },
      autoResize: {
        type: Boolean,
        default: true
      }
    },
    data() {
      return {
        chart: null
      }
    },
    mounted() {
      this.$nextTick(() => {
        this.initChart()
      })
    },
    methods: {
      initChart() {
        this.chart = echarts.init(this.$el, 'macarons')
        this.setOptions()
      },
      setOptions() {
        this.chart.setOption({
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'shadow'
            }
          },
          grid: {
            left: 30,
            right: 50,
            bottom: 20,
            top: 30,
            containLabel: true
          },
          legend: {
            data: this.xAxis
          },
          toolbox: {
            show: true,
            orient: 'vertical',
            left: 'right',
            top: 'center',
            feature: {
              magicType: {show: true, type: ['line', 'bar']}
            }
          },
          xAxis: [
            {
              type: 'category',
              axisTick: {show: false},
              data: this.xAxis
            }
          ],
          yAxis: [
            {
              type: 'value'
            }
          ],
          series: this.formatBarData()
        })
      },
      formatBarData() {
        let series = []
        this.chartData.map(item => {
          series.push({
            name: item.name,
            barGap: 0,
            type: 'bar',
            data: item.data,
            emphasis: {
              focus: 'series'
            },
          })
        })
        return series
      },
    }
  }
</script>

<style scoped>

</style>