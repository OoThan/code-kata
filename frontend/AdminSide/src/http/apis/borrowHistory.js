import http from "../axios";

/**
 * @Borrow_History
 */

export const getBorrowHistory = (data) =>
  http({ url: "/borrow/get-history", method: "GET", params: data });

export const addBorrow = (data) =>
  http({
    url: "/borrow/create",
    method: "POST",
    data,
  });

export const changeStatus = (data) =>
  http({
    url: "/borrow/update-status",
    method: "POST",
    data,
  });

export const reBorrow = (data) =>
  http({
    url: "/borrow/re-borrow",
    method: "POST",
    data,
  });
