import {asyncRoutes, constantRoutes} from '@/router'
import {rolesHasSuperAdmin} from '@/utils/permission'

/**
 * Use meta.role to determine if the current user has permission
 * @param permissions
 * @param route
 */
function hasPermission(permissions, route) {
  if (route.meta && route.meta.auth) {
    let auth = "/api/" + route.meta.auth
    return permissions.some(p => auth.includes(p))
  } else {
    return true
  }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param permissions
 */
export function filterAsyncRoutes(routes, permissions) {
  const res = []
  routes.forEach(route => {
    const tmp = {...route}
    if (hasPermission(permissions, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, permissions)
      }
      if (tmp.children && tmp.children.length === 0 && tmp.redirect) {
        // 没有子菜单不写入，JS 不能 continue
      } else {
        res.push(tmp)
      }
    }
  })

  return res
}

const state = {
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  }
}

const actions = {
  generateRoutes({commit}, auths) {
    return new Promise(resolve => {
      const {roles, permissions} = auths
      let accessedRoutes
      // if (rolesHasSuperAdmin(roles)) {
        accessedRoutes = asyncRoutes || []
      // } else {
      //   accessedRoutes = filterAsyncRoutes(asyncRoutes, permissions)
      // }

      commit('SET_ROUTES', accessedRoutes)
      resolve(accessedRoutes)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
