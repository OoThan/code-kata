import { admin } from "./admin";
import { article } from "./aritcle";
import { bookManagement } from "./bookManagement";
import { userManagement } from "./userManagement";
import { borrowHistory } from "./borrowHistory";
import { report } from "./report";
import { config } from "./config";

export const routes = [
  {
    name: "login",
    path: "/login",
    component: () => import("@/pages/login/index.vue"),
    meta: {
      title: "login",
      auth: false,
    },
  },

  {
    name: "Home",
    redirect: "/home",
    path: "/",
    meta: {
      title: "Home",
      auth: true,
      layout: "admin",
      tagsView: false,
    },
  },
  {
    name: "dashboard",
    path: "/home",
    component: () => import("@/pages/home/index.vue"),
    meta: {
      title: "dashboard",
      auth: false,
      layout: "admin",
      tagsView: true,
    },
  },
  ...admin,
  ...article,
  ...config,
  ...bookManagement,
  ...userManagement,
  ...borrowHistory,
  ...report,
];
