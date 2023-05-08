<template>
  <div class="container">
    <div class="login-container">
      <div class="login">
        <div class="login-header">
          <!-- <h3>{{ appName }}</h3> -->
        </div>
        <div class="login-wrapper">
          <div>
            <div class="login-body">
              <div class="login-title">
                <h3 class="title">{{ t("common.login") }}</h3>
              </div>
              <el-form
                class="login-form"
                label-position="top"
                :model="form"
                ref="loginFromRef"
              >
                <el-form-item
                  prop="username"
                  label="Email"
                  :rules="[
                    {
                      required: true,
                      message: t('common.require'),
                      trigger: 'blur',
                    },
                  ]"
                >
                  <el-input
                    :placeholder="t('table.userName')"
                    v-model="form.username"
                    type="text"
                  />
                </el-form-item>

                <el-form-item
                  prop="password"
                  label="Password"
                  :rules="[
                    {
                      required: true,
                      message: t('common.require'),
                      trigger: 'blur',
                    },
                  ]"
                >
                  <el-input
                    :placeholder="t('table.password')"
                    v-model="form.password"
                    type="password"
                    show-password
                  ></el-input>
                </el-form-item>

                <el-button
                  class="login-button"
                  type="primary"
                  style="width: 100%"
                  :loading="loading"
                  @click="handleLogin"
                >
                  {{ t("common.login") }}
                </el-button>
              </el-form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, reactive, onMounted } from "vue";
import { useStore } from "vuex";
import { useRouter, useRoute } from "vue-router";

import { ElMessage } from "element-plus";
import http from "@/http";
import { setToken, removeToken, setRefreshToken } from "@/utils/cookie";
import Cookies from "js-cookie";
import QrcodeVue from "qrcode.vue";

import { changeLang } from "@/locales/i18n";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const form = ref({
  username: "",
  password: "",
  googleauthcode: "",
});

const loading = ref(false);
const loginFromRef = ref(null);
const store = useStore();
const router = useRouter();
const route = useRoute();
const query = computed(() => route.query);
const qr = ref("");
const state = reactive({
  googleAuthDialog: false,
  language: store.state.user.lang == "en" ? "English" : "Chinese",
  configData: {},
  isNeedGoogleAuth: false,
});

const validatePassword = () => {
  return (rule, value, callback) => {
    if (value.length < 6) {
      callback(new Error("Password length should have more than 5"));
    } else {
      callback();
    }
  };
};

const appName = computed(() => store.state.app.appName);

const handleLogin = () => {
  loginFromRef.value.validate((valid) => {
    if (!valid) return;
    loading.value = true;
    let param = {
      email_username: form.value.username,
      password: form.value.password,
    };

    http.user
      .login(param)
      .then((res) => {
        if (res.data.err_code == 0) {
          console.log("Here is token", res.data.data.accessToken);
          setToken(res.data.data.accessToken);
          // setRefreshToken(res.data.data.refresh_token);

          Cookies.set("timer_access", res.data.data.accessToken, {
            expires: 2,
          });
          Cookies.set("userInfo", res.data.data.username);

          // Cookies.set("timer_refresh", res.data.data.refresh_token, {
          //   expires: 0.0416667,
          // });
          ElMessage.success(res.data.err_msg);
          loading.value = false;
          router.push(query.value.redirect ? query.value.redirect : "/home");
        } else {
          // ElMessage.error(res.data.err_msg);
          loading.value = false;
          removeToken();
        }
      })
      .catch((err) => {
        loading.value = false;
        removeToken();
      });
  });
};

onMounted(() => {});
</script>
<style lang="scss" scoped>
@import "@/styles/auth.scss";
.qr-con {
  text-align: center;
}
.dialog-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  .google-auth-input {
    width: 300px;
    border: 1px solid #999;
  }
}
</style>
