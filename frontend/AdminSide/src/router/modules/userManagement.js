export const userManagement = [
  {
    name: "user",
    path: "/auth/user",
    component: () => import("@/pages/userManagement/index.vue"),
    meta: {
      title: "User Management",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
