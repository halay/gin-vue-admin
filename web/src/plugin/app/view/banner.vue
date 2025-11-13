
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
      
            <el-form-item label="标题" prop="title">
  <el-input v-model="searchInfo.title" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="位置标识" prop="position">
    <el-tree-select v-model="formData.position" placeholder="请选择位置标识" :data="banner_positionOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="状态" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="banner_statusOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="开始时间" prop="startTime">
  <template #label>
    <span>
      开始时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.startTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
           
            <el-form-item label="结束时间" prop="endTime">
  <template #label>
    <span>
      结束时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.endTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
           
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
        
            <el-table-column align="left" label="标题" prop="title" width="120" />

            <el-table-column label="图片地址" prop="imageUrl" width="200">
    <template #default="scope">
      <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.imageUrl)" fit="cover"/>
    </template>
</el-table-column>
            <el-table-column align="left" label="位置标识" prop="position" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.position,banner_positionOptions) }}
    </template>
</el-table-column>
            <el-table-column sortable align="left" label="排序" prop="sort" width="120" />

            <el-table-column align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,banner_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="开始时间" prop="startTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="结束时间" prop="endTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateBannerFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="标题:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入标题" />
</el-form-item>
             <el-form-item label="图片地址:" prop="imageUrl">
    <SelectImage
     v-model="formData.imageUrl"
     file-type="image"
    />
</el-form-item>
             <el-form-item label="跳转链接:" prop="linkUrl">
    <el-input v-model="formData.linkUrl" :clearable="true" placeholder="请输入跳转链接" />
</el-form-item>
             <el-form-item label="位置标识:" prop="position">
    <el-tree-select v-model="formData.position" placeholder="请选择位置标识" :data="banner_positionOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="排序:" prop="sort">
    <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入排序" />
</el-form-item>
             <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="banner_statusOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="开始时间:" prop="startTime">
    <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="结束时间:" prop="endTime">
    <el-date-picker v-model="formData.endTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="标题">
    {{ detailForm.title }}
</el-descriptions-item>
                 <el-descriptions-item label="图片地址">
    <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailForm.imageUrl)" :src="getUrl(detailForm.imageUrl)" fit="cover" />
</el-descriptions-item>
                 <el-descriptions-item label="跳转链接">
    {{ detailForm.linkUrl }}
</el-descriptions-item>
                 <el-descriptions-item label="位置标识">
    {{ detailForm.position }}
</el-descriptions-item>
                 <el-descriptions-item label="状态">
    {{ detailForm.status }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createBanner,
  deleteBanner,
  deleteBannerByIds,
  updateBanner,
  findBanner,
  getBannerList
} from '@/plugin/app/api/banner'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'Banner'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const banner_positionOptions = ref([])
const banner_statusOptions = ref([])
const formData = ref({
            title: '',
            imageUrl: "",
            linkUrl: '',
            position: '',
            sort: 0,
            status: '',
            startTime: new Date(),
            endTime: new Date(),
        })



// 验证规则
const rule = reactive({
               title : [{
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
               imageUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               position : [{
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
            sort: 'sort',
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
  const table = await getBannerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    banner_positionOptions.value = await getDictFunc('banner_position')
    banner_statusOptions.value = await getDictFunc('banner_status')
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
            deleteBannerFunc(row)
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
      const res = await deleteBannerByIds({ IDs })
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
const updateBannerFunc = async(row) => {
    const res = await findBanner({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBannerFunc = async (row) => {
    const res = await deleteBanner({ ID: row.ID })
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
        title: '',
        imageUrl: "",
        linkUrl: '',
        position: '',
        sort: 0,
        status: '',
        startTime: new Date(),
        endTime: new Date(),
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
                  res = await createBanner(formData.value)
                  break
                case 'update':
                  res = await updateBanner(formData.value)
                  break
                default:
                  res = await createBanner(formData.value)
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
  const res = await findBanner({ ID: row.ID })
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

</style>
