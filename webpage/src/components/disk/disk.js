import echarts from 'echarts'
import {color1, color2, lineItemStyle, isupdated, areaStyleMem, areaStyleSwap, tooltipPercent, tooltipDefault} from '../../consts/echarts.js'
import {getDiskInfo} from '../../models/disk.js'

export default {
  name: 'disk',
  data () {
    return {}
  },
  created () {
  },
  computed: {
    listendiskinfo() {
      return this.$store.state.diskinfo.latestx
    }
  },
  watch:{
    listendiskinfo:function(newd, old){
      if (newd <= old) {
        return
      }
      var _disk = this.$store.state.diskinfo
      this.drawLinedisk(_disk)
    }
  },
  mounted(){
    this.diskusage = echarts.init(this.$refs.diskusage)
    this.diskioKBps = echarts.init(this.$refs.diskiokbps)
    this.diskioKRps = echarts.init(this.$refs.diskiokrps)
    window.addEventListener("resize", this.diskusage.resize)
    window.addEventListener("resize", this.diskioKBps.resize)
    window.addEventListener("resize", this.diskioKRps.resize)

    this.diskusage.showLoading({text: '加载中...'})
    this.diskioKBps.showLoading({text: '加载中...'})
    this.diskioKRps.showLoading({text: '加载中...'})
  },
  beforeDestroy () {
    window.removeEventListener("resize", this.diskusage.resize)
    window.removeEventListener("resize", this.diskioKBps.resize)
    window.removeEventListener("resize", this.diskioKRps.resize)
    this.diskusage.clear()
    this.diskioKBps.clear()
    this.diskioKRps.clear()
    this.diskusage.dispose()
    this.diskioKBps.dispose()
    this.diskioKRps.dispose()
  },
  methods: {
    drawLinedisk(info){
      var legendData1 = []
      var legendData2 = []
      var yData1 = []
      var yData2 = []
      var yData3 = []
      var index = 0
      for (var key in info.y) {
        yData1.push({name: key, color: color1[index % color1.length], data: info.y[key].usage, symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
        legendData1.push(key)

        yData2.push({name: key + '_read', color: color1[index % color1.length], data: info.y[key].rkbps, symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
        yData2.push({name: key + '_write', color: color2[index % color2.length], data: info.y[key].wkbps, symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
        yData3.push({name: key + '_read', color: color1[index % color1.length], data: info.y[key].rrps, symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
        yData3.push({name: key + '_write', color: color2[index % color2.length], data: info.y[key].wrps, symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
        legendData2.push(key + '_read')
        legendData2.push(key + '_write')
        index ++
      }

      var chartoptusage = {
        title: {text: '硬盘使用率'},
        animation: false,
        xAxis: {boundaryGap: false, data: info.x},
        yAxis: {max: 100},
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData1},
        tooltip: tooltipPercent,
        series: yData1
      }
      this.diskusage.hideLoading()
      this.diskusage.setOption(chartoptusage)

      var chartoptIOKBps = {
        title: {text: '磁盘读写数据量(KBps)'},
        animation: false,
        xAxis: {boundaryGap: false, data: info.x},
        yAxis: {max: parseInt(info.maxkbps * 1.2)},
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData2},
        tooltip: tooltipDefault,
        series: yData2
      }
      this.diskioKBps.hideLoading()
      this.diskioKBps.setOption(chartoptIOKBps)
   
      var chartoptIORps = {
        title: {text: '磁盘读写请求次数(Rps)'},
        animation: false,
        xAxis: {boundaryGap: false, data: info.x},
        yAxis: {max: parseInt(info.maxrps * 1.2)},
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData2},
        tooltip: tooltipDefault,
        series: yData3
      }
      this.diskioKRps.hideLoading()
      this.diskioKRps.setOption(chartoptIORps)
    }
  }
}
