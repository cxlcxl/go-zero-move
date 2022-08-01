<template>
  <div class="profile-container">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="基础信息" name="profile">
        <div style="padding: 30px;">
          <profile-form ref="profile"/>
        </div>
      </el-tab-pane>
      <el-tab-pane label="密码修改" name="pass">
        <div style="padding: 30px;">
          <pass-reset ref="pass"/>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
  import ProfileForm from './components/profile'
  import PassReset from './components/pass'

  export default {
    name: 'Profile',
    components: {ProfileForm, PassReset},
    data() {
      return {
        activeName: 'profile'
      }
    },
    methods: {
      handleClick(tab, event) {

      }
    }
  }
</script>
