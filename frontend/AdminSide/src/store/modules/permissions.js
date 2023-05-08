import http from "@/http";
import { useStore } from "vuex";
const store = useStore;
const getDefaultState = () => {
  return {
    permissions: JSON.parse(localStorage.getItem("permissionList")),
  };
};

const state = getDefaultState();

const getters = {
  permissions: (state) => state.permissions,
};

const mutations = {
  SET_PERMISSIONS: (state, permissions) => {
    state.permissions = permissions;
    // localStorage.setItem("permissionList", JSON.stringify(permissions));
  },
  SET_SECRETKEY(state, secret_key) {
    state.secret_key = secret_key;
  },
};

const actions = {
  // setPermissions({ commit }) {
  //   return new Promise((resolve, reject) => {
  //     http.user
  //       .getPermissions()
  //       .then((response) => {
  //         const { data } = response;
  //         commit("SET_PERMISSIONS", data.permissions);
  //         location.reload();
  //         resolve();
  //       })
  //       .catch((error) => {
  //         reject(error);
  //       });
  //   });
  // },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
