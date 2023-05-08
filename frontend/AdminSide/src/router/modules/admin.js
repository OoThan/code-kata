export const admin = [
  {
    name: "adminUser",
    path: "/auth/admin",
    component: () => import("@/pages/auth/admin/index.vue"),
    meta: {
      title: "Admin User",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
