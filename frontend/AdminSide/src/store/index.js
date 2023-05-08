import { createStore } from "vuex";
import app from "@/store/modules/app.js";
import tagsView from "@/store/modules/tags-view.js";
import user from "@/store/modules/user.js";
import permissions from "@/store/modules/permissions.js";
import { sidebarItem } from "./menu/index.js";
import createPersistedState from "vuex-persistedstate";

export default createStore({
  modules: {
    app,
    tagsView,
    user,
    permissions,
    sidebarItem,
  },
  // plugins: [
  //   createPersistedState({
  //     storage: window.localStorage,
  //     key: "permissionList",
  //     paths: ["permissions"],
  //   }),
  // ],
  devtools: true,
});
