<template>
  <el-pagination background :current-page="page" :page-size="limit" :total="total" hide-on-single-page
      layout="prev, pager, next, jumper, total, sizes" prev-text="上页" next-text="下页"
      @current-change="handlePage" @size-change="handleSizeChange"
  />
</template>

<script>
  export default {
    props: {
      page: {
        type: Number,
        default: 1
      },
      total: {
        type: Number,
        default: 0
      },
      limit: {
        type: Number,
        default: 10
      }
    },
    methods: {
      handlePage(p) {
        this.$emit('current-change', p)
      },
      handleSizeChange(s) {
        this.$emit('size-change', s)
      }
    }
  }
</script>
