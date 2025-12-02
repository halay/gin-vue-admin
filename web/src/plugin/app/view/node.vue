
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
      
            <el-form-item label="节点名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="系统层级ID" prop="systemLevel">
    <el-tree-select v-model="searchInfo.systemLevel" placeholder="请选择系统层级ID" :data="system_levelOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="限额数量(0为不限)" prop="limitCount">
  <el-input v-model.number="searchInfo.limitCount" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="加盟费用" prop="joinFee">
  <el-input v-model.number="searchInfo.joinFee" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="结算货币" prop="settlementCurrency">
  <el-input v-model="searchInfo.settlementCurrency" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="一线城市/特定区域加收比例(%)" prop="citySurchargePercent">
  <el-input v-model.number="searchInfo.citySurchargePercent" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="代币配送倍数(X倍)" prop="tokenDistributionX">
  <el-input v-model.number="searchInfo.tokenDistributionX" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="释放周期(周)" prop="releaseWeeks">
  <el-input v-model.number="searchInfo.releaseWeeks" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="平台消费现金券赠送(%)" prop="cashCouponPercent">
  <el-input v-model.number="searchInfo.cashCouponPercent" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="消费级差PV分润(%)" prop="pvBonusPercent">
  <el-input v-model.number="searchInfo.pvBonusPercent" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="RWA品牌上市费用分成(%)" prop="rwaBrandFeeSharePercent">
  <el-input v-model.number="searchInfo.rwaBrandFeeSharePercent" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="入驻年费级差分成(%)" prop="annualFeeSharePercent">
  <el-input v-model.number="searchInfo.annualFeeSharePercent" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="在线交易手续费分成(%)" prop="onlineTradeFeeSharePercent">
  <el-input v-model.number="searchInfo.onlineTradeFeeSharePercent" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="股权(%)" prop="equityPercent">
  <el-input v-model.number="searchInfo.equityPercent" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="状态" prop="status">
    <el-tree-select v-model="searchInfo.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="排序" prop="sort">
  <el-input v-model.number="searchInfo.sort" placeholder="搜索条件" />
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
            <ExportTemplate  template-id="app_Node" />
            <ExportExcel  template-id="app_Node" filterDeleted/>
            <ImportExcel  template-id="app_Node" @on-success="getTableData" />
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
        
            <el-table-column sortable align="left" label="节点名称" prop="name" width="120" />

            <el-table-column align="left" label="系统层级ID" prop="systemLevel" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.systemLevel,system_levelOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="限额数量(0为不限)" prop="limitCount" width="120" />

            <el-table-column align="left" label="加盟费用" prop="joinFee" width="120" />

            <el-table-column align="left" label="结算货币" prop="settlementCurrency" width="120" />

            <el-table-column align="left" label="一线城市/特定区域加收比例(%)" prop="citySurchargePercent" width="120" />

            <el-table-column align="left" label="代币配送倍数(X倍)" prop="tokenDistributionX" width="120" />

            <el-table-column align="left" label="释放周期(周)" prop="releaseWeeks" width="120" />

            <el-table-column align="left" label="平台消费现金券赠送(%)" prop="cashCouponPercent" width="120" />

            <el-table-column align="left" label="消费级差PV分润(%)" prop="pvBonusPercent" width="120" />

            <el-table-column align="left" label="RWA品牌上市费用分成(%)" prop="rwaBrandFeeSharePercent" width="120" />

            <el-table-column align="left" label="入驻年费级差分成(%)" prop="annualFeeSharePercent" width="120" />

            <el-table-column align="left" label="在线交易手续费分成(%)" prop="onlineTradeFeeSharePercent" width="120" />

            <el-table-column align="left" label="股权(%)" prop="equityPercent" width="120" />

            <el-table-column align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="排序" prop="sort" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateNodeFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="节点名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入节点名称" />
</el-form-item>
             <el-form-item label="系统层级ID:" prop="systemLevel">
    <el-tree-select v-model="formData.systemLevel" placeholder="请选择系统层级ID" :data="system_levelOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="限额数量(0为不限):" prop="limitCount">
    <el-input v-model.number="formData.limitCount" :clearable="false" placeholder="请输入限额数量(0为不限)" />
</el-form-item>
             <el-form-item label="加盟费用:" prop="joinFee">
    <el-input-number v-model="formData.joinFee" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="结算货币:" prop="settlementCurrency">
    <el-input v-model="formData.settlementCurrency" :clearable="false" placeholder="请输入结算货币" />
</el-form-item>
             <el-form-item label="一线城市/特定区域加收比例(%):" prop="citySurchargePercent">
    <el-input-number v-model="formData.citySurchargePercent" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="代币配送倍数(X倍):" prop="tokenDistributionX">
    <el-input-number v-model="formData.tokenDistributionX" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="释放周期(周):" prop="releaseWeeks">
    <el-input v-model.number="formData.releaseWeeks" :clearable="false" placeholder="请输入释放周期(周)" />
</el-form-item>
             <el-form-item label="平台消费现金券赠送(%):" prop="cashCouponPercent">
    <el-input-number v-model="formData.cashCouponPercent" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="消费级差PV分润(%):" prop="pvBonusPercent">
    <el-input-number v-model="formData.pvBonusPercent" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="RWA品牌上市费用分成(%):" prop="rwaBrandFeeSharePercent">
    <el-input-number v-model="formData.rwaBrandFeeSharePercent" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="入驻年费级差分成(%):" prop="annualFeeSharePercent">
    <el-input-number v-model="formData.annualFeeSharePercent" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="在线交易手续费分成(%):" prop="onlineTradeFeeSharePercent">
    <el-input-number v-model="formData.onlineTradeFeeSharePercent" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="股权(%):" prop="equityPercent">
    <el-input-number v-model="formData.equityPercent" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="排序:" prop="sort">
    <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入排序" />
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
  createNode,
  deleteNode,
  deleteNodeByIds,
  updateNode,
  findNode,
  getNodeList
} from '@/plugin/app/api/node'

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
    name: 'Node'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const system_levelOptions = ref([])
const statusOptions = ref([])
const formData = ref({
            name: '',
            systemLevel: '',
            limitCount: 0,
            joinFee: 0,
            settlementCurrency: '',
            citySurchargePercent: 0,
            tokenDistributionX: 0,
            releaseWeeks: 0,
            cashCouponPercent: 0,
            pvBonusPercent: 0,
            rwaBrandFeeSharePercent: 0,
            annualFeeSharePercent: 0,
            onlineTradeFeeSharePercent: 0,
            equityPercent: 0,
            status: '',
            sort: 0,
        })



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
               systemLevel : [{
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
            name: 'name',
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
  const table = await getNodeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    system_levelOptions.value = await getDictFunc('system_level')
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
            deleteNodeFunc(row)
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
      const res = await deleteNodeByIds({ IDs })
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
const updateNodeFunc = async(row) => {
    const res = await findNode({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteNodeFunc = async (row) => {
    const res = await deleteNode({ ID: row.ID })
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
        systemLevel: '',
        limitCount: 0,
        joinFee: 0,
        settlementCurrency: '',
        citySurchargePercent: 0,
        tokenDistributionX: 0,
        releaseWeeks: 0,
        cashCouponPercent: 0,
        pvBonusPercent: 0,
        rwaBrandFeeSharePercent: 0,
        annualFeeSharePercent: 0,
        onlineTradeFeeSharePercent: 0,
        equityPercent: 0,
        status: '',
        sort: 0,
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
                  res = await createNode(formData.value)
                  break
                case 'update':
                  res = await updateNode(formData.value)
                  break
                default:
                  res = await createNode(formData.value)
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
  const res = await findNode({ ID: row.ID })
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
