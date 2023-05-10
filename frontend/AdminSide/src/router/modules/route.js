import { admin } from "./admin";
import { userManagement } from "./userManagement";
import { loanManagement } from "./loanManagement";

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
    name: "Dashboard",
    path: "/home",
    component: () => import("@/pages/home/index.vue"),
    meta: {
      title: "Dashboard",
      auth: false,
      layout: "admin",
      tagsView: true,
    },
  },
  ...admin,
  ...userManagement,
  ...loanManagement,
];
