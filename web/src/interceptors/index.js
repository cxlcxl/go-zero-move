import { Message } from '_element-ui@2.13.2@element-ui'
import store from '@/store'
import { getToken, removeToken } from '@/utils/auth'

export function responseSuccessful(response) {
  const res = response.data
  if (res.code !== 0) {
    Message({
      message: res.msg || 'Error',
      type: 'error',
      duration: 5 * 1000
    })

    return Promise.reject(new Error(res.msg || 'Error'))
  } else {
    return res
  }
}

export function responseError(error) {
  const err = error.response
  if (err === undefined) {
    //removeToken()
    Message({
      message: '登陆信息已过期，请刷新页面重新登录',
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(err)
  }
  const errMsg = err.data
    ? '状态码：' + err.status + ' - ' + err.data
    : '接口调用异常，请重试'
  Message({
    message: errMsg,
    type: 'error',
    duration: 5 * 1000
  })
  return Promise.reject(err)
}

export function requestConfigs(config) {
  if (store.getters.token) {
    config.headers['Authorization'] = 'Bearer ' + getToken()
  }
  return config
}

export function requestError(error) {
  return Promise.reject(error)
}
