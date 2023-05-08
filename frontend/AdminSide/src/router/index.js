import { createRouter, createWebHashHistory } from "vue-router";
import NProgress from "nprogress";
import { routes } from "./modules/route";
import Cookies from "js-cookie";
import {
  setToken,
  removeToken,
  setRefreshToken,
  removeRefreshToken,
} from "@/utils/cookie";
import http from "@/http";
import store from "@/store";

const constantRoute = [
  ...routes,
  {
    name: "404",
    path: "/404",
    component: () => import("@/pages/404/index.vue"),
    meta: {
      title: "404页面未找到",
      auth: false,
      tagsView: false,
      tagsViewAffix: false,
    },
  },
  {
    path: "/:patchAll(\\S+)",
    redirect: "/404",
  },
];

const router = createRouter({
  routes: constantRoute,
  history: createWebHashHistory(),
});

const whiteList = [
  "/login",
  "/transferOpenTrns",
  "/transferOpenTrnsFilter",
  "apiCallLists",
];

router.beforeEach((to, from, next) => {
  NProgress.start();

  if (to.meta && to.meta.title) {
    document.title = to.meta.title;
  }
  let accessToken =
    typeof Cookies.get("timer_access") != "undefined"
      ? Cookies.get("timer_access")
      : false;

  if (accessToken) {
    if (whiteList.includes(to.path)) {
      next(from.path);
      NProgress.done();
      return;
    }
  }
  if (
    !accessToken &&
    to.meta.title !== "login" &&
    to.meta.title !== "transferOpenTrns" &&
    to.meta.title !== "transferOpenTrnsFilter" &&
    to.meta.title !== "apiCallLists"
  ) {
    removeToken();
    removeRefreshToken();
    Cookies.remove("userInfo");
    router.push({ name: "login", query: { redirect: to.path } });
    return;
  } else {
    next();
  }
});

router.afterEach(() => {
  NProgress.done();
});

export function createRoute() {}

createRoute();

export function resetRouter() {
  const newRouter = createRouter({
    routes: constantRoute,
    history: createWebHashHistory(),
  });
  router.matcher = newRouter.matcher;
}

export default router;
