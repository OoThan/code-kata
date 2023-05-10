import { checkPermissions } from "@/utils/permissions";
export const loanManagement = [
  {
    id: 500,
    name: "Loan Management",
    icon: "fa-solid fa-file-invoice",
    url: "/loanManagement",
    type: "group",
    permission: true,
    perName: "report",
    children: [
      {
        id: 101,
        name: "Loan Packages",
        icon: "fa-regular fa-circle",
        type: "page",
        permission: true,
        perName: "loanPackages",
        url: "/loanManagement/loanPackages",
      },
      {
        id: 102,
        name: "Loan Package Logs",
        icon: "fa-regular fa-circle",
        type: "page",
        permission: true,
        perName: "loanPackageLog",
        url: "/loanManagement/loanPackageLog",
      },
    ],
  },
];
