import http from "../axios";

export const getBankList = (data) =>
  http({ url: "/api/admin/payment/list", method: "GET", params: data });

export const createBankList = (data) =>
  http({
    url: "/admin/payment/create",
    method: "POST",
    data,
  });

export const updateBankList = (data) =>
  http({
    url: `/admin/payment/update`,
    method: "POST",
    data,
  });

export const deleteBankList = (data) =>
  http({
    url: `/admin/payment/delete`,
    method: "POST",
    data,
  });

export const getUserBankList = (data) =>
  http({ url: "/api/admin/user/bank/list", method: "GET", params: data });

export const updateUserBankList = (id, data) =>
  http({
    url: `/admin/user/bank/update`,
    method: "POST",
    data,
  });

export const deleteUserBankList = (data) =>
  http({
    url: `/admin/user/bank/delete`,
    method: "POST",
    data,
  });
