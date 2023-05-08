import http from "../axios";

/**
 * @Config
 */

export const getConfigData = (data) =>
  http({ url: "system-config/get", method: "GET", params: data });

export const updateConfigData = (data) =>
  http({
    url: "/system-config/update",
    method: "POST",
    data,
  });

export const getPunishment = (data) =>
  http({ url: "punishment/get", method: "GET", params: data });

export const addPunishment = (data) =>
  http({
    url: "/punishment/create",
    method: "POST",
    data,
  });

export const updatePunishment = (data) =>
  http({
    url: "/punishment/update",
    method: "POST",
    data,
  });
