
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
      
            <el-form-item label="商品名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="商品类型" prop="type">
    <el-tree-select v-model="formData.type" placeholder="请选择商品类型" :data="product_typeOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="状态" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="product_statusOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="所属分类" prop="categoryId">
  <el-select v-model="searchInfo.categoryId" filterable placeholder="请选择所属分类" :clearable="true">
    <el-option v-for="(item,key) in dataSource.categoryId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="封面图" prop="coverImage">
  <el-input v-model="searchInfo.coverImage" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="轮播图" prop="carouselImages">
  <el-input v-model="searchInfo.carouselImages" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="视频" prop="video">
  <el-input v-model="searchInfo.video" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="商品详情" prop="detail">
  <el-input v-model="searchInfo.detail" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="是否多规格" prop="hasVariants">
  <el-select v-model="searchInfo.hasVariants" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
          
          <el-form-item label="开启积分支付" prop="enablePoints">
  <el-select v-model="searchInfo.enablePoints" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
          
          <el-form-item label="商品价格" prop="price">
  <el-input v-model.number="searchInfo.price" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="所需积分" prop="points">
  <el-input v-model.number="searchInfo.points" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="库存" prop="stock">
  <el-input v-model.number="searchInfo.stock" placeholder="搜索条件" />
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
            <ExportTemplate  template-id="app_Product" />
            <ExportExcel  template-id="app_Product" filterDeleted/>
            <ImportExcel  template-id="app_Product" @on-success="getTableData" />
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
        
            <el-table-column sortable align="left" label="商品名称" prop="name" width="120" />

            <el-table-column sortable align="left" label="商品类型" prop="type" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.type,product_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column sortable align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,product_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column sortable align="left" label="所属分类" prop="categoryId">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.categoryId,scope.row.categoryId) }}</span>
    </template>
</el-table-column>
            <el-table-column label="封面图" prop="coverImage" width="200">
    <template #default="scope">
      <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.coverImage)" fit="cover"/>
    </template>
</el-table-column>
            <el-table-column sortable align="left" label="是否多规格" prop="hasVariants" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.hasVariants) }}</template>
</el-table-column>
            <el-table-column sortable align="left" label="开启积分支付" prop="enablePoints" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.enablePoints) }}</template>
</el-table-column>
            <el-table-column sortable align="left" label="商品价格" prop="price" width="120" />

            <el-table-column sortable align="left" label="所需积分" prop="points" width="120" />

            <el-table-column sortable align="left" label="库存" prop="stock" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateProductFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="商品名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入商品名称" />
</el-form-item>
             <el-form-item label="商品类型:" prop="type">
    <el-tree-select v-model="formData.type" placeholder="请选择商品类型" :data="product_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="product_statusOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
             <el-form-item label="所属分类:" prop="categoryId">
    <el-select v-model="formData.categoryId" placeholder="请选择所属分类" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.categoryId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="封面图:" prop="coverImage">
    <SelectImage
     v-model="formData.coverImage"
     file-type="image"
    />
</el-form-item>
            <el-form-item label="轮播图:" prop="carouselImages">
    <SelectImage
     multiple
     v-model="formData.carouselImages"
     file-type="image"
     />
</el-form-item>
             <el-form-item label="详情图片模式:" prop="detailImages">
    <SelectImage
     multiple
     v-model="formData.detailImages"
     file-type="image"
     />
</el-form-item>
             <el-form-item label="视频:" prop="video">
    <SelectImage
    v-model="formData.video"
    file-type="video"
    />
</el-form-item>
             <el-form-item label="详情图文模式:" prop="detail">
    <RichEdit v-model="formData.detail"/>
</el-form-item>
             <el-form-item label="是否多规格:" prop="hasVariants">
    <el-switch v-model="formData.hasVariants" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
             <el-form-item label="开启积分支付:" prop="enablePoints">
    <el-switch v-model="formData.enablePoints" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
             <el-form-item label="商品价格:" prop="price">
    <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="所需积分:" prop="points">
    <el-input v-model.number="formData.points" :clearable="true" placeholder="请输入所需积分" />
</el-form-item>
             <el-form-item label="库存:" prop="stock">
    <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入库存" />
</el-form-item>
        <el-form-item label="税点(%)" prop="taxRate">
          <el-input-number v-model="formData.taxRate" :precision="2" :step="0.1" :min="0" :max="100" />
        </el-form-item>
        <el-form-item label="赠送积分" prop="giftPoints">
          <el-input-number v-model="formData.giftPoints" :min="0" />
        </el-form-item>
        <el-form-item label="关联商户" prop="merchantId" v-show="false">
          <el-input v-model.number="formData.merchantId" disabled placeholder="请输入关联商户" />
        </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="轮播图">
    <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailForm.carouselImages)" :initial-index="index" v-for="(item,index) in detailForm.carouselImages" :key="index" :src="getUrl(item)" fit="cover" />
</el-descriptions-item>
                 <el-descriptions-item label="详情图片模式">
    <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailForm.detailImages)" :initial-index="index" v-for="(item,index) in detailForm.detailImages" :key="index" :src="getUrl(item)" fit="cover" />
</el-descriptions-item>
                 <el-descriptions-item label="视频">
    {{ detailForm.video }}
</el-descriptions-item>
                 <el-descriptions-item label="详情图文模式">
    <RichView v-model="detailForm.detail" />
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getProductDataSource,
  createProduct,
  deleteProduct,
  deleteProductByIds,
  updateProduct,
  findProduct,
  getProductList
} from '@/plugin/app/api/product'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

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


import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'Product'
})

const userStore = useUserStore()


// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const product_statusOptions = ref([])
const product_typeOptions = ref([])
  const formData = ref({
            name: '',
            type: '',
            status: '',
            categoryId: undefined,
            coverImage: "",
            carouselImages: [],
            detailImages: [],
            video: "",
            detail: '',
            hasVariants: false,
            enablePoints: false,
            price: 0,
            points: 0,
            stock: 0,
            taxRate: 0,
            giftPoints: 0,
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getProductDataSource()
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
               type : [{
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
               categoryId : [{
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"created_at",
    ID:"id",
            name: 'name',
            type: 'type',
            status: 'status',
            categoryId: 'category_id',
            hasVariants: 'has_variants',
            enablePoints: 'enable_points',
            price: 'price',
            points: 'points',
            stock: 'stock',
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
    if (searchInfo.value.hasVariants === ""){
        searchInfo.value.hasVariants=null
    }
    if (searchInfo.value.enablePoints === ""){
        searchInfo.value.enablePoints=null
    }
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
  const table = await getProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    product_statusOptions.value = await getDictFunc('product_status')
    product_typeOptions.value = await getDictFunc('product_type')
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
            deleteProductFunc(row)
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
      const res = await deleteProductByIds({ IDs })
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
const updateProductFunc = async(row) => {
    const res = await findProduct({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteProductFunc = async (row) => {
    const res = await deleteProduct({ ID: row.ID })
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
    formData.value.merchantId = userStore.userInfo.merchantId
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
  formData.value = {
        name: '',
        type: '',
        status: '',
        categoryId: undefined,
        coverImage: "",
        carouselImages: [],
        detailImages: [],
        video: "",
        detail: '',
        hasVariants: false,
        enablePoints: false,
        price: 0,
       points: 0,
      stock: 0,
      taxRate: 0,
      giftPoints: 0,
      merchantId: userStore.userInfo.merchantId,
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
                  res = await createProduct(formData.value)
                  break
                case 'update':
                  res = await updateProduct(formData.value)
                  break
                default:
                  res = await createProduct(formData.value)
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
  const res = await findProduct({ ID: row.ID })
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
