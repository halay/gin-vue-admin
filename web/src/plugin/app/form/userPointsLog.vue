
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户:" prop="userId">
        <el-select  v-model="formData.userId" placeholder="请选择用户" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="变动值:" prop="change">
          <el-input v-model.number="formData.change" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="变动后余额:" prop="balanceAfter">
          <el-input v-model.number="formData.balanceAfter" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="原因:" prop="reason">
          <el-input v-model="formData.reason" :clearable="false"  placeholder="请输入原因" />
       </el-form-item>
        <el-form-item label="关联订单号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="false"  placeholder="请输入关联订单号" />
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
    getUserPointsLogDataSource,
  createUserPointsLog,
  updateUserPointsLog,
  findUserPointsLog
} from '@/plugin/app/api/userPointsLog'

defineOptions({
    name: 'UserPointsLogForm'
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
            userId: undefined,
            change: 0,
            balanceAfter: 0,
            reason: '',
            orderNo: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               change : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               balanceAfter : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getUserPointsLogDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserPointsLog({ ID: route.query.id })
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
               res = await createUserPointsLog(formData.value)
               break
             case 'update':
               res = await updateUserPointsLog(formData.value)
               break
             default:
               res = await createUserPointsLog(formData.value)
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
