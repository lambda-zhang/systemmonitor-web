<template>
<div class="rootdiv">
  <div class="tableholder" >
    <table class="hovertable">
      <tr><th>基本信息</th><th> </th></tr>
      <tr><td class=titleclass>系统启动时间</td><td>{{osinfo.uptime}}</td></tr>
      <tr><td class=titleclass>系统启动时刻</td><td>{{osinfo.starttime}}</td></tr>
      <tr><td class=titleclass>操作系统</td><td>{{osinfo.os}}</td></tr>
      <tr><td class=titleclass>内核版本</td><td>{{osinfo.kernel}}</td></tr>
      <tr><td class=titleclass>HostName</td><td>{{osinfo.hostname}}</td></tr>
      <tr><td class=titleclass>CPU</td><td>{{osinfo.arch}} / {{osinfo.numcpu}}核</td></tr>
      <tr><td class=titleclass>平均使用率</td><td>{{osinfo.cpuusage}}%</td></tr>
      <tr><td class=titleclass>内存/SWAP</td><td>{{osinfo.memtotalmb}} ({{osinfo.memusage}}%) / {{osinfo.swaptotalmb}} ({{osinfo.swapusage}}%)</td></tr>
    </table>
  </div>
  <div class="tableholder" >
    <table class="hovertable">
      <tr><th>配置信息</th><th> </th></tr>
      <tr v-for="item in osinfo.netinfo" :key='item.name'>
        <td class=titleclass>{{item.name}} 发/收</td>
        <td>数据:({{item.outkb}} / {{item.inkb}})<br />数据包:({{item.oukp}}K / {{item.inkp}}K)</td>
      </tr>
      <tr v-for="item in osinfo.diskinfo" :key='item.name'><td class=titleclass>{{item.name}}</td><td>{{item.totalmb}} (已使用: {{item.usage}}%)</td></tr>
    </table>
  </div>

  <div class="charthloder" >
  <div class="chartclass" id="cpuusage" ref="thermalstate" ></div>
  </div>
</div>
</template>
<script scoped src="./thermal.js"></script>
<style scoped src="./thermal.css"></style>
