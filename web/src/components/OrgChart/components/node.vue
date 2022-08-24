<template>
  <table>
    <tbody>
    <tr>
      <td :colspan="datasource.children && datasource.children.length ? datasource.children.length * 2 : null">
        <div class="node" :id="datasource.id">
          <slot :node-data="datasource">
            <div class="node-box">
              <el-tooltip :content="'[' + datasource.allow_act + '] ' + datasource.permission" placement="top">
                <el-button type="primary" plain @click.stop="handleClick(datasource, 'default')">{{ datasource.desc }}</el-button>
              </el-tooltip>
              <div class="operates">
                <i class="el-icon-plus text-primary" @click.stop="handleClick(datasource, 'create')"/>
                <i class="el-icon-edit text-primary" @click.stop="handleClick(datasource, 'update')" v-show="datasource.id > 0"/>
                <i class="el-icon-delete text-error" @click.stop="handleClick(datasource, 'delete')" v-show="datasource.id > 0"/>
              </div>
            </div>
          </slot>
        </div>
      </td>
    </tr>
    <template v-if="datasource.children && datasource.children.length">
      <tr class="lines">
        <td :colspan="datasource.children.length * 2">
          <div class="downLine"></div>
        </td>
      </tr>
      <tr class="lines">
        <td class="rightLine"></td>
        <template v-for="n in datasource.children.length - 1">
          <td class="leftLine topLine"></td>
          <td class="rightLine topLine"></td>
        </template>
        <td class="leftLine"></td>
      </tr>
      <tr class="nodes">
        <td colspan="2" v-for="child in datasource.children" :key="child.id">
          <node :datasource="child" :handle-click="handleClick">
            <template v-for="slot in Object.keys($scopedSlots)" :slot="slot" slot-scope="scope">
              <slot :name="slot" v-bind="scope" />
            </template>
          </node>
        </td>
      </tr>
    </template>
    </tbody>
  </table>
</template>

<script>
  export default {
    name: "node",
    props: {
      datasource: Object,
      handleClick: Function,
    }
  };
</script>
