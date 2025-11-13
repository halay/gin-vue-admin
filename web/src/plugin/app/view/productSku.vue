
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
      
            <el-form-item label="商品" prop="productId">
  <el-select v-model="searchInfo.productId" filterable placeholder="请选择商品" :clearable="true">
    <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
            <el-form-item label="SKU编码" prop="skuCode">
  <el-input v-model="searchInfo.skuCode" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="价格" prop="price">
  <el-input v-model.number="searchInfo.price" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="库存" prop="stock">
  <el-input v-model.number="searchInfo.stock" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="状态" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="sku_statusOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="属性组合" prop="attrs">
  <el-input v-model="searchInfo.attrs" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="积分" prop="points">
  <el-input v-model.number="searchInfo.points" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="图片" prop="image">
  <el-input v-model="searchInfo.image" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="关联商户" prop="merchantId">
  <el-input v-model.number="searchInfo.merchantId" placeholder="搜索条件" />
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
            <ExportTemplate  template-id="app_ProductSku" />
            <ExportExcel  template-id="app_ProductSku" filterDeleted/>
            <ImportExcel  template-id="app_ProductSku" @on-success="getTableData" />
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
        
            <el-table-column sortable align="left" label="商品" prop="productId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.productId,scope.row.productId) }}</span>
    </template>
</el-table-column>
            <el-table-column sortable align="left" label="SKU编码" prop="skuCode" width="120" />

            <el-table-column sortable align="left" label="价格" prop="price" width="120" />

            <el-table-column sortable align="left" label="积分" prop="points" width="120" />

            <el-table-column label="图片" prop="image" width="200">
    <template #default="scope">
      <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.image)" fit="cover"/>
    </template>
</el-table-column>
            <el-table-column sortable align="left" label="库存" prop="stock" width="120" />

            <el-table-column sortable align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,sku_statusOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateProductSkuFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="商品:" prop="productId">
    <el-select v-model="formData.productId" placeholder="请选择商品" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="SKU编码:" prop="skuCode">
    <el-input v-model="formData.skuCode" :clearable="true" placeholder="请输入SKU编码" />
</el-form-item>
             <el-form-item label="属性组合:" prop="attrs">
    <div class="w-full">
      <div class="flex items-center mb-2">
        <el-button size="small" type="primary" @click="addAttrRow">添加一行</el-button>
        <el-button size="small" class="ml-2" @click="importFromJson">从JSON导入</el-button>
        <el-button size="small" class="ml-2" @click="exportToJson">导出JSON</el-button>
      </div>
      <div>
        <div v-for="(row, idx) in kvList" :key="idx" class="flex items-center mb-2 gap-2">
          <el-input v-model="row.key" placeholder="属性名，如 颜色" style="width: 220px" />
          <el-input v-model="row.value" placeholder="属性值，如 红" style="width: 220px" />
          <el-button type="danger" link @click="removeAttrRow(idx)">删除</el-button>
        </div>
      </div>
    </div>
  </el-form-item>
             <el-form-item label="价格:" prop="price">
    <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="积分:" prop="points">
    <el-input v-model.number="formData.points" :clearable="true" placeholder="请输入积分" />
</el-form-item>
             <el-form-item label="图片:" prop="image">
    <SelectImage
     v-model="formData.image"
     file-type="image"
    />
</el-form-item>
             <el-form-item label="库存:" prop="stock">
    <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入库存" />
</el-form-item>
             <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="sku_statusOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="属性组合">
    <pre style="white-space:pre-wrap;word-break:break-all;">
{{ typeof detailForm.attrs === 'string' ? detailForm.attrs : JSON.stringify(detailForm.attrs||{}, null, 2) }}
    </pre>
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getProductSkuDataSource,
  createProductSku,
  deleteProductSku,
  deleteProductSkuByIds,
  updateProductSku,
  findProductSku,
  getProductSkuList
} from '@/plugin/app/api/productsku'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'ProductSku'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const sku_statusOptions = ref([])
const formData = ref({
            productId: undefined,
            skuCode: '',
            attrs: {},
            price: 0,
            points: 0,
            image: "",
            stock: 0,
            status: '',
        })
// 属性组合编辑（以键值对形式编辑，最终绑定到 formData.attrs 对象）
const kvList = ref([{ key: '', value: '' }])
const rebuildAttrsFromKvList = () => {
  const obj = {}
  kvList.value.forEach(r => {
    if ((r.key ?? '').toString().trim() !== '') {
      obj[r.key] = r.value
    }
  })
  formData.value.attrs = obj
}
const initKvListFromAttrs = () => {
  const entries = Object.entries(formData.value.attrs || {})
  if (entries.length === 0) {
    kvList.value = [{ key: '', value: '' }]
  } else {
    kvList.value = entries.map(([k, v]) => ({ key: k, value: v }))
  }
}
const addAttrRow = () => {
  kvList.value.push({ key: '', value: '' })
}
const removeAttrRow = (idx) => {
  kvList.value.splice(idx, 1)
  if (kvList.value.length === 0) kvList.value.push({ key: '', value: '' })
}
const importFromJson = async () => {
  try {
    const text = await ElMessageBox.prompt('粘贴 JSON 对象（如 {"颜色":"红","尺码":"M"}）', '从JSON导入', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputValue: JSON.stringify(formData.value.attrs || {}, null, 2),
      inputType: 'textarea'
    })
    if (text && text.value) {
      const obj = JSON.parse(text.value)
      if (obj && typeof obj === 'object' && !Array.isArray(obj)) {
        formData.value.attrs = obj
        initKvListFromAttrs()
      } else {
        ElMessage.error('必须是对象类型，如 {"颜色":"红"}')
      }
    }
  } catch {}
}
const exportToJson = () => {
  rebuildAttrsFromKvList()
  ElMessageBox.alert(`<pre>${JSON.stringify(formData.value.attrs || {}, null, 2)}</pre>`, '导出JSON', {
    dangerouslyUseHTMLString: true
  })
}
watch(kvList, rebuildAttrsFromKvList, { deep: true })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getProductSkuDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               productId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               price : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               stock : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
            productId: 'product_id',
            skuCode: 'sku_code',
            price: 'price',
            points: 'points',
            stock: 'stock',
            status: 'status',
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
  const table = await getProductSkuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    sku_statusOptions.value = await getDictFunc('sku_status')
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
            deleteProductSkuFunc(row)
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
      const res = await deleteProductSkuByIds({ IDs })
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
const updateProductSkuFunc = async(row) => {
    const res = await findProductSku({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        initKvListFromAttrs()
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteProductSkuFunc = async (row) => {
    const res = await deleteProductSku({ ID: row.ID })
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
    kvList.value = [{ key: '', value: '' }]
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        productId: undefined,
        skuCode: '',
        attrs: {},
        price: 0,
        points: 0,
        image: "",
        stock: 0,
        status: '',
        }
    kvList.value = [{ key: '', value: '' }]
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createProductSku(formData.value)
                  break
                case 'update':
                  res = await updateProductSku(formData.value)
                  break
                default:
                  res = await createProductSku(formData.value)
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
  const res = await findProductSku({ ID: row.ID })
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
