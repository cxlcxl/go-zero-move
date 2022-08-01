<template>
  <div :class="className" :style="{height: height, width: width}"/>
</template>

<script>
  import echarts from 'echarts'

  require('echarts/theme/macarons') // echarts theme
  import resize from './mixins/resize'
  import {month_between} from "../report"

  const easings = ['linear', 'bounceOut', 'quadraticOut', 'cubicInOut', 'exponentialOut', 'backOut', 'quadraticOut', 'elasticInOut']

  export default {
    mixins: [resize],
    props: {
      chartData: {
        type: Array,
        required: true
      },
      xAxis: { // 横坐标
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
        default: '350px'
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
    watch: {
      chartData: {
        deep: true,
        handler(val) {
          this.setOptions(val)
        }
      }
    },
    mounted() {
      this.$nextTick(() => {
        this.initChart()
      })
    },
    beforeDestroy() {
      if (!this.chart) {
        return
      }
      this.chart.dispose()
      this.chart = null
    },
    methods: {
      initChart() {
        this.chart = echarts.init(this.$el, 'macarons')
        this.setOptions()
      },
      setOptions() {
        this.chart.setOption({
          xAxis: {
            data: month_between(this.xAxis),
            boundaryGap: false,
            axisTick: {
              show: false
            }
          },
          legend: {
            type: 'scroll',
            padding: [3, 0],
          },
          grid: {
            left: 30,
            right: 50,
            bottom: 20,
            top: 30,
            containLabel: true
          },
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'cross'
            },
            padding: [5, 10],
            position: (point, params, dom, rect, size) => {
              return [point[0], 0]
            }
          },
          yAxis: {
            axisTick: {
              show: false
            }
          },
          series: this.formatLineData()
        })
      },
      formatLineData() {
        let series = []
        this.chartData.map(item => {
          series.push({
            name: item.name,
            smooth: true,
            type: 'line',
            data: item.data,
            animationDuration: 1500,
            animationEasing: this.easing()
          })
        })
        return series
      },
      easing() {
        return easings[Math.floor(Math.random() * easings.length)]
      }
    }
  }
</script>