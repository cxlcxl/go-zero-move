<template>
  <div v-loading="pageLoading" style="width: 100%;height: 100%;"
       :element-loading-text="loadingMsg"
       element-loading-background="rgba(0, 0, 0, 0.8)"/>
</template>

<script>

import {getAccessToken} from "@/api/account"

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
      if (this.$route.query.error) {
        this.$message.error(this.$route.query.error_description)
        return
      }
      if (!this.$route.query.authorization_code || this.$route.query.authorization_code === '') {
        //window.close()
        this.$message.error("回调请求失败，请重试")
        return
      }
      getAccessToken(this.$route.query).then(res => {
        this.$message.success("回调成功，3s 后自动关闭页面，您可手动关闭")
        this.loadingMsg = "回调成功，3s 后自动关闭页面，您可手动关闭"
        setTimeout(function () {
          //window.close()
        }, 5000)
      }).catch(err => {
        //window.close()
      })
    },
  }
}
</script>
