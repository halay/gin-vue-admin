
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商家名:" prop="merchantName">
          <el-input v-model="formData.merchantName" :clearable="false"  placeholder="请输入商家名" />
       </el-form-item>
        <el-form-item label="商户logo:" prop="logo">
          <SelectImage v-model="formData.logo" file-type="image"/>
       </el-form-item>
        <el-form-item label="创始人:"  prop="founder" >
            <el-input v-model="formData.founder" :clearable="true"  placeholder="请输入创始人" />
        </el-form-item>
        <el-form-item label="创始人电话:"  prop="founderPhone" >
            <el-input v-model="formData.founderPhone" :clearable="true"  placeholder="请输入创始人电话" />
        </el-form-item>
        <el-form-item label="创始人简介:"  prop="founderDescriptions" >
            <RichEdit v-model="formData.founderDescriptions"/>
        </el-form-item>
        <el-form-item label="企业介绍:"  prop="enterpriseDesc" >
            <RichEdit v-model="formData.enterpriseDesc"/>
        </el-form-item>
        <el-form-item label="企业宣传片:"  prop="corporateVideo" >
            <SelectImage
            v-model="formData.corporateVideo"
            file-type="video"
            />
        </el-form-item>
        <el-form-item label="荣誉证书:"  prop="certificate" >
            <SelectImage
            multiple
            v-model="formData.certificate"
            file-type="image"
            />
        </el-form-item>
        <el-form-item label="数字资产名称:"  prop="digitalAssetName" >
            <el-input v-model="formData.digitalAssetName" :clearable="true"  placeholder="请输入数字资产名称" />
        </el-form-item>
        <el-form-item label="交易对:"  prop="tradingPair" >
            <el-input v-model="formData.tradingPair" :clearable="true"  placeholder="请输入交易对" />
        </el-form-item>
        <el-form-item label="经营范围描述:" prop="businessScope">
          <el-input v-model="formData.businessScope" :clearable="false"  placeholder="请输入经营范围描述" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="false"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="联系人:" prop="contactPerson">
          <el-input v-model="formData.contactPerson" :clearable="false"  placeholder="请输入联系人" />
       </el-form-item>
        <el-form-item label="联系电话:" prop="contactPhone">
          <el-input v-model="formData.contactPhone" :clearable="false"  placeholder="请输入联系电话" />
       </el-form-item>
        <el-form-item label="评分:" prop="rating">
          <el-input-number v-model="formData.rating" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
       </el-form-item>
        <el-form-item label="营业执照:" prop="businessLicense">
          <SelectImage v-model="formData.businessLicense" file-type="image"/>
       </el-form-item>
        <el-form-item label="身份证照片:" prop="idCardImage">
           <SelectImage v-model="formData.idCardImage" multiple file-type="image"/>
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
  createMerchants,
  updateMerchants,
  findMerchants
} from '@/plugin/app/api/merchants'

defineOptions({
    name: 'MerchantsForm'
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
const statusOptions = ref([])
const formData = ref({
            merchantName: '',
            logo: "",
            businessScope: '',
            address: '',
            contactPerson: '',
            contactPhone: '',
            rating: 0,
            email: '',
            businessLicense: "",
            idCardImage: [],
            status: '',
            founder: '',
            founderPhone: '',
            founderDescriptions: '',
            enterpriseDesc: '',
            corporateVideo: "",
            certificate: [],
            digitalAssetName: '',
            tradingPair: '',
        })
// 验证规则
const rule = reactive({
               merchantName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               logo : [{
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
      const res = await findMerchants({ ID: route.query.id })
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
               res = await createMerchants(formData.value)
               break
             case 'update':
               res = await updateMerchants(formData.value)
               break
             default:
               res = await createMerchants(formData.value)
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
