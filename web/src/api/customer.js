import request from '../utils/request'

export function customerList(params) {
  return request({url: '/customer/list', method: 'get', params})
}

export function customerCreate(data) {
  return request({url: '/customer/create', method: 'post', data})
}

export function customerUpdate(data) {
  return request({url: '/customer/update', method: 'post', data})
}

export function searchCustomer(customer_name) {
  return request({url: '/customer/search', method: 'get', params: {customer_name}})
}

export function hasManyContract(customer_id) {
  return request({url: '/customer/contract', method: 'get', params: {customer_id}})
}

export function hasManyAccount(customer_id) {
  return request({url: '/customer/account/' + customer_id, method: 'get'})
}

export function getPlatformCustomers() {
  return request({url: '/customer/platform', method: 'get'})
}

export function getCmCustomers() {
  return request({url: '/customer/cm', method: 'get'})
}

export function getDefaultCustomers() {
  return request({url: '/customer/default', method: 'get'})
}

export function accountCreate(data) {
  return request({url: '/account/create', method: 'post', data})
}

export function getCustomerTypes() {
  return request({url: '/customer/types', method: 'get'})
}

export function getAllCustomer() {
  return request({url: '/customer/all', method: 'get'})
}

export function flowBelongsAccount(params) {
  return request({url: '/customer/flow-account', method: 'get', params})
}

export function customerBalance(params) {
  return request({url: '/customer/balance', method: 'get', params})
}

