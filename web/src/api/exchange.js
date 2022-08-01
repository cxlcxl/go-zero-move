import request from '@/utils/request'

export function exchangeList(params) {
  return request({
    url: '/exchange/list',
    method: 'get',
    params
  })
}

export function exchangeCreate(data) {
  return request({
    url: '/exchange/create',
    method: 'post',
    data
  })
}

export function exchangeUpdate(data) {
  return request({
    url: '/exchange/update',
    method: 'post',
    data
  })
}

export function exchangeDestroy(id) {
  return request({
    url: '/exchange/destroy',
    method: 'post',
    data: {
      id
    }
  })
}
