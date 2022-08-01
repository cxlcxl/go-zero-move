import request from '@/utils/request'

export const excelTemplate = process.env.VUE_APP_BASE_API + '/excel/template/'

export function getImportData(params) {
  return request({
    url: '/excel/list',
    type: 'get',
    params
  })
}

export function dataDestroy(id) {
  return request({
    url: '/excel/destroy',
    method: 'post',
    data: {id}
  })
}

export function flushImportData(module_name) {
  return request({
    url: '/excel/flush',
    method: 'post',
    data: {module_name}
  })
}

export function saveImportData(module_name) {
  return request({
    url: '/excel/save',
    method: 'post',
    data: {module_name}
  })
}
