
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商户ID:" prop="merchantId">
          <el-input v-model.number="formData.merchantId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="积分全称:" prop="tokenName">
          <el-input v-model="formData.tokenName" :clearable="false"  placeholder="请输入积分全称" />
       </el-form-item>
        <el-form-item label="积分缩写(Symbol):" prop="symbol">
          <el-input v-model="formData.symbol" :clearable="false"  placeholder="请输入积分缩写(Symbol)" />
       </el-form-item>
        <el-form-item label="邀请人奖励积分:" prop="inviterReward">
          <el-input v-model.number="formData.inviterReward" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="被邀请人奖励积分:" prop="inviteeReward">
          <el-input v-model.number="formData.inviteeReward" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="注册奖励积分:" prop="registerReward">
          <el-input v-model.number="formData.registerReward" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分发受众策略:" prop="coverageMode">
          <el-input v-model="formData.coverageMode" :clearable="false"  placeholder="请输入分发受众策略" />
       </el-form-item>
        <el-form-item label="每日发放上限(人次):" prop="dailyLimit">
          <el-input v-model.number="formData.dailyLimit" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="说明:" prop="remark">
          <el-input v-model="formData.remark" :clearable="false"  placeholder="请输入说明" />
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
  createMerchantPointsSettings,
  updateMerchantPointsSettings,
  findMerchantPointsSettings
} from '@/plugin/app/api/merchantPointsSettings'

defineOptions({
    name: 'MerchantPointsSettingsForm'
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
const statusOptions = ref([])
const formData = ref({
            merchantId: 0,
            tokenName: '',
            symbol: '',
            inviterReward: 0,
            inviteeReward: 0,
            registerReward: 0,
            coverageMode: '',
            dailyLimit: 0,
            status: '',
            sort: 0,
            remark: '',
        })
// 验证规则
const rule = reactive({
               merchantId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tokenName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               symbol : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               inviterReward : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               inviteeReward : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               registerReward : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               coverageMode : [{
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
      const res = await findMerchantPointsSettings({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
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
               res = await createMerchantPointsSettings(formData.value)
               break
             case 'update':
               res = await updateMerchantPointsSettings(formData.value)
               break
             default:
               res = await createMerchantPointsSettings(formData.value)
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
