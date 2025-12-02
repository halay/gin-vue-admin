
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="节点名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false"  placeholder="请输入节点名称" />
       </el-form-item>
        <el-form-item label="系统层级ID:" prop="systemLevel">
           <el-select v-model="formData.systemLevel" placeholder="请选择系统层级ID" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in system_levelOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="限额数量(0为不限):" prop="limitCount">
          <el-input v-model.number="formData.limitCount" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="加盟费用:" prop="joinFee">
          <el-input-number v-model="formData.joinFee" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="结算货币:" prop="settlementCurrency">
          <el-input v-model="formData.settlementCurrency" :clearable="false"  placeholder="请输入结算货币" />
       </el-form-item>
        <el-form-item label="一线城市/特定区域加收比例(%):" prop="citySurchargePercent">
          <el-input-number v-model="formData.citySurchargePercent" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="代币配送倍数(X倍):" prop="tokenDistributionX">
          <el-input-number v-model="formData.tokenDistributionX" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="释放周期(周):" prop="releaseWeeks">
          <el-input v-model.number="formData.releaseWeeks" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="平台消费现金券赠送(%):" prop="cashCouponPercent">
          <el-input-number v-model="formData.cashCouponPercent" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="消费级差PV分润(%):" prop="pvBonusPercent">
          <el-input-number v-model="formData.pvBonusPercent" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="RWA品牌上市费用分成(%):" prop="rwaBrandFeeSharePercent">
          <el-input-number v-model="formData.rwaBrandFeeSharePercent" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="入驻年费级差分成(%):" prop="annualFeeSharePercent">
          <el-input-number v-model="formData.annualFeeSharePercent" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="在线交易手续费分成(%):" prop="onlineTradeFeeSharePercent">
          <el-input-number v-model="formData.onlineTradeFeeSharePercent" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="股权(%):" prop="equityPercent">
          <el-input-number v-model="formData.equityPercent" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="false" placeholder="请输入" />
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
  createNode,
  updateNode,
  findNode
} from '@/plugin/app/api/node'

defineOptions({
    name: 'NodeForm'
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
const system_levelOptions = ref([])
const statusOptions = ref([])
const formData = ref({
            name: '',
            systemLevel: '',
            limitCount: 0,
            joinFee: 0,
            settlementCurrency: '',
            citySurchargePercent: 0,
            tokenDistributionX: 0,
            releaseWeeks: 0,
            cashCouponPercent: 0,
            pvBonusPercent: 0,
            rwaBrandFeeSharePercent: 0,
            annualFeeSharePercent: 0,
            onlineTradeFeeSharePercent: 0,
            equityPercent: 0,
            status: '',
            sort: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               systemLevel : [{
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
      const res = await findNode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    system_levelOptions.value = await getDictFunc('system_level')
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
               res = await createNode(formData.value)
               break
             case 'update':
               res = await updateNode(formData.value)
               break
             default:
               res = await createNode(formData.value)
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
