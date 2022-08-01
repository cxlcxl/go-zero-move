import { formatMoney } from '@/utils'

/**
 * month_between 返回两个月份中间的所有月份
 * @param month_range ['2021-01', '2021-12']
 * @returns {Array}
 */
export function month_between(month_range) {
  const result = [
    month_range[0]
  ]
  const s = month_range[0].split('-')
  const e = month_range[1].split('-')
  const min = new Date()
  const max = new Date()
  min.setFullYear(s[0], s[1])
  max.setFullYear(e[0], e[1])
  const curr = min
  let str = ''
  while (curr <= max) {
    const month = curr.getMonth()
    if (month === 0) {
      str = (curr.getFullYear() - 1) + '-' + 12
    } else {
      str = curr.getFullYear() + '-' + (month < 10 ? ('0' + month) : month)
    }
    result.push(str)
    curr.setMonth(month + 1)
  }
  return result
}

/**
 * formatLineChart 充值/消耗 格式化图表数据
 * @param topic
 * @param rows
 * @param month_range
 * @param has_customer
 * @returns {Array}
 */
export function formatLineChart(topic, rows, month_range, has_customer) {
  const customerIds = []; const chartLines = []; const customerMap = {}
  if (rows && Array.isArray(rows)) {
    if (has_customer === 1) {
      rows.map(item => {
        if (!customerIds.includes(item.customer_id)) {
          customerIds.push(item.customer_id)
          customerMap[item.customer_id] = item.customer_name
        }
      })

      for (let i = 0; i < customerIds.length; i++) {
        const data = []
        month_between(month_range).map(dateMonth => {
          let hasData = false
          rows.map(row => {
            if (row.trans_month === dateMonth && row.customer_id === customerIds[i]) {
              data.push(row[topic])
              hasData = true
            }
          })
          if (!hasData) {
            // 无数据填充 0
            data.push(0)
          }
        })

        chartLines.push({ name: customerMap[customerIds[i]], data })
      }
    } else {
      const data = []
      month_between(month_range).map(dateMonth => {
        let hasData = false
        rows.map(row => {
          if (row.trans_month === dateMonth) {
            data.push(row[topic])
            hasData = true
          }
        })
        if (!hasData) {
          // 无数据填充 0
          data.push(0)
        }
      })
      chartLines.push({ name: '', data })
    }
  }

  return chartLines
}

/**
 * formatBarChart 充值/消耗 格式化图表数据
 * @param topic
 * @param rows
 * @param has_customer
 * @returns {Array}
 */
export function formatBarChart(topic, rows, has_customer) {
  const customerIds = []; const chartLines = []; const customerMap = {}
  if (rows && Array.isArray(rows)) {
    if (has_customer === 1) {

    } else {
      for (const t in topic) {
        const data = []
        for (let i = rows.length; i > 0; i--) {
          data.push(rows[i - 1][t])
        }
        chartLines.push({ name: topic[t], data })
      }
    }
  }
  return chartLines
}

export const quarters = [{ value: 1, label: 'Q1' }, { value: 2, label: 'Q2' }, { value: 3, label: 'Q3' }, {value: 4, label: 'Q4'}]

/**
 * 应收应付字段
 * @param date_type
 * @param has_customer
 * @returns {Array}
 */
export function profitColumns(date_type, has_customer) {
  const columns = []
  if (date_type === 1) {
    columns.push({ title: '时间', align: 'center', key: 'date_quarter', minWidth: 90 })
  } else {
    columns.push({ title: '时间', align: 'center', key: 'date_month', minWidth: 90 })
  }
  if (has_customer === 1) {
    columns.push({ title: '客户', align: 'center', key: 'customer_name', minWidth: 110 })
  }
  columns.push({
    title: '应收（垫付未还）', align: 'right', key: 'repaid', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['repaid']))
    }
  })
  columns.push({
    title: '现金账', children: [
      {
        title: '充值', align: 'right', key: 'c_pay_recharge', minWidth: 110, render: (h, { row }) => {
          return h('span', {}, formatMoney(row['c_pay_recharge']))
        }
      }, {
        title: '还款', align: 'right', key: 'c_pay_repayment', minWidth: 110, render: (h, { row }) => {
          return h('span', {}, formatMoney(row['c_pay_repayment']))
        }
      }, {
        title: '服务费', align: 'right', key: 'c_pay_service', minWidth: 110, render: (h, { row }) => {
          return h('span', {}, formatMoney(row['c_pay_service']))
        }
      }, {
        title: '小计', align: 'right', key: 'c_pay_total', minWidth: 110, render: (h, { row }) => {
          return h('span', {}, formatMoney(row['c_pay_total']))
        }
      }
    ]
  })
  columns.push({
    title: '已拨付', align: 'right', key: 'allocate', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['allocate']))
    }
  })
  columns.push({
    title: '未拨付', align: 'right', key: 'no_allocate', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['no_allocate']))
    }
  })
  columns.push({
    title: '消耗', align: 'right', key: 'consume', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['consume']))
    }
  })
  columns.push({
    title: '消耗返', align: 'right', key: 'consume_rebate', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['consume_rebate']))
    }
  })
  /*columns.push({
    title: '充返', align: 'right', key: 'recharge_rebate', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['recharge_rebate']))
    }
  })*/
  columns.push({
    title: '归档', align: 'center', key: 'is_archive', minWidth: 70, render: (h, { row }) => {
      return h('span', { style: { color: row['is_archive'] === 1 ? '#13ce66' : '#606266' }}, (row['is_archive'] === 1 ? '已归档' : '未归档'))
    }
  })
  return columns
}

/**
 * 充值流水收支
 * @param platform_types
 * @param types
 * @returns {Array}
 */
export function rechargeColumns(platform_types, types) {
  const receive = []; const pay = []
  for (const t in platform_types) {
    pay.push({
      title: platform_types[t], align: 'right', key: 'pay_' + t, minWidth: 110, render: (h, { row }) => {
        return h('span', {}, formatMoney(row['pay_' + t]))
      }
    })
  }
  for (const t in types) {
    receive.push({
      title: types[t], align: 'right', key: 'receive_' + t, minWidth: 110, render: (h, { row }) => {
        return h('span', {}, formatMoney(row['receive_' + t]))
      }
    })
  }
  receive.push({
    title: '小计', align: 'right', key: 'receive_sum', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['receive_sum']))
    }
  })
  pay.push({
    title: '小计', align: 'right', key: 'pay_sum', minWidth: 110, render: (h, { row }) => {
      return h('span', {}, formatMoney(row['pay_sum']))
    }
  })
  return [
    { align: 'center', key: 'date_quarter', minWidth: 90, title: '时间' },
    { title: '收入', children: receive },
    { title: '支出', children: pay },
    {
      title: '合计', align: 'right', key: 'profit', minWidth: 110, render: (h, { row }) => {
        return h('span', {}, formatMoney(row['profit']))
      }
    }
  ]
}

/**
 * 转账流水收支
 * @param accounts
 * @returns {Array}
 */
export function transferColumns(accounts) {
  const receive = []; const pay = []; const profit = []
  for (const t in accounts) {
    pay.push({
      title: accounts[t], align: 'right', key: 'pay_' + t, minWidth: 110, render: (h, { row }) => {
        return h('span', {}, formatMoney(row['pay_' + t]))
      }
    })
    profit.push({
      title: accounts[t], align: 'right', key: 'profit_' + t, minWidth: 110, render: (h, { row }) => {
        return h('span', {}, formatMoney(row['profit_' + t]))
      }
    })
    receive.push({
      title: accounts[t], align: 'right', key: 'receive_' + t, minWidth: 110, render: (h, { row }) => {
        return h('span', {}, formatMoney(row['receive_' + t]))
      }
    })
  }
  return [
    { align: 'center', key: 'date_quarter', minWidth: 90, title: '时间' },
    { title: '收入', children: receive },
    { title: '支出', children: pay },
    { title: '汇总', children: profit }
  ]
}

export const managementExportHeaders = {
  'payment': ['待还', '客户打款(充值)', '客户打款(还款)', '客户打款(服务费)', '客户打款小计', '已拨付', '未拨付', '消耗', '消耗返', '充返', '归档'],
  'recharge': ['时间', '现金充值(收)', '预付还款(收)', '服务费(收)', '小计(收)', '现金充值(支)', '预付还款(支)', '服务费(支)', '小计(支)', '总计'],
  'transfer': ['时间', '现金(收)', '授信(收)', '纯赠(收)', '充返(收)', '返利(收)', '现金(支)', '授信(支)', '纯赠(支)', '充返(支)', '返利(支)', '现金(总)', '授信(总)', '纯赠(总)', '充返(总)', '返利(总)']
}
export const managementExportHeaderCodes = {
  payment: [
    { code: 'repaid' }, { code: 'c_pay_recharge' }, { code: 'c_pay_repayment' }, { code: 'c_pay_service' },
    { code: 'c_pay_total' }, { code: 'allocate' }, { code: 'no_allocate' }, { code: 'consume' }, { code: 'consume_rebate' },
    { code: 'recharge_rebate' }, { code: 'is_archive' }
  ],
  recharge: [
    { code: 'date_quarter' }, { code: 'receive_1' }, { code: 'receive_2' }, { code: 'receive_3' }, { code: 'receive_sum' },
    { code: 'pay_1' }, { code: 'pay_2' }, { code: 'pay_3' }, { code: 'pay_sum' }, { code: 'profit' }
  ],
  transfer: [
    { code: 'date_quarter' }, { code: 'receive_CASH' }, { code: 'receive_CREDIT' }, { code: 'receive_GIFT' }, { code: 'receive_REBATE' }, { code: 'receive_URGE' },
    { code: 'pay_CASH' }, { code: 'pay_CREDIT' }, { code: 'pay_GIFT' }, { code: 'pay_REBATE' }, { code: 'pay_URGE' },
    { code: 'profit_CASH' }, { code: 'profit_CREDIT' }, { code: 'profit_GIFT' }, { code: 'profit_REBATE' }, { code: 'profit_URGE' }
  ]
}

// 一到两层，超过两层需要改动table和导出方法
export const profitHeaders = [
  {title: '日期', prop: 'date', minWidth: 90, align: 'center', desc: '', isFormat: false},
  {title: '利润', prop: 'total_profit', minWidth: 110, align: 'right', desc: '现金账利润 + 系统账利润', isFormat: true, highlight: true},
  {
    title: '现金账', prop: 'null', minWidth: 0, align: 'center', desc: '', isFormat: false, children: [
      {title: '现金账利润', prop: 'cash_profit', minWidth: 110, align: 'right', desc: '收支结余 + 应收 - 应付', isFormat: true, highlight: true},
      {title: '收支结余', prop: 'cash_balance', minWidth: 110, align: 'right', desc: '客户打款与转出平台余额', isFormat: true},
      {title: '应收', prop: 'cash_receivable', minWidth: 110, align: 'right', desc: '垫付未还(转账为基准)', isFormat: true},
      {title: '应付', prop: 'cash_payable', minWidth: 110, align: 'right', desc: '授信消耗未还平台', isFormat: true},
    ]
  },
  {
    title: '系统账', prop: 'null', minWidth: 0, align: 'center', desc: '', isFormat: false, children: [
      {title: '系统账利润', prop: 'sys_profit', minWidth: 110, align: 'right', desc: '现金账户结余 + 纯赠账户结余 + 充赠账户结余 + 返利账户结余 - 未拨付 + 应收平台返点 - 应付客户返点', isFormat: true, highlight: true},
      {title: '现金账户结余', prop: 'sys_cash', minWidth: 110, align: 'right', desc: '', isFormat: true},
      {title: '纯赠账户结余', prop: 'sys_gift', minWidth: 110, align: 'right', desc: '', isFormat: true},
      {title: '充赠账户结余', prop: 'sys_recharge_gift', minWidth: 110, align: 'right', desc: '', isFormat: true},
      {title: '返利账户结余', prop: 'sys_rebate', minWidth: 110, align: 'right', desc: '', isFormat: true},
      {title: '授信账户结余', prop: 'sys_credit', minWidth: 110, align: 'right', desc: '不纳入系统账户利润计算', isFormat: true},
      {title: '未拨付', prop: 'no_allocate', minWidth: 110, align: 'right', desc: '客户充值未拨付', isFormat: true},
      {title: '应收平台返点', prop: 'receivable_rebate', minWidth: 110, align: 'right', desc: '系统试算平台消耗返点', isFormat: true},
      {title: '应付客户返点', prop: 'payable_rebate', minWidth: 110, align: 'right', desc: '系统试算客户消耗返点', isFormat: true},
    ]
  },
]
