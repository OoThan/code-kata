// 会员管理
export const article = [
  {
    name: "Article",
    path: "/article",
    component: () => import("@/pages/article/index.vue"),
    meta: {
      title: "articleManagement",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
