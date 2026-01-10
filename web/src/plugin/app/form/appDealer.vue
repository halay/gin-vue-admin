
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="经销商名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false"  placeholder="请输入经销商名称" />
       </el-form-item>
        <el-form-item label="销售提成:" prop="salesCommission">
          <el-input-number v-model="formData.salesCommission" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="提成类型:" prop="commissionType">
          <el-input v-model.number="formData.commissionType" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="费用补贴:" prop="expenseAllowance">
          <el-input-number v-model="formData.expenseAllowance" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="补贴类型:" prop="allowanceType">
          <el-input v-model.number="formData.allowanceType" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联商户:" prop="merchantId">
        <el-select  v-model="formData.merchantId" placeholder="请选择关联商户" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.merchantId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="false"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getAppDealerDataSource,
  createAppDealer,
  updateAppDealer,
  findAppDealer
} from '@/plugin/app/api/appDealer'

defineOptions({
    name: 'AppDealerForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const commission_typeOptions = ref([])
const allowance_typeOptions = ref([])
const formData = ref({
            name: '',
            salesCommission: 0,
            commissionType: 0,
            expenseAllowance: 0,
            allowanceType: 0,
            merchantId: undefined,
            remark: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               commissionType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               allowanceType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               merchantId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getAppDealerDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppDealer({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    commission_typeOptions.value = await getDictFunc('commission_type')
    allowance_typeOptions.value = await getDictFunc('allowance_type')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createAppDealer(formData.value)
               break
             case 'update':
               res = await updateAppDealer(formData.value)
               break
             default:
               res = await createAppDealer(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
