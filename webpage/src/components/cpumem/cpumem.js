import echarts from 'echarts'

import {getSysInfo} from '../../models/cpumem.js'
import {color1, color2, lineItemStyle, isupdated, areaStyleMem, areaStyleSwap, tooltipPercent, tooltipDefault} from '../../consts/echarts.js'

export default {
  name: 'cpumem',
  data () {
    return {}
  },
  created () {
  },
  mounted(){
    this.cpuchart = echarts.init(this.$refs.cpuusage)
    this.memchart = echarts.init(this.$refs.memusage)
    this.loadchart = echarts.init(this.$refs.loadusage)
    window.addEventListener("resize", this.cpuchart.resize)
    window.addEventListener("resize", this.memchart.resize)
    window.addEventListener("resize", this.loadchart.resize)

    this.cpuchart.showLoading({text: '加载中...'})
    this.memchart.showLoading({text: '加载中...'})
    this.loadchart.showLoading({text: '加载中...'})
  },
  beforeDestroy () {
    window.removeEventListener("resize", this.cpuchart.resize)
    window.removeEventListener("resize", this.memchart.resize)
    window.removeEventListener("resize", this.loadchart.resize)
    this.cpuchart.clear()
    this.memchart.clear()
    this.loadchart.clear()
    this.cpuchart.dispose()
    this.memchart.dispose()
    this.loadchart.dispose()
  },
  computed: {
    listencpumeminfo() {
      return this.$store.state.cpumeminfo.latestx
    }
  },
  watch:{
    listencpumeminfo:function(newd, old){
      if (newd <= old) {
        return
      }
      var _cpumem = this.$store.state.cpumeminfo
      this.drawLineCpu(_cpumem)
    }
  },
  methods: {
    drawLineCpu(info){
      var chartoptcpu = {
        title: { text: 'CPU使用率' },
        animation: false,
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['CPU']},
        tooltip: tooltipPercent,
        xAxis: {boundaryGap: false, data: info.updateat},
        yAxis: {max: 100},
        series: [{name: 'CPU', symbol: 'none', type: 'line', color: ['#66AEDE'], smooth: 0.3, data: info.cpuusage, itemStyle: lineItemStyle}]
      }
      this.cpuchart.hideLoading()
      this.cpuchart.setOption(chartoptcpu)

      var chartoptmem = {
        title: {text: '内存使用率'},
        animation: false,
        xAxis: {boundaryGap: false, data: info.updateat},
        yAxis: {max: 100},
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['内存', 'SWAP']},
        tooltip: tooltipPercent,
        series: [{name: '内存', symbol: 'none', type: 'line', itemStyle: lineItemStyle, color: color2[0], smooth: 0.3, areaStyle: areaStyleMem, data: info.memusage},
          {name: 'SWAP', symbol: 'none', type: 'line', itemStyle: lineItemStyle, color: color2[1], smooth: 0.3, areaStyle: areaStyleSwap, data: info.swapusage}]
      }
      this.memchart.hideLoading()
      this.memchart.setOption(chartoptmem)

      var cpunum = 1
      if (this.$store.state.osinfo.numcpu) {
        cpunum = this.$store.state.osinfo.numcpu
      }
      var chartoptload = {
        title: {text: '系统平均负载 (' + cpunum + '核)'},
        animation: false,
        calculable: true,
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['load_1m', 'load_5m', 'load_15m']},
        tooltip: tooltipDefault,
        xAxis: {boundaryGap: false, data: info.updateat},
        yAxis: {max: parseInt(info.maxavg) + 2},
        series: [
          {name: 'load_1m', symbol: 'none', type: 'line', color: color1[0], smooth: 0.3, data: info.avg1min, itemStyle: lineItemStyle},
          {name: 'load_5m', symbol: 'none', type: 'line', color: color1[1], smooth: 0.3, data: info.avg5min, itemStyle: lineItemStyle},
          {name: 'load_15m', symbol: 'none', type: 'line', color: color1[2], smooth: 0.3, data: info.avg15min, itemStyle: lineItemStyle}
        ]
      }
      this.loadchart.hideLoading()
      this.loadchart.setOption(chartoptload)
    }
  }
}
