
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品:" prop="productId">
        <el-select  v-model="formData.productId" placeholder="请选择商品" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="SKU编码:" prop="skuCode">
          <el-input v-model="formData.skuCode" :clearable="true"  placeholder="请输入SKU编码" />
       </el-form-item>
        <el-form-item label="属性组合:" prop="attrs">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.attrs 后端会按照json的类型进行存取
          {{ formData.attrs }}
       </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="积分:" prop="points">
          <el-input v-model.number="formData.points" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片:" prop="image">
          <SelectImage v-model="formData.image" file-type="image"/>
       </el-form-item>
        <el-form-item label="库存:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in sku_statusOptions" :key="key" :label="item.label" :value="item.value" />
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
    getProductSkuDataSource,
  createProductSku,
  updateProductSku,
  findProductSku
} from '@/plugin/app/api/productSku'

defineOptions({
    name: 'ProductSkuForm'
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
const sku_statusOptions = ref([])
const formData = ref({
            productId: undefined,
            skuCode: '',
            attrs: {},
            price: 0,
            points: 0,
            image: "",
            stock: 0,
            status: '',
        })
// 验证规则
const rule = reactive({
               productId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               price : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               stock : [{
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
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getProductSkuDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProductSku({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    sku_statusOptions.value = await getDictFunc('sku_status')
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
               res = await createProductSku(formData.value)
               break
             case 'update':
               res = await updateProductSku(formData.value)
               break
             default:
               res = await createProductSku(formData.value)
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
