<template>
  <div class="rootdiv">
    <headerbar></headerbar>

    <div class="system">
      <thermal></thermal>
      <cpumem></cpumem>
      <net></net>
      <disk></disk>
    </div>

    <footerbar></footerbar>
  </div>
</template>

<script scoped src="./system.js"></script>
<style scoped src="./system.css"></style>
