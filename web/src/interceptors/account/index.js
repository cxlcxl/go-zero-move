import axios from 'axios'
import {requestConfigs, requestError, responseError, responseSuccessful} from "@/interceptors"

const accountService = axios.create({
  baseURL: process.env.VUE_APP_ACCOUNT_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 10000 // request timeout
})

// request interceptor
accountService.interceptors.request.use(requestConfigs, requestError)

// response interceptor
accountService.interceptors.response.use(responseSuccessful, responseError)

export default accountService
