import dayjs from "dayjs";

export function arrayToObject(
  data = [],
  format = { keyName: "id", valueName: "name" }
) {
  let map = {};
  if (data.length == 0) return;
  if (format && format.keyName && format.valueName) {
    data.forEach((item) => {
      map[item[format.keyName]] = item[format.valueName];
    });
    return map;
  } else if (format && format.keyName) {
    data.forEach((item) => {
      map[item[format.keyName]] = item;
    });
    return map;
  } else {
    console.error("请输入keyName或valueName");
    return;
  }
}

export function timeFormat(time, formatStr = "YYYY-MM-DD HH:mm") {
  if (!time) return "--";
  return time ? dayjs(time).format(formatStr) : "";
}

export function getType(value) {
  let type = Object.prototype.toString.call(value);
  type = type.replace(/^\[object\s/, "");
  type = type.replace(/\]$/, "");
  return type;
}

export function traverseArrayTree(
  data = [],
  childNodeKey = "children",
  callBack
) {
  if (Array.isArray(data)) {
    data.forEach((item) => {
      typeof callBack == "function" && callBack(item);
      if (item[childNodeKey] && Array.isArray(item[childNodeKey])) {
        traverseArrayTree(item[childNodeKey], childNodeKey, callBack);
      }
    });
  } else {
    // eslint-disable-next-line no-undef
    throw new error("请传入数组");
  }
}

export function throttle(fn, wait) {
  var pre = Date.now();
  return function () {
    // eslint-disable-next-line @typescript-eslint/no-this-alias
    var context = this;
    var args = arguments;
    var now = Date.now();
    if (now - pre >= wait) {
      fn.apply(context, args);
      pre = Date.now();
    }
  };
}

export function debounce(fn, delay) {
  var timer = null;
  return function () {
    var context = this;
    var args = arguments;
    clearTimeout(timer);
    timer = setTimeout(function () {
      fn.apply(context, args);
    }, delay);
  };
}

export const getClientRect = () => {
  let currentClientHeight =
    document.documentElement.clientHeight || document.body.clientHeight;
  let currentClientWidth =
    document.documentElement.clientWidth || document.body.clientWidth;

  return {
    clientWidth: currentClientWidth,
    clientHeight: currentClientHeight,
  };
};
