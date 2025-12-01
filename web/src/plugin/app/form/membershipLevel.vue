
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="等级名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false"  placeholder="请输入等级名称" />
       </el-form-item>
        <el-form-item label="等级编码:" prop="code">
          <el-input v-model="formData.code" :clearable="false"  placeholder="请输入等级编码" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="要求描述:" prop="requirement">
          <el-input v-model="formData.requirement" :clearable="false"  placeholder="请输入要求描述" />
       </el-form-item>
        <el-form-item label="要求值:" prop="requirementValue">
          <el-input-number v-model="formData.requirementValue" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="权益说明:" prop="benefits">
          <RichEdit v-model="formData.benefits"/>
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
  createMembershipLevel,
  updateMembershipLevel,
  findMembershipLevel
} from '@/plugin/app/api/membershipLevel'

defineOptions({
    name: 'MembershipLevelForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const statusOptions = ref([])
const formData = ref({
            name: '',
            code: '',
            status: '',
            sort: 0,
            requirement: '',
            requirementValue: 0,
            benefits: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               code : [{
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
      const res = await findMembershipLevel({ ID: route.query.id })
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
               res = await createMembershipLevel(formData.value)
               break
             case 'update':
               res = await updateMembershipLevel(formData.value)
               break
             default:
               res = await createMembershipLevel(formData.value)
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
