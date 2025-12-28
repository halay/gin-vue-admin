
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="false"  placeholder="请输入姓名" />
       </el-form-item>
        <el-form-item label="账号ID:" prop="accountID">
          <el-input v-model="formData.accountID" :clearable="false"  placeholder="请输入账号ID" />
       </el-form-item>
        <el-form-item label="结算周期:" prop="settlementPeriod">
          <el-input v-model="formData.settlementPeriod" :clearable="false"  placeholder="请输入结算周期" />
       </el-form-item>
        <el-form-item label="补贴金额:" prop="subsidyAmount">
          <el-input-number v-model="formData.subsidyAmount" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="结算状态:" prop="status">
        <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="false">
          <el-option v-for="item in ['待结算','已结算','已驳回']" :key="item" :label="item" :value="item" />
        </el-select>
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
  createSettlement,
  updateSettlement,
  findSettlement
} from '@/plugin/app/api/settlementModel'

defineOptions({
    name: 'SettlementForm'
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
const formData = ref({
            name: '',
            accountID: '',
            settlementPeriod: '',
            subsidyAmount: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               accountID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               settlementPeriod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               subsidyAmount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSettlement({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createSettlement(formData.value)
               break
             case 'update':
               res = await updateSettlement(formData.value)
               break
             default:
               res = await createSettlement(formData.value)
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
