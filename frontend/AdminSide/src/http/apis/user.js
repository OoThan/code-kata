import http from "../axios";

export const getConfig = (data) =>
  http({
    url: "/system-configs/check2fa",
    method: "POST",
    data,
  });

export const login = (data) =>
  http({
    url: "/auth/login",
    method: "POST",
    data,
  });

export const logout = (data) =>
  http({
    url: "/auth/logout",
    method: "POST",
    data,
  });

export const getExchangeList = (data) =>
  http({
    url: "/finance/exchanges",
    method: "POST",
    data,
  });

export const getDashboard = (data) =>
  http({
    url: "/admin-users/dashboard",
    method: "GET",
    params: data,
  });

export const activate2fa = (data) =>
  http({
    url: "/admin/activate2fa",
    method: "POST",
    data,
  });

export const getPermissions = (data) =>
  http({
    url: "/permission/list",
    method: "POST",
    data,
  });

export const refreshToken = (data) =>
  http({
    url: "/admin/refresh",
    method: "POST",
    data,
  });

// export const register = (data) =>
// 	http({
// 		url: '/register',
// 		method: 'POST',
// 		data
// 	});

export const appUpload = (data, progress) =>
  http({
    url: "media/upload",
    method: "POST",
    data,
    onUploadProgress: progress,
  });
