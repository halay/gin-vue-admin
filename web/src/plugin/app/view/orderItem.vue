
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
      
            <el-form-item label="订单" prop="orderId">
  <el-select v-model="searchInfo.orderId" filterable placeholder="请选择订单" :clearable="false">
    <el-option v-for="(item,key) in dataSource.orderId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
            <el-form-item label="商品" prop="productId">
  <el-select v-model="searchInfo.productId" filterable placeholder="请选择商品" :clearable="false">
    <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
            <el-form-item label="SKU" prop="skuId">
  <el-select v-model="searchInfo.skuId" filterable placeholder="请选择SKU" :clearable="false">
    <el-option v-for="(item,key) in dataSource.skuId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
            <el-form-item label="商品名快照" prop="productName">
  <el-input v-model="searchInfo.productName" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="单价" prop="price">
  <el-input v-model.number="searchInfo.price" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="单件积分" prop="points">
  <el-input v-model.number="searchInfo.points" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="数量" prop="quantity">
  <el-input v-model.number="searchInfo.quantity" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="小计金额" prop="totalAmount">
  <el-input v-model.number="searchInfo.totalAmount" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="小计积分" prop="totalPoints">
  <el-input v-model.number="searchInfo.totalPoints" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="图片" prop="coverImage">
  <el-input v-model="searchInfo.coverImage" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="商户" prop="merchantId">
  <el-input v-model.number="searchInfo.merchantId" placeholder="搜索条件" />
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="SKU属性" prop="skuAttrs">
  <el-input v-model="searchInfo.skuAttrs" placeholder="搜索条件" />
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
            <ExportTemplate  template-id="app_OrderItem" />
            <ExportExcel  template-id="app_OrderItem" filterDeleted/>
            <ImportExcel  template-id="app_OrderItem" @on-success="getTableData" />
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
        
        <el-table-column align="left" label="用户邮箱" prop="userEmail" width="180" />

            <el-table-column align="left" label="订单" prop="orderId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.orderId,scope.row.orderId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="商品" prop="productId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.productId,scope.row.productId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="SKU" prop="skuId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.skuId,scope.row.skuId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="商品名快照" prop="productName" width="120" />

            <el-table-column align="left" label="单价" prop="price" width="120" />

            <el-table-column align="left" label="单件积分" prop="points" width="120" />

            <el-table-column align="left" label="数量" prop="quantity" width="120" />

            <el-table-column align="left" label="小计金额" prop="totalAmount" width="120" />

            <el-table-column align="left" label="小计积分" prop="totalPoints" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateOrderItemFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="订单:" prop="orderId">
    <el-select v-model="formData.orderId" placeholder="请选择订单" filterable style="width:100%" :clearable="false">
        <el-option v-for="(item,key) in dataSource.orderId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="商品:" prop="productId">
    <el-select v-model="formData.productId" placeholder="请选择商品" filterable style="width:100%" :clearable="false">
        <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="SKU:" prop="skuId">
    <el-select v-model="formData.skuId" placeholder="请选择SKU" filterable style="width:100%" :clearable="false">
        <el-option v-for="(item,key) in dataSource.skuId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="商品名快照:" prop="productName">
    <el-input v-model="formData.productName" :clearable="false" placeholder="请输入商品名快照" />
</el-form-item>
             <el-form-item label="SKU属性:" prop="skuAttrs">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.skuAttrs 后端会按照json的类型进行存取
    {{ formData.skuAttrs }}
</el-form-item>
             <el-form-item label="单价:" prop="price">
    <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="单件积分:" prop="points">
    <el-input v-model.number="formData.points" :clearable="false" placeholder="请输入单件积分" />
</el-form-item>
             <el-form-item label="数量:" prop="quantity">
    <el-input v-model.number="formData.quantity" :clearable="false" placeholder="请输入数量" />
</el-form-item>
             <el-form-item label="小计金额:" prop="totalAmount">
    <el-input-number v-model="formData.totalAmount" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="小计积分:" prop="totalPoints">
    <el-input v-model.number="formData.totalPoints" :clearable="false" placeholder="请输入小计积分" />
</el-form-item>
             <el-form-item label="图片:" prop="coverImage">
    <SelectImage
     v-model="formData.coverImage"
     file-type="image"
    />
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
    getOrderItemDataSource,
  createOrderItem,
  deleteOrderItem,
  deleteOrderItemByIds,
  updateOrderItem,
  findOrderItem,
  getOrderItemList
} from '@/plugin/app/api/orderItem'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

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
    name: 'OrderItem'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            orderId: undefined,
            productId: undefined,
            skuId: undefined,
            productName: '',
            skuAttrs: {},
            price: 0,
            points: 0,
            quantity: 0,
            totalAmount: 0,
            totalPoints: 0,
            coverImage: "",
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getOrderItemDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               orderId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               productId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               skuId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               productName : [{
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
               price : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               quantity : [{
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
  const table = await getOrderItemList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteOrderItemFunc(row)
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
      const res = await deleteOrderItemByIds({ IDs })
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
const updateOrderItemFunc = async(row) => {
    const res = await findOrderItem({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOrderItemFunc = async (row) => {
    const res = await deleteOrderItem({ ID: row.ID })
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
        orderId: undefined,
        productId: undefined,
        skuId: undefined,
        productName: '',
        skuAttrs: {},
        price: 0,
        points: 0,
        quantity: 0,
        totalAmount: 0,
        totalPoints: 0,
        coverImage: "",
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
                  res = await createOrderItem(formData.value)
                  break
                case 'update':
                  res = await updateOrderItem(formData.value)
                  break
                default:
                  res = await createOrderItem(formData.value)
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
  const res = await findOrderItem({ ID: row.ID })
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
