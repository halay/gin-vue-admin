import axios from 'axios'

const baseURL = import.meta.env.DEV ? 'https://api.coze.cn' : 'https://api.coze.cn'
const token = 'sat_NTcKXkyzHAvdF1RcmFRRxH4VhVFiDpxrPfqV7E6B48vhk7nho3oLAmCQDrpiTcKy'
export const apiKey = '2af09cae-106a-49db-9361-00b7b0b20ca8'
const workflowIds = ['7579514432432472127', '7579514463738167323']

const axiosInstance = axios.create({
  baseURL,
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  }
})
function isBusinessError(data) {
  // 这里根据后端返回的业务错误格式进行判断
  return data != null && typeof data === 'object' && 'code' in data && data.code !== 0;
}
function handleAxiosError(error) {
  // 网络错误
  if (!error.response) {
    return {
      code: -1,
      msg: '网络错误，请检查您的网络连接'
    }
  }

  const { response } = error;
  const status = response.status;

  // 从响应数据中提取错误信息
  const responseData = response.data;
  // 安全地访问响应数据
  const data = responseData;
  const errorMessage = data?.msg || `请求失败，状态码: ${status}`;

  const errorCode = data?.code || status;

  return {
    code: errorCode,
    msg: errorMessage
  };
}
axiosInstance.interceptors.response.use(
  (response) => {
    const data = response.data;

    // 检查是否为业务错误
    if (isBusinessError(data)) {
      return Promise.reject({
        code: data.code,
        msg: data.msg
      });
    }

    // 直接返回响应数据
    return data;
  },
  (error) => {
    // 处理HTTP错误
    const apiError = handleAxiosError(error);
    // globalErrorHandler(apiError);
    return Promise.reject(apiError);
  }
);

// 上传图片到Coze API
export const uploadImage = (file) => {
  const formData = new FormData()
  formData.append('file', file)
  return axios({
    method: 'post',
    baseURL,
    url: '/v1/files/upload',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': `multipart/form-data;`
    },
    data: formData
  }).then(res => res.data)
}

// 执行工作流
export const executeWorkflow = (wId, params, isAsync = true) => {
  const data = {
    workflow_id: workflowIds[wId],
    is_async: isAsync,
    parameters: {
      ...params
    }
  }
  return axiosInstance.post('/v1/workflow/run', data)
}
// 查看工作流异步执行状态
export const checkExecuteStatus = (wId, executeId) => {
  return axiosInstance.get(`/v1/workflows/${workflowIds[wId]}/run_histories/${executeId}`)
}
