import axios from 'axios'
import { requestConfigs, requestError, responseError, responseSuccessful } from '@/interceptors'

const marketingService = axios.create({
  baseURL: process.env.VUE_APP_MARKETING_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 10000 // request timeout
})

// request interceptor
marketingService.interceptors.request.use(
  requestConfigs, requestError
)

// response interceptor
marketingService.interceptors.response.use(
  responseSuccessful, responseError
)

export default marketingService
