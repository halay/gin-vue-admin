
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
      
            <el-form-item label="用户邮箱" prop="email">
  <el-input v-model="searchInfo.email" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="用户昵称" prop="nickname">
  <el-input v-model="searchInfo.nickname" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="用户状态" prop="status">
    <el-tree-select v-model="searchInfo.status" placeholder="请选择用户状态" :data="statusOptions" style="width:220px" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
            <el-form-item label="所属商家" prop="merchantId">
  <el-select v-model="searchInfo.merchantId" clearable filterable placeholder="请选择商家" style="width:220px">
    <el-option v-for="(m,idx) in merchantOptions" :key="idx" :label="m.name" :value="m.ID" />
  </el-select>
</el-form-item>
            <el-form-item label="会员等级" prop="membershipLevelId">
  <el-select v-model="searchInfo.membershipLevelId" clearable placeholder="请选择" style="width:220px">
    <el-option v-for="(lv,idx) in membershipLevels" :key="idx" :label="lv.name" :value="lv.ID" />
  </el-select>
</el-form-item>
            <el-form-item label="所属节点" prop="nodeId">
  <el-select v-model="searchInfo.nodeId" clearable placeholder="请选择" style="width:220px">
    <el-option v-for="(n,idx) in nodeOptions" :key="idx" :label="n.name" :value="n.ID" />
  </el-select>
</el-form-item>
           
            <el-form-item label="最后登录时间" prop="lastLoginTime">
  <template #label>
    <span>
      最后登录时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.lastLoginTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
           
            <el-form-item label="邮箱是否已验证" prop="emailVerified">
  <el-select v-model="searchInfo.emailVerified" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
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
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="app_AppUsers" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="app_AppUsers" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="app_AppUsers" @on-success="getTableData" />
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
        <el-table-column align="left" label="用户ID" prop="ID" width="80" />
        <el-table-column align="left" label="用户邮箱" prop="email" width="180" />
        <el-table-column align="left" label="股东身份" prop="shareholderProfitId" width="140">
              <template #default="scope">
                {{ formatShareholder(scope.row.shareholderProfitId) }}
              </template>
            </el-table-column>
            <el-table-column align="left" label="会员等级" prop="membershipLevelId" width="80">
              <template #default="scope">
                {{ formatLevel(scope.row.membershipLevelId) }}
              </template>
            </el-table-column>
            <el-table-column align="left" label="所属节点" prop="nodeId" width="120">
              <template #default="scope">
                {{ formatNode(scope.row.nodeId) }}
              </template>
            </el-table-column>
                        <el-table-column align="left" label="所属商家" prop="merchantId" width="160">
              <template #default="scope">
                {{ formatMerchant(scope.row.merchantId) }}
              </template>
            </el-table-column>
            <el-table-column align="left" label="用户昵称" prop="nickname" width="120" />

            <el-table-column align="left" label="用户头像URL" prop="avatar" width="120" />

            <el-table-column align="left" label="用户手机号" prop="phone" width="120" />

            <el-table-column align="left" label="用户状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="最后登录时间" prop="lastLoginTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastLoginTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="最后登录IP" prop="lastLoginIp" width="120" />
            <el-table-column align="left" label="邀请码" prop="inviteCode" width="140" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button type="primary" link icon="user" class="table-button" @click="openIdentityDialog(scope.row)">编辑身份</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateAppUsersFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
             <el-form-item label="用户昵称:" prop="nickname">
    <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入用户昵称" />
</el-form-item>
             <el-form-item label="用户手机号:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入用户手机号" />
</el-form-item>
             <el-form-item label="会员等级:" prop="membershipLevelId">
    <el-select v-model="formData.membershipLevelId" placeholder="请选择会员等级" filterable clearable>
      <el-option v-for="(lv,idx) in membershipLevels" :key="idx" :label="lv.name" :value="lv.ID" />
    </el-select>
</el-form-item>
             <el-form-item label="所属节点:" prop="nodeId">
    <el-select v-model="formData.nodeId" placeholder="请选择节点" filterable clearable>
      <el-option v-for="(n,idx) in nodeOptions" :key="idx" :label="n.name" :value="n.ID" />
    </el-select>
</el-form-item>
            <el-form-item label="所属商家:" prop="merchantId">
    <el-select v-model="formData.merchantId" placeholder="请选择商家" filterable clearable>
      <el-option v-for="(m,idx) in merchantOptions" :key="idx" :label="m.merchantName" :value="m.ID" />
    </el-select>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-dialog v-model="identityDialogVisible" title="编辑身份" width="500px">
      <el-form :model="identityForm" label-width="100px">
        <el-form-item label="股东身份">
          <el-select v-model="identityForm.shareholderProfitId" placeholder="请选择股东身份" clearable style="width: 100%">
            <el-option v-for="item in shareholderOptions" :key="item.ID" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="identityDialogVisible = false">取消</el-button>
          <el-button type="primary" :loading="btnLoading" @click="updateIdentityFunc">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="用户邮箱">
    {{ detailForm.email }}
</el-descriptions-item>
                 <el-descriptions-item label="用户昵称">
    {{ detailForm.nickname }}
</el-descriptions-item>
                 <el-descriptions-item label="用户头像URL">
    {{ detailForm.avatar }}
</el-descriptions-item>
                 <el-descriptions-item label="用户手机号">
    {{ detailForm.phone }}
</el-descriptions-item>
                 <el-descriptions-item label="用户状态">
    {{ filterDict(detailForm.status,statusOptions) }}
</el-descriptions-item>
                 <el-descriptions-item label="最后登录时间">
    {{ detailForm.lastLoginTime }}
</el-descriptions-item>
                 <el-descriptions-item label="最后登录IP">
    {{ detailForm.lastLoginIp }}
</el-descriptions-item>
                 <el-descriptions-item label="邮箱是否已验证">
    {{ detailForm.emailVerified }}
</el-descriptions-item>
                <el-descriptions-item label="用户角色ID">
    {{ detailForm.authorityId }}
</el-descriptions-item>
                <el-descriptions-item label="邀请码">
    {{ detailForm.inviteCode }}
</el-descriptions-item>
                <el-descriptions-item label="所有上级">
    {{ Array.isArray(detailForm.ancestors) ? detailForm.ancestors.join(', ') : detailForm.ancestors }}
</el-descriptions-item>
                <el-descriptions-item label="所有下级">
    {{ Array.isArray(detailForm.descendants) ? detailForm.descendants.join(', ') : detailForm.descendants }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAppUsers,
  deleteAppUsers,
  deleteAppUsersByIds,
  updateAppUsers,
  findAppUsers,
  getAppUsersList
} from '@/plugin/app/api/appUsers'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { getMembershipLevelPublic } from '@/plugin/app/api/membershipLevel'
import { getMerchantsList } from '@/plugin/app/api/merchants'
import { getNodeList } from '@/plugin/app/api/node'
import { getShareholderProfitList } from '@/plugin/app/api/shareholder'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'AppUsers'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const statusOptions = ref([])
const membershipLevels = ref([])
const merchantOptions = ref([])
const nodeOptions = ref([])
const shareholderOptions = ref([])
const formData = ref({
            nickname: '',
            phone: '',
            membershipLevelId: undefined,
            nodeId: undefined,
            merchantId: undefined,
        })



// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()
const formatLevel = (id) => {
  const lv = membershipLevels.value.find(i => i.ID === id)
  return lv ? lv.name : ''
}
const formatNode = (id) => {
  const n = nodeOptions.value.find(i => i.ID === id)
  return n ? n.name : ''
}
const formatMerchant = (id) => {
  const m = merchantOptions.value.find(i => i.ID === id)
  return m ? m.merchantName : ''
}
const formatShareholder = (id) => {
  const s = shareholderOptions.value.find(i => i.ID === id)
  return s ? s.name : ''
}
const loadMembershipLevels = async () => {
  const res = await getMembershipLevelPublic()
  if (res.code === 0) membershipLevels.value = res.data || []
}
loadMembershipLevels()
const loadMerchants = async () => {
  const res = await getMerchantsList({ page:1, pageSize:9999 })
  if (res.code === 0) merchantOptions.value = res.data.list || []
}
loadMerchants()
const loadNodes = async () => {
  const res = await getNodeList({ page:1, pageSize:9999 })
  if (res.code === 0) nodeOptions.value = res.data.list || []
}
loadNodes()
const loadShareholders = async () => {
  const res = await getShareholderProfitList({ page:1, pageSize:9999 })
  if (res.code === 0) shareholderOptions.value = res.data.list || []
}
loadShareholders()

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
    if (searchInfo.value.emailVerified === ""){
        searchInfo.value.emailVerified=null
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
  const table = await getAppUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteAppUsersFunc(row)
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
      const res = await deleteAppUsersByIds({ IDs })
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
const updateAppUsersFunc = async(row) => {
    const res = await findAppUsers({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppUsersFunc = async (row) => {
    const res = await deleteAppUsers({ ID: row.ID })
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
        nickname: '',
        phone: '',
        merchantId: undefined,
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
                  res = await createAppUsers(formData.value)
                  break
                case 'update':
                  res = await updateAppUsers(formData.value)
                  break
                default:
                  res = await createAppUsers(formData.value)
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
  const res = await findAppUsers({ ID: row.ID })
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

const identityDialogVisible = ref(false)
const identityForm = ref({
  ID: 0,
  shareholderProfitId: undefined
})

const openIdentityDialog = (row) => {
  identityForm.value.ID = row.ID
  identityForm.value.shareholderProfitId = row.shareholderProfitId
  identityDialogVisible.value = true
}

const updateIdentityFunc = async () => {
  btnLoading.value = true
  // 我们只更新 ID 和 shareholderProfitId，后端 UpdateAppUsers 应该支持局部更新
  // 或者我们需要先 findAppUsers 获取全量数据再更新？
  // GORM 的 Updates 方法通常只更新非零值，但如果传的是结构体，零值字段会被忽略。
  // 前端 API updateAppUsers 传递的是整个对象。
  // 为了安全起见，先获取详情，修改字段，再保存。
  const res = await findAppUsers({ ID: identityForm.value.ID })
  if (res.code === 0) {
    const data = res.data
    data.shareholderProfitId = identityForm.value.shareholderProfitId
    const updateRes = await updateAppUsers(data)
    if (updateRes.code === 0) {
      ElMessage({
        type: 'success',
        message: '更新成功'
      })
      identityDialogVisible.value = false
      getTableData()
    }
  }
  btnLoading.value = false
}


</script>

