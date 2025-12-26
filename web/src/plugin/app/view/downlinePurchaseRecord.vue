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
            value-format="YYYY-MM-DD HH:mm:ss"
            :default-time="[new Date(2000, 1, 1, 0, 0, 0), new Date(2000, 1, 1, 23, 59, 59)]"
          />
        </el-form-item>

        <el-form-item label="订单号" prop="orderNo">
          <el-input v-model="searchInfo.orderNo" placeholder="搜索订单号" />
        </el-form-item>

        <el-form-item label="购买人ID" prop="purchaseUserId">
          <el-input v-model.number="searchInfo.purchaseUserId" placeholder="搜索购买人ID" />
        </el-form-item>
        
        <el-form-item label="购买人邮箱" prop="purchaseUserEmail">
          <el-input v-model="searchInfo.purchaseUserEmail" placeholder="搜索购买人邮箱" />
        </el-form-item>
        
        <el-form-item label="受益人ID" prop="uplineUserId">
          <el-input v-model.number="searchInfo.uplineUserId" placeholder="搜索受益人ID" />
        </el-form-item>

        <el-form-item label="受益人邮箱" prop="uplineUserEmail">
          <el-input v-model="searchInfo.uplineUserEmail" placeholder="搜索受益人邮箱" />
        </el-form-item>

        <el-form-item label="支付状态" prop="payStatus">
            <el-select v-model="searchInfo.payStatus" placeholder="请选择支付状态" clearable>
                <el-option label="已支付" value="paid" />
                <el-option label="未支付" value="unpaid" />
                <el-option label="失败" value="failed" />
            </el-select>
        </el-form-item>
        
        <el-form-item label="是否直推" prop="isDirect">
             <el-select v-model="searchInfo.isDirect" placeholder="请选择" clearable>
                <el-option label="是" :value="true" />
                <el-option label="否" :value="false" />
            </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @sort-change="sortChange"
      >
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="订单号" prop="orderNo" width="180" />
        
        <el-table-column align="left" label="购买人" min-width="200">
             <template #default="scope">
                <div>ID: {{ scope.row.purchaseUserId }}</div>
                <div>{{ scope.row.purchaseUserNickname }}</div>
                <div>{{ scope.row.purchaseUserPhone }}</div>
                <div>{{ scope.row.purchaseUserEmail }}</div>
             </template>
        </el-table-column>
        
        <el-table-column align="left" label="受益人" min-width="200">
             <template #default="scope">
                <div>ID: {{ scope.row.uplineUserId }}</div>
                <div>{{ scope.row.uplineUserEmail }}</div>
             </template>
        </el-table-column>
        
        <el-table-column align="left" label="商户" prop="merchantName" min-width="120" />

        <el-table-column align="left" label="直推" prop="isDirect" width="80">
            <template #default="scope">
                <el-tag :type="scope.row.isDirect ? 'success' : 'info'">{{ scope.row.isDirect ? '是' : '否' }}</el-tag>
            </template>
        </el-table-column>
        
        <el-table-column align="left" label="层级" prop="inviteLevel" width="80" />
        
        <el-table-column align="left" label="金额" prop="totalAmount" width="100" />
        <el-table-column align="left" label="积分" prop="totalPoints" width="100" />
        
        <el-table-column align="left" label="支付状态" prop="payStatus" width="100">
             <template #default="scope">
                <el-tag v-if="scope.row.payStatus === 'paid'" type="success">已支付</el-tag>
                <el-tag v-else-if="scope.row.payStatus === 'unpaid'" type="warning">未支付</el-tag>
                <el-tag v-else type="danger">{{ scope.row.payStatus }}</el-tag>
            </template>
        </el-table-column>
        
        <el-table-column align="left" label="订单状态" prop="orderStatus" width="100" />

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
  </div>
</template>

<script setup>
import {
  getDownlinePurchaseRecordList
} from '@/plugin/app/api/downlinePurchaseRecord'
import { formatDate } from '@/utils/format'
import { ref } from 'vue'

defineOptions({
  name: 'DownlinePurchaseRecord'
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const elSearchFormRef = ref()

// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt: 'created_at',
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
    // 处理时间范围
    if (searchInfo.value.createdAtRange && searchInfo.value.createdAtRange.length === 2) {
        searchInfo.value.startCreatedAt = searchInfo.value.createdAtRange[0]
        searchInfo.value.endCreatedAt = searchInfo.value.createdAtRange[1]
    } else {
        searchInfo.value.startCreatedAt = undefined
        searchInfo.value.endCreatedAt = undefined
    }

  const table = await getDownlinePurchaseRecordList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
</script>

<style>
</style>
