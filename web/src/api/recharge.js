import request from '../utils/request'

export function rechargeList(params) {
  return request({url: '/recharge/list', method: 'get', params})
}


export function rechargeDetailList(params) {
  return request({url: '/recharge/detail', method: 'get', params})
}

export function rechargeListDownload(params) {
  return request({url: '/recharge/download', method: 'get', params})
}

export function rechargeDetailDownload(params) {
  return request({url: '/recharge/detail-download', method: 'get', params})
}

export function rechargeCreate(data) {
  return request({url: '/recharge/create', method: 'post', data})
}

export function rechargeDestroy(id) {
  return request({url: '/recharge/destroy', method: 'post', data: {id}})
}

export function changeRechargeState(data) {
  return request({url: '/recharge/state', method: 'post', data})
}

export function rechargeDetail(recharge_id) {
  return request({url: '/recharge/' + recharge_id, method: 'get'})
}

export function rechargeForTrans(id) {
  return request({url: '/recharge/for-trans', method: 'get', params:{customer_id: id}})
}

export function rechargePlatformList(params) {
  return request({url: '/recharge-platform/list', method: 'get', params})
}

export function rechargePlatformListDownload(params) {
  return request({url: '/recharge-platform/download', method: 'get', params})
}

export function rechargePlatformCreate(data) {
  return request({url: '/recharge-platform/create', method: 'post', data})
}

export function rechargePlatformDestroy(id) {
  return request({url: '/recharge-platform/destroy', method: 'post', data:{id}})
}
