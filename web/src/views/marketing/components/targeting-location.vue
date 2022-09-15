<template>
  <drawer-panel title="国家/地区选择" confirm-text="完成选择" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="550px">
    <div class="location-selector">
      <p class="location-search">
        <el-input placeholder="请输入搜索国家/地区 - 切换大洲会自动过滤关键词" v-model="keyword" size="small" @change="doSearch" clearable>
          <el-button slot="append" icon="el-icon-search" @click="doSearch"></el-button>
        </el-input>
      </p>

      <el-tabs type="border-card" v-model="tabIdx">
        <el-tab-pane :label="item.c_name" v-for="item in continents" :name="item.id.toString()"/>
        <div class="location-table">
          <p class="location-header">
            <span class="location-item" style="width: 300px;">国家/地区名</span>
            <span class="location-item flex-item">
              <el-tag size="mini" type="success" @click="includeAll">全部包含</el-tag>
            </span>
            <span class="location-item flex-item">
              <el-tag size="mini" type="danger" @click="excludeAll">全部排除</el-tag>
            </span>
          </p>
          <div class="location-body">
            <p class="location-line" v-for="item in (showSearch ? searchList :countries)" v-show="item.continent_id === Number(tabIdx)">
              <span class="location-item" style="width: 300px;">{{ item.c_name }}</span>
              <span class="location-item flex-item">
                <el-tag size="mini" type="success" @click="handleInclude(item.c_id)" v-show="!selectorValue._include.includes(item.c_id)">包含</el-tag>
              </span>
              <span class="location-item flex-item">
                <el-tag size="mini" type="danger" @click="handleExclude(item.c_id)" v-show="!selectorValue._exclude.includes(item.c_id)">排除</el-tag>
              </span>
            </p>
          </div>
        </div>
      </el-tabs>
    </div>
  </drawer-panel>
</template>

<script>
import DrawerPanel from '@c/Drawer'

export default {
  components: {
    DrawerPanel
  },
  props: {
    continents: {
      type: Array,
      required: true
    },
    countries: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      visible: false,
      loading: false,
      showSearch: false,
      tabIdx: '1',
      keyword: '',
      selectorValue: {
        _include: [],
        _exclude: []
      },
      searchList: []
    }
  },
  methods: {
    initCreate(_in, _ex) {
      if (this.continents.length > 0) {
        this.tabIdx = this.continents[0].id.toString()
      }
      if (!Array.isArray(_in) || !Array.isArray(_ex)) {
        this.$message.error("参数错误")
        return
      }
      this.$set(this.selectorValue, '_include', _in)
      this.$set(this.selectorValue, '_exclude', _ex)
      this.visible = true
    },
    cancel() {
      this.visible = false
    },
    add() {
      this.$emit('handle-select', this.selectorValue)
      this.visible = false
    },
    handleInclude(c_id) {
      let idx = this.selectorIndexOf(this.selectorValue._exclude, c_id)
      if (idx !== -1) {
        this.selectorValue._exclude.splice(idx ,1)
      }
      this.selectorValue._include.push(c_id)
    },
    handleExclude(c_id) {
      let idx = this.selectorIndexOf(this.selectorValue._include, c_id)
      if (idx !== -1) {
        this.selectorValue._include.splice(idx ,1)
      }
      this.selectorValue._exclude.push(c_id)
    },
    includeAll() {
      let includeSrc = this.showSearch ? this.searchList : this.countries
      for (let i = 0; i < includeSrc.length; i++) {
        if (Number(this.tabIdx) === includeSrc[i].continent_id) {
          let c_id = includeSrc[i].c_id
          let includeIdx = this.selectorIndexOf(this.selectorValue._include, c_id)
          if (includeIdx === -1) { // 未被包含的才可以被加入包含数组
            let idx = this.selectorIndexOf(this.selectorValue._exclude, c_id)
            if (idx !== -1) {
              this.selectorValue._exclude.splice(idx ,1)
            }
            this.selectorValue._include.push(c_id)
          }
        }
      }
    },
    excludeAll() {
      let excludeSrc = this.showSearch ? this.searchList : this.countries
      for (let i = 0; i < excludeSrc.length; i++) {
        if (Number(this.tabIdx) === excludeSrc[i].continent_id) {
          let c_id = excludeSrc[i].c_id
          let excludeIdx = this.selectorIndexOf(this.selectorValue._exclude, c_id)
          if (excludeIdx === -1) { // 未被排除的才可以被加入排除数组
            let idx = this.selectorIndexOf(this.selectorValue._include, c_id)
            if (idx !== -1) {
              this.selectorValue._include.splice(idx ,1)
            }
            this.selectorValue._exclude.push(c_id)
          }
        }
      }
    },
    selectorIndexOf(arr, v) {
      for (let i = 0; i < arr.length; i++) {
        if (arr[i] === v) {
          return i
        }
      }
      return -1
    },
    doSearch() {
      this.searchList = []
      if (this.keyword === '') {
        this.showSearch = false
        return
      }
      for (let i = 0; i < this.countries.length; i++) {
        let cName = this.countries[i].c_name
        if (cName.indexOf(this.keyword) >= 0) {
          this.searchList.push(this.countries[i])
        }
      }
      this.showSearch = true
    }
  }
}
</script>

<style lang="scss">
.location-selector {
  padding: 0 10px;
  margin-bottom: 20px;

  .location-table {
    .location-header {
      display: flex;
      height: 35px;
      line-height: 35px;
      border-bottom: 1px solid #e8e8e8;
      margin-bottom: 5px;
      font-weight: 600;

      .flex-item {
        flex: 1;
        cursor: pointer;
        text-align: center;
      }
    }
    .location-body {
      max-height: 300px;
      overflow-y: scroll;

      .location-line {
        height: 30px;
        line-height: 30px;
        display: flex;
        padding: 0 5px;

        &:hover {
          background: #f3f3f3;
        }

        .location-item {
          display: inline-block;
        }
        .flex-item {
          flex: 1;
          cursor: pointer;
          text-align: center;
        }
      }
    }
  }

  .location-search {
    margin-bottom: 8px;
  }
}
</style>
