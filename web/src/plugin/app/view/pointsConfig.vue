<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAtRange">
          <template #label>
            <span>创建日期<el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）"><el-icon><QuestionFilled /></el-icon></el-tooltip></span>
          </template>
          <el-date-picker v-model="searchInfo.createdAtRange" class="!w-380px" type="datetimerange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="searchInfo.type" clearable placeholder="请选择" style="width:220px">
            <el-option label="首单优惠" value="first_order" />
            <el-option label="积分充值" value="recharge" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-tree-select v-model="searchInfo.status" placeholder="请选择" :data="statusOptions" style="width:220px" filterable :clearable="true" check-strictly />
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
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180"><template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template></el-table-column>
        <el-table-column align="left" label="类型" prop="type" width="120"><template #default="scope">{{ formatType(scope.row.type) }}</template></el-table-column>
        <el-table-column align="left" label="积分数" prop="points" width="100" />
        <el-table-column align="left" label="赠送积分" prop="bonusPoints" width="100" />
        <el-table-column align="left" label="现价(¥)" prop="currentPrice" width="100" />
        <el-table-column align="left" label="原价(¥)" prop="originalPrice" width="100" />
        <el-table-column align="left" label="排序" prop="sort" width="80" />
        <el-table-column align="left" label="限时" prop="limited" width="80"><template #default="scope">{{ formatBoolean(scope.row.limited) }}</template></el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100"><template #default="scope">{{ filterDict(scope.row.status, statusOptions) }}</template></el-table-column>
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
        <el-form-item label="类型:" prop="type">
          <el-radio-group v-model="formData.type">
            <el-radio label="first_order">首单优惠</el-radio>
            <el-radio label="recharge">积分充值</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="积分数:" prop="points"><el-input v-model.number="formData.points" placeholder="请输入积分数"/></el-form-item>
        <el-form-item label="赠送积分数:" prop="bonusPoints"><el-input v-model.number="formData.bonusPoints" placeholder="请输入赠送积分"/></el-form-item>
        <el-form-item label="现价(¥):" prop="currentPrice"><el-input v-model.number="formData.currentPrice" placeholder="请输入现价"/></el-form-item>
        <el-form-item label="原价(¥):" prop="originalPrice"><el-input v-model.number="formData.originalPrice" placeholder="请输入原价"/></el-form-item>
        <el-form-item label="显示顺序:" prop="sort"><el-input v-model.number="formData.sort" placeholder="请输入排序"/></el-form-item>
        <el-form-item label="限时抢购:" prop="limited"><el-switch v-model="formData.limited" /></el-form-item>
        <el-form-item v-if="formData.limited" label="开始时间:" prop="startAt"><el-date-picker v-model="formData.startAt" type="datetime" style="width:100%" /></el-form-item>
        <el-form-item v-if="formData.limited" label="结束时间:" prop="endAt"><el-date-picker v-model="formData.endAt" type="datetime" style="width:100%" /></el-form-item>
        <el-form-item label="是否上架:" prop="status"><el-tree-select v-model="formData.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="true" check-strictly /></el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="类型">{{ formatType(detailForm.type) }}</el-descriptions-item>
        <el-descriptions-item label="积分数">{{ detailForm.points }}</el-descriptions-item>
        <el-descriptions-item label="赠送积分数">{{ detailForm.bonusPoints }}</el-descriptions-item>
        <el-descriptions-item label="现价(¥)">{{ detailForm.currentPrice }}</el-descriptions-item>
        <el-descriptions-item label="原价(¥)">{{ detailForm.originalPrice }}</el-descriptions-item>
        <el-descriptions-item label="显示顺序">{{ detailForm.sort }}</el-descriptions-item>
        <el-descriptions-item label="限时抢购">{{ formatBoolean(detailForm.limited) }}</el-descriptions-item>
        <el-descriptions-item label="开始时间">{{ formatDate(detailForm.startAt) }}</el-descriptions-item>
        <el-descriptions-item label="结束时间">{{ formatDate(detailForm.endAt) }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ filterDict(detailForm.status, statusOptions) }}</el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import { createPointsConfig, deletePointsConfig, deletePointsConfigByIds, updatePointsConfig, findPointsConfig, getPointsConfigList } from '@/plugin/app/api/pointsConfig'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'

defineOptions({ name: 'PointsConfig' })
const btnLoading = ref(false)
const elFormRef = ref()
const elSearchFormRef = ref()
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const statusOptions = ref([])
const setOptions = async () => { statusOptions.value = await getDictFunc('banner_status') }
setOptions()

const formatType = (t) => t === 'first_order' ? '首单优惠' : (t === 'recharge' ? '积分充值' : (t || ''))

const onReset = () => { searchInfo.value = {}; getTableData() }
const onSubmit = () => { elSearchFormRef.value?.validate(async(v) => { if (!v) return; page.value = 1; getTableData() }) }
const handleSizeChange = (val) => { pageSize.value = val; getTableData() }
const handleCurrentChange = (val) => { page.value = val; getTableData() }
const getTableData = async() => {
  const table = await getPointsConfigList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) { tableData.value = table.data.list; total.value = table.data.total; page.value = table.data.page; pageSize.value = table.data.pageSize }
}
getTableData()

const multipleSelection = ref([])
const handleSelectionChange = (val) => { multipleSelection.value = val }
const deleteRow = (row) => { ElMessageBox.confirm('确定要删除吗?', '提示', { confirmButtonText:'确定', cancelButtonText:'取消', type:'warning' }).then(() => { deleteFunc(row) }) }
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', { confirmButtonText:'确定', cancelButtonText:'取消', type:'warning' }).then(async() => {
    const IDs = multipleSelection.value.map(i => i.ID)
    if (!IDs.length) { ElMessage({ type:'warning', message:'请选择要删除的数据' }); return }
    const res = await deletePointsConfigByIds({ IDs })
    if (res.code === 0) { ElMessage({ type:'success', message:'删除成功' }); if (tableData.value.length === IDs.length && page.value > 1) { page.value-- }; getTableData() }
  })
}

const type = ref('')
const updateFunc = async(row) => { const res = await findPointsConfig({ ID: row.ID }); type.value='update'; if (res.code === 0) { formData.value = res.data; dialogFormVisible.value = true } }
const deleteFunc = async(row) => { const res = await deletePointsConfig({ ID: row.ID }); if (res.code === 0) { ElMessage({ type:'success', message:'删除成功' }); if (tableData.value.length === 1 && page.value > 1) { page.value-- }; getTableData() } }
const dialogFormVisible = ref(false)
const openDialog = () => { type.value='create'; dialogFormVisible.value = true }
const closeDialog = () => { dialogFormVisible.value = false; formData.value = initForm() }
const enterDialog = async () => {
  btnLoading.value = true
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return btnLoading.value = false
    let res
    res = type.value === 'update' ? await updatePointsConfig(formData.value) : await createPointsConfig(formData.value)
    btnLoading.value = false
    if (res.code === 0) { ElMessage({ type:'success', message:'创建/更改成功' }); closeDialog(); getTableData() }
  })
}

const initForm = () => ({ type:'recharge', points:0, bonusPoints:0, currentPrice:0, originalPrice:0, sort:0, limited:false, startAt:null, endAt:null, status:'enabled' })
const formData = ref(initForm())
const rule = reactive({ type:[{ required:true, message:'请选择类型', trigger:['change'] }], points:[{ required:true, message:'请输入积分数', trigger:['blur','change'] }], currentPrice:[{ required:true, message:'请输入现价', trigger:['blur','change'] }] })

const detailForm = ref({})
const detailShow = ref(false)
const openDetailShow = () => { detailShow.value = true }
const getDetails = async(row) => { const res = await findPointsConfig({ ID: row.ID }); if (res.code === 0) { detailForm.value = res.data; openDetailShow() } }
const closeDetailShow = () => { detailShow.value = false; detailForm.value = {} }
</script>

<style>
</style>
