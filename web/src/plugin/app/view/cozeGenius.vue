<script setup>
import { ref, computed, useTemplateRef, nextTick } from 'vue'
import { Upload, Download, Picture, Delete } from '@element-plus/icons-vue'
import { useCozeGeniusStore } from '@/pinia/modules/cozeGenius'

const useCozeGenius = useCozeGeniusStore()

const posterGeniusElement = useTemplateRef('posterGeniusRef')

// 状态管理
const outputRatio = ref('9:16')
const promptText = ref('')
const imgFile = ref(null);
const rawImgFile = ref(null); // 新增：保存原始文件对象

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

// 判断生成按钮是否可用
const isGenerateDisabled = computed(() => {
  return useCozeGenius.isGenerating || !rawImgFile.value
})

// 上传文件处理
const handleUpload = (file) => {
  const fileData = file.raw
  // 释放之前的URL对象
  if (imgFile.value) {
    URL.revokeObjectURL(imgFile.value)
  }
  imgFile.value = URL.createObjectURL(fileData)
  rawImgFile.value = fileData // 保存原始文件
  
  // 移除FileReader相关逻辑，因为不再需要前端转Base64
  return false // 阻止自动上传
}

// 删除上传图片
const handleRemoveImage = () => {
  if (imgFile.value) {
    URL.revokeObjectURL(imgFile.value)
  }
  imgFile.value = null
  rawImgFile.value = null
}

// 生成海报处理
const handleGenerate = () => {
  if (isGenerateDisabled.value) return
  const prompt = promptText.value || `根据我上传的图片，帮我生成一张${outputRatio.value}商品类型的宣传海报，这个海报要符合商品图的风格，商品图要非常完美的融入到海报中，海报整体设计要高级，专业，光影，材质，氛围感要完全体现在海报中，适合在各类电商平台进行宣传，能吸引用户点击查看`
  const params = {
    model: 'nano-banana-2',
    prompt,
    aspect_ratio: outputRatio.value,
    file: rawImgFile.value
  }
  useCozeGenius.createPosterTask(params);
  nextTick(() => {
    posterGeniusElement.value?.scrollIntoView({ behavior: 'smooth' })
  })
}

// 获取任务历史记录
if (!useCozeGenius.list?.length) {
  useCozeGenius.getHistory()
}


</script>

<template>
  <div>
    <div class="poster-genius pt-8 pb-12 px-4 max-w-3xl mx-auto">
      <div class="text-center mb-10">
        <h1 class="text-3xl font-bold mb-2">AI智能海报系统</h1>
        <p class="text-gray-400">输入产品风格关键词，让AI自动为您布置场景</p>
      </div>
      <div class="p-10 rounded-10 shadow-sm  bg-white border-1 border-solid border-gray-100">
        <div class="flex items-center justify-center mb-10">
          <el-upload
            v-if="!imgFile"
            class="upload-box w-1/2 !opacity-100"
            drag
            accept="image/*"
            :disabled="useCozeGenius.isGenerating"
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
          <el-radio-group v-model="outputRatio" :disabled="useCozeGenius.isGenerating" class="ratio-group flex items-center justify-center gap-3">
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
            :loading="useCozeGenius.isGenerating"
            @click="handleGenerate"
            class="w-1/2 py-6 rounded-xl text-base font-medium"
          >
            <span v-if="useCozeGenius.isGenerating" class="ml-2">正在生成智能海报，请稍候...</span>
            <span v-else>生成智能海报</span>
          </el-button>
        </div>
      </div>
    </div>
    <div ref="posterGeniusRef" class="max-w-4xl min-h-[200px] mx-auto">
      <div class="flex justify-between items-center border-b border-gray-200 border-b-solid mb-5 pb-2">
        <div class="">
          <h2 class="text-lg font-semibold m-0">我的作品库</h2>
          <p class="my-1 text-sm text-gray-500">管理您生成的所有智能海报</p>
        </div>
        <span class="text-gray-500 text-[12px] flex items-center cursor-pointer" @click="useCozeGenius.sortHistory">
          按时间{{ useCozeGenius.sort === 'desc' ? '倒序' : '正序' }}排序
          <el-icon class="text-gray-500  ml-1"><ArrowDown /></el-icon>
        </span>
      </div>
      <div
        class="poster-genius mb-2xl w-full min-h-[200px] bg-white p-xl rounded-5 border-2 border-dashed border-gray-100"
        v-for="(item, index) in useCozeGenius.list" :key="item.id"
      >
        <p class="text-gray-400 mb-2">{{ item.task_time }}</p>
        <p class="text-gray-600 mb-2">做一个海报，内容：{{ item.text }}</p>
        <div class="w-full grid grid-cols-4 gap-4">
          <div
            class="h-[280px] rounded-4 border-1 border-solid border-gray-100 overflow-hidden relative"
            v-for="(img, i) in item.result" :key="img.index"
            v-loading="img.status === 'loading'"
            element-loading-text="努力生成中，马上就好"
          >
            <el-image
              v-if="img.url"
              :src="img.url"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="item.result.map(i => i.url)"
              show-progress
              :initial-index="i"
              fit="cover"
              class="w-full h-full"
              lazy
            >
            <template #placeholder>
              <div class="w-full h-full flex items-center justify-center">图片加载中...</div>
            </template>
            </el-image>
            <div
              v-if="img.status === 'error'"
              class="relative h-full flex flex-col items-center justify-center gap-3"
            >
              <div
                class="absolute top-0 left-0 bg-red-500 text-white text-xs px-2 py-1 rounded-br-4"
              >生成失败</div>
              <el-icon class="!text-2xl"><Picture/></el-icon>
              <p class="text-center text-sm text-black-500 m0">居然失败了，再试一下吧！</p>
              <el-tooltip
                :content="img.error"
              >
                <a class="cursor-pointer text-sm text-red-500">失败原因</a>
              </el-tooltip>
            </div>
            <div
              v-if="img.status === 'success'"
              class="flex justify-center gap-2 absolute bottom-2 left-1/2 transform -translate-x-1/2"
            >
              <el-button
                type="default"
                size="small"
                round
                class="flex items-center gap-1 !bg-white"
                @click="useCozeGenius.saveImage(img.url)"
              >
                <el-icon class="mr-2"><Download /></el-icon>
                下载
              </el-button>
              <el-popconfirm title="确定要删除吗？" @confirm="useCozeGenius.deleteImage(item.id, index, img.index, i)">
                <template #reference>
                  <el-button
                    type="default"
                    size="small"
                    round
                    class="flex items-center gap-1 !bg-white"
                  >
                    <el-icon class="mr-2"><Delete /></el-icon>
                    删除
                  </el-button>
                </template>
              </el-popconfirm>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-2xl mb-6xl text-center">
        <el-button
          v-if="useCozeGenius.more"
          type="primary"
          text
          @click="useCozeGenius.getHistory"
          :loading="useCozeGenius.loading"
        >
          {{!useCozeGenius.loading ? '点击加载更多' : '加载中...'}}
        </el-button>
        <span v-else>没有更多了</span>
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