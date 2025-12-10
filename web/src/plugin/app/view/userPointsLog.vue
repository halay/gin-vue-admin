
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
      <el-form-item label="记录类型" prop="type">
        <el-select v-model="searchInfo.type" clearable placeholder="请选择类型" style="width:220px">
          <el-option label="积分充值" value="recharge_purchase" />
          <el-option label="积分赠送" value="bonus" />
          <el-option label="购买实体商品" value="buy_physical" />
          <el-option label="购买数字商品" value="buy_digital" />
        </el-select>
      </el-form-item>
      <el-form-item label="交易状态" prop="status">
        <el-select v-model="searchInfo.status" clearable placeholder="请选择状态" style="width:220px">
          <el-option label="成功" value="success" />
          <el-option label="失败" value="failed" />
        </el-select>
      </el-form-item>
      
            <el-form-item label="用户" prop="userId">
  <el-select v-model="searchInfo.userId" filterable placeholder="请选择用户" :clearable="false">
    <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
            <el-form-item label="变动值" prop="change">
  <el-input v-model.number="searchInfo.change" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="变动后余额" prop="balanceAfter">
  <el-input v-model.number="searchInfo.balanceAfter" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="原因" prop="reason">
  <el-input v-model="searchInfo.reason" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="关联订单号" prop="orderNo">
  <el-input v-model="searchInfo.orderNo" placeholder="搜索条件" />
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
            <ExportTemplate  template-id="app_UserPointsLog" />
            <ExportExcel  template-id="app_UserPointsLog" filterDeleted/>
            <ImportExcel  template-id="app_UserPointsLog" @on-success="getTableData" />
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
        
            <el-table-column align="left" label="用户" prop="userId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.userId,scope.row.userId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="变动值" prop="change" width="120" />

            <el-table-column align="left" label="变动后余额" prop="balanceAfter" width="120" />

            <el-table-column align="left" label="类型" prop="type" width="120" />
            <el-table-column align="left" label="状态" prop="status" width="120" />
            <el-table-column align="left" label="原因" prop="reason" width="160" />

            <el-table-column align="left" label="关联订单号" prop="orderNo" width="120" />

            <el-table-column align="left" label="备注" prop="remark" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateUserPointsLogFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="用户:" prop="userId">
    <el-select v-model="formData.userId" placeholder="请选择用户" filterable style="width:100%" :clearable="false">
        <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="变动值:" prop="change">
    <el-input v-model.number="formData.change" :clearable="false" placeholder="请输入变动值" />
</el-form-item>
             <el-form-item label="变动后余额:" prop="balanceAfter">
    <el-input v-model.number="formData.balanceAfter" :clearable="false" placeholder="请输入变动后余额" />
</el-form-item>
             <el-form-item label="原因:" prop="reason">
    <el-input v-model="formData.reason" :clearable="false" placeholder="请输入原因" />
</el-form-item>
             <el-form-item label="关联订单号:" prop="orderNo">
    <el-input v-model="formData.orderNo" :clearable="false" placeholder="请输入关联订单号" />
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
    getUserPointsLogDataSource,
  createUserPointsLog,
  deleteUserPointsLog,
  deleteUserPointsLogByIds,
  updateUserPointsLog,
  findUserPointsLog,
  getUserPointsLogList
} from '@/plugin/app/api/userPointsLog'

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
    name: 'UserPointsLog'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            userId: undefined,
            change: 0,
            balanceAfter: 0,
            reason: '',
            orderNo: '',
            remark: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getUserPointsLogDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               change : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               balanceAfter : [{
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
  const table = await getUserPointsLogList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteUserPointsLogFunc(row)
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
      const res = await deleteUserPointsLogByIds({ IDs })
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
const updateUserPointsLogFunc = async(row) => {
    const res = await findUserPointsLog({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteUserPointsLogFunc = async (row) => {
    const res = await deleteUserPointsLog({ ID: row.ID })
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
        userId: undefined,
        change: 0,
        balanceAfter: 0,
        reason: '',
        orderNo: '',
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
                  res = await createUserPointsLog(formData.value)
                  break
                case 'update':
                  res = await updateUserPointsLog(formData.value)
                  break
                default:
                  res = await createUserPointsLog(formData.value)
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
  const res = await findUserPointsLog({ ID: row.ID })
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
