import request from '@/utils/request'

export function contractList(params) {
  return request({
    url: '/contract/list',
    method: 'get',
    params
  })
}

export function contractCreate(data) {
  return request({
    url: '/contract/create',
    method: 'post',
    data
  })
}

export function contractInfo(id) {
  return request({
    url: '/contract/' + id,
    method: 'get'
  })
}

export function contractUpdate(data) {
  return request({
    url: '/contract/update',
    method: 'post',
    data
  })
}

// 协议/附件 API
export function agreementCreate(data) {
  return request({
    url: '/agreement/create',
    method: 'post',
    data
  })
}

export function agreementUpdate(data) {
  return request({
    url: '/agreement/update',
    method: 'post',
    data
  })
}

export function agreementDestroy(id) {
  return request({
    url: '/agreement/destroy',
    method: 'post',
    data: {id}
  })
}

