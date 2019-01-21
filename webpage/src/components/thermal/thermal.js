import echarts from 'echarts'

import {getThermalInfo} from '../../models/thermal.js'
import {getOsInfo} from '../../models/osinfo.js'
import {color1, lineItemStyle, isupdated, tooltipDefault} from '../../consts/echarts.js'

export default {
  name: 'thermal',
  data () {
    return {
      msg: '404 Not Found',
      thermal: {},
      osinfo: {
        uptime: "...",
        starttime: "...",
        cpuusage: "...",
        arch: "...",
        os: "...",
        kernel: "...",
        hostname: "...",
        numcpu: 1,
        memtotalmb: 0,
        memusage: 0,
        swaptotalmb: 0,
        swapusage: 0,
        netinfo: [],
        diskinfo: []
      },
    }
  },
  created () {
  },
  mounted(){
    this.thermalstate = echarts.init(this.$refs.thermalstate)
    window.addEventListener("resize", this.thermalstate.resize)
    this.thermalstate.showLoading({text: '加载中...'})
  },
  beforeDestroy () {
    window.removeEventListener("resize", this.thermalstate.resize)
    this.thermalstate.clear()
    this.thermalstate.dispose()
  },
  computed: {
    listenosinfo() {
      return this.$store.state.osinfo.uptime
    },
    listenthermalinfo() {
      return this.$store.state.thermalinfo.latestx
    }
  },
  watch:{
    listenosinfo:function(newd, old){
      if (newd <= old) {
        return
      }
      this.osinfo = this.$store.state.osinfo
    },
    listenthermalinfo:function(newd, old){
      if (newd <= old) {
        return
      }
      var _thermal = this.$store.state.thermalinfo
      if (!_thermal.x || !_thermal.y) {
        return
      }
      this.thermal = _thermal
      this.drawLineThermal(this.thermal)
    }
  },
  methods: {
    drawLineThermal(thermalinfo){
      var yData = []
      var legendData = []
      var index = 0

      for (var key in thermalinfo.y) {
        index ++
        yData.push({name: key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: thermalinfo.y[key], itemStyle: lineItemStyle})
        legendData.push(key)
      }

      var chartopt = {
        title: { text: '主板温度(℃)' },
        animation: false,
        calculable: true,
        legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData},
        tooltip: tooltipDefault,
        xAxis: {boundaryGap: false, data: thermalinfo.x},
        yAxis: {max: (parseInt(thermalinfo.maxTemp) + 5)},
        series: yData
      }
      this.thermalstate.hideLoading()
      this.thermalstate.setOption(chartopt)
    }
  }
}
