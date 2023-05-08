export const bookManagement = [
  {
    name: "categories",
    path: "/book/categories",
    component: () => import("@/pages/bookManagement/category/index.vue"),
    meta: {
      title: "Categories",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
  {
    name: "subCategories",
    path: "/book/subCategories",
    component: () => import("@/pages/bookManagement/subCategory/index.vue"),
    meta: {
      title: "Sub Categories",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
  {
    name: "books",
    path: "/book/books",
    component: () => import("@/pages/bookManagement/book/index.vue"),
    meta: {
      title: "Books",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
