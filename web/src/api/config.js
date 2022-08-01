import request from '@/utils/request'

export function configList(params) {
  return request({
    url: '/conf/list',
    method: 'get',
    params
  })
}

export function confClose(id) {
  return request({
    url: '/conf/close',
    method: 'post',
    params: {id}
  })
}

export function confCreate(data) {
  return request({
    url: '/conf/create',
    method: 'post',
    data
  })
}

export function confUpdate(data) {
  return request({
    url: '/conf/update',
    method: 'post',
    data
  })
}

