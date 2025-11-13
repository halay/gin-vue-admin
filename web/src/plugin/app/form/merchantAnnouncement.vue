
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公告标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入公告标题" />
       </el-form-item>
        <el-form-item label="公告内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="公告状态:" prop="status">
        <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [0]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="发布时间:" prop="publishAt">
          <el-date-picker v-model="formData.publishAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="过期时间:" prop="expireAt">
          <el-date-picker v-model="formData.expireAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
    getMerchantAnnouncementDataSource,
  createMerchantAnnouncement,
  updateMerchantAnnouncement,
  findMerchantAnnouncement
} from '@/plugin/app/api/merchantAnnouncement'

defineOptions({
    name: 'MerchantAnnouncementForm'
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
const merchant_announcement_statusOptions = ref([])
const formData = ref({
            title: '',
            content: '',
            publishAt: new Date(),
            expireAt: new Date(),
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
    const res = await getMerchantAnnouncementDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMerchantAnnouncement({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    merchant_announcement_statusOptions.value = await getDictFunc('merchant_announcement_status')
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
               res = await createMerchantAnnouncement(formData.value)
               break
             case 'update':
               res = await updateMerchantAnnouncement(formData.value)
               break
             default:
               res = await createMerchantAnnouncement(formData.value)
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
