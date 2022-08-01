<template>
  <div v-loading="pageLoading" style="width: 100%;height: 100%;"
       :element-loading-text="loadingMsg"
       element-loading-background="rgba(0, 0, 0, 0.8)"/>
</template>

<script>

import {getAccessToken} from "@/api/marketing"

export default {
  name: "Callback",
  data() {
    return {
      pageLoading: true,
      loadingMsg: '请求回调中...'
    }
  },
  mounted() {
    this.initCallback()
  },
  methods: {
    initCallback() {
      console.log(this.$route.query)
      if (this.$route.query.error) {
        this.$message.error(this.$route.query.error_description)
        return
      }
      this.$message.success(this.$route.query.authorization_code)
      if (!this.$route.query.authorization_code || this.$route.query.authorization_code === '') {
        this.$message.error("回调请求失败，请重试")
        return
      }
      getAccessToken(this.$route.query).then(res => {
        console.log(res)
      }).catch(err => {
        console.log(err)
      })
    },
  }
}
</script>
