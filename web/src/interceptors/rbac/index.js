import axios from 'axios'
import {requestConfigs, requestError, responseError, responseSuccessful} from "@/interceptors";

const rbacService = axios.create({
  baseURL: process.env.VUE_APP_RBAC_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 10000 // request timeout
})

// request interceptor
rbacService.interceptors.request.use(
  requestConfigs, requestError
)

// response interceptor
rbacService.interceptors.response.use(
  responseSuccessful, responseError
)

export default rbacService
