<template>
  <div class="dashboard-container">
    <panel-group />
    <transaction-table style="padding-right:8px;margin-bottom:30px;" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import PanelGroup from './components/PanelGroup'
import TransactionTable from './components/TransactionTable'

export default {
  name: 'Dashboard',
  components: {
    PanelGroup,
    TransactionTable
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  methods: {
  }
}
</script>

<style lang="scss" scoped>
.dashboard-container {
  height: 100vh;
  padding: 32px;
  background-color: rgb(240, 242, 245);
}
</style>
