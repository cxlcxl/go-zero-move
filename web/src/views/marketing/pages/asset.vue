<template>
  <div class="resource-container" v-loading="loading">
    <p class="search-container">
      <el-radio-group v-model="search.asset_type" size="mini">
        <el-radio-button label="">全部</el-radio-button>
        <el-radio-button label="CREATIVE_ASSET_PICTURE">图片素材</el-radio-button>
        <el-radio-button label="CREATIVE_ASSET_VIDEO">视频素材</el-radio-button>
      </el-radio-group>&nbsp;&nbsp;
      <el-input v-model="search.width" size="mini" class="w70" clearable placeholder="素材宽"/>&nbsp;&nbsp;
      <el-input v-model="search.height" size="mini" class="w70" clearable placeholder="素材高"/>&nbsp;&nbsp;
      <el-button size="mini" @click="fetchData" type="primary" icon="el-icon-search">筛选</el-button>
      <el-button size="mini" @click="refreshResource" plain type="primary" icon="el-icon-refresh" v-show="search.app_id !== ''">从华为同步</el-button>
      <template v-if="search.app_id === '' || optRelation">
        <el-button size="mini" @click="doRelation(true)" plain icon="el-icon-setting" v-show="!optRelation">批量关联应用</el-button>
        <el-button size="mini" @click="doRelation(false)" plain icon="el-icon-setting" v-show="optRelation">取消</el-button>
        <el-button size="mini" @click="saveRelation" plain icon="el-icon-check" v-show="optRelation">提交关联</el-button>
      </template>
    </p>
    <el-card class="opt" shadow="hover" v-show="search.app_id !== '' && !optRelation">
      <div class="upload-resource" @click="uploadResource"><i class="el-icon-upload"/></div>
    </el-card>
    <el-card shadow="hover" class="resource" v-for="item in assets.list" v-if="assets.total > 0">
      <p class="asset-title">
        <i class="el-icon-check text-success" v-if="item.app_id !== ''"/>&nbsp;
        <i class="text-primary">{{item.width}}*{{item.height}}</i>
        <span class="asset-name">{{item.asset_name}}</span>
      </p>
      <img class="preview__file" :src="item.file_url" v-if="item.asset_type === 'CREATIVE_ASSET_PICTURE'"/>
      <video class="preview__file" :type="item.file_format" :src="item.file_url" v-else/>
      <div class="view" @click="showFile(item)">
        <i class="el-icon-search" v-if="item.asset_type === 'CREATIVE_ASSET_PICTURE'"/>
        <i class="el-icon-video-play" v-else/>
      </div>
      <p v-if="optRelation === true" @click="selectAsset(item.asset_id)"
         :class="{'check': true, 'check-active': relationForm.asset_ids.includes(item.asset_id)}">
        <i class="el-icon-check"/>
      </p>
      <el-popconfirm icon="el-icon-info" class="remove" icon-color="red" title="确认删除此素材吗?" v-else
                     @onConfirm="removeResource(item.asset_id, item.account_id, item.advertiser_id)">
        <i class="el-icon-delete" slot="reference"> 删除</i>
      </el-popconfirm>
    </el-card>
    <p class="pages-opt">
      <el-pagination background :current-page="search.page" :page-size="25" :total="assets.total" hide-on-single-page
                     layout="prev, pager, next, jumper, total" prev-text="上页" next-text="下页"
                     @current-change="handlePage" @size-change="handlePageSize"/>
    </p>

    <file-show ref="fileShow"/>
    <asset-upload ref="upload" @upload-success="fetchData"/>
  </div>
</template>

<script>
import FileShow from './play-show'
import AssetUpload from './asset-upload'
import { syncAsset, assetList, assetDelete, bindAsset } from '@/api/marketing-resource'

export default {
  components: {FileShow, AssetUpload},
  data() {
    return {
      loading: false,
      optRelation: false,
      search: {
        app_id: '',
        asset_type: 'CREATIVE_ASSET_PICTURE',
        width: '',
        height: '',
        page: 1,
        page_size: 25,
      },
      assets: {
        total: 25,
        list: [],
        asset_type: {}
      },
      relationForm: {
        app_id: '',
        asset_ids: []
      }
    }
  },
  methods: {
    initData(app_id) {
      this.$set(this.search, 'app_id', app_id)
      if (this.optRelation) {
        this.$set(this.relationForm, 'app_id', app_id)
        return
      }
      this.$nextTick(() => {
        this.fetchData()
      })
    },
    doSearch(assetType) {
      this.search.asset_type = assetType
      this.fetchData()
    },
    fetchData() {
      this.loading = true
      assetList(this.search).then(res => {
        this.loading = false
        this.assets = res.data
      }).catch(err => {
        this.loading = false
      })
    },
    doRelation(b) {
      this.optRelation = b
      if (b === false) {
        this.fetchData()
      }
    },
    saveRelation() {
      if (this.relationForm.app_id === '') {
        this.$message.error('需要选择应用')
        return
      }
      if (this.relationForm.asset_ids.length === 0) {
        this.$message.error('需要选择素材')
        return
      }
      this.$confirm('已关联的素材也会被修改关联，确定提交关联吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }).then(() => {
        this.loading = true
        bindAsset(this.relationForm).then(() => {
          this.loading = false
          this.doRelation(false)
        }).catch(() => {
          this.loading = false
        })
      }).catch(() => {})
    },
    selectAsset(asset_id) {
      if (this.relationForm.asset_ids.includes(asset_id)) {
        for (let i = 0; i < this.relationForm.asset_ids.length; i++) {
          if (this.relationForm.asset_ids[i] === asset_id) {
            this.relationForm.asset_ids.splice(i, 1)
          }
        }
      } else {
        this.relationForm.asset_ids.push(asset_id)
      }
    },
    removeResource(asset_id, account_id, advertiser_id) {
      this.loading = true
      assetDelete({asset_ids: [asset_id], account_id, advertiser_id}).then(res => {
        this.loading = false
        this.fetchData()
      }).catch(() => {
        this.loading = false
      })
    },
    uploadResource() {
      this.$refs.upload.initUpload(this.search.app_id)
    },
    refreshResource() {
      this.$confirm('因华为没有应用与素材的绑定关系，所以此操作只能同步关联到账户级别?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).then(() => {
        this.$message.info("素材同步中")
        this.loading = true
        syncAsset(this.search).then(() => {
          this.loading = false
          this.$message.success("素材同步完成")
        }).catch(() => {
          this.loading = false
        })
      }).catch(() => {})
    },
    showFile(f) {
      this.$refs.fileShow.initPlay(f)
    },
    handlePage(p) {
      this.search.page = p
      this.fetchData()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.fetchData()
    }
  }
}
</script>

<style lang="scss">
.resource-container {
  .w70 {
    width: 70px;
  }
  .pages-opt {
    text-align: center;
    height: 40px;
    line-height: 40px;
  }
  .search-container {
    margin-bottom: 10px;
  }
  .el-card {
    cursor: pointer;
    display: inline-block;
    width: 150px;
    height: 120px;
    margin-right: 10px;
    margin-bottom: 10px;

    .el-card__body {
      padding: 0;
    }
  }
  .opt {
    font-weight: 800;
    color: #909399;
    text-align: center;

    .sync-resource {
      height: 40px;
      line-height: 40px;
    }
    .upload-resource {
      font-size: 60px;
      line-height: 110px;

      &:hover {
        color: #1890ff;
      }
    }
  }
  .resource {
    position: relative;
    text-align: center;

    .asset-title {
      height: 25px;
      line-height: 25px;
      color: #606266;
      font-size: 12px;
      text-align: left !important;
      overflow: hidden;
      padding: 0 5px;
      word-break: break-all;

      .asset-name {
        display: inline-block;
        margin-left: 5px;
      }
    }

    .view {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 95px;
      line-height: 95px;
      font-size: 30px;
      background: rgba(144, 147, 153, 0.5);
      color: #ffffff;
      font-weight: 600;
      z-index: 999;
      display: none;
    }
    .preview__file {
      max-width: 148px;
      max-height: 95px;
    }
    .remove {
      color: #ffffff;
      height: 25px;
      line-height: 25px;
      background: rgba(255, 0, 0, 1);
      position: absolute;
      width: 100%;
      bottom:  0;
      display: none;
      font-size: 12px;
    }
    .check {
      height: 25px;
      line-height: 25px;
      background: rgba(208, 208, 208, 0.8);
      position: absolute;
      width: 100%;
      bottom:  0;
    }
    .check-active {
      background: #1890ff !important;
      color: #ffffff !important;
      font-weight: 600;
    }
    &:hover {
      .remove,.view {
        display: block;
      }
    }
  }
}
</style>
