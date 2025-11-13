
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
      
            <el-form-item label="订单号" prop="orderNo">
  <el-input v-model="searchInfo.orderNo" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="下单用户" prop="userId">
  <el-select v-model="searchInfo.userId" filterable placeholder="请选择下单用户" :clearable="false">
    <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
            <el-form-item label="支付方式" prop="payMethod">
    <el-tree-select v-model="formData.payMethod" placeholder="请选择支付方式" :data="pay_methodOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="支付状态" prop="payStatus">
    <el-tree-select v-model="formData.payStatus" placeholder="请选择支付状态" :data="pay_statusOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="订单状态" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择订单状态" :data="order_statusOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="发货状态" prop="deliveryStatus">
    <el-tree-select v-model="formData.deliveryStatus" placeholder="请选择发货状态" :data="delivery_statusOptions" style="width:100%" filterable :clearable="false" check-strictly ></el-tree-select>
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="商户" prop="merchantId">
  <el-input v-model.number="searchInfo.merchantId" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="总金额" prop="totalAmount">
  <el-input v-model.number="searchInfo.totalAmount" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="总积分" prop="totalPoints">
  <el-input v-model.number="searchInfo.totalPoints" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="收货人" prop="consigneeName">
  <el-input v-model="searchInfo.consigneeName" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="收货人手机号" prop="consigneePhone">
  <el-input v-model="searchInfo.consigneePhone" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="收货地址" prop="address">
  <el-input v-model="searchInfo.address" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="快递公司" prop="expressName">
  <el-input v-model="searchInfo.expressName" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="快递单号" prop="expressNo">
  <el-input v-model="searchInfo.expressNo" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="备注" prop="remark">
  <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
</el-form-item>
          
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
            <ExportTemplate  template-id="app_Order" />
            <ExportExcel  template-id="app_Order" filterDeleted/>
            <ImportExcel  template-id="app_Order" @on-success="getTableData" />
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
        
            <el-table-column sortable align="left" label="订单号" prop="orderNo" width="120" />

            <el-table-column align="left" label="下单用户" prop="userId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.userId,scope.row.userId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="商户" prop="merchantId" width="120" />

            <el-table-column align="left" label="总金额" prop="totalAmount" width="120" />

            <el-table-column align="left" label="总积分" prop="totalPoints" width="120" />

            <el-table-column align="left" label="支付方式" prop="payMethod" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.payMethod,pay_methodOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="支付状态" prop="payStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.payStatus,pay_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="订单状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,order_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="发货状态" prop="deliveryStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.deliveryStatus,delivery_statusOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateOrderFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="订单号:" prop="orderNo">
    <el-input v-model="formData.orderNo" :clearable="false" placeholder="请输入订单号" />
</el-form-item>
             <el-form-item label="下单用户:" prop="userId">
    <el-select v-model="formData.userId" placeholder="请选择下单用户" filterable style="width:100%" :clearable="false">
        <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="总金额:" prop="totalAmount">
    <el-input-number v-model="formData.totalAmount" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="总积分:" prop="totalPoints">
    <el-input v-model.number="formData.totalPoints" :clearable="false" placeholder="请输入总积分" />
</el-form-item>
             <el-form-item label="支付方式:" prop="payMethod">
    <el-tree-select v-model="formData.payMethod" placeholder="请选择支付方式" :data="pay_methodOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="支付状态:" prop="payStatus">
    <el-tree-select v-model="formData.payStatus" placeholder="请选择支付状态" :data="pay_statusOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="订单状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择订单状态" :data="order_statusOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="收货人:" prop="consigneeName">
    <el-input v-model="formData.consigneeName" :clearable="false" placeholder="请输入收货人" />
</el-form-item>
             <el-form-item label="收货人手机号:" prop="consigneePhone">
    <el-input v-model="formData.consigneePhone" :clearable="false" placeholder="请输入收货人手机号" />
</el-form-item>
             <el-form-item label="收货地址:" prop="address">
    <el-input v-model="formData.address" :clearable="false" placeholder="请输入收货地址" />
</el-form-item>
             <el-form-item label="发货状态:" prop="deliveryStatus">
    <el-tree-select v-model="formData.deliveryStatus" placeholder="请选择发货状态" :data="delivery_statusOptions" style="width:100%" filterable :clearable="false" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="快递公司:" prop="expressName">
    <el-input v-model="formData.expressName" :clearable="false" placeholder="请输入快递公司" />
</el-form-item>
             <el-form-item label="快递单号:" prop="expressNo">
    <el-input v-model="formData.expressNo" :clearable="false" placeholder="请输入快递单号" />
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
    getOrderDataSource,
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList
} from '@/plugin/app/api/order'

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
    name: 'Order'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const pay_methodOptions = ref([])
const pay_statusOptions = ref([])
const delivery_statusOptions = ref([])
const order_statusOptions = ref([])
const formData = ref({
            orderNo: '',
            userId: undefined,
            totalAmount: 0,
            totalPoints: 0,
            payMethod: '',
            payStatus: '',
            status: '',
            consigneeName: '',
            consigneePhone: '',
            address: '',
            deliveryStatus: '',
            expressName: '',
            expressNo: '',
            remark: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getOrderDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               orderNo : [{
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
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               payMethod : [{
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
               payStatus : [{
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
            orderNo: 'order_no',
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
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    pay_methodOptions.value = await getDictFunc('pay_method')
    pay_statusOptions.value = await getDictFunc('pay_status')
    delivery_statusOptions.value = await getDictFunc('delivery_status')
    order_statusOptions.value = await getDictFunc('order_status')
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
            deleteOrderFunc(row)
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
      const res = await deleteOrderByIds({ IDs })
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
const updateOrderFunc = async(row) => {
    const res = await findOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOrderFunc = async (row) => {
    const res = await deleteOrder({ ID: row.ID })
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
        orderNo: '',
        userId: undefined,
        totalAmount: 0,
        totalPoints: 0,
        payMethod: '',
        payStatus: '',
        status: '',
        consigneeName: '',
        consigneePhone: '',
        address: '',
        deliveryStatus: '',
        expressName: '',
        expressNo: '',
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
                  res = await createOrder(formData.value)
                  break
                case 'update':
                  res = await updateOrder(formData.value)
                  break
                default:
                  res = await createOrder(formData.value)
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
  const res = await findOrder({ ID: row.ID })
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
