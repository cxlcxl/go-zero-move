import request from '@/interceptors/account'

export function accountUpdate(data) {
  return request({
    url: '/account/update',
    method: 'post',
    data
  })
}

export function changeContract(data) {
  return request({
    url: '/account/change-contract',
    method: 'post',
    data
  })
}

export function accountCreate(data) {
  return request({
    url: '/account/create',
    method: 'post',
    data
  })
}

export function accountList(data) {
  return request({
    url: '/account/list',
    method: 'post',
    data
  })
}

export function accountInfo(account_id) {
  return request({
    url: '/account/info',
    method: 'post',
    params: {id: account_id}
  })
}

export function getAccounts(account_id) {
  return request({
    url: '/account-balances/' + account_id,
    method: 'post'
  })
}

export function accountDestroy(account_id) {
  return request({
    url: '/account/destroy',
    method: 'post',
    data: {
      id: account_id
    }
  })
}
