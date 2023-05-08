import { createI18n } from "vue-i18n";
import store from "../store";
import messages from "./getLang";
import ElementPlus from "element-plus";
import locale_cn from "element-plus/lib/locale/lang/zh-cn";
import locale_en from "element-plus/lib/locale/lang/en";

const lang = store.state.user.lang;
const i18n = createI18n({
  legacy: false,
  locale: lang,
  fallbackLocale: "en",
  messages: {
    en: { ...messages.en },
    cn: { ...messages.cn },
  },
});
// Locals(i18n.global.fallbackLocale.value);
let changedLang;

export function changeLang(changesLan) {
  i18n.global.locale.value = changesLan;
  changedLang = changesLan;
  console.log("here change in usei18", changedLang);
  store.dispatch("user/setLang", changesLan);
  // setLangElementPlus(i18n.global.locale.value);
}

export function setLangElementPlus(app) {
  let locale;
  switch (i18n.global.locale.value) {
    case "en":
      locale = locale_en;
      break;
    case "cn":
      locale = locale_cn;
      break;
    default:
      locale = locale_cn;
  }
  app.use(ElementPlus, { locale });
}

export default i18n;
