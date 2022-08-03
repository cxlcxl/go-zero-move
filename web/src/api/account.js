import request from '@/interceptors/marketing'

export function accountUpdate(data) {
  return request({
    url: '/marketing/account/update',
    method: 'post',
    data
  })
}

export function accountCreate(data) {
  return request({
    url: '/marketing/account/create',
    method: 'post',
    data
  })
}

export function accountList(data) {
  return request({
    url: '/marketing/account/list',
    method: 'post',
    data
  })
}

export function accountInfo(account_id) {
  return request({
    url: '/marketing/account/info',
    method: 'post',
    data: {id: account_id}
  })
}

export function accountAuth(account_id) {
  return request({
    url: '/marketing/account/auth',
    method: 'post',
    data: {id: account_id}
  })
}
