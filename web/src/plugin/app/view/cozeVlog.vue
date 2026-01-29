<script setup>
import { ref, computed } from 'vue'
import { UploadFilled, MagicStick, PictureFilled, Shop, StarFilled, VideoCameraFilled, Download, Refresh, WarningFilled } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { defineCozeWorkflowStore } from '@/pinia/modules/cozeWorkflow'
const useVlogWorkflowStore = defineCozeWorkflowStore('vlog');

const imageList = computed(() => useVlogWorkflowStore.imageList)
const isGenerating = computed(() => useVlogWorkflowStore.isGenerating)
const isGenerateDisabled = computed(() => {
  const imgCount = imageList.value.filter(item => item.id).length
  return shopName.value.trim() === '' || imgCount !== 4 || isGenerating.value
})

const videoUrl = computed(() => useVlogWorkflowStore.videoUrl)

// 创建ref数组存储所有上传组件引用
const uploadRefs = ref([])
// 店铺名称
const shopName = ref(useVlogWorkflowStore.text || '')

const beforeUpload = (file) => {
  // 限制图片大小最大512M
  const maxSize = 512 * 1024 * 1024
  if (file.size > maxSize) {
    ElMessage.error('图片大小不能超过512MB')
    return false
  }
  return true
}
// 上传图片
const handleUpload = (file, index) => {
  useVlogWorkflowStore.uploadImages(index, { file })
}

// 生成视频
const handleGenerate = async () => {
  if (isGenerateDisabled.value) {
    ElMessage.error('请检查店铺名称和图片是否上传完成')
    return
  }
  useVlogWorkflowStore.text = shopName.value
  useVlogWorkflowStore.createTask({
    img_1: JSON.stringify({ file_id: imageList.value[0].id }),
    img_2: JSON.stringify({ file_id: imageList.value[1].id }),
    img_3: JSON.stringify({ file_id: imageList.value[2].id }),
    img_4: JSON.stringify({ file_id: imageList.value[3].id }),
    text: shopName.value,
  })
}

const handleReset = () => {
  useVlogWorkflowStore.onResetCreate()
  handleGenerate()
}

// 获取任务历史记录
if (useVlogWorkflowStore.list.length === 0) {
  useVlogWorkflowStore.getHistory()
}

</script>

<template>
  <div>
  <div class="min-h-screen bg-gray-50 py-10">
    <div class="text-center mb-10">
      <h1 class="mb-2 text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-[#4C52EB] to-[#6C44EA]">AI 探店视频智能生成</h1>
      <p class="text-gray-500 my-0">上传图片，一键生成 5 秒精美探店短视频</p>
    </div>
    <div class="flex justify-center gap-10 max-w-6xl mx-auto">
      <div class="bg-white rounded-2xl shadow-lg p-6 w-2/5">
        <div class="flex items-center gap-2 mb-4">
          <div class="w-1 h-6 bg-blue-500 rounded"></div>
          <h2 class="text-lg font-semibold">视频配置</h2>
        </div>
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">店铺名称</label>
          <el-input
            v-model="shopName"
            placeholder="请输入店铺完整名称..."
            size="large"
            :disabled="isGenerating"
          />
        </div>
        <div class="mb-8">
          <div class="flex justify-between items-center mb-2">
            <label class="block text-sm font-medium text-gray-700">上传展示图片</label>
            <span class="text-xs text-gray-500">需上传 4 张图片</span>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <template v-for="(item, index) in ['封面', '细节', '环境', '特色']" :key="index">
              <div v-if="imageList[index]?.id" class="uploader-box-img !h-[186px]">
                <img
                  :src="imageList[index].url"
                  class="w-full h-full object-cover"
                  :zoom-rate="1.2"
                  :max-scale="7"
                  :min-scale="0.2"
                  show-progress
                />
                <div
                  v-if="!isGenerating"
                  class="absolute top-2 right-2 bg-red-500 text-white text-xs px-1 rounded-full" @click="useVlogWorkflowStore.deleteImage(index)"
                >删除</div>
              </div>
              <el-upload
                v-else
                v-loading="imageList[index].status === 'loading'"
                :ref="el => uploadRefs[index] = el"
                class="uploader-box h-[186px]"
                :show-file-list="false"
                accept="image/*"
                :http-request="(options) => handleUpload(options.file, index)"
                :before-upload="(file) => beforeUpload(file, index)"
                :limit="1"
              >
                <div class="flex flex-col items-center justify-center py-4">
                  <el-icon class="text-lg text-[#9CA3AF] mb-2 w-8 h-8 rounded-full bg-[#F0F2F5] flex items-center justify-center">
                    <UploadFilled v-if="index === 0" />
                    <PictureFilled v-if="index === 1" />
                    <Shop v-if="index === 2" />
                    <StarFilled v-if="index === 3" />
                  </el-icon>
                  <div class="text-gray-600 text-sm">第{{index + 1}}张({{item}})</div>
                </div>
              </el-upload>
            </template>
          </div>
        </div>
        <div class="space-y-2 flex flex-col items-center justify-center">
          <el-button
            type="primary"
            size="large"
            :disabled="isGenerateDisabled"
            :loading="isGenerating"
            class="w-full py-6 bg-blue-600 !hover:bg-blue-700 text-white rounded-xl text-base font-medium"
            @click="handleGenerate"
          >
            <el-icon v-if="!isGenerating" class="mr-2"><MagicStick /></el-icon>
            开始生成5秒短视频
          </el-button>
          <p class="text-xs text-gray-400 text-center">生成过程约需5分钟左右，请耐心等待</p>
        </div>
      </div>

      <!-- 右侧视频预览 -->
      <div class="relative">
        <div class="bg-black rounded-[20px] p-2 w-[320px] h-[670px] relative shadow-2xl">
          <div class="absolute top-4 left-1/2 transform -translate-x-1/2 w-20 h-4 bg-gray-800 rounded-full"></div>
          <div class="w-full h-full bg-gray-900 rounded-lg flex flex-col items-center justify-center">
            <div v-if="!videoUrl" class="text-center w-2/3">
              <div class="w-16 h-16 rounded-full bg-gray-800 flex items-center justify-center mx-auto mb-4">
                <el-icon class="text-gray-500 text-2xl"><VideoCameraFilled /></el-icon>
              </div>
              <h3 class="text-white font-medium mb-2">等待生成预览</h3>
              <p v-if="!isGenerating" class="text-gray-500 text-sm">在左侧输入店铺名称并上传图片，生成的视频将在此处播放</p>
              <p v-else class="text-gray-500 text-sm">视频正在生成中，可处理其它工作，完成后将在下方显示</p>
              <!-- <div v-if="useVlogWorkflowStore.taskStatus === 'failed' && useVlogWorkflowStore.taskId" class="mt-2xl">
                <a class="text-red-500 text-sm mt-2 cursor-pointer border-b border-red-500" @click="useVlogWorkflowStore.retryTask(useVlogWorkflowStore.taskId)">
                  视频生成失败，点击此处重新生成
                </a>
              </div> -->
            </div>
            <video
              v-else
              ref="videoPlayer"
              class="w-full h-full object-cover rounded-lg"
              :src="videoUrl"
              controls
            ></video>
          </div>
        </div>
        
        <!-- 手机底部按钮 -->
        <div class="flex justify-center gap-4 mt-8">
          <a
            v-if="videoUrl"
            class="flex items-center gap-1 bg-white rounded-full px-4 border border-solid border-gray-200 text-sm text-gray-600 font-medium no-underline"
            :disabled="!videoUrl"
            :link="true"
            :href="videoUrl"
            target="_blank"
          >
            <el-icon class="mr-2"><Download /></el-icon>
            下载视频
          </a>
          <a
            v-else
            class="flex items-center gap-1 bg-white rounded-full px-4 border border-solid border-gray-200 text-sm text-gray-400 font-medium no-underline cursor-not-allowed"
          >
            <el-icon class="mr-2"><Download /></el-icon>
            下载视频
          </a>
          <el-button
            type="default"
            round
            class="flex items-center gap-1 bg-white"
            :disabled="!videoUrl"
            @click="handleReset()"
          >
            <el-icon class="mr-2"><Refresh /></el-icon>
            重新生成
          </el-button>
        </div>
      </div>
    </div>
    <!-- 历史记录 -->
    <div class="max-w-6xl mx-auto mt-20">
      <div class="flex justify-between items-center border-b border-gray-200 border-b-solid mb-5 pb-2">
        <div class="">
          <h2 class="text-lg font-semibold m-0">我的作品库</h2>
          <p class="my-1 text-sm text-gray-500">管理您生成的所有商户宣传视频</p>
        </div>
        <span class="text-gray-500 text-[12px] flex items-center cursor-pointer" @click="useVlogWorkflowStore.sortHistory">
          按时间{{ useVlogWorkflowStore.sort === 'desc' ? '倒序' : '正序' }}排序
          <el-icon class="text-gray-500 cursor-pointer ml-1"><ArrowDown /></el-icon>
        </span>
      </div>
      <div class="grid grid-cols-2 gap-5">
        <div
          v-for="(item, index) in useVlogWorkflowStore.list"
          :key="index"
          class="rounded-lg shadow-md overflow-hidden relative"
        >
          <div
            class="relative w-full h-[320px] bg-black/50 rounded-lg"
            v-loading="item.status === 'running'"
            element-loading-text="努力生成中，马上就好"
          >
            <video
              v-if="item.result"
              class="w-full h-full"
              :src="item.result"
              controls
            ></video>
            <div v-if="item.status === 'failed'" class="w-full h-full flex flex-col items-center justify-center">
              <div class="text-white/50 text-4xl">
                <el-icon><WarningFilled /></el-icon>
              </div>
              <p class="text-white/80 text-sm mt-3 px-3">{{ item.error }}</p>
              <!-- <a class="text-white/80 text-sm cursor-pointer mt-6" @click="useVlogWorkflowStore.retryTask(item.id)">
                点击此处重新生成
              </a> -->
            </div>
          </div>
          <div class=" p-3 bg-white text-black flex flex-col justify-end">
            <h4 class="text-sm mt-0 mb-1">{{ item.text }}</h4>
            <div class="flex justify-between items-center text-[12px]">
              <span>{{ item.task_time }}</span>
              <span>{{ item.status === 'running' ? '生成中' : '已完成' }}</span>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-2xl mb-6xl text-center">
        <el-button
          v-if="useVlogWorkflowStore.more"
          type="primary"
          text
          @click="useVlogWorkflowStore.getHistory"
          :loading="useVlogWorkflowStore.loading"
        >
          {{!useVlogWorkflowStore.loading ? '点击加载更多' : '加载中...'}}
        </el-button>
        <div v-else>
          <el-empty v-if="useVlogWorkflowStore.list.length === 0" description="暂无历史记录" />
          <span v-else>没有更多了</span>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<style scoped>
:deep(.el-input__wrapper) {
  --el-input-border-radius: 0.75rem;
}
:deep(.el-input--large) {
  --el-input-inner-height: 48px;
}

.uploader-box :deep(.el-upload),
.uploader-box-img {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px dashed var(--el-border-color);
  border-radius: 0.75rem;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}
.uploader-box-img {
  border: 2px solid var(--el-border-color);
}
.uploader-box :deep(.el-upload:hover) {
  border-color: var(--el-color-primary);
}

</style>