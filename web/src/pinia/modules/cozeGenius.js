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
    }
  })
}
// 获取图片生成历史记录
const getBLCTYImageHistory = (params) => {
  return service({
    url: '/extAiTask/getExtAiTaskList',
    method: 'get',
    params,
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
  const list = ref([]) // 历史记录列表
  const loading = ref(false)
  const page = ref(1)
  const more = ref(true)
  const pageSize = ref(10)

  // 发起生成请求的 Action
  const createPosterTask = async (params) => {
    isGenerating.value = true;
    try {
      const { code, data } = await createBLCTYImages(params);
      if (code !== 0) {
        ElMessage.error(data.msg)
        throw new Error(data.msg)
      }
      const taskData = {
        id: data.id,
        status: 'loading',
        text: params.prompt,
        task_time: useDateFormat(new Date(), 'YYYY-MM-DD HH:mm:ss'),
        result: initImages()
      }
      list.value.unshift(taskData)
      // 开启轮询
      pollTaskStatus(data.id);
      ElMessage.success('智能海报生成中，生成结果将在历史记录中展示')
      return taskData;
    } catch (error) {
      ElMessage.error('智能海报生成失败')
      throw error;
    } finally {
      isGenerating.value = false
    }
  };

  // 轮询逻辑 (递归 setTimeout)
  const delay = (ms) => new Promise(resolve => setTimeout(resolve, ms));
  const pollTaskStatus = async (taskId) => {
    try {
      const { data } = await getBLCTYImageResult(taskId);
      const { status, images } = data;

      const oldIndex = list.value.findIndex((i) => i.id === taskId)
      const oldData = oldIndex !== -1 ? list.value[oldIndex] : {}
      if (oldIndex === -1) {
        ElMessage.error(`获取结果失败，任务不存在: ${taskId}`)
        return;
      }
      // 更新图片列表状态
      const imagesList = oldData.result;
      if (images?.length) {
        const result = imagesList.map((item, i) => {
          const imgData = images[i] || item;
          let status = imgData.url ? 'success' : 'loading'
          if (imgData.error) {
            status = 'error'
          }
          const data = {
            url: imgData.url,
            index: imgData.index,
            status,
            error: imgData?.error
          };
          return data
        })
        const newData = {
          ...oldData,
          status,
          result
        };
        list.value[oldIndex] = newData
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
        await delay(15000);
        pollTaskStatus(taskId);
      }
    } catch (error) {
      console.error('获取结果失败，5秒后重试:', error);
      // 遇到网络报错不中断，延长间隔继续重试
      await delay(5000);
      pollTaskStatus(taskId);
    }
  };

  // 删除图片
  const deleteImage = (taskId, index, imgIndex, imgIndexInResult) => {
    deleteBLCTYImage(taskId, imgIndex).then(() => {
      list.value[index].result[imgIndexInResult] = {
        ...list.value[index].result[imgIndexInResult],
        status: 'error',
        error: '图片已删除',
        url: ''
      }
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

  // 获取历史记录
  const getHistory = async () => {
    loading.value = true
    try {
      const { code, data, msg } = await getBLCTYImageHistory({
        type: 'image',
        page: page.value,
        pageSize: pageSize.value
      })
      if (code !== 0) {
        ElMessage.error(msg)
        throw new Error(msg)
      }
      const newList = data?.list?.map((item) => {
        const options = JSON.parse(item.options || '{}');
        const result = JSON.parse(item.result || '[]');
        const status = item.status;
        let imgUrlNum = 0;
        const resultList = result.map((i) => {
          let imgStatus = i?.url ? 'success' : 'loading'
          if (i?.error) {
            imgStatus = 'error'
          }
          if (i?.url) {
            imgUrlNum++
          }
          return {
            index: i?.index,
            url: i?.url || '',
            status: imgStatus,
            error: i?.error || '',
          }
        })
        if (status === 'running' && imgUrlNum < 4) {
          pollTaskStatus(item.Id)
        }
        return {
          id: item.Id,
          status,
          task_time: useDateFormat(new Date(item.CreatedAt), 'YYYY-MM-DD HH:mm:ss'),
          text: options.prompt || '',
          result: resultList,
        }
      }) || [];
      if (page.value === 1) {
        list.value = [...newList]
      } else {
        list.value = [...list.value, ...newList]
      }
      more.value = list.value.length < data.total
      page.value++
    } catch (error) {
      console.log(error)
      ElMessage.error(error.msg || error || '获取历史记录失败')
    } finally {
      loading.value = false
    }
  }

  return {
    isGenerating,
    list,
    loading,
    more,
    createPosterTask,
    deleteImage,
    saveImage,
    getHistory
  }
})
