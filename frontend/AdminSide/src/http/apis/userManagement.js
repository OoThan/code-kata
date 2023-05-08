import http from "../axios";

export const getUserList = (data) =>
  http({
    url: "/user/list",
    method: "POST",
    data,
  });

export const addUser = (data) =>
  http({
    url: "/user/add",
    method: "POST",
    data,
  });

export const editUser = (data) =>
  http({
    url: "/user/edit",
    method: "POST",
    data,
  });

export const deleteUser = (data) =>
  http({
    url: "/user/delete",
    method: "POST",
    data,
  });
