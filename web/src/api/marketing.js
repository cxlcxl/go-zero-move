import request from '../interceptors/marketing'

export function getAccessToken(data) {
  return request({url: '/marketing/token', method: 'post', data})
}
