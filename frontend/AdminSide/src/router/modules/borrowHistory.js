export const borrowHistory = [
  {
    name: "borrowHistory",
    path: "/borrow/borrowHistory",
    component: () => import("@/pages/borrowHistory/index.vue"),
    meta: {
      title: "Borrow History",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
