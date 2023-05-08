export const report = [
  {
    name: "popularBooks",
    path: "/report/popularBooks",
    component: () => import("@/pages/report/popularBook/index.vue"),
    meta: {
      title: "Popular Books",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
  {
    name: "booksSummary",
    path: "/report/booksSummary",
    component: () => import("@/pages/report/bookSummary/index.vue"),
    meta: {
      title: "Books Summary",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
