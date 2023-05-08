import Cookies from "js-cookie";

const TokenKey = "access_token";
const RefreshKey = "refresh_token";

export function getToken() {
  return Cookies.get(TokenKey);
}

export function setToken(token) {
  return Cookies.set(TokenKey, "Bearer " + token, { expires: 2 });
}

export function removeToken() {
  return Cookies.remove(TokenKey);
}

export function getRefreshToken() {
  return Cookies.get(RefreshKey);
}

export function setRefreshToken(token) {
  return Cookies.set(RefreshKey, token);
}

export function removeRefreshToken() {
  return Cookies.remove(RefreshKey);
}
