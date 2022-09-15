<template>
  <el-row :gutter="8">
    <el-col :span="12" class="location-view">
      <p class="location-header">
        <span class="title text-success">包含列表</span>
        <span class="item"><el-tag size="mini" type="danger" @click="excludeAll">全部排除</el-tag></span>
        <span class="item"><i class="el-icon-close text-error" @click="handleRemoveIncludeAll"></i></span>
      </p>

      <div class="location-body">
        <p class="location-line" v-for="item in countries" v-show="selectorValue._include.includes(item.c_id)">
          <span class="location-item" style="width: 220px;">{{ item.c_name }}</span>
          <span class="location-item flex-item">
            <el-tag size="mini" type="danger" @click="handleExclude(item.c_id)">排除</el-tag>
          </span>
          <span class="location-item flex-item">
            <i class="el-icon-close text-error" @click="handleRemoveInclude(item.c_id)"></i>
          </span>
        </p>
      </div>
    </el-col>
    <el-col :span="12" class="location-view">
      <p class="location-header">
        <span class="title text-error">排除列表</span>
        <span class="item"><el-tag size="mini" type="success" @click="includeAll">全部包含</el-tag></span>
        <span class="item"><i class="el-icon-close text-error" @click="handleRemoveExcludeAll"></i></span>
      </p>

      <div class="location-body">
        <p class="location-line" v-for="item in countries" v-show="selectorValue._exclude.includes(item.c_id)">
          <span class="location-item" style="width: 220px;">{{ item.c_name }}</span>
          <span class="location-item flex-item">
            <el-tag size="mini" type="success" @click="handleInclude(item.c_id)">包含</el-tag>
          </span>
          <span class="location-item flex-item">
            <i class="el-icon-close text-error" @click="handleRemoveExclude(item.c_id)"></i>
          </span>
        </p>
      </div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  props: {
    countries: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      selectorValue: {
        _include: [],
        _exclude: []
      },
    }
  },
  methods: {
    watchChange(v) {
      this.$set(this.selectorValue, '_include', v._include)
      this.$set(this.selectorValue, '_exclude', v._exclude)
    },
    handleRemoveInclude(c_id) {
      let idx = this.selectorIndexOf(this.selectorValue._include, c_id)
      if (idx !== -1) {
        this.selectorValue._include.splice(idx ,1)
      }
      this.$emit('handle-change', this.selectorValue)
    },
    handleRemoveExclude(c_id) {
      let idx = this.selectorIndexOf(this.selectorValue._exclude, c_id)
      if (idx !== -1) {
        this.selectorValue._exclude.splice(idx ,1)
      }
      this.$emit('handle-change', this.selectorValue)
    },
    handleInclude(c_id) {
      let idx = this.selectorIndexOf(this.selectorValue._exclude, c_id)
      if (idx !== -1) {
        this.selectorValue._exclude.splice(idx ,1)
      }
      this.selectorValue._include.push(c_id)
      this.$emit('handle-change', this.selectorValue)
    },
    handleExclude(c_id) {
      let idx = this.selectorIndexOf(this.selectorValue._include, c_id)
      if (idx !== -1) {
        this.selectorValue._include.splice(idx ,1)
      }
      this.selectorValue._exclude.push(c_id)
      this.$emit('handle-change', this.selectorValue)
    },
    includeAll() {
      let all = this.selectorValue._include.concat(this.selectorValue._exclude)
      this.$set(this.selectorValue, '_include', all)
      this.$set(this.selectorValue, '_exclude', [])

      this.$emit('handle-change', this.selectorValue)
    },
    excludeAll() {
      let all = this.selectorValue._include.concat(this.selectorValue._exclude)
      this.$set(this.selectorValue, '_include', [])
      this.$set(this.selectorValue, '_exclude', all)

      this.$emit('handle-change', this.selectorValue)
    },
    handleRemoveExcludeAll() {
      this.$set(this.selectorValue, '_exclude', [])
      this.$emit('handle-change', this.selectorValue)
    },
    handleRemoveIncludeAll() {
      this.$set(this.selectorValue, '_include', [])
      this.$emit('handle-change', this.selectorValue)
    },
    selectorIndexOf(arr, v) {
      for (let i = 0; i < arr.length; i++) {
        if (arr[i] === v) {
          return i
        }
      }
      return -1
    },
  }
}
</script>

<style scoped lang="scss">
.location-view {
  margin-top: 5px;
  .location-header {
    display: flex;

    .title,.item {
      display: inline-block;
    }
    .title {
      width: 210px;
    }
    .item {
      flex: 1;
      text-align: center;
      cursor: pointer;
    }
  }
  .location-body {
    height: 200px;
    overflow-y: scroll;
    border: 1px solid #eee;
    border-radius: 4px;
    padding: 5px;

    .location-line {
      height: 25px;
      line-height: 25px;
      display: flex;
      font-size: 12px;
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
</style>
