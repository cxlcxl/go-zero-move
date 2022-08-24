<template>
  <el-row>
    <el-col :span="24">
      <upload ref="upload" :post-data="postData" accept=".xlsx" @success="handleSuccess">
        <el-select slot="opt" style="margin-left: 8px;" placeholder="选择导入模块" v-model="postData.module"
                   clearable @change="importDataList">
          <el-option v-for="m in modules" :label="m._name" :value="m._val" :key="m._val"/>
        </el-select>
        <el-button type="primary" @click="submit" slot="opt" icon="el-icon-upload" style="margin-left: 8px;" v-permission="'upload'">上传
        </el-button>
        <el-button @click="importDataList" slot="opt" icon="el-icon-refresh" style="margin-left: 8px;">刷新列表
        </el-button>
        <el-button @click="downloadTemplate" slot="opt" icon="el-icon-download" style="margin-left: 8px;">下载模板
        </el-button>
        <el-button @click="saveData" v-show="dataList.list && dataList.list.length > 0" v-permission="'excel/save'"
                   type="primary" icon="el-icon-check" slot="opt" style="float: right;">确认无误保存数据</el-button>
        <el-button @click="flushData" v-show="dataList.list && dataList.list.length > 0" v-permission="'excel/flush'"
                   type="danger" icon="el-icon-delete" slot="opt" style="float: right;">清空</el-button>
      </upload>
    </el-col>

    <el-col :span="24" style="margin-top: 15px;">
      <p style="margin-bottom: 10px;">
      </p>

      <el-table v-loading="dataLoading" :data="dataList.list" border>
        <el-table-column align="center" label="操作" width="60" v-permission="'excel/destroy'">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="danger" plain @click.native.prevent="destroyRow(scope.$index, scope.row)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
        <el-table-column prop="module_name" label="模块" width="120">
          <template slot-scope="scope">
            {{scope.row.module_name|moduleName(modules)}}
          </template>
        </el-table-column>
        <el-table-column v-for="column in dataList.columns" :prop="column._val.toLowerCase()" min-width="100"
                         :label="(column.bak2 ? column.bak2 : column.bak1)+' '+column._val"/>
        <el-table-column label="保存失败提示" min-width="150">
          <template slot-scope="scope">
            <span class="text-error">{{scope.row.error_info}}</span>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" @current-change="handlePage" @size-change="handlePageSize" :total="dataList.total" :limit="search.limit"/>
    </el-col>
  </el-row>
</template>

<script>
  import Upload from '@c/Upload'
  import Page from '@c/Page'
  import {getImportData, excelTemplate, dataDestroy, saveImportData, flushImportData} from "@a/excel-tool"
  import {getConfigs} from "@a/global"

  export default {
    name: "import",
    components: {
      Upload, Page
    },
    data() {
      return {
        desc: "表格的 Sheet1 请勿重命名，日期的单元格需要修改成文本类型后再填入数据",
        dataLoading: false,
        postData: {
          module: ''
        },
        modules: [],
        search: {
          module_name: '',
          page: 1,
          limit: 10
        },
        dataList: {
          total: 0,
          list: [],
          columns: []
        }
      }
    },
    filters: {
      moduleName(m, modules) {
        for (let i = 0; i < modules.length; i++) {
          if (m === modules[i]._val) {
            return modules[i]._name
          }
        }

        return m
      }
    },
    mounted() {
      this.getImportModules()
    },
    methods: {
      getImportModules() {
        getConfigs('excel_import_module').then(res => {
          this.modules = res.data
        }).catch(() => {})
      },
      submit() {
        if (this.search.module_name === '') {
          this.$message.error('请先选择“导入模板”')
          return false
        }
        this.$refs.upload.submit()
      },
      importDataList() {
        this.search.module_name = this.postData.module
        if (this.postData.module === '') {
          this.dataList = {total: 0, list: [], columns: []}
          this.search.page = 1
          return
        }
        this.dataLoading = true
        getImportData(this.search).then(res => {
          this.dataLoading = false
          this.dataList = res.data
        }).catch(() => {
          this.dataLoading = false
        })
      },
      handlePage(p) {
        this.search.page = p
        this.importDataList()
      },
      handlePageSize(s) {
        this.search.limit = s
        this.importDataList()
      },
      handleSuccess(data) {
        this.$message.success('数据文件上传成功')
        this.dataLoading = true
        this.$refs.upload.fileList = []
        setTimeout(this.importDataList, 1000)
      },
      downloadTemplate() {
        if (this.search.module_name === '') {
          this.$message.error('请先选择“导入模板”')
          return false
        }
        window.open(excelTemplate + this.search.module_name)
      },
      destroyRow(idx, row) {
        this.$confirm('确认删除此行数据么?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.dataLoading = true
          dataDestroy(row.id).then(() => {
            this.$message.success('删除成功')
            this.dataList.list.splice(idx, 1)
            this.dataLoading = false
          }).catch(err => {
            this.dataLoading = false
          })
        }).catch(() => {})
      },
      flushData() {
        this.$confirm('确认清空导入的数据么?', '操作提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.dataLoading = true
          flushImportData(this.search.module_name).then(() => {
            this.$message.success('清除成功')
            this.handlePage(1)
          }).catch(err => {
            this.dataLoading = false
          })
        }).catch(() => {})
      },
      saveData() {
        this.$confirm('如提示保存成功，但列表还存在数据则表示列表的数据存在匹配错误，确认保存数据么?', '操作提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'info'
        }).then(() => {
          this.dataLoading = true
          saveImportData(this.search.module_name).then(() => {
            this.$message.success('保存成功')
            this.handlePage(1)
          }).catch(err => {
            this.handlePage(1)
          })
        }).catch(() => {})
      }
    }
  }
</script>
