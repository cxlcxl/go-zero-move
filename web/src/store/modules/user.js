import { login, logout, getInfo } from "@a/user";
import { getToken, setToken, removeToken } from "@/utils/auth";
import router, { resetRouter } from "@/router";

const state = {
  token: getToken(),
  user_id: 0,
  name: "",
  email: "",
  avatar: "",
  introduction: "",
  roles: [],
  permissions: [],
};

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token;
  },
  SET_LOGIN_INFO: (state, data) => {
    const { id, roles, username, avatar, email, permissions } = data;
    state.user_id = id;
    state.roles = roles;
    state.name = username;
    state.avatar = avatar;
    state.email = email;
    state.permissions = permissions;
  },
};

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { email, pass } = userInfo;
    return new Promise((resolve, reject) => {
      login({ email: email.trim(), pass })
        .then((response) => {
          const { data } = response;
          commit("SET_TOKEN", data.token);
          setToken(data.token);
          resolve();
        })
        .catch((error) => {
          reject(error);
        });
    });
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(state.token)
        .then((response) => {
          const { data } = response;

          if (!data) {
            reject("获取用户信息失败，请重新登陆.");
          }
          data.roles = [data.role];
          commit("SET_LOGIN_INFO", data);
          resolve(data);
        })
        .catch((error) => {
          reject(error);
        });
    });
  },

  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      logout()
        .then(() => {
          commit("SET_TOKEN", "");
          commit("SET_LOGIN_INFO", {
            id: 0,
            roles: [],
            permissions: [],
            username: "",
            email: "",
            avatar: "",
          });
          removeToken();
          resetRouter();

          // reset visited views and cached views
          // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
          dispatch("tagsView/delAllViews", null, { root: true });

          resolve();
        })
        .catch((error) => {
          reject(error);
        });
    });
  },

  // remove token
  resetToken({ commit }) {
    return new Promise((resolve) => {
      commit("SET_TOKEN", "");
      commit("SET_LOGIN_INFO", {
        id: 0,
        roles: [],
        permissions: [],
        username: "",
        email: "",
        avatar: "",
      });
      removeToken();
      resolve();
    });
  },

  // dynamically modify permissions
  async changeRoles({ commit, dispatch }, role) {
    const token = role + "-token";

    commit("SET_TOKEN", token);
    setToken(token);

    const { roles } = await dispatch("getInfo");

    resetRouter();

    // generate accessible routes map based on roles
    const accessRoutes = await dispatch("permission/generateRoutes", roles, {
      root: true,
    });
    // dynamically add accessible routes
    router.addRoutes(accessRoutes);

    // reset visited views and cached views
    dispatch("tagsView/delAllViews", null, { root: true });
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
};
