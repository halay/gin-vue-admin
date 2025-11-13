
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
       </el-form-item>
        <el-form-item label="图片地址:" prop="imageUrl">
          <SelectImage v-model="formData.imageUrl" file-type="image"/>
       </el-form-item>
        <el-form-item label="跳转链接:" prop="linkUrl">
          <el-input v-model="formData.linkUrl" :clearable="true"  placeholder="请输入跳转链接" />
       </el-form-item>
        <el-form-item label="位置标识:" prop="position">
           <el-select v-model="formData.position" placeholder="请选择位置标识" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in banner_positionOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in banner_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="开始时间:" prop="startTime">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结束时间:" prop="endTime">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createBanner,
  updateBanner,
  findBanner
} from '@/plugin/app/api/banner'

defineOptions({
    name: 'BannerForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const banner_positionOptions = ref([])
const banner_statusOptions = ref([])
const formData = ref({
            title: '',
            imageUrl: "",
            linkUrl: '',
            position: '',
            sort: 0,
            status: '',
            startTime: new Date(),
            endTime: new Date(),
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               imageUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               position : [{
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
      const res = await findBanner({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    banner_positionOptions.value = await getDictFunc('banner_position')
    banner_statusOptions.value = await getDictFunc('banner_status')
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
               res = await createBanner(formData.value)
               break
             case 'update':
               res = await updateBanner(formData.value)
               break
             default:
               res = await createBanner(formData.value)
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
