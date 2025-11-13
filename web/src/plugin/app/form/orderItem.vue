
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单:" prop="orderId">
        <el-select  v-model="formData.orderId" placeholder="请选择订单" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.orderId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="商品:" prop="productId">
        <el-select  v-model="formData.productId" placeholder="请选择商品" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="SKU:" prop="skuId">
        <el-select  v-model="formData.skuId" placeholder="请选择SKU" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.skuId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="商品名快照:" prop="productName">
          <el-input v-model="formData.productName" :clearable="false"  placeholder="请输入商品名快照" />
       </el-form-item>
        <el-form-item label="SKU属性:" prop="skuAttrs">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.skuAttrs 后端会按照json的类型进行存取
          {{ formData.skuAttrs }}
       </el-form-item>
        <el-form-item label="单价:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="单件积分:" prop="points">
          <el-input v-model.number="formData.points" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="小计金额:" prop="totalAmount">
          <el-input-number v-model="formData.totalAmount" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="小计积分:" prop="totalPoints">
          <el-input v-model.number="formData.totalPoints" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片:" prop="coverImage">
          <SelectImage v-model="formData.coverImage" file-type="image"/>
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
    getOrderItemDataSource,
  createOrderItem,
  updateOrderItem,
  findOrderItem
} from '@/plugin/app/api/orderItem'

defineOptions({
    name: 'OrderItemForm'
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
const formData = ref({
            orderId: undefined,
            productId: undefined,
            skuId: undefined,
            productName: '',
            skuAttrs: {},
            price: 0,
            points: 0,
            quantity: 0,
            totalAmount: 0,
            totalPoints: 0,
            coverImage: "",
        })
// 验证规则
const rule = reactive({
               orderId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               productId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               skuId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               productName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               price : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               quantity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getOrderItemDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrderItem({ ID: route.query.id })
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
               res = await createOrderItem(formData.value)
               break
             case 'update':
               res = await updateOrderItem(formData.value)
               break
             default:
               res = await createOrderItem(formData.value)
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
