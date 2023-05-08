import { checkPermissions } from "@/utils/permissions";

export const admin = [
  {
    id: 94,
    name: "Admin User",
    icon: "fa-solid fa-users",
    url: "/auth/admin",
    type: "page",
    permission: true,
    perName: "auth",
  },
];
