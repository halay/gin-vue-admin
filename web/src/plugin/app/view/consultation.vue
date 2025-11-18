<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAtRange">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）"><el-icon><QuestionFilled /></el-icon></el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.createdAtRange" class="!w-380px" type="datetimerange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="标题" prop="title"><el-input v-model="searchInfo.title" placeholder="搜索条件" /></el-form-item>
        <el-form-item label="状态" prop="status">
          <el-tree-select v-model="searchInfo.status" placeholder="请选择状态" :data="statusOptions" style="width:240px" filterable :clearable="true" check-strictly />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>
        <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID" @selection-change="handleSelectionChange" @sort-change="sortChange">
        <el-table-column type="selection" width="55" />
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180"><template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template></el-table-column>
        <el-table-column align="left" label="标题" prop="title" width="240" />
        <el-table-column label="封面图" prop="coverImage" width="160"><template #default="scope"><el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.coverImage)" fit="cover"/></template></el-table-column>
        <el-table-column sortable align="left" label="排序" prop="sort" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120"><template #default="scope">{{ filterDict(scope.row.status, statusOptions) }}</template></el-table-column>
        <el-table-column align="left" label="发布时间" prop="publishAt" width="180"><template #default="scope">{{ formatDate(scope.row.publishAt) }}</template></el-table-column>
        <el-table-column align="left" label="点击" prop="clicks" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange" @size-change="handleSizeChange" />
      </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center"><span class="text-lg">{{type==='create'?'新增':'编辑'}}</span><div><el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button><el-button @click="closeDialog">取 消</el-button></div></div>
      </template>
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title"><el-input v-model="formData.title" clearable placeholder="请输入标题"/></el-form-item>
        <el-form-item label="封面图:" prop="coverImage"><SelectImage v-model="formData.coverImage" file-type="image"/></el-form-item>
        <el-form-item label="跳转链接:" prop="jumpUrl"><el-input v-model="formData.jumpUrl" clearable placeholder="可选跳转链接"/></el-form-item>
        <el-form-item label="摘要:" prop="excerpt"><el-input type="textarea" v-model="formData.excerpt" placeholder="可选摘要"/></el-form-item>
        <el-form-item label="内容:" prop="content"><RichEdit v-model="formData.content"/></el-form-item>
        <el-form-item label="状态:" prop="status"><el-tree-select v-model="formData.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="true" check-strictly /></el-form-item>
        <el-form-item label="排序:" prop="sort"><el-input v-model.number="formData.sort" placeholder="请输入排序"/></el-form-item>
        <el-form-item label="发布时间:" prop="publishAt"><el-date-picker v-model="formData.publishAt" type="datetime" style="width:100%" placeholder="选择时间"/></el-form-item>
        <el-form-item label="过期时间:" prop="expireAt"><el-date-picker v-model="formData.expireAt" type="datetime" style="width:100%" placeholder="选择时间"/></el-form-item>
      </el-form>
    </el-drawer>
    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="标题">{{ detailForm.title }}</el-descriptions-item>
        <el-descriptions-item label="封面图"><el-image style="width: 50px; height: 50px" :src="getUrl(detailForm.coverImage)" fit="cover" /></el-descriptions-item>
        <el-descriptions-item label="跳转链接">{{ detailForm.jumpUrl }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ filterDict(detailForm.status, statusOptions) }}</el-descriptions-item>
        <el-descriptions-item label="发布时间">{{ formatDate(detailForm.publishAt) }}</el-descriptions-item>
        <el-descriptions-item label="点击">{{ detailForm.clicks }}</el-descriptions-item>
        <el-descriptions-item label="内容"><RichView :content="detailForm.content"/></el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import { createConsultation, deleteConsultation, deleteConsultationByIds, updateConsultation, findConsultation, getConsultationList } from '@/plugin/app/api/consultation'
import { getUrl } from '@/utils/image'
import SelectImage from '@/components/selectImage/selectImage.vue'
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { formatDate, getDictFunc, filterDict } from '@/utils/format'

defineOptions({ name: 'AppConsultation' })
const btnLoading = ref(false)
const elFormRef = ref()
const elSearchFormRef = ref()
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const showAllQuery = ref(false)

const sortChange = ({ prop, order }) => {
  const sortMap = { CreatedAt:'created_at', ID:'id', sort:'sort', publishAt:'publish_at', clicks:'clicks' }
  let sort = sortMap[prop] || prop.replace(/[A-Z]/g, m => `_${m.toLowerCase()}`)
  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
const onReset = () => { searchInfo.value = {}; getTableData() }
const onSubmit = () => { elSearchFormRef.value?.validate(async(v) => { if (!v) return; page.value = 1; getTableData() }) }
const handleSizeChange = (val) => { pageSize.value = val; getTableData() }
const handleCurrentChange = (val) => { page.value = val; getTableData() }
const getTableData = async() => {
  const table = await getConsultationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) { tableData.value = table.data.list; total.value = table.data.total; page.value = table.data.page; pageSize.value = table.data.pageSize }
}
getTableData()

const statusOptions = ref([])
const setOptions = async () => {
  statusOptions.value = await getDictFunc('banner_status')
}
setOptions()

const multipleSelection = ref([])
const handleSelectionChange = (val) => { multipleSelection.value = val }
const deleteRow = (row) => { ElMessageBox.confirm('确定要删除吗?', '提示', { confirmButtonText:'确定', cancelButtonText:'取消', type:'warning' }).then(() => { deleteFunc(row) }) }
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', { confirmButtonText:'确定', cancelButtonText:'取消', type:'warning' }).then(async() => {
    const IDs = multipleSelection.value.map(i => i.ID)
    if (!IDs.length) { ElMessage({ type:'warning', message:'请选择要删除的数据' }); return }
    const res = await deleteConsultationByIds({ IDs })
    if (res.code === 0) { ElMessage({ type:'success', message:'删除成功' }); if (tableData.value.length === IDs.length && page.value > 1) { page.value-- }; getTableData() }
  })
}

const type = ref('')
const updateFunc = async(row) => { const res = await findConsultation({ ID: row.ID }); type.value='update'; if (res.code === 0) { formData.value = res.data; dialogFormVisible.value = true } }
const deleteFunc = async(row) => { const res = await deleteConsultation({ ID: row.ID }); if (res.code === 0) { ElMessage({ type:'success', message:'删除成功' }); if (tableData.value.length === 1 && page.value > 1) { page.value-- }; getTableData() } }
const dialogFormVisible = ref(false)
const openDialog = () => { type.value='create'; dialogFormVisible.value = true }
const closeDialog = () => { dialogFormVisible.value = false; formData.value = initForm() }
const enterDialog = async () => {
  btnLoading.value = true
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return btnLoading.value = false
    let res
    res = type.value === 'update' ? await updateConsultation(formData.value) : await createConsultation(formData.value)
    btnLoading.value = false
    if (res.code === 0) { ElMessage({ type:'success', message:'创建/更改成功' }); closeDialog(); getTableData() }
  })
}

const initForm = () => ({ title:'', coverImage:'', content:'', jumpUrl:'', status:'', sort:0, publishAt:new Date(), expireAt:null, excerpt:'' })
const formData = ref(initForm())
const rule = reactive({ title:[{ required:true, message:'请输入标题', trigger:['input','blur'] },{ whitespace:true, message:'不能只输入空格', trigger:['input','blur']}], coverImage:[{ required:true, message:'请选择封面图', trigger:['change','blur'] }], content:[{ required:true, message:'请输入内容', trigger:['change','blur'] }] })

const detailForm = ref({})
const detailShow = ref(false)
const openDetailShow = () => { detailShow.value = true }
const getDetails = async(row) => { const res = await findConsultation({ ID: row.ID }); if (res.code === 0) { detailForm.value = res.data; openDetailShow() } }
const closeDetailShow = () => { detailShow.value = false; detailForm.value = {} }
</script>

<style>
</style>