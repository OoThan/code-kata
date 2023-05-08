export const userManagement = [
  {
    name: "teacher",
    path: "/user/teacher",
    component: () => import("@/pages/userManagement/teacher/index.vue"),
    meta: {
      title: "Teacher",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
  {
    name: "student",
    path: "/user/student",
    component: () => import("@/pages/userManagement/student/index.vue"),
    meta: {
      title: "Student",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
  {
    name: "staff",
    path: "/user/staff",
    component: () => import("@/pages/userManagement/staff/index.vue"),
    meta: {
      title: "Staff",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
