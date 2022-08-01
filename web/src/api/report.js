import request from '../utils/request'

export function rechargeReport(data) {
  return request({ url: '/report/recharge', method: 'post', data })
}

export function consumeReport(data) {
  return request({ url: '/report/consume', method: 'post', data })
}

export function consumeRebate(data) {
  return request({ url: '/report/consume-rebate', method: 'post', data })
}

export function refreshConsumeRebate(data) {
  return request({ url: '/report/rebate-refresh', method: 'post', data })
}

export function incomeProfit(data) {
  return request({ url: '/report/income-profit', method: 'post', data })
}

export function incomeRecharge(data) {
  return request({ url: '/report/income-recharge', method: 'post', data })
}

export function incomeTransfer(data) {
  return request({ url: '/report/income-transfer', method: 'post', data })
}

export function archiveQuarters() {
  return request({ url: '/report/archive-qs', method: 'get' })
}

export function saveArchive(data) {
  return request({ url: '/report/archive', method: 'post', data })
}

export function refreshManagementReport() {
  return request({ url: '/report/profit-refresh', method: 'post' })
}

export function refreshTransferReport() {
  return request({ url: '/report/transfer-refresh', method: 'post' })
}

export function dashboardProfit() {
  return request({method: 'get', url: '/dashboard/profit'})
}

export function profit(data) {
  return request({method: 'post', url: '/report/profit', data})
}
