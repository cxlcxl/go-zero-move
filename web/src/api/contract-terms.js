import request from '@/utils/request'

export function termsList(params) {
  return request({
    url: '/terms/list',
    method: 'get',
    params
  })
}

export function termsClose(id) {
  return request({
    url: '/terms/close',
    method: 'post',
    params: {id}
  })
}

export function termsCreate(data) {
  return request({
    url: '/terms/create',
    method: 'post',
    data
  })
}

export function termsUpdate(data) {
  return request({
    url: '/terms/update',
    method: 'post',
    data
  })
}

