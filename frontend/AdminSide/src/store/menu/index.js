import { admin } from "./modules/admin";
import { userManagement } from "./modules/userManagement";
import { loanManagement } from "./modules/loanManagement";

export const sidebarItem = [
  {
    id: 89,
    name: "Dashboard",
    icon: "fa-solid fa-gauge-high",
    type: "page",
    permission: true,
    url: "/home",
    path: [],
    tagsViewAffix: true,
    children: [
      {
        id: 17,
        name: "list",
      },
    ],
  },
  ...admin,
  ...loanManagement,
  ...userManagement,
];

localStorage.setItem("menu", JSON.stringify(sidebarItem));
