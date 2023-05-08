import { checkPermissions } from "@/utils/permissions";
export const report = [
  {
    id: 500,
    name: "Report",
    icon: "fa-solid fa-file-invoice",
    url: "/report",
    type: "group",
    permission: true,
    perName: "report",
    children: [
      {
        id: 101,
        name: "Popular Books",
        icon: "fa-regular fa-circle",
        type: "page",
        permission: true,
        perName: "popularBooks",
        url: "/report/popularBooks",
      },
      {
        id: 102,
        name: "Books Summary",
        icon: "fa-regular fa-circle",
        type: "page",
        permission: true,
        perName: "booksSummary",
        url: "/report/booksSummary",
      },
    ],
  },
];
