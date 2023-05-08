import http from "@/http";
import {
  getToken,
  setToken,
  removeToken,
  getRefreshToken,
  setRefreshToken,
  removeRefreshToken,
} from "@/utils/cookie";
import Cookies from "js-cookie";
import { ElMessage } from "element-plus";

const getDefaultState = () => {
  return {
    token: getToken(),
    name: "",
    avatar: "",
    lang: localStorage.getItem("lang") || "en",
  };
};

const state = getDefaultState();

const getters = {
  token: (state) => state.token,
};

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState());
  },
  SET_TOKEN: (state, token) => {
    state.token = token;
  },
  SET_NAME: (state, name) => {
    state.name = name;
  },
  SET_LANG(state, data) {
    state.lang = data;
    localStorage.setItem("lang", data);
  },
};

const actions = {
  // user logout
  logout({ commit }) {
    commit("SET_TOKEN", "");
    commit("SET_NAME", "");
    removeToken();
    //Cookies.remove("userInfo");
    Cookies.remove("timer_access");
    localStorage.removeItem("permissions");
  },

  refreshToken({ commit }) {
    return new Promise((resolve, reject) => {
      http.user
        .refreshToken({ refresh_token: Cookies.get("refresh_token") })
        .then((response) => {
          const { data } = response;
          // console.log("Here is call refresh token api", data.data.access_token);
          setToken(data.data.access_token);
          setRefreshToken(data.data.refresh_token);

          Cookies.set("userInfo", Cookies.get("userInfo"));

          Cookies.set("timer_access", data.data.access_token, {
            expires: 0.003472222,
          });
          //5minutes

          Cookies.set("timer_refresh", data.data.refresh_token, {
            expires: 0.0416667,
            // expires: 3,
          });
          //1hour
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
      removeToken(); // must remove  token  first
      commit("RESET_STATE");
      resolve();
    });
  },

  setLang({ commit }, data) {
    commit("SET_LANG", data);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
