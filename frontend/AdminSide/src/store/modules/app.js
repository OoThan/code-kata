import logo from "../../assets/images/logo.png";
import { getClientRect } from "@/others/utils";
import { sidebarItem } from "../menu/index";
import Cookies from "js-cookie";

const getCachePageIndicator = () => {
  let val = localStorage.getItem("pageIndicator");
  return val;
};
const getCacheMenuCollapse = () => {
  let val = localStorage.getItem("menuCollapse");
  return val;
};
const getCachePageKeepAlive = () => {
  let val = localStorage.getItem("getCachePageKeepAlive");
  return val;
};

export default {
  namespaced: true,
  state: () => ({
    appName: `${import.meta.env.VITE_APP_NAME}`,

    appLogo: logo,

    device: "desktop",

    pageIndicator: getCachePageIndicator() || "面包屑",

    menuCollapse: getCacheMenuCollapse() || 0,

    pageKeepAlive: getCachePageKeepAlive() || 0,

    windowRect: {
      clientHeight: getClientRect().clientHeight,
      clientWidth: getClientRect().clientWidth,
    },
  }),

  getters: {
    userInfo() {
      let userInfo = Cookies.get("userInfo");
      return userInfo;
    },
    menu() {
      let menu = [...sidebarItem];
      return menu;
    },
  },

  mutations: {
    SET_DEVICE(state, device) {
      state.device = device;
      if (device === "mobile") this.commit("app/TOGGLE_MENU_COLLAPSE", 1);
      if (device === "desktop") this.commit("app/TOGGLE_MENU_COLLAPSE", 0);
    },

    TOGGLE_MENU_COLLAPSE(state, collapse) {
      if (state.menuCollapse === collapse) return;
      state.menuCollapse = collapse;
      localStorage.setItem("menuCollapse", state.menuCollapse);
    },

    SET_MENU_TAG_SWITCHER(state, val = "面包屑") {
      state.pageIndicator = val;
      localStorage.setItem("pageIndicator", state.pageIndicator);
    },

    SET_PAGE_KEEP_ALIVE(state, val = 0) {
      state.pageKeepAlive = val;
      localStorage.setItem("pageKeepAlive", state.pageKeepAlive);
    },

    SET_WINDOW_RECT(state, val) {
      state.windowRect = {
        clientHeight: val.clientHeight || 0,
        clientWidth: val.clientWidth || 0,
      };
    },
  },

  actions: {},
};
