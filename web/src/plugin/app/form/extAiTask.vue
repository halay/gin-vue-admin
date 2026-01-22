
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务ID:" prop="taskId">
          <el-input v-model="formData.taskId" :clearable="true"  placeholder="请输入任务ID" />
       </el-form-item>
        <el-form-item label="任务参数:" prop="options">
          <el-input v-model="formData.options" :clearable="true"  placeholder="请输入任务参数" />
       </el-form-item>
        <el-form-item label="任务类型:" prop="type">
           <el-select v-model="formData.type" placeholder="请选择任务类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ai_task_typeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="任务状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择任务状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ai_task_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="备注说明:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注说明" />
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
  createExtAiTask,
  updateExtAiTask,
  findExtAiTask
} from '@/plugin/app/api/extAiTask'

defineOptions({
    name: 'ExtAiTaskForm'
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
const ai_task_typeOptions = ref([])
const ai_task_statusOptions = ref([])
const formData = ref({
            id: undefined,
            createdAt: new Date(),
            userId: undefined,
            taskId: '',
            options: '',
            type: '',
            status: '',
            description: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findExtAiTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ai_task_typeOptions.value = await getDictFunc('ai_task_type')
    ai_task_statusOptions.value = await getDictFunc('ai_task_status')
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
               res = await createExtAiTask(formData.value)
               break
             case 'update':
               res = await updateExtAiTask(formData.value)
               break
             default:
               res = await createExtAiTask(formData.value)
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
