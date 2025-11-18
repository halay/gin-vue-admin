import service from '@/utils/request'

export const createConsultation = (data) => {
  return service({ url: '/CN/createConsultation', method: 'post', data })
}

export const deleteConsultation = (params) => {
  return service({ url: '/CN/deleteConsultation', method: 'delete', params })
}

export const deleteConsultationByIds = (params) => {
  return service({ url: '/CN/deleteConsultationByIds', method: 'delete', params })
}

export const updateConsultation = (data) => {
  return service({ url: '/CN/updateConsultation', method: 'put', data })
}

export const findConsultation = (params) => {
  return service({ url: '/CN/findConsultation', method: 'get', params })
}

export const getConsultationList = (params) => {
  return service({ url: '/CN/getConsultationList', method: 'get', params })
}

export const getConsultationPublic = (params) => {
  return service({ url: '/CN/getConsultationPublic', method: 'get', params })
}

export const findConsultationPublic = (params) => {
  return service({ url: '/CN/findConsultationPublic', method: 'get', params })
}