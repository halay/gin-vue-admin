
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="代理资格(万):" prop="qualificationAmount">
          <el-input v-model.number="formData.qualificationAmount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分红范围:" prop="dividendScope">
           <el-select v-model="formData.dividendScope" placeholder="请选择分红范围" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in bonusRangeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="业务提成:" prop="cfRate1">
          <el-input-number v-model="formData.cfRate1" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="销售费用补贴:" prop="cfRate2">
          <el-input-number v-model="formData.cfRate2" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="经销补贴:" prop="cfRate3">
          <el-input-number v-model="formData.cfRate3" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="分红比例:" prop="cfRate4">
          <el-input-number v-model="formData.cfRate4" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="额外补贴:" prop="cfRate5">
          <el-input-number v-model="formData.cfRate5" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
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
    getAgentLevelDataSource,
  createAgentLevel,
  updateAgentLevel,
  findAgentLevel
} from '@/plugin/app/api/agentLevel'

defineOptions({
    name: 'AgentLevelForm'
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
const bonusRangeOptions = ref([])
const statusOptions = ref([])
const formData = ref({
            id: undefined,
            name: '',
            qualificationAmount: undefined,
            dividendScope: '',
            cfRate1: 0,
            cfRate2: 0,
            cfRate3: 0,
            cfRate4: 0,
            cfRate5: 0,
            status: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getAgentLevelDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAgentLevel({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    bonusRangeOptions.value = await getDictFunc('bonusRange')
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
               res = await createAgentLevel(formData.value)
               break
             case 'update':
               res = await updateAgentLevel(formData.value)
               break
             default:
               res = await createAgentLevel(formData.value)
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
