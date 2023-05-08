import { checkPermissions } from "@/utils/permissions";

export const config = [
  {
    id: 500,
    name: "Config",
    icon: "fa-solid fa-gears",
    url: "/config",
    type: "group",
    permission: true,
    perName: "book",
    children: [
      {
        id: 501,
        name: "System Configs",
        icon: "fa-solid fa-gear",
        type: "page",
        permission: true,
        perName: "systemConfigs",
        url: "/config/systemConfigs",
      },
      {
        id: 502,
        name: "Manage Punishment",
        icon: "fa-solid fa-gavel",
        type: "page",
        permission: true,
        perName: "books",
        url: "/cofig/managementPunishment",
      },
    ],
  },
];
