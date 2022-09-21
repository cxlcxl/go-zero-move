<template>
<div>adgroup</div>
</template>

<script>
export default {
  data() {
    return {
      search: {
        app_id: ''
      }
    }
  },
  methods: {
    initData(app_id) {
      this.$set(this.search, 'app_id', app_id)
    },
  }
}
</script>

<style scoped>

</style>
