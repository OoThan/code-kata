export const loanManagement = [
  {
    name: "loanPackages",
    path: "/loanManagement/loanPackages",
    component: () => import("@/pages/loanManagement/loanPackage/index.vue"),
    meta: {
      title: "Loan Packages",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
  {
    name: "loanPackageLog",
    path: "/loanManagement/loanPackageLog",
    component: () => import("@/pages/loanManagement/loanPackageLog/index.vue"),
    meta: {
      title: "Loan Package Logs",
      auth: true,
      layout: "admin",
      tagsView: true,
    },
  },
];
