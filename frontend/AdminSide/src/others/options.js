import { arrayToObject } from "@/others/utils.js";

export const userTypeOptions = [
  {
    id: 1,
    name: "普通人员",
  },
  {
    id: 2,
    name: "管理员",
  },
];

export const userTypeOptionsMap = arrayToObject(userTypeOptions, {
  keyName: "id",
});

export const sexOptions = [
  {
    id: 1,
    name: "男",
  },
  {
    id: 2,
    name: "女",
  },
];

export const sexOptionsMap = arrayToObject(sexOptions, {
  keyName: "id",
});
