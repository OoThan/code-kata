import { checkPermissions } from "@/utils/permissions";

export const userManagement = [
  {
    id: 200,
    name: "User",
    icon: "fa-solid fa-users",
    url: "/user",
    type: "group",
    permission: true,
    perName: "user",
    children: [
      {
        id: 201,
        name: "Teacher",
        icon: "fa-solid fa-user",
        type: "page",
        permission: true,
        perName: "teacher",
        url: "/user/teacher",
      },
      {
        id: 202,
        name: "Student",
        icon: "fa-solid fa-user",
        type: "page",
        permission: true,
        perName: "student",
        url: "/user/student",
      },
      {
        id: 203,
        name: "Staff",
        icon: "fa-solid fa-user",
        type: "page",
        permission: true,
        perName: "staff",
        url: "/user/staff",
      },
    ],
  },
];
