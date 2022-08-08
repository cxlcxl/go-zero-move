import axios from 'axios'
import { requestConfigs, requestError, responseError, responseSuccessful } from '@/interceptors'

const appService = axios.create({
  baseURL: process.env.VUE_APP_APPLICATION_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 10000 // request timeout
})

// request interceptor
appService.interceptors.request.use(
  requestConfigs, requestError
)

// response interceptor
appService.interceptors.response.use(
  responseSuccessful, responseError
)

export default appService
