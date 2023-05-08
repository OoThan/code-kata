export const config = [
  {
    name: "systemConfigs",
    path: "/config/systemConfigs",
    component: () => import("@/pages/config/systemConfig/index.vue"),
    meta: {
      title: "System Configs",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
  {
    name: "managementPunishment",
    path: "/cofig/managementPunishment",
    component: () => import("@/pages/config/managePunishment/index.vue"),
    meta: {
      title: "Punishments",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
