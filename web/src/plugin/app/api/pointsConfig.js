import service from '@/utils/request'

export const createPointsConfig = (data) => service({ url: '/pcfg/createPointsConfig', method: 'post', data })
export const deletePointsConfig = (params) => service({ url: '/pcfg/deletePointsConfig', method: 'delete', params })
export const deletePointsConfigByIds = (params) => service({ url: '/pcfg/deletePointsConfigByIds', method: 'delete', params })
export const updatePointsConfig = (data) => service({ url: '/pcfg/updatePointsConfig', method: 'put', data })
export const findPointsConfig = (params) => service({ url: '/pcfg/findPointsConfig', method: 'get', params })
export const getPointsConfigList = (params) => service({ url: '/pcfg/getPointsConfigList', method: 'get', params })
export const getPointsConfigPublic = (params) => service({ url: '/pcfg/getPointsConfigPublic', method: 'get', params })

