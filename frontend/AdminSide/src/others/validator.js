import { Message } from "element-ui";

export const validPhone = (phone, errMsg = "") => {
  const phoneRule =
    /^(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-7|9])|(?:5[0-3|5-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[1|8|9]))\d{8}$/;
  let status = phoneRule.test(phone);
  if (!status) {
    if (errMsg) Message({ type: "error", message: errMsg });
  }

  return status;
};

export const validEmail = (email, errMsg = "") => {
  const emailRule =
    // eslint-disable-next-line no-useless-escape
    /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
  let status = emailRule.test(email);
  if (!status) {
    if (errMsg) Message({ type: "error", message: errMsg });
  }

  return status;
};

export const validNotEmpty = (content, errMsg = "") => {
  if (content == null) {
    if (errMsg) Message({ type: "error", message: errMsg });
    return false;
  }
  const notEmptyRule = /\S/;
  let status = notEmptyRule.test(content);
  if (!status) {
    if (errMsg) Message({ type: "error", message: errMsg });
  }

  return status;
};

export const validArrayNotEmpty = (contentList = [], errMsg = "") => {
  let status = contentList.length == 0 ? false : true;
  if (!status) {
    if (errMsg) Message({ type: "error", message: errMsg });
  }

  return status;
};
