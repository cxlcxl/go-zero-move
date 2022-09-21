<template>
  <div v-loading="loading" class="app-box">
    <div class="app-card-opts">
      <p style="margin-bottom: 3px;">
        <el-input class="w100" size="mini" placeholder="应用名搜索" clearable v-model="pages.app_name" @change="doSearch"/>
      </p>

      <el-input-number size="mini" :min="1" v-model="pages.page" :max="maxPage" @change="changePage"/>
    </div>
    <div :class="{'app-card': true, 'app-active': AppId === ''}" @click="handleAppChange('')">
      <img :src="allPng"/>
      <p class="app-name">全部应用</p>
    </div>
    <div :class="{'app-card': true, 'app-active': AppId === app.app_id}"
         @click="handleAppChange(app.app_id)" v-for="app in appList.list">
      <img :src="!app.icon_url ? imageUnknown : app.icon_url"/>
      <p class="app-name">{{ app.app_name }}</p>
    </div>
  </div>
</template>

<script>
import allPng from '@/assets/images/all.png'
import imageUnknown from '@/assets/images/image-unknown.png'
import { campaignAppList } from '@/api/app'

export default {
  props: {
    AppId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      allPng,
      imageUnknown,
      loading: false,
      appList: {
        total: 0,
        list: []
      },
      maxPage: 1,
      pages: {
        app_name: '',
        page: 1,
        page_size: 10
      }
    }
  },
  mounted() {
    this.getAppList()
  },
  methods: {
    handleAppChange(v) {
      this.$emit("handle-change", v)
    },
    doSearch() {
      this.pages.page = 1
      this.getAppList()
    },
    changePage(v) {
      this.pages.page = v
      this.getAppList()
    },
    getAppList() {
      this.loading = true
      campaignAppList(this.pages).then(res => {
        this.appList.list = res.data
        this.appList.total = res.total
        this.maxPage = Math.ceil(res.total/this.pages.page_size)
        this.loading = false
      }).catch(err => {
        this.loading = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.app-box {
  text-align: center;
  width: 120px;
  height: 550px;
  overflow-y: scroll;
  padding: 5px 10px 5px 13px;

  &::-webkit-scrollbar {
    /*滚动条整体样式*/
    width: 3px;
    background-color: #aaa; /* or add it to the track */
  }

  &::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    background: #d5d3d3;
    border-radius: 10px;
  }

  &::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    background: #ffffff;
  }

  .app-card {
    width: 64px;
    padding: 0 !important;
    border: 2px solid transparent;
    border-radius: 5px;
    text-align: center;
    cursor: pointer;
    margin: 0 auto 5px;

    img {
      width: 60px;
      height: 60px;
      margin: 0;
      padding: 0;
      display: block;
      border-radius: 13px;
    }

    p {
      margin: 3px 0 0 0;
      padding: 0;
      line-height: 18px;
      color: #606266;
      font-size: 12px;
      overflow: hidden;
    }

    &:hover {
      border-color: #1890ff;

      p {
        color: #1890ff;
      }
    }
  }

  .app-active {
    border-color: #1890ff !important;
    p {
      color: #1890ff !important;
    }
  }
}
</style>

<style lang="scss">
.app-card-opts {
  margin-bottom: 5px;

  .el-button--mini {
    padding: 3px 0;
    width: 30px;
  }

  .el-input-number--mini {
    width: 100px;

    .el-input-number__decrease, .el-input-number__increase {
      width: 22px !important;
    }
  }
}
</style>
