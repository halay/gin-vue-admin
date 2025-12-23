
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
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
            <el-table-column align="left" label="品牌方" prop="merchantId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.merchantId,scope.row.merchantId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="名称" prop="name" width="120" />

            <el-table-column align="left" label="代理资格(万)" prop="qualificationAmount" width="120" />

            <el-table-column align="left" label="分红范围" prop="dividendScope" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.dividendScope,bonusRangeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="业务提成" prop="cfRate1" width="120" />

            <el-table-column align="left" label="销售费用补贴" prop="cfRate2" width="120" />

            <el-table-column align="left" label="经销补贴" prop="cfRate3" width="120" />

            <el-table-column align="left" label="分红比例" prop="cfRate4" width="120" />

            <el-table-column align="left" label="额外补贴" prop="cfRate5" width="120" />

            <el-table-column align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,statusOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateAgentLevelFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入名称" />
</el-form-item>
             <el-form-item label="代理资格(万):" prop="qualificationAmount">
    <el-input v-model.number="formData.qualificationAmount" :clearable="true" placeholder="请输入代理资格(万)" />
</el-form-item>
             <el-form-item label="分红范围:" prop="dividendScope">
    <el-tree-select v-model="formData.dividendScope" placeholder="请选择分红范围" :data="bonusRangeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="业务提成:" prop="cfRate1">
    <el-input-number v-model="formData.cfRate1" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="销售费用补贴:" prop="cfRate2">
    <el-input-number v-model="formData.cfRate2" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="经销补贴:" prop="cfRate3">
    <el-input-number v-model="formData.cfRate3" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="分红比例:" prop="cfRate4">
    <el-input-number v-model="formData.cfRate4" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="额外补贴:" prop="cfRate5">
    <el-input-number v-model="formData.cfRate5" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="品牌方">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.merchantId,detailForm.merchantId) }}</span>
    </template>
</el-descriptions-item>
                 <el-descriptions-item label="名称">
    {{ detailForm.name }}
</el-descriptions-item>
                 <el-descriptions-item label="代理资格(万)">
    {{ detailForm.qualificationAmount }}
</el-descriptions-item>
                 <el-descriptions-item label="分红范围">
    {{ detailForm.dividendScope }}
</el-descriptions-item>
                 <el-descriptions-item label="业务提成">
    {{ detailForm.cfRate1 }}
</el-descriptions-item>
                 <el-descriptions-item label="销售费用补贴">
    {{ detailForm.cfRate2 }}
</el-descriptions-item>
                 <el-descriptions-item label="经销补贴">
    {{ detailForm.cfRate3 }}
</el-descriptions-item>
                 <el-descriptions-item label="分红比例">
    {{ detailForm.cfRate4 }}
</el-descriptions-item>
                 <el-descriptions-item label="额外补贴">
    {{ detailForm.cfRate5 }}
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
    getAgentLevelDataSource,
  createAgentLevel,
  deleteAgentLevel,
  deleteAgentLevelByIds,
  updateAgentLevel,
  findAgentLevel,
  getAgentLevelList
} from '@/plugin/app/api/agentLevel'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'AgentLevel'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const bonusRangeOptions = ref([])
const statusOptions = ref([])
const formData = ref({
            id: undefined,
            name: '',
            qualificationAmount: undefined,
            dividendScope: '',
            cfRate1: 0,
            cfRate2: 0,
            cfRate3: 0,
            cfRate4: 0,
            cfRate5: 0,
            status: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getAgentLevelDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



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
  const table = await getAgentLevelList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    bonusRangeOptions.value = await getDictFunc('bonusRange')
    statusOptions.value = await getDictFunc('status')
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
            deleteAgentLevelFunc(row)
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
      const res = await deleteAgentLevelByIds({ ids })
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
const updateAgentLevelFunc = async(row) => {
    const res = await findAgentLevel({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAgentLevelFunc = async (row) => {
    const res = await deleteAgentLevel({ id: row.id })
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
        id: undefined,
        name: '',
        qualificationAmount: undefined,
        dividendScope: '',
        cfRate1: 0,
        cfRate2: 0,
        cfRate3: 0,
        cfRate4: 0,
        cfRate5: 0,
        status: '',
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
                  res = await createAgentLevel(formData.value)
                  break
                case 'update':
                  res = await updateAgentLevel(formData.value)
                  break
                default:
                  res = await createAgentLevel(formData.value)
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
  const res = await findAgentLevel({ id: row.id })
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
