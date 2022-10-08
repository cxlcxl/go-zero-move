import request from '../interceptors/marketing'

export function positionList(params) {
  return request({
    url: '/marketing/position/query',
    method: 'get',
    params
  })
}

export function positionPrice(params) {
  return request({
    url: '/marketing/position/price',
    method: 'get',
    params
  })
}

export function positionCategory(params) {
  return request({
    url: '/marketing/position/category',
    method: 'get'
  })
}

export function positionPlacement(params) {
  return request({
    url: '/marketing/position/placement',
    method: 'get',
    params
  })
}

export function positionElements(data) {
  return request({
    url: '/marketing/position/element',
    method: 'post',
    data
  })
}
