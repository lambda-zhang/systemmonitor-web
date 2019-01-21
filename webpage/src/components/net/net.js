import echarts from 'echarts'
import {getNetIfInfo, getTcpInfo} from '../../models/net.js'
import {color1, color2, lineItemStyle, isupdated, areaStyleMem, areaStyleSwap, tooltipPercent, tooltipDefault} from '../../consts/echarts.js'

export default {
  name: 'net',
  data () {
    return {
      msg: '404 Not Found',
    }
  },
  created () {
  },
  computed: {
    listennetworkinfo() {
      return this.$store.state.networkinfo.latestx
    },
    listennetifinfo() {
      return this.$store.state.netifinfo.latestx
    }
  },
  watch:{
    listennetworkinfo:function(newd, old){
      if (newd <= old) {
        return
      }
      var _network = this.$store.state.networkinfo
      this.drawLineNetWork(_network)
    },
    listennetifinfo:function(newd, old){
      if (newd <= old) {
        return
      }
      var _netif = this.$store.state.netifinfo
      this.drawLineNetIf(_netif)
    }
  },
  mounted(){
    this.netchart = echarts.init(this.$refs.netifusage)
    this.packageschart = echarts.init(this.$refs.netifpackagesusage)
    this.tcpchart = echarts.init(this.$refs.tcpusage)
    window.addEventListener("resize", this.netchart.resize)
    window.addEventListener("resize", this.packageschart.resize)
    window.addEventListener("resize", this.tcpchart.resize)

    this.netchart.showLoading({text: '加载中...'})
    this.packageschart.showLoading({text: '加载中...'})
    this.tcpchart.showLoading({text: '加载中...'})
  },
  beforeDestroy () {
    window.removeEventListener("resize", this.netchart.resize)
    window.removeEventListener("resize", this.packageschart.resize)
    window.removeEventListener("resize", this.tcpchart.resize)
    this.netchart.clear()
    this.packageschart.clear()
    this.tcpchart.clear()
    this.netchart.dispose()
    this.packageschart.dispose()
    this.tcpchart.dispose()
  },
  methods: {
    drawLineNetIf(info){
      var legendData = []
      var yData = []
      var yData2 = []
      var index = 0
      for (var key in info.y) {
        index ++
        yData.push({name: 'in_' + key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: info.y[key].inkb, itemStyle: lineItemStyle})
        yData2.push({name: 'in_' + key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: info.y[key].inp, itemStyle: lineItemStyle})
        legendData.push('in_' + key)
        index ++
        yData.push({name: 'out_' + key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: info.y[key].outkb, itemStyle: lineItemStyle})
        yData2.push({name: 'out_' + key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: info.y[key].outp, itemStyle: lineItemStyle})
        legendData.push('out_' + key)
      }
      var chartopt = {
        title: { text: '网络流入流出速率 KBps' },
        animation: false,
        calculable: true,
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData},
        tooltip: tooltipDefault,
        xAxis: {boundaryGap: false, data: info.x},
        yAxis: {max: parseInt(info.maxkb * 1.2)},
        series: yData
      }
      this.netchart.hideLoading()
      this.netchart.setOption(chartopt)

      var chartopt2 = {
        title: { text: '网络流入流出数据包数 pps' },
        animation: false,
        calculable: true,
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData},
        tooltip: tooltipDefault,
        xAxis: {boundaryGap: false, data: info.x},
        yAxis: {max: parseInt(info.maxp * 1.2)},
        series: yData2
      }
      this.packageschart.hideLoading()
      this.packageschart.setOption(chartopt2)
    },
    drawLineNetWork(info){
      var chartopt = {
        title: { text: 'TCP连接数' },
        animation: false,
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['TCP_total', 'Established']},
        tooltip: tooltipDefault,
        xAxis: {boundaryGap: false, data: info.updateat},
        yAxis: {max: parseInt(info.maxtcp * 1.2)},
        series: [
          {name: 'TCP_total', symbol: 'none', type: 'line', color: color1[0], smooth: 0.3, data: info.tcpall, itemStyle: lineItemStyle},
          {name: 'Established', symbol: 'none', type: 'line', color: color1[1], smooth: 0.3, data: info.tcpest, itemStyle: lineItemStyle},
          {name: 'TcpListen', symbol: 'none', type: 'line', color: color1[2], smooth: 0.3, data: info.tcplis, itemStyle: lineItemStyle}
        ]
      }
      this.tcpchart.hideLoading()
      this.tcpchart.setOption(chartopt)
    }
  }
}
