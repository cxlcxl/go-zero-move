import request from '../interceptors/account'

export function getAccessToken(data) {
  return request({url: '/marketing/token', method: 'post', data})
}
