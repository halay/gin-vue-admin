
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="版本名:" prop="versionName">
          <el-input v-model="formData.versionName" :clearable="false"  placeholder="请输入版本名" />
       </el-form-item>
        <el-form-item label="版本号:" prop="versionCode">
          <el-input v-model.number="formData.versionCode" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="平台:" prop="platform">
           <el-select v-model="formData.platform" placeholder="请选择平台" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in app_platformOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in release_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="强制更新:" prop="forceUpdate">
          <el-switch v-model="formData.forceUpdate" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="最低支持版本号:" prop="minSupport">
          <el-input v-model.number="formData.minSupport" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="安卓包名:" prop="packageName">
          <el-input v-model="formData.packageName" :clearable="false"  placeholder="请输入安卓包名" />
       </el-form-item>
        <el-form-item label="iOS Bundle ID:" prop="bundleId">
          <el-input v-model="formData.bundleId" :clearable="false"  placeholder="请输入iOS Bundle ID" />
       </el-form-item>
        <el-form-item label="安卓安装包:" prop="androidPackage">
          <SelectFile v-model="formData.androidPackage" />
       </el-form-item>
        <el-form-item label="iOS安装包/链接:" prop="iosPackage">
          <SelectFile v-model="formData.iosPackage" />
       </el-form-item>
        <el-form-item label="安卓下载链接:" prop="androidUrl">
          <el-input v-model="formData.androidUrl" :clearable="false"  placeholder="请输入安卓下载链接" />
       </el-form-item>
        <el-form-item label="iOS下载链接:" prop="iosUrl">
          <el-input v-model="formData.iosUrl" :clearable="false"  placeholder="请输入iOS下载链接" />
       </el-form-item>
        <el-form-item label="更新日志:" prop="changelog">
          <RichEdit v-model="formData.changelog"/>
       </el-form-item>
        <el-form-item label="发布时间:" prop="releaseAt">
          <el-date-picker v-model="formData.releaseAt" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
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
  createAppRelease,
  updateAppRelease,
  findAppRelease
} from '@/plugin/app/api/appRelease'

defineOptions({
    name: 'AppReleaseForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const app_platformOptions = ref([])
const release_statusOptions = ref([])
const formData = ref({
            versionName: '',
            versionCode: 0,
            platform: '',
            status: '',
            forceUpdate: false,
            minSupport: 0,
            packageName: '',
            bundleId: '',
            androidPackage: [],
            iosPackage: [],
            androidUrl: '',
            iosUrl: '',
            changelog: '',
            releaseAt: new Date(),
        })
// 验证规则
const rule = reactive({
               versionName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               versionCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               platform : [{
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
      const res = await findAppRelease({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    app_platformOptions.value = await getDictFunc('app_platform')
    release_statusOptions.value = await getDictFunc('release_status')
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
               res = await createAppRelease(formData.value)
               break
             case 'update':
               res = await updateAppRelease(formData.value)
               break
             default:
               res = await createAppRelease(formData.value)
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
