
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="!w-380px"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
       </el-form-item>
      
            <el-form-item label="版本名" prop="versionName">
  <el-input v-model="searchInfo.versionName" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="版本号" prop="versionCode">
  <el-input v-model.number="searchInfo.versionCode" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="平台" prop="platform">
    <el-tree-select v-model="searchInfo.platform" placeholder="请选择平台" :data="app_platformOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="状态" prop="status">
    <el-tree-select v-model="searchInfo.status" placeholder="请选择状态" :data="release_statusOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="强制更新" prop="forceUpdate">
  <el-select v-model="searchInfo.forceUpdate" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
           
            <el-form-item label="最低支持版本号" prop="minSupport">
  <el-input v-model.number="searchInfo.minSupport" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="安卓包名" prop="packageName">
  <el-input v-model="searchInfo.packageName" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="iOS Bundle ID" prop="bundleId">
  <el-input v-model="searchInfo.bundleId" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="安卓安装包" prop="androidPackage">
  <el-input v-model="searchInfo.androidPackage" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="iOS安装包/链接" prop="iosPackage">
  <el-input v-model="searchInfo.iosPackage" placeholder="搜索条件" />
</el-form-item>
           
            
           
            <el-form-item label="更新日志" prop="changelog">
  <el-input v-model="searchInfo.changelog" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="发布时间" prop="releaseAt">
  <template #label>
    <span>
      发布时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.releaseAtRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="app_AppRelease" />
            <ExportExcel  template-id="app_AppRelease" filterDeleted/>
            <ImportExcel  template-id="app_AppRelease" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column sortable align="left" label="版本名" prop="versionName" width="120" />

            <el-table-column sortable align="left" label="版本号" prop="versionCode" width="120" />

            <el-table-column align="left" label="平台" prop="platform" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.platform,app_platformOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,release_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="强制更新" prop="forceUpdate" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.forceUpdate) }}</template>
</el-table-column>
            <el-table-column align="left" label="最低支持版本号" prop="minSupport" width="120" />

            <el-table-column align="left" label="安卓包名" prop="packageName" width="120" />

            <el-table-column align="left" label="iOS Bundle ID" prop="bundleId" width="120" />

            <el-table-column align="left" label="安卓下载链接" prop="androidUrl" width="120" />

            <el-table-column align="left" label="iOS下载链接" prop="iosUrl" width="120" />

            <el-table-column align="left" label="发布时间" prop="releaseAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.releaseAt) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateAppReleaseFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
             <el-form-item label="版本名:" prop="versionName">
    <el-input v-model="formData.versionName" :clearable="false" placeholder="请输入版本名" />
</el-form-item>
             <el-form-item label="版本号:" prop="versionCode">
    <el-input v-model.number="formData.versionCode" :clearable="false" placeholder="请输入版本号" />
</el-form-item>
             <el-form-item label="平台:" prop="platform">
    <el-tree-select v-model="formData.platform" placeholder="请选择平台" :data="app_platformOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="release_statusOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="强制更新:" prop="forceUpdate">
    <el-switch v-model="formData.forceUpdate" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
             <el-form-item label="最低支持版本号:" prop="minSupport">
    <el-input v-model.number="formData.minSupport" :clearable="false" placeholder="请输入最低支持版本号" />
</el-form-item>
             <el-form-item label="安卓包名:" prop="packageName">
    <el-input v-model="formData.packageName" :clearable="false" placeholder="请输入安卓包名" />
</el-form-item>
             <el-form-item label="iOS Bundle ID:" prop="bundleId">
    <el-input v-model="formData.bundleId" :clearable="false" placeholder="请输入iOS Bundle ID" />
</el-form-item>
             <el-form-item label="安卓安装包:" prop="androidPackage">
    <SelectFile v-model="formData.androidPackage" />
</el-form-item>
             <el-form-item label="iOS安装包/链接:" prop="iosPackage">
    <SelectFile v-model="formData.iosPackage" />
</el-form-item>
             
             <el-form-item label="更新日志:" prop="changelog">
    <RichEdit v-model="formData.changelog"/>
</el-form-item>
             <el-form-item label="发布时间:" prop="releaseAt">
    <el-date-picker v-model="formData.releaseAt" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAppRelease,
  deleteAppRelease,
  deleteAppReleaseByIds,
  updateAppRelease,
  findAppRelease,
  getAppReleaseList
} from '@/plugin/app/api/appRelease'
import { getUrl } from '@/utils/image'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'AppRelease'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const app_platformOptions = ref([])
const release_statusOptions = ref([])
const formData = ref({
            versionName: '',
            versionCode: 0,
            platform: '',
            status: '',
            forceUpdate: false,
            minSupport: 0,
            packageName: '',
            bundleId: '',
            androidPackage: [],
            iosPackage: [],
            androidUrl: '',
            iosUrl: '',
            changelog: '',
            releaseAt: new Date(),
        })



// 验证规则
const rule = reactive({
               versionName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               versionCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               platform : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"created_at",
    ID:"id",
            versionName: 'version_name',
            versionCode: 'version_code',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.forceUpdate === ""){
        searchInfo.value.forceUpdate=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAppReleaseList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    app_platformOptions.value = await getDictFunc('app_platform')
    release_statusOptions.value = await getDictFunc('release_status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteAppReleaseFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteAppReleaseByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAppReleaseFunc = async(row) => {
    const res = await findAppRelease({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppReleaseFunc = async (row) => {
    const res = await deleteAppRelease({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        versionName: '',
        versionCode: 0,
        platform: '',
        status: '',
        forceUpdate: false,
        minSupport: 0,
        packageName: '',
        bundleId: '',
        androidPackage: [],
        iosPackage: [],
        androidUrl: '',
        iosUrl: '',
        changelog: '',
        releaseAt: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createAppRelease(formData.value)
                  break
                case 'update':
                  res = await updateAppRelease(formData.value)
                  break
                default:
                  res = await createAppRelease(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findAppRelease({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}

</style>
