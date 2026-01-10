
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
      
            <el-form-item label="经销商名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="销售提成" prop="salesCommission">
  <el-input v-model.number="searchInfo.salesCommission" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="提成类型" prop="commissionType">
    <el-tree-select v-model="searchInfo.commissionType" placeholder="请选择提成类型" :data="commission_typeOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="费用补贴" prop="expenseAllowance">
  <el-input v-model.number="searchInfo.expenseAllowance" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="补贴类型" prop="allowanceType">
    <el-tree-select v-model="searchInfo.allowanceType" placeholder="请选择补贴类型" :data="allowance_typeOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="关联商户" prop="merchantId">
  <el-select v-model="searchInfo.merchantId" filterable placeholder="请选择关联商户" :clearable="false">
    <el-option v-for="(item,key) in dataSource.merchantId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
            <el-form-item label="备注" prop="remark">
  <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
</el-form-item>
           
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
            <ExportTemplate  template-id="app_AppDealer" />
            <ExportExcel  template-id="app_AppDealer" filterDeleted/>
            <ImportExcel  template-id="app_AppDealer" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="经销商名称" prop="name" width="120" />

            <el-table-column align="left" label="销售提成" prop="salesCommission" width="120" />

            <el-table-column align="left" label="提成类型" prop="commissionType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.commissionType,commission_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="费用补贴" prop="expenseAllowance" width="120" />

            <el-table-column align="left" label="补贴类型" prop="allowanceType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.allowanceType,allowance_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="关联商户" prop="merchantId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.merchantId,scope.row.merchantId) }}</span>
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateAppDealerFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="经销商名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入经销商名称" />
</el-form-item>
             <el-form-item label="销售提成:" prop="salesCommission">
    <el-input-number v-model="formData.salesCommission" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="提成类型:" prop="commissionType">
    <el-select v-model="formData.commissionType" placeholder="请选择提成类型" style="width:100%" :clearable="true">
      <el-option v-for="(item,key) in commission_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="费用补贴:" prop="expenseAllowance">
    <el-input-number v-model="formData.expenseAllowance" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="补贴类型:" prop="allowanceType">
    <el-select v-model="formData.allowanceType" placeholder="请选择补贴类型" style="width:100%" :clearable="true">
      <el-option v-for="(item,key) in allowance_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="false" placeholder="请输入备注" />
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
    getAppDealerDataSource,
  createAppDealer,
  deleteAppDealer,
  deleteAppDealerByIds,
  updateAppDealer,
  findAppDealer,
  getAppDealerList
} from '@/plugin/app/api/appDealer'

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
    name: 'AppDealer'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const commission_typeOptions = ref([])
const allowance_typeOptions = ref([])
const formData = ref({
            name: '',
            salesCommission: 0,
            commissionType: 0,
            expenseAllowance: 0,
            allowanceType: 0,
            remark: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getAppDealerDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               name : [{
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
               commissionType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               allowanceType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
  const table = await getAppDealerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    commission_typeOptions.value = await getDictFunc('commission_type')
    allowance_typeOptions.value = await getDictFunc('allowance_type')
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
            deleteAppDealerFunc(row)
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
      const res = await deleteAppDealerByIds({ IDs })
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
const updateAppDealerFunc = async(row) => {
    const res = await findAppDealer({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppDealerFunc = async (row) => {
    const res = await deleteAppDealer({ ID: row.ID })
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
        name: '',
        salesCommission: 0,
        commissionType: 0,
        expenseAllowance: 0,
        allowanceType: 0,
        remark: '',
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
                  res = await createAppDealer(formData.value)
                  break
                case 'update':
                  res = await updateAppDealer(formData.value)
                  break
                default:
                  res = await createAppDealer(formData.value)
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
  const res = await findAppDealer({ ID: row.ID })
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
