<template>
  <dialog-panel title="配置添加" confirm-text="添加" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading">
    <el-form :model="confForm" ref="confForm" label-width="100px" size="small">
      <el-form-item label="配置名称" prop="_name" :rules="{required: true, message: '请填写配置名称'}">
        <el-input v-model="confForm._name"/>
      </el-form-item>
      <el-form-item label="配置代码" prop="_key" :rules="{required: true, message: '请填写配置代码'}">
        <el-input v-model="confForm._key" placeholder="仅支持字符开头的大小写字母数字下划线组合 [50 位以内]"/>
      </el-form-item>
      <el-form-item label="配置值" prop="_val" :rules="{required: true, message: '请填写配置值'}">
        <el-input v-model="confForm._val"/>
      </el-form-item>
      <el-form-item label="扩展1" prop="bak1">
        <el-input v-model="confForm.bak1"/>
      </el-form-item>
      <el-form-item label="扩展2" prop="bak2">
        <el-input v-model="confForm.bak2"/>
      </el-form-item>
      <el-form-item label="扩展3" prop="bak3">
        <el-input v-model="confForm.bak3"/>
      </el-form-item>
      <el-form-item label="配置描述" prop="desc">
        <el-input v-model="confForm.desc"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import {confCreate} from '@a/config'

  export default {
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        confForm: {
          _name: '',
          _key: '',
          desc: '',
          bak1: '',
          bak2: '',
          bak3: '',
          _val: ''
        },
      }
    },
    methods: {
      cancel() {
        this.$refs.confForm.resetFields()
        this.visible = false
      },
      add() {
        this.$refs.confForm.validate(v => {
          if (v) {
            this.loading = true
            confCreate(this.confForm).then(res => {
              this.$message.success("创建成功")
              this.$emit('success')
              this.loading = false
              this.cancel()
            }).catch(err => {
              this.loading = false
              console.log(err)
            })
          } else {
            return false
          }
        })
      }
    }
  }
</script>
