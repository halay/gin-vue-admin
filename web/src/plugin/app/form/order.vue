
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="false"  placeholder="请输入订单号" />
       </el-form-item>
        <el-form-item label="下单用户:" prop="userId">
        <el-select  v-model="formData.userId" placeholder="请选择下单用户" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="总金额:" prop="totalAmount">
          <el-input-number v-model="formData.totalAmount" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="总积分:" prop="totalPoints">
          <el-input v-model.number="formData.totalPoints" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付方式:" prop="payMethod">
           <el-select v-model="formData.payMethod" placeholder="请选择支付方式" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in pay_methodOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="支付状态:" prop="payStatus">
           <el-select v-model="formData.payStatus" placeholder="请选择支付状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in pay_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="订单状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择订单状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in order_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="收货人:" prop="consigneeName">
          <el-input v-model="formData.consigneeName" :clearable="false"  placeholder="请输入收货人" />
       </el-form-item>
        <el-form-item label="收货人手机号:" prop="consigneePhone">
          <el-input v-model="formData.consigneePhone" :clearable="false"  placeholder="请输入收货人手机号" />
       </el-form-item>
        <el-form-item label="收货地址:" prop="address">
          <el-input v-model="formData.address" :clearable="false"  placeholder="请输入收货地址" />
       </el-form-item>
        <el-form-item label="发货状态:" prop="deliveryStatus">
           <el-select v-model="formData.deliveryStatus" placeholder="请选择发货状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in delivery_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="快递公司:" prop="expressName">
          <el-input v-model="formData.expressName" :clearable="false"  placeholder="请输入快递公司" />
       </el-form-item>
        <el-form-item label="快递单号:" prop="expressNo">
          <el-input v-model="formData.expressNo" :clearable="false"  placeholder="请输入快递单号" />
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
    getOrderDataSource,
  createOrder,
  updateOrder,
  findOrder
} from '@/plugin/app/api/order'

defineOptions({
    name: 'OrderForm'
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
const pay_methodOptions = ref([])
const pay_statusOptions = ref([])
const delivery_statusOptions = ref([])
const order_statusOptions = ref([])
const formData = ref({
            orderNo: '',
            userId: undefined,
            totalAmount: 0,
            totalPoints: 0,
            payMethod: '',
            payStatus: '',
            status: '',
            consigneeName: '',
            consigneePhone: '',
            address: '',
            deliveryStatus: '',
            expressName: '',
            expressNo: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               orderNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               merchantId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               payMethod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               payStatus : [{
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
    const res = await getOrderDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    pay_methodOptions.value = await getDictFunc('pay_method')
    pay_statusOptions.value = await getDictFunc('pay_status')
    delivery_statusOptions.value = await getDictFunc('delivery_status')
    order_statusOptions.value = await getDictFunc('order_status')
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
               res = await createOrder(formData.value)
               break
             case 'update':
               res = await updateOrder(formData.value)
               break
             default:
               res = await createOrder(formData.value)
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
