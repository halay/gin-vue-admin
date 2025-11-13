
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入商品名称" />
       </el-form-item>
        <el-form-item label="商品类型:" prop="type">
           <el-select v-model="formData.type" placeholder="请选择商品类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in product_typeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in product_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="所属分类:" prop="categoryId">
        <el-select  v-model="formData.categoryId" placeholder="请选择所属分类" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.categoryId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="封面图:" prop="coverImage">
          <SelectImage v-model="formData.coverImage" file-type="image"/>
       </el-form-item>
        <el-form-item label="轮播图:" prop="carouselImages">
           <SelectImage v-model="formData.carouselImages" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="视频:" prop="video">
          <SelectImage v-model="formData.video" file-type="video"/>
       </el-form-item>
        <el-form-item label="商品详情:" prop="detail">
          <RichEdit v-model="formData.detail"/>
       </el-form-item>
        <el-form-item label="是否多规格:" prop="hasVariants">
          <el-switch v-model="formData.hasVariants" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="开启积分支付:" prop="enablePoints">
          <el-switch v-model="formData.enablePoints" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="商品价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="所需积分:" prop="points">
          <el-input v-model.number="formData.points" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="库存:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入" />
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
    getProductDataSource,
  createProduct,
  updateProduct,
  findProduct
} from '@/plugin/app/api/product'

defineOptions({
    name: 'ProductForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const product_statusOptions = ref([])
const product_typeOptions = ref([])
const formData = ref({
            name: '',
            type: '',
            status: '',
            categoryId: undefined,
            coverImage: "",
            carouselImages: [],
            video: "",
            detail: '',
            hasVariants: false,
            enablePoints: false,
            price: 0,
            points: 0,
            stock: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               categoryId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getProductDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProduct({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    product_statusOptions.value = await getDictFunc('product_status')
    product_typeOptions.value = await getDictFunc('product_type')
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
               res = await createProduct(formData.value)
               break
             case 'update':
               res = await updateProduct(formData.value)
               break
             default:
               res = await createProduct(formData.value)
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
