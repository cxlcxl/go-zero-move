import request from '@/utils/request'

export const fileUpload = process.env.VUE_APP_BASE_API + "/upload"
export const fileView = process.env.VUE_APP_BASE_API + "/file/download?file_id="
export const getFile = process.env.VUE_APP_BASE_API + '/file'

export function getConfigs(key) {
  return request({
    method: 'get',
    url: '/conf/' + key
  })
}

export function getConfigMap(key) {
  return request({
    method: 'get',
    url: '/conf-map/' + key
  })
}

export function rechargeSummary() {
  return request({
    method: 'get',
    url: '/dashboard/recharge'
  })
}

export function balance() {
  return request({
    method: 'get',
    url: '/dashboard/balance'
  })
}

export function platformBalance() {
  return request({
    method: 'get',
    url: '/dashboard/platform-balance'
  })
}

export function savePlatformBalance(data) {
  return request({
    method: 'post',
    url: '/dashboard/platform-balance',
    data
  })
}

export function variable(params) {
  return request({
    method: 'get',
    url: '/variable',
    params
  })
}

export function boundFlows(id, flow_type) {
  return request({url: '/flow/bound', method: 'get', params:{flow_id: id, flow_type}})
}
