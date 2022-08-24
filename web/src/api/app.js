import request from '@/interceptors/app'

export function appList(params) {
  return request({
    url: '/app/list',
    method: 'get',
    params
  })
}

export function appCreate(data) {
  return request({
    url: '/app/create',
    method: 'post',
    data
  })
}

export function appUpdate(data) {
  return request({
    url: '/app/update',
    method: 'post',
    data
  })
}

export function appInfo(id) {
  return request({
    url: '/app/info',
    method: 'get',
    params: {id}
  })
}

