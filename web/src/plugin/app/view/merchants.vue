
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
      
            <el-form-item label="商家名" prop="merchantName">
  <el-input v-model="searchInfo.merchantName" placeholder="搜索条件" />
</el-form-item>
           <el-form-item label="创始人" prop="founder">
  <el-input v-model="searchInfo.founder" placeholder="搜索条件" />
</el-form-item>
    
<el-form-item label="创始人电话" prop="founderPhone">
  <el-input v-model="searchInfo.founderPhone" placeholder="搜索条件" />
</el-form-item>
    
<el-form-item label="数字资产名称" prop="digitalAssetName">
  <el-input v-model="searchInfo.digitalAssetName" placeholder="搜索条件" />
</el-form-item>
            <el-form-item label="经营范围描述" prop="businessScope">
  <el-input v-model="searchInfo.businessScope" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="联系人" prop="contactPerson">
  <el-input v-model="searchInfo.contactPerson" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="联系电话" prop="contactPhone">
  <el-input v-model="searchInfo.contactPhone" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="邮箱" prop="email">
  <el-input v-model="searchInfo.email" placeholder="搜索条件" />
</el-form-item>
           
<el-form-item label="状态" prop="status">
  <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
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
            <ExportTemplate  template-id="app_Merchants" />
            <ExportExcel  template-id="app_Merchants" filterDeleted/>
            <ImportExcel  template-id="app_Merchants" @on-success="getTableData" />
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
        
            <el-table-column sortable align="left" label="商家名" prop="merchantName" width="120" />

            <el-table-column label="商户logo" prop="logo" width="200">
    <template #default="scope">
      <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.logo)" fit="cover"/>
    </template>
</el-table-column>
            <el-table-column align="left" label="经营范围描述" prop="businessScope" width="120" />

            <el-table-column align="left" label="地址" prop="address" width="120" />

            <el-table-column align="left" label="联系人" prop="contactPerson" width="120" />

            <el-table-column align="left" label="联系电话" prop="contactPhone" width="120" />

            <el-table-column sortable align="left" label="评分" prop="rating" width="120" />

            <el-table-column align="left" label="邮箱" prop="email" width="120" />

            <el-table-column label="营业执照" prop="businessLicense" width="200">
    <template #default="scope">
      <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.businessLicense)" fit="cover"/>
    </template>
</el-table-column>
            <el-table-column label="身份证照片" prop="idCardImage" width="200">
   <template #default="scope">
      <div class="multiple-img-box">
         <el-image preview-teleported v-for="(item,index) in scope.row.idCardImage" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
     </div>
   </template>
</el-table-column>
<el-table-column align="left" label="创始人" prop="founder" width="120" />

       <el-table-column align="left" label="创始人电话" prop="founderPhone" width="120" />

       <el-table-column label="企业宣传片" prop="corporateVideo" width="200">
   <template #default="scope">
    <video
       style="width: 100px; height: 100px"
       muted
       preload="metadata"
       >
         <source :src="getUrl(scope.row.corporateVideo) + '#t=1'">
       </video>
   </template>
</el-table-column>
       <el-table-column align="left" label="数字资产名称" prop="digitalAssetName" width="120" />

       <el-table-column align="left" label="交易对" prop="tradingPair" width="120" />
            <el-table-column align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,statusOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateMerchantsFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="商家名:" prop="merchantName">
    <el-input v-model="formData.merchantName" :clearable="false" placeholder="请输入商家名" />
</el-form-item>
             <el-form-item label="商户logo:" prop="logo">
    <SelectImage
     v-model="formData.logo"
     file-type="image"
    />
</el-form-item>
             <el-form-item label="经营范围描述:" prop="businessScope">
    <el-input v-model="formData.businessScope" :clearable="false" placeholder="请输入经营范围描述" />
</el-form-item>
             <el-form-item label="地址:" prop="address">
    <el-input v-model="formData.address" :clearable="false" placeholder="请输入地址" />
</el-form-item>
             <el-form-item label="联系人:" prop="contactPerson">
    <el-input v-model="formData.contactPerson" :clearable="false" placeholder="请输入联系人" />
</el-form-item>
             <el-form-item label="联系电话:" prop="contactPhone">
    <el-input v-model="formData.contactPhone" :clearable="false" placeholder="请输入联系电话" />
</el-form-item>
             <el-form-item label="评分:" prop="rating">
    <el-input-number v-model="formData.rating" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
             <el-form-item label="邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
</el-form-item>
             <el-form-item label="营业执照:" prop="businessLicense">
    <SelectImage
     v-model="formData.businessLicense"
     file-type="image"
    />
</el-form-item>
             <el-form-item label="身份证照片:" prop="idCardImage">
    <SelectImage
     multiple
     v-model="formData.idCardImage"
     file-type="image"
     />
</el-form-item>
<el-form-item label="创始人:" prop="founder">
    <el-input v-model="formData.founder" :clearable="true" placeholder="请输入创始人" />
</el-form-item>
     <el-form-item label="创始人电话:" prop="founderPhone">
    <el-input v-model="formData.founderPhone" :clearable="true" placeholder="请输入创始人电话" />
</el-form-item>
     <el-form-item label="创始人简介:" prop="founderDescriptions">
    <RichEdit v-model="formData.founderDescriptions"/>
</el-form-item>
     <el-form-item label="企业介绍:" prop="enterpriseDesc">
    <RichEdit v-model="formData.enterpriseDesc"/>
</el-form-item>
     <el-form-item label="企业宣传片:" prop="corporateVideo">
    <SelectImage
    v-model="formData.corporateVideo"
    file-type="video"
    />
</el-form-item>
     <el-form-item label="荣誉证书:" prop="certificate">
    <SelectImage
     multiple
     v-model="formData.certificate"
     file-type="image"
     />
</el-form-item>
     <el-form-item label="数字资产名称:" prop="digitalAssetName">
    <el-input v-model="formData.digitalAssetName" :clearable="true" placeholder="请输入数字资产名称" />
</el-form-item>
     <el-form-item label="交易对:" prop="tradingPair">
    <el-input v-model="formData.tradingPair" :clearable="true" placeholder="请输入交易对" />
</el-form-item>
            <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
            <el-form-item label="是否推荐:" prop="isRecommended">
    <el-switch v-model="formData.isRecommended" active-text="是" inactive-text="否" />
</el-form-item>
            <el-form-item label="所属分类:" prop="categoryId">
    <el-select v-model="formData.categoryId" placeholder="请选择分类" filterable clearable style="width:100%">
      <el-option v-for="(item,key) in merchantCategories" :key="key" :label="item.Name" :value="item.ID" />
    </el-select>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="商家名">
    {{ detailForm.merchantName }}
</el-descriptions-item>
                 <el-descriptions-item label="商户logo">
    <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailForm.logo)" :src="getUrl(detailForm.logo)" fit="cover" />
</el-descriptions-item>
                 <el-descriptions-item label="经营范围描述">
    {{ detailForm.businessScope }}
</el-descriptions-item>
                 <el-descriptions-item label="地址">
    {{ detailForm.address }}
</el-descriptions-item>
                 <el-descriptions-item label="联系人">
    {{ detailForm.contactPerson }}
</el-descriptions-item>
                 <el-descriptions-item label="联系电话">
    {{ detailForm.contactPhone }}
</el-descriptions-item>
                 <el-descriptions-item label="评分">
    {{ detailForm.rating }}
</el-descriptions-item>
                 <el-descriptions-item label="邮箱">
    {{ detailForm.email }}
</el-descriptions-item>
                 <el-descriptions-item label="营业执照">
    <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailForm.businessLicense)" :src="getUrl(detailForm.businessLicense)" fit="cover" />
</el-descriptions-item>
                 <el-descriptions-item label="身份证照片">
    <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailForm.idCardImage)" :initial-index="index" v-for="(item,index) in detailForm.idCardImage" :key="index" :src="getUrl(item)" fit="cover" />
</el-descriptions-item>
 <el-descriptions-item label="创始人">
    {{ detailForm.founder }}
</el-descriptions-item>
    <el-descriptions-item label="创始人电话">
    {{ detailForm.founderPhone }}
</el-descriptions-item>
    <el-descriptions-item label="创始人简介">
    <RichView v-model="detailForm.founderDescriptions" />
</el-descriptions-item>
    <el-descriptions-item label="企业介绍">
    <RichView v-model="detailForm.enterpriseDesc" />
</el-descriptions-item>
    <el-descriptions-item label="企业宣传片">
    {{ detailForm.corporateVideo }}
</el-descriptions-item>
    <el-descriptions-item label="荣誉证书">
    <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailForm.certificate)" :initial-index="index" v-for="(item,index) in detailForm.certificate" :key="index" :src="getUrl(item)" fit="cover" />
</el-descriptions-item>
    <el-descriptions-item label="数字资产名称">
    {{ detailForm.digitalAssetName }}
</el-descriptions-item>
    <el-descriptions-item label="交易对">
    {{ detailForm.tradingPair }}
</el-descriptions-item>
                 <el-descriptions-item label="状态">
    {{ filterDict(detailForm.status,statusOptions) }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createMerchants,
  deleteMerchants,
  deleteMerchantsByIds,
  updateMerchants,
  findMerchants,
  getMerchantsList
} from '@/plugin/app/api/merchants'
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
import RichView from '@/components/richtext/rich-view.vue'
import RichEdit from '@/components/richtext/rich-edit.vue'

defineOptions({
    name: 'Merchants'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const statusOptions = ref([])
const merchantCategories = ref([])
const formData = ref({
            merchantName: '',
            logo: "",
            businessScope: '',
            address: '',
            contactPerson: '',
            contactPhone: '',
            rating: 0,
            email: '',
            businessLicense: "",
            idCardImage: [],
            status: '',
             founder: '',
            founderPhone: '',
            founderDescriptions: '',
            enterpriseDesc: '',
            corporateVideo: "",
            certificate: [],
            digitalAssetName: '',
            tradingPair: '',
            isRecommended: false,
            categoryId: undefined,
        })



// 验证规则
const rule = reactive({
               merchantName : [{
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
               logo : [{
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
            merchantName: 'merchant_name',
            rating: 'rating',
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
  const table = await getMerchantsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteMerchantsFunc(row)
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
      const res = await deleteMerchantsByIds({ IDs })
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
const updateMerchantsFunc = async(row) => {
    const res = await findMerchants({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        loadMerchantCategories()
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMerchantsFunc = async (row) => {
    const res = await deleteMerchants({ ID: row.ID })
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
    loadMerchantCategories()
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        merchantName: '',
        logo: "",
        businessScope: '',
        address: '',
        contactPerson: '',
        contactPhone: '',
        rating: 0,
        email: '',
        businessLicense: "",
        idCardImage: [],
        status: '',
        isRecommended: false,
        categoryId: undefined,
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
                  res = await createMerchants(formData.value)
                  break
                case 'update':
                  res = await updateMerchants(formData.value)
                  break
                default:
                  res = await createMerchants(formData.value)
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
  const res = await findMerchants({ ID: row.ID })
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

import { getMerchantCategoryList } from '@/plugin/app/api/merchants'
const loadMerchantCategories = async () => {
  try {
    const res = await getMerchantCategoryList()
    if (res.code === 0) merchantCategories.value = res.data || []
  } catch {}
}
loadMerchantCategories()
</script>

<style>

</style>

