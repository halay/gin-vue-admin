<script setup>
import { ref, computed, useTemplateRef } from 'vue'
import { Upload, Download, Folder } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import {
    getBLCTYImages
} from '@/plugin/app/api/cozeGenius'
const posterGeniusElement = useTemplateRef('posterGeniusRef')

// 状态管理
const outputRatio = ref('9:16')
const promptText = ref('')
const imgFile = ref(null);
const rawImgFile = ref(null); // 新增：保存原始文件对象
const imgBase64 = ref('')
const isReading = ref(false)
const isGenerating = ref(false)

const outcome = ref([])

// 输出比例选项
const ratioOptions = [
  { label: '4:3', value: '4:3' },
  { label: '3:4', value: '3:4' },
  { label: '16:9', value: '16:9' },
  { label: '9:16', value: '9:16' },
  { label: '2:3', value: '2:3' },
  { label: '3:2', value: '3:2' },
  { label: '1:1', value: '1:1' },
  { label: '21:9', value: '21:9' }
]

// 上传文件处理
const handleUpload = (file) => {
  const fileData = file.raw
  // 释放之前的URL对象
  if (imgFile.value) {
    URL.revokeObjectURL(imgFile.value)
  }
  isReading.value = false // 不需要读取base64了，直接使用file
  imgFile.value = URL.createObjectURL(fileData)
  rawImgFile.value = fileData // 保存原始文件
  
  // 移除FileReader相关逻辑，因为不再需要前端转Base64
  return false // 阻止自动上传
}

// 判断生成按钮是否可用
const isGenerateDisabled = computed(() => {
  return isGenerating.value || isReading.value
  // return !isUploaded.value || !promptText.value.trim() || isGenerating.value
})

// 生成海报处理
const handleGenerate = () => {
  if (isGenerateDisabled.value) return
  isGenerating.value = true
  outcome.value = []
  posterGeniusElement.value?.scrollIntoView({ behavior: 'smooth' })
  const formData = new FormData()
  formData.append('model', 'nano-banana-2')
  formData.append('prompt', promptText.value || `根据我上传的图片，帮我生成一张${outputRatio.value}商品类型的宣传海报，这个海报要符合商品图的风格，商品图要非常完美的融入到海报中，海报整体设计要高级，专业，光影，材质，氛围感要完全体现在海报中，适合在各类电商平台进行宣传，能吸引用户点击查看`)
  
  if (outputRatio.value) {
    formData.append('aspect_ratio', outputRatio.value)
  }
  
  if (rawImgFile.value) {
    formData.append('file', rawImgFile.value)
  }

  getBLCTYImage(formData)
//   request(params).then(response => {
//     // 检查响应结构并提取图片URL
//     isGenerating.value = false
//     outcome.value = response.urls
//     ElMessage.success('智能海报生成成功')
//   }).catch(err => {
//     console.log('API请求错误:', err)
//     isGenerating.value = false
//     ElMessage.error('智能海报生成失败')
//   })
}
// 删除行
const getBLCTYImage = async (row) => {
    const res = await getBLCTYImages(row)
    if (res.code === 0) {
        isGenerating.value = false
        outcome.value = res.data.urls
        ElMessage.success('智能海报生成成功')
    }else{
        ElMessage.error('智能海报生成失败')
    }
}
// 下载图片处理
const handleDownload = (url) => {
  const suffix = url.slice(url.lastIndexOf('.'))
  const filename = Date.now() + suffix;
  ElMessage.primary('图片下载中，请稍后...')
  fetch(url)
    .then((response) => response.blob())
    .then((blob) => {
      const blobUrl = URL.createObjectURL(new Blob([blob]))
      const link = document.createElement('a')
      link.href = blobUrl
      link.download = filename
      document.body.appendChild(link)
      link.click()
      URL.revokeObjectURL(blobUrl)
      link.remove()
      ElMessage.success('图片下载成功')
    })
}
// 保存图片处理
const handleSave = (url) => {
  handleDownload(url)
}

const handleRemoveImage = () => {
  if (imgFile.value) {
    URL.revokeObjectURL(imgFile.value)
  }
  imgFile.value = null
  rawImgFile.value = null
  imgBase64.value = ''
}
</script>

<template>
  <div>
    <div class="poster-genius pt-8 pb-12 px-4 max-w-3xl mx-auto">
      <div class="text-center mb-10">
        <h1 class="text-3xl font-bold mb-2">AI智能海报系统</h1>
        <p class="text-gray-400">输入产品风格关键词，让AI自动为您布置场景</p>
      </div>
      <div class="p-10 rounded-10 shadow-sm  bg-[var(--bg-color)] border-1 border-solid border-gray-100">
        <div class="flex items-center justify-center mb-10">
          <el-upload
            v-if="!imgFile"
            class="upload-box w-1/2 !opacity-100"
            drag
            accept="image/*"
            :disabled="isGenerating"
            :show-file-list="false"
            :auto-upload="false"
            :on-change="handleUpload"
            :limit="1"
          >
            <div class="flex flex-col items-center justify-center py-4">
              <el-icon class="text-4xl text-[var(--primary-color)] mb-3"><Upload /></el-icon>
              <div class="text-gray-600 text-base">点击上传产品原图</div>
            </div>
          </el-upload>
          <div v-if="imgFile" class="relative border-1 border-solid border-gray-300 rounded-4 overflow-hidden">
            <img class="block w-auto h-[220px] object-cover" :src="imgFile" alt="">
            <div class="absolute top-2 right-2 bg-red-500 text-white text-xs px-1 rounded-full cursor-pointer" @click="handleRemoveImage">删除</div>
          </div>
        </div>
        
        <div class="flex items-center justify-between gap-2 mb-3">
          <div class="flex items-center gap-2 text-[var(--text-color)]">
            <span class="font-medium text-sm">创意描述（PROMPT）</span>
          </div>
          <div class="text-[var(--primary-color)] text-sm">
            建议包含：场景、灯光、背景材质
          </div>
        </div>
        <el-input
          type="textarea"
          placeholder="例如：极简白色磨砂台面，侧面柔和阳光，背景虚化，点缀少量干花..."
          :rows="5"
          v-model="promptText"
          resize="none"
        ></el-input>
        <el-divider class="my-4xl" content-position="center">输出比例</el-divider>
        <div class="flex items-center justify-center">
          <el-radio-group v-model="outputRatio" :disabled="isGenerating" class="ratio-group flex items-center justify-center gap-3">
            <el-radio-button
              v-for="option in ratioOptions"
              :key="option.value"
              :value="option.value"
            >
              {{ option.label }}
            </el-radio-button>
          </el-radio-group>
        </div>

        <!-- 生成按钮 -->
        <div class="pt-3xl flex items-center justify-center">
          <el-button
            type="primary"
            size="large"
            :disabled="isGenerateDisabled"
            :loading="isGenerating"
            @click="handleGenerate"
            class="w-1/2 py-6 rounded-xl text-base font-medium"
          >
            <span v-if="isGenerating" class="ml-2">正在生成智能海报，请稍候...</span>
            <span v-else>生成智能海报</span>
          </el-button>
        </div>
      </div>
    </div>
    <div ref="posterGeniusRef" class="poster-genius mb-6xl max-w-4xl min-h-[200px] bg-[#FDFDFE] mx-auto p-xl rounded-5 border-2 border-dashed border-gray-100 flex items-center justify-center">
      <p v-if="!outcome.length" class="text-center text-sm text-gray-400">
        {{isGenerating ? '正在生成智能海报，请稍候...' : '等待生成任务'}}
      </p>
      <div v-else class="w-full grid grid-cols-4 gap-4">
        <div
          class="h-[280px] rounded-4 border-1 border-solid border-gray-100 overflow-hidden relative"
          v-for="(item, index) in outcome" :key="item">
          <el-image
            :src="item"
            :zoom-rate="1.2"
            :max-scale="7"
            :min-scale="0.2"
            :preview-src-list="outcome"
            show-progress
            :initial-index="index"
            fit="cover"
            class="w-full h-full"
          />
          <div class="flex justify-center gap-2 absolute bottom-2 left-1/2 transform -translate-x-1/2">
            <el-button
              type="default"
              size="small"
              round
              class="flex items-center gap-1 bg-white"
              @click="handleDownload(item)"
            >
              <el-icon class="mr-2"><Download /></el-icon>
              下载
            </el-button>
            <!-- <el-button
              type="default"
              size="small"
              round
              class="flex items-center gap-1 bg-white"
              @click="handleSave(item)"
            >
              <el-icon class="mr-2"><Folder /></el-icon>
              保存
            </el-button> -->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.poster-genius {
  --bg-color: #F8FAFC;
  --text-color: #98A6BB;
  --primary-color: #2563EB;
}
:deep(.el-upload-dragger) {
  border-width: 2px;
  border-radius: 1rem;
}
:deep(.el-textarea__inner) {
  border-radius: 0.75rem;
  padding: 1rem;
}
:deep(.el-divider) {
  --el-bg-color: var(--bg-color);
  --el-text-color-primary: var(--text-color);
}
:deep(.el-button) {
  --el-button-bg-color: var(--primary-color);
}
:deep(.el-radio-button) {
  --el-radio-button-checked-bg-color: var(--primary-color);
}
:deep(.el-radio-button__inner) {
  width: 60px;
}
:deep(.el-radio-button__inner),
:deep(.el-radio-button:first-child .el-radio-button__inner),
:deep(.el-radio-button:last-child .el-radio-button__inner) {
  border-radius: 12px;
}
</style>