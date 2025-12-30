<script setup>
import { ref, computed, onBeforeUnmount } from 'vue'
import { GoodsFilled, MagicStick, UserFilled, VideoCameraFilled, Download, Refresh, SwitchFilled, Delete, CaretRight, ArrowDown } from '@element-plus/icons-vue'
import { executeWorkflow, uploadImage, checkExecuteStatus, apiKey } from '@/plugin/app/api/cozeProduct'
import { ElMessage } from 'element-plus'


// 素材
const materialListConfig = ref([
  {
    name: '商品主图',
    desc: '清晰白底或产品抠图'
  },
  {
    name: '模特/场景图',
    desc: '上身效果或使用场景'
  }
]);
// 效果
const effectListConfig = ref([
  {
    label: '高清画质',
    value: '1'
  },
  {
    label: '动态光效',
    value: '2'
  },
  {
    label: 'AI运镜',
    value: '3'
  }
]);

// 创建ref数组存储所有上传组件引用
const uploadRefs = ref([])

const isGenerating = ref(false)

// 商品名称
const productName = ref('')

// 上传图片列表
const imageList = ref([])
// 图片上传中状态
const imgUploading = ref([])

const videoUrl = ref('')

// 选中的效果
const checkedEffects = ref(['1','2','3'])

const isGenerateDisabled = computed(() => {
  return false
  const imgCount = imageList.value.filter(item => item).length
  return productName.value.trim() === '' || imgCount !== 2 || isGenerating.value
})

const beforeUpload = (file, index) => {
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
  imgUploading.value[index] = true
  uploadImage(file).then(res => {
    if (res.code !== 0) {
      // 上传失败，清除文件
      uploadRefs.value[index]?.clearFiles()
      imgUploading.value[index] = false
      ElMessage.error(res.msg || '上传失败')
      return
    }
    const id = res.data.id
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = (res) => {
      imageList.value[index] = {
        id,
        url: reader.result
      }
      imgUploading.value[index] = false
    }
  }).catch(err => {
    ElMessage.error(err.msg || '图片上传错误')
    imgUploading.value[index] = false
    uploadRefs.value[index]?.clearFiles()
  })
}

// 生成视频
const handleGenerate = async () => {
  isGenerating.value = true
  try {
    // 创建工作流
    const res = await executeWorkflow(0, {
      api_token: apiKey,
      product_img: JSON.stringify({ file_id: imageList.value[0].id }),
      person_img: JSON.stringify({ file_id: imageList.value[1].id }),
      product: productName.value,
    });
    console.log('工作流结果:', res);
    const executeId = res.execute_id
    // 轮询查询工作流创建结果
    const result = await pollVideoStatus(0, executeId)
    console.log('工作流结果:', result, result.task_Id);
    // // 创建任务工作流轮训查询结果
    const taskRes = await pollWorkflowStatus({
      task_id: result.task_Id,
    });
    videoUrl.value = taskRes
    ElMessage.success('视频生成成功')
    isGenerating.value = false
  } catch (error) {
    console.error(error)
    ElMessage.error(error.msg || error || '视频生成错误')
    isGenerating.value = false
  }
}
const pollWorkflowStatus = async (params, options = {}) => {
  // 默认参数配置
  const { maxRetries = 40, pollInterval = 20000 } = options;
  
  let attempts = 0;

  while (attempts < maxRetries && isGenerating.value) {
    try {
      // 调用API检查执行状态
      const res = await executeWorkflow(1, params, false);
      // 解析输出数据，添加错误处理
      const data = JSON.parse(res.data || '{}')
      if (data.video_url && data.video_url.includes('https://')) {
        return data.video_url;
      }
      attempts++;
      await new Promise(resolve => setTimeout(resolve, pollInterval));
    } catch (error) {
      console.error('轮询出错:', error.message || error);
      return Promise.reject({msg: error.message || error}); // 抛出错误供外部捕获
    }
  }
  // 轮询超时
  return Promise.reject({msg: '查询超时，请稍后在作品列表中查看'});
};
// 轮训工作流结果
const pollVideoStatus = async (wId, executeId, options = {}) => {
  // 默认参数配置
  const { maxRetries = 30, pollInterval = 10000 } = options;
  const TIMEOUT_ERROR = '轮询超时，请稍后在作品列表中查看';
  const PROCESSING_LOG = '生成中... 第 %d 次轮询';
  
  let attempts = 0;

  while (attempts < maxRetries && isGenerating.value) {
    try {
      // 调用API检查执行状态
      const res = await checkExecuteStatus(wId, executeId);
      
      // 验证API返回数据结构
      if (!res.data || !Array.isArray(res.data) || res.data.length === 0) {
        return Promise.reject({msg: '无效的API返回数据'});
      }
      
      const { execute_status, output, error_message } = res.data[0];
      attempts++;

      // 处理失败状态
      if (execute_status === 'Fail') {
        return Promise.reject({msg: '视频生成失败'});
      }

      // 处理成功状态
      if (execute_status === 'Success') {
        console.log('任务工作流成功:', execute_status);
        console.log('任务工作流结果:', output);
        
        try {
          // 解析输出数据，添加错误处理
          const outputData = JSON.parse(output);
          const outputContent = JSON.parse(outputData.Output);         
          // 轮询成功，返回结果
          return outputContent;
        } catch (parseError) {
          return Promise.reject({msg: `数据解析错误: ${parseError.message}`});
        }
      }

      // 处理进行中状态
      console.log(PROCESSING_LOG, attempts);
      await new Promise(resolve => setTimeout(resolve, pollInterval));

    } catch (error) {
      console.error('轮询出错:', error.message || error);
      return Promise.reject({msg: error.message || error}); // 抛出错误供外部捕获
    }
  }
  // 轮询超时
  return Promise.reject({msg: TIMEOUT_ERROR});
};


// 重新生成视频
const handleRefresh = () => {
  isGenerating.value = false
  videoUrl.value = ''
  handleGenerate()
}

// 历史视频列表
const historyList = ref([])
const tabList = ref(false)

// 播放视频
const handlePlay = (item) => {
  console.log(item)
}
// 下载视频
const handleDownload = async () => {
  if (!videoUrl.value) {
    ElMessage.error('视频URL为空')
    return
  }
  ElMessage.primary('视频下载中，请稍后...')
  downLoading.value = true;
  try {
    const response = await fetch(videoUrl.value);
    if (!response.ok) {
       ElMessage.error('网络响应错误')
       return
    }
    // 将响应转换为二进制数据流 (Blob)
    const blob = await response.blob();
    const url = window.URL.createObjectURL(blob);
    
    // 创建虚拟链接并触发点击
    const a = document.createElement('a');
    a.href = url;
    const suffix = url.slice(url.lastIndexOf('.'))
    a.download = 'video_' + new Date().getTime() + suffix;
    document.body.appendChild(a);
    a.click();
    
    // 释放内存
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);
  } catch (error) {
    console.error('下载失败:', error);
    ElMessage.error('下载失败，请检查网络或跨域设置')
  } finally {
    downLoading.value = false;
  }
}
// 删除视频
const handleDelete = (item) => {
  console.log(item)
}

// 组件卸载前清除接口轮训
onBeforeUnmount(() => {
  isGenerating.value = false
})
</script>

<template>
  <div>
  <div class="min-h-screen bg-gray-50 py-10">
    <div class="text-center mb-10">
      <h1 class="mb-2 text-3xl font-bold">
        AI 商品宣传片智能生成
      </h1>
      <p class="text-gray-500 my-0">只需两张图片，AI为您的商品打造电影级5秒竖屏宣传短片</p>
    </div>
    <div class="flex justify-center gap-10 max-w-6xl mx-auto">
      <div class="bg-white rounded-2xl shadow-lg p-6 w-2/5">
        <div class="flex items-center gap-2 mb-4">
          <div class="w-1 h-6 bg-blue-500 rounded"></div>
          <h2 class="text-lg font-semibold">素材配置</h2>
        </div>
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">商品名称</label>
          <el-input
            v-model="productName"
            placeholder="例如：法式复古丝绒连衣裙"
            size="large"
          />
        </div>
        <div class="mb-10">
          <div class="flex justify-between items-center mb-2">
            <label class="block text-sm font-medium text-gray-700">核心素材</label>
            <span class="text-xs text-gray-500">支持 JPG, PNG 格式</span>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <template v-for="(item, index) in materialListConfig" :key="index">
              <div v-if="imageList[index]" class="uploader-box-img">
                <img
                  :src="imageList[index].url"
                  class="w-full h-full object-cover"
                  :zoom-rate="1.2"
                  :max-scale="7"
                  :min-scale="0.2"
                  show-progress
                />
                <div class="absolute top-2 right-2 bg-red-500 text-white text-xs px-1 rounded-full" @click="imageList[index] = ''">删除</div>
              </div>
              <el-upload
                v-else
                v-loading="imgUploading[index]"
                :ref="el => uploadRefs[index] = el"
                class="uploader-box h-[186px]"
                :show-file-list="false"
                accept="image/*"
                :http-request="(options) => handleUpload(options.file, index)"
                :before-upload="(file) => beforeUpload(file, index)"
                :limit="1"
              >
                <div class="flex flex-col items-center justify-center py-4">
                  <el-icon class="text-xl text-[#9CA3AF] mb-2 w-10 h-10 rounded-lg bg-[#F0F2F5] flex items-center justify-center">
                    <GoodsFilled v-if="index === 0" />
                    <UserFilled v-if="index === 1" />
                  </el-icon>
                  <div class="text-gray-600 text-[14px]">{{ item.name }}</div>
                  <div class="text-gray-400 text-[12px]">{{ item.desc }}</div>
                </div>
              </el-upload>
            </template>
          </div>
        </div>
        <div class="space-y-2 flex flex-col items-center justify-center">
          <el-button
            type="default"
            size="large"
            :loading="isGenerating"
            :disabled="isGenerateDisabled"
            class="w-full py-6 bg-white rounded-xl text-base font-medium shadow-[0_5px_10px_#FDE998] border-none hover:bg-[#FDE998] hover:text-white mb-2"
            @click="handleGenerate"
          >
            <el-icon v-if="!isGenerating" class="mr-2"><MagicStick /></el-icon>
            生成产品视频
          </el-button>
          <div class="flex items-center justify-center">
            <el-checkbox-group v-model="checkedEffects" disabled>
              <el-checkbox v-for="item in effectListConfig" :key="item.value" :label="item.value" :value="item.value">
                {{ item.label }}
              </el-checkbox>
            </el-checkbox-group>
          </div>
        </div>
      </div>

      <!-- 右侧视频预览 -->
      <div class="relative">
        <div class="bg-black rounded-[20px] p-2 w-[320px] h-[540px] relative shadow-2xl">
          <div class="absolute top-4 left-1/2 transform -translate-x-1/2 w-20 h-4 bg-gray-800 rounded-full"></div>
          <div class="w-full h-full bg-gray-900 rounded-lg flex flex-col items-center justify-center">
            <div v-if="!videoUrl" class="text-center w-2/3">
              <div class="w-16 h-16 rounded-full bg-gray-800 flex items-center justify-center mx-auto mb-4">
                <el-icon class="text-gray-500 text-2xl"><VideoCameraFilled /></el-icon>
              </div>
              <h3 class="text-white font-medium mb-2">等待生成宣传片</h3>
              <p class="text-gray-500 text-sm">AI将自动识别商品特征，为您匹配最佳的运镜模板与背景音乐</p>
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
            @click="handleRefresh"
          >
            <el-icon class="mr-2"><Refresh /></el-icon>
            重新生成
          </el-button>
        </div>
      </div>
    </div>
    <!-- 历史记录 -->
    <div class="max-w-6xl mx-auto mt-20" v-if="historyList.length">
      <div class="flex justify-between items-center border-b border-gray-200 border-b-solid mb-5 pb-2">
        <div class="">
          <h2 class="text-lg font-semibold m-0">我的作品库</h2>
          <p class="my-1 text-sm text-gray-500">管理您生成的所有商品宣传视频</p>
        </div>
        <span class="text-gray-500 text-[12px] flex items-center">
          按时间排序
          <el-icon class="text-gray-500 cursor-pointer ml-1"><ArrowDown /></el-icon>
        </span>
      </div>
      <div class="grid grid-cols-5 gap-3">
        <div
          v-for="(item, index) in historyList"
          :key="index"
          class="rounded-lg shadow-md overflow-hidden relative min-h-[280px]">
          <el-image
            :src="item"
            class="w-full h-[186px]"
            fit="cover"
          />
          <div class="absolute bottom-0 left-0 right-0 top-0 p-3 text-white bg-[linear-gradient(180deg,rgba(0,0,0,0)30%,rgba(0,0,0,0.2)100%)] text-white flex flex-col justify-end">
            <h4 class="text-sm mt-0 mb-1">{{ '遇见咖啡' }}</h4>
            <div class="flex justify-between items-center text-[12px]">
              <span>10月24日 14:20</span>
              <span>5s</span>
            </div>
          </div>
          <div class="absolute top-0 right-0 left-0 bottom-0 flex flex-col justify-center items-center opacity-0 hover:opacity-100 hover:transition-all">
            <el-icon class="w-12 h-12 text-2xl rounded-full bg-gray-100 opacity-80 flex items-center justify-center text-[#F49D0B] shadow cursor-pointer" @click="handlePlay(item)"><CaretRight /></el-icon>
            <div class="flex items-center justify-center gap-8 mt-4">
              <el-icon class="w-8 h-8 rounded-full bg-[rgba(0,0,0,0.1)] text-white shadow text-4 flex items-center justify-center cursor-pointer" @click="handleDownload(item)"><Download /></el-icon>
              <el-icon class="w-8 h-8 rounded-full bg-[rgba(0,0,0,0.1)] text-white shadow text-4 flex items-center justify-center cursor-pointer" @click="handleDelete(item)"><Delete /></el-icon>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<style scoped>
 :deep(.el-button) {
  --el-button-bg-color: var(--primary-color);
 }
:deep(.el-input__wrapper) {
  --el-input-border-radius: 0.75rem;
}
:deep(.el-input--large) {
  --el-input-inner-height: 48px;
}
:deep(.el-button) {
  --el-button-bg-color: var(--primary-color);
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

:deep(.el-checkbox__inner) {
  --el-checkbox-border-radius: 1rem;
}
:deep(.el-checkbox__input.is-disabled.is-checked .el-checkbox__inner) {
  --el-checkbox-disabled-checked-input-fill: var(--el-color-primary);
  --el-checkbox-disabled-checked-input-border-color: var(--el-color-primary);
  --el-checkbox-disabled-checked-icon-color: #fff;
}
:deep(.el-checkbox__input.is-disabled+span.el-checkbox__label) {
  --el-disabled-text-color: var(--el-color-primary);
}
</style>