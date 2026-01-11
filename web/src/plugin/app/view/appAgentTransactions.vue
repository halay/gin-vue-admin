
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
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
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="id字段" prop="id" width="120" />

            <el-table-column align="left" label="createdAt字段" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="updatedAt字段" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="deletedAt字段" prop="deletedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.deletedAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="订单号" prop="orderNo" width="120" />

            <el-table-column align="left" label="订单ID" prop="orderId" width="120" />

            <el-table-column align="left" label="订单金额" prop="orderAmount" width="120" />

            <el-table-column align="left" label="分润基数(87%)" prop="baseAmount" width="120" />

            <el-table-column align="left" label="商户ID" prop="merchantId" width="120" />

            <el-table-column align="left" label="受益人ID" prop="beneficiaryId" width="120" />

            <el-table-column align="left" label="来源用户ID" prop="sourceUserId" width="120" />

            <el-table-column align="left" label="触发的代理等级ID" prop="agentLevelId" width="120" />

            <el-table-column align="left" label="触发的代理等级名称" prop="agentLevelName" width="120" />

            <el-table-column align="left" label="分红范围" prop="dividendScope" width="120" />

            <el-table-column align="left" label="比例1" prop="rate1" width="120" />

            <el-table-column align="left" label="金额1" prop="amount1" width="120" />

            <el-table-column align="left" label="比例2" prop="rate2" width="120" />

            <el-table-column align="left" label="金额2" prop="amount2" width="120" />

            <el-table-column align="left" label="比例3" prop="rate3" width="120" />

            <el-table-column align="left" label="金额3" prop="amount3" width="120" />

            <el-table-column align="left" label="比例4" prop="rate4" width="120" />

            <el-table-column align="left" label="金额4" prop="amount4" width="120" />

            <el-table-column align="left" label="比例5" prop="rate5" width="120" />

            <el-table-column align="left" label="金额5" prop="amount5" width="120" />

            <el-table-column align="left" label="总分润金额" prop="totalAmount" width="120" />

            <el-table-column align="left" label="描述" prop="description" width="120" />

            <el-table-column align="left" label="来源id" prop="sourceId" width="120" />

            <el-table-column align="left" label="来源" prop="source" width="120" />

            <el-table-column align="left" label="提现状态" prop="reflectStatus" width="120" />
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="id字段">
    {{ detailForm.id }}
</el-descriptions-item>
                 <el-descriptions-item label="createdAt字段">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                 <el-descriptions-item label="updatedAt字段">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
                 <el-descriptions-item label="deletedAt字段">
    {{ detailForm.deletedAt }}
</el-descriptions-item>
                 <el-descriptions-item label="订单号">
    {{ detailForm.orderNo }}
</el-descriptions-item>
                 <el-descriptions-item label="订单ID">
    {{ detailForm.orderId }}
</el-descriptions-item>
                 <el-descriptions-item label="订单金额">
    {{ detailForm.orderAmount }}
</el-descriptions-item>
                 <el-descriptions-item label="分润基数(87%)">
    {{ detailForm.baseAmount }}
</el-descriptions-item>
                 <el-descriptions-item label="商户ID">
    {{ detailForm.merchantId }}
</el-descriptions-item>
                 <el-descriptions-item label="受益人ID">
    {{ detailForm.beneficiaryId }}
</el-descriptions-item>
                 <el-descriptions-item label="来源用户ID">
    {{ detailForm.sourceUserId }}
</el-descriptions-item>
                 <el-descriptions-item label="触发的代理等级ID">
    {{ detailForm.agentLevelId }}
</el-descriptions-item>
                 <el-descriptions-item label="触发的代理等级名称">
    {{ detailForm.agentLevelName }}
</el-descriptions-item>
                 <el-descriptions-item label="分红范围">
    {{ detailForm.dividendScope }}
</el-descriptions-item>
                 <el-descriptions-item label="比例1">
    {{ detailForm.rate1 }}
</el-descriptions-item>
                 <el-descriptions-item label="金额1">
    {{ detailForm.amount1 }}
</el-descriptions-item>
                 <el-descriptions-item label="比例2">
    {{ detailForm.rate2 }}
</el-descriptions-item>
                 <el-descriptions-item label="金额2">
    {{ detailForm.amount2 }}
</el-descriptions-item>
                 <el-descriptions-item label="比例3">
    {{ detailForm.rate3 }}
</el-descriptions-item>
                 <el-descriptions-item label="金额3">
    {{ detailForm.amount3 }}
</el-descriptions-item>
                 <el-descriptions-item label="比例4">
    {{ detailForm.rate4 }}
</el-descriptions-item>
                 <el-descriptions-item label="金额4">
    {{ detailForm.amount4 }}
</el-descriptions-item>
                 <el-descriptions-item label="比例5">
    {{ detailForm.rate5 }}
</el-descriptions-item>
                 <el-descriptions-item label="金额5">
    {{ detailForm.amount5 }}
</el-descriptions-item>
                 <el-descriptions-item label="总分润金额">
    {{ detailForm.totalAmount }}
</el-descriptions-item>
                 <el-descriptions-item label="描述">
    {{ detailForm.description }}
</el-descriptions-item>
                 <el-descriptions-item label="来源id">
    {{ detailForm.sourceId }}
</el-descriptions-item>
                 <el-descriptions-item label="来源">
    {{ detailForm.source }}
</el-descriptions-item>
                 <el-descriptions-item label="提现状态">
    {{ detailForm.reflectStatus }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAppAgentTransactions,
  deleteAppAgentTransactions,
  deleteAppAgentTransactionsByIds,
  updateAppAgentTransactions,
  findAppAgentTransactions,
  getAppAgentTransactionsList
} from '@/plugin/app/api/appAgentTransactions'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'AppAgentTransactions'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        })



// 验证规则
const rule = reactive({
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
  const table = await getAppAgentTransactionsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteAppAgentTransactionsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteAppAgentTransactionsByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAppAgentTransactionsFunc = async(row) => {
    const res = await findAppAgentTransactions({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppAgentTransactionsFunc = async (row) => {
    const res = await deleteAppAgentTransactions({ id: row.id })
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
                  res = await createAppAgentTransactions(formData.value)
                  break
                case 'update':
                  res = await updateAppAgentTransactions(formData.value)
                  break
                default:
                  res = await createAppAgentTransactions(formData.value)
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
  const res = await findAppAgentTransactions({ id: row.id })
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
