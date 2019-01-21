import cpumem from '@/components/cpumem/cpumem.vue'
import net from '@/components/net/net.vue'
import disk from '@/components/disk/disk.vue'
import thermal from '@/components/thermal/thermal.vue'
import headerbar from '@/components/headerbar/headerbar.vue'
import footerbar from '@/components/footerbar/footerbar.vue'

export default {
  name: 'system',
  data () {
    return {
      tscpu: 0,
    }
  },
  components: {
      headerbar: headerbar,
      footerbar: footerbar,
      cpumem: cpumem,
      net: net,
      disk: disk,
      thermal: thermal
  },
  created () {
  },
  mounted(){
  },
  beforeDestroy () {
  },
  methods: {
  }
}
