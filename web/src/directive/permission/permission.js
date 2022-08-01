import store from '@/store'
import {rolesHasSuperAdmin} from "@/utils/permission"

function checkPermission(el, binding) {
  const {value} = binding
  const permissions = store.getters && store.getters.permissions
  const roles = store.getters && store.getters.roles

  if (value) {
    const hasPermission = permissions.includes('/api/' + value)
    if (!rolesHasSuperAdmin(roles) && !hasPermission) {
      el.parentNode && el.parentNode.removeChild(el)
    }
  } else {
    throw new Error(`need permission! Like v-permission="'user/list'"`)
  }
}

export default {
  inserted(el, binding) {
    checkPermission(el, binding)
  },
  update(el, binding) {
    checkPermission(el, binding)
  }
}
