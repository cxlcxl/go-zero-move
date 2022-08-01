import request from '../utils/request'

export function transList(params) {
  return request({url: '/trans/list', method: 'get', params})
}

export function transListDownload(params) {
  return request({url: '/trans/download', method: 'get', params})
}

export function transCreate(data) {
  return request({url: '/trans/create', method: 'post', data})
}

export function transDestroy(id) {
  return request({url: '/trans/destroy', method: 'post', data:{id}})
}

export function getAdvanceTransfers(customer_id) {
  return request({url: '/trans/advance-paid', method: 'get', params: {customer_id}})
}

export function transPlatformList(params) {
  return request({url: '/trans-platform/list', method: 'get', params})
}

export function transPlatformListDownload(params) {
  return request({url: '/trans-platform/download', method: 'get', params})
}

export function transPlatformCreate(data) {
  return request({url: '/trans-platform/create', method: 'post', data})
}

export function transPlatformUpdate(data) {
  return request({url: '/trans-platform/update', method: 'post', data})
}

export function transPlatformDestroy(id) {
  return request({url: '/trans-platform/destroy', method: 'post', data:{id}})
}
