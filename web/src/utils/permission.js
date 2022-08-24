import store from "@/store";

/**
 * @param {Array} value
 * @returns {Boolean}
 * @example see @v/permission/directive.vue
 */
export function checkPermission(value) {
  if (value && Array.isArray(value) && value.length > 0) {
    const permissions = store.getters && store.getters.permissions;
    const permissionRoles = value;

    const hasPermission = permissions.some((role) => {
      return permissionRoles.includes(role);
    });
    return hasPermission;
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`);
    return false;
  }
}

// 检查用户是否包含超管角色
export function rolesHasSuperAdmin(roles) {
  if (Array.isArray(roles)) {
    if (roles.length === 0) {
      return false;
    }
    for (let i = 0; i < roles.length; i++) {
      if (roles[i].id === 1) {
        return true;
      }
    }
  } else if (Object.prototype.toString.call(roles) === "[Object Object]") {
    if (roles.id === 1) {
      return true;
    }
  }
  return false;
}
