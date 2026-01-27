import service from '@/utils/request'
import { ElMessage, ElNotification } from 'element-plus'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useDateFormat } from '@vueuse/core'

// 创建图片生成任务
const createBLCTYImages = (body) => {
  let formData = new FormData()
  if (!(body instanceof FormData)) {
    Object.keys(body).forEach((key) => {
      formData.append(key, body[key])
    })
  } else {
    formData = body
  }
  return service({
    url: '/yApi/getBLCTYImages',
    method: 'post',
    data: formData,
    donNotShowLoading: true
  })
}
// 获取图片生成结果
const getBLCTYImageResult = (id) => {
  return service({
    url: '/yApi/getBLCTYImageResult',
    method: 'get',
    params: {
      id
    },
    donNotShowLoading: true
  })
}
// 删除图片生成
const deleteBLCTYImage = (id, indexes) => {
  return service({
    url: '/yApi/deleteBLCTYImage',
    method: 'put',
    params: {
      id,
      indexes
    },
    // 不显示 loading 动画
    donNotShowLoading: true
  })
}

const initImages = () => {
  const list = new Array(4).fill({}).map((item, i) => ({
    status: 'loading',
    url: '',
    index: i,
    error: ''
  }))
  return list
}

export const useCozeGeniusStore = defineStore('cozeGenius', () => {

  const isGenerating = ref(false) // 是否正在生成中
  const taskId = ref(0) // 图片生成任务 ID
  const imageList = ref([]) // 图片列表
  const promptContent = ref('') // 提示内容
  const taskTime = ref('') // 任务创建时间

  // 发起生成请求的 Action
  const createPosterTask = async (params) => {
    isGenerating.value = true;
    try {
      imageList.value = initImages()
      promptContent.value = params.prompt
      taskTime.value = useDateFormat(new Date(), 'YYYY-MM-DD HH:mm:ss')
      const { code, data } = await createBLCTYImages(params);
      if (code !== 0) {
        ElMessage.error(data.msg)
        throw new Error(data.msg)
      }
      taskId.value = data.id;
      // 开启轮询
      pollTaskStatus(data.id);
      return taskId.value;
    } catch (error) {
      ElMessage.error('智能海报生成失败')
      imageList.value = []
      throw error;
    } finally {
      isGenerating.value = false;
      
    }
  };

  // 轮询逻辑 (递归 setTimeout)
  const pollTaskStatus = async (taskId) => {
    try {
      const { data } = await getBLCTYImageResult(taskId);
      const { status, images } = data;
      // 更新图片列表状态
      if (images?.length) {
        imageList.value = imageList.value.map((item, index) => {
          const data = item;
          const img = images[index] || {};
          let status = img?.url ? 'success' : 'loading';
          if (img.error) {
            status = 'error'
          }
          data.status = status;
          data.url = img?.url || '';
          data.index = img?.index;
          data.error = img?.error || '';
          return item
        })
      }

      if (status === 'completed' && images?.length === 4) {
        ElNotification({
          title: 'AI海报生成完成',
          message: `4 张海报已全部生成`,
          type: 'success',
          duration: 5000
        });
      } else {
        // 未完成：3秒后继续下一次轮询
        setTimeout(() => pollTaskStatus(taskId), 15000);
      }
    } catch (error) {
      console.error('获取结果失败，5秒后重试:', error);
      // 遇到网络报错不中断，延长间隔继续重试
      setTimeout(() => pollTaskStatus(taskId), 5000);
    }
  };

  // 删除图片
  const deleteImage = (item) => {
    deleteBLCTYImage(taskId.value, item.index).then(() => {
      imageList.value = imageList.value.map((i) => {
        if (i.index === item.index) {
          i.status = 'error'
          i.error = '图片已删除'
          i.url = ''
        }
        return i
      })
      ElMessage.success('图片删除成功')
    })
  }

  // 下载图片
  const saveImage = (url) => {
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

  return {
    isGenerating,
    imageList,
    promptContent,
    taskTime,
    createPosterTask,
    pollTaskStatus,
    deleteImage,
    saveImage
  }
})
