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
        
        <el-form-item label="受益人ID" prop="beneficiaryId">
          <el-input v-model.number="searchInfo.beneficiaryId" placeholder="搜索受益人ID" />
        </el-form-item>

        <el-form-item label="来源用户ID" prop="sourceUserId">
          <el-input v-model.number="searchInfo.sourceUserId" placeholder="搜索来源用户ID" />
        </el-form-item>
        
        <el-form-item label="代理等级" prop="agentLevelName">
          <el-input v-model="searchInfo.agentLevelName" placeholder="搜索代理等级" />
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
        
        <el-table-column align="left" label="总分润金额" prop="totalAmount" width="120">
             <template #default="scope">
                <span style="color: red; font-weight: bold;">{{ scope.row.totalAmount }}</span>
             </template>
        </el-table-column>

        <el-table-column align="left" label="触发等级" prop="agentLevelName" width="120" />
        <el-table-column align="left" label="分红范围" prop="dividendScope" width="100" />
        
        <el-table-column align="left" label="受益人ID" prop="beneficiaryId" width="100" />
        <el-table-column align="left" label="来源用户ID" prop="sourceUserId" width="100" />
        
        <el-table-column align="left" label="订单金额" prop="orderAmount" width="100" />
        <el-table-column align="left" label="分润基数" prop="baseAmount" width="100" />
        
        <el-table-column align="left" label="分润明细" min-width="200">
             <template #default="scope">
                <div v-if="scope.row.amount1 > 0">等级1: {{scope.row.rate1}}% = {{scope.row.amount1}}</div>
                <div v-if="scope.row.amount2 > 0">等级2: {{scope.row.rate2}}% = {{scope.row.amount2}}</div>
                <div v-if="scope.row.amount3 > 0">等级3: {{scope.row.rate3}}% = {{scope.row.amount3}}</div>
                <div v-if="scope.row.amount4 > 0">等级4: {{scope.row.rate4}}% = {{scope.row.amount4}}</div>
                <div v-if="scope.row.amount5 > 0">等级5: {{scope.row.rate5}}% = {{scope.row.amount5}}</div>
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
  </div>
</template>

<script setup>
import {
  getAgentTransactionList
} from '@/plugin/app/api/agentTransaction'
import { formatDate } from '@/utils/format'
import { ref } from 'vue'

defineOptions({
  name: 'AgentTransaction'
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

  const table = await getAgentTransactionList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
