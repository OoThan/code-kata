<template>
  <header class="mobile-admin-header">
    <div class="header--left">
      <img
        src="/right-arrow.png"
        alt="right arrow"
        @click="onShowMobileMenu"
        class="fold-btn"
      />
    </div>
    <div class="header--right">
      <!-- <el-popover placement="bottom-end" :width="100" trigger="click">
        <template #reference>
          <div class="language-change" v-if="language == 'English'">
            <img src="/us.png" alt="" />
            <span>{{ language }}</span>
          </div>

          <div class="language-change" v-else>
            <img src="/china.png" alt="" />
            <span>{{ language }}</span>
          </div>
        </template>
        <ul class="language-child">
          <li>
            <div class="language-change" @click="changeLan('en')">
              <img src="/us.png" alt="" />
              <span>English</span>
            </div>
          </li>
          <li>
            <div class="language-change" @click="changeLan('cn')">
              <img src="/china.png" alt="" />
              <span>Chinese</span>
            </div>
          </li>
        </ul>
      </el-popover> -->

      <font-awesome-icon
        icon="fa-solid fa-arrows-rotate"
        @click="refreshScreen"
        class="refresh"
      />
      <section class="user-info-section">
        <!-- <el-popover
          placement="bottom"
          :width="300"
          trigger="hover"
          class="header-icon"
          style="padding: 0"
        >
          <template #reference>
            <el-badge
              :value="noticeList.length"
              :max="99"
              type="success"
              :hidden="noticeList.length === 0"
              style="margin-right: 20px"
            >
              <i class="el-icon-message-solid" style="font-size: 20px"></i>
            </el-badge>
          </template>
          <el-tabs>
            <el-tab-pane label="未读消息" name="0"></el-tab-pane>
          </el-tabs>
          <div class="notice-popover-panel" v-if="noticeList.length !== 0">
            <div v-for="item in noticeList" :key="item.id" class="notice-item">
              <i class="el-icon-message-solid message-icon"></i>
              <div class="notify-content">
                <h2>{{ item.content }}</h2>
                <p>{{ item.time }}</p>
              </div>
            </div>
          </div>
          <el-empty description="暂无消息" v-else></el-empty>
          <div class="panel-btn-group" v-if="noticeList.length !== 0">
            <el-button type="text" size="mini" @click="onClearNoticeClick"
              >清除消息</el-button
            >
          </div>
        </el-popover> -->
        <el-popover
          placement="bottom-end"
          :width="150"
          trigger="click"
          popper-class="info-popper"
        >
          <template #reference>
            <div class="avatar">
              <img src="/user.png" alt="User" />
              <font-awesome-icon
                icon="fa-solid fa-caret-down"
                style="margin-left: 5px"
              />
            </div>
          </template>
          <div class="userInfo-popover-panel">
            <div class="panel-title" style="text-align: center">
              <el-avatar
                :size="40"
                src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
              />

              <div>{{ userInfo || "--" }}</div>
              <!-- <el-tag size="mini" style="margin-top: 5px">{{
                userInfo.role || "--"
              }}</el-tag> -->
            </div>
            <!-- <div class="info-item">
              <el-button
                type="text"
                size="mini"
                @click="showQr"
                style="color: #e6a23c"
                >绑定谷歌验证</el-button
              >
            </div> -->
            <div class="info-item">
              <el-button type="" @click="onLogoutClick" style="color: red" link>
                <font-awesome-icon
                  icon="fa-solid fa-circle-exclamation"
                  style="margin-right: 10px"
                  class="fa-lg"
                />
                退出登录</el-button
              >
            </div>
          </div>
        </el-popover>

        <el-dialog v-model="qrDialog">
          <el-form>
            <el-form-item label="" align="center">
              <qrcode-vue :value="qr" :size="150" level="H" />
            </el-form-item>
          </el-form>
        </el-dialog>
      </section>
    </div>
    <MobileAppSideBar ref="mobileAppSideBar"></MobileAppSideBar>
  </header>
</template>

<script>
import { useStore } from "vuex";
import { resetRouter } from "@/router";
import { getType, traverseArrayTree } from "@/others/utils.js";
import {
  defineComponent,
  computed,
  ref,
  reactive,
  watch,
  onMounted,
  toRefs,
} from "vue";
import { useRouter } from "vue-router";
import http from "@/http";
import QrcodeVue from "qrcode.vue";
import { changeLang } from "@/locales/i18n";

export default defineComponent({
  name: "AdminHeader",
  components: {
    QrcodeVue,
  },
  setup() {
    const router = useRouter();

    const store = useStore();

    let state = reactive({
      language: store.state.user.lang == "en" ? "English" : "Chinese",
    });

    //通知信息
    const noticeList = ref([
      {
        id: 1,
        content: "这是一条短通知。",
        time: "2021-6-10",
      },
      {
        id: 2,
        content: "这是一条短通知。",
        time: "2021-6-10",
      },
      {
        id: 3,
        content:
          "这是一条长长长长长长长长长长长长长长长长长长长长长长长长通知。",
        time: "2021-6-10",
      },
      {
        id: 4,
        content:
          "这是一条长长长长长长长长长长长长长长长长长长长长长长长长通知。",
        time: "2021-6-10",
      },
      {
        id: 5,
        content:
          "这是一条长长长长长长长长长长长长长长长长长长长长长长长长通知。",
        time: "2021-6-10",
      },
    ]);

    const qr = ref("");
    const qrDialog = ref(false);

    const onClearNoticeClick = () => {
      noticeList.value.splice(0, noticeList.value.length);
    };
    const noticeTabActive = ref("0");
    const onNoticeTabChange = (tab) => {
      noticeTabActive.value = tab.paneName;
    };

    const onLogoutClick = () => {
      http.user.logout().then((res) => {
        if (res.data.err_code == 0) {
          store.commit("tagsView/delAllVisitedRoutes");
          store.dispatch("user/logout").then(() => {
            resetRouter();
            router.push({
              path: "/login",
            });
          });
        }
      });
    };

    const mobileAppSideBar = ref(null);
    const onShowMobileMenu = () => {
      mobileAppSideBar.value.show();
    };

    const showQr = () => {
      qrDialog.value = true;
    };

    const changeLan = (lan) => {
      if (lan == "en") {
        state.language = "English";
        changeLang("en");
      } else {
        state.language = "Chinese";
        changeLang("cn");
      }
    };

    const refreshScreen = () => {
      router.go();
    };

    onMounted(() => {
      // http.user.getPermissions().then((res) => {
      //   if (res.err_code == 0) {
      //     qr.value = `otpauth://totp/${res.data.username}?secret=${res.data.secret_key}`;
      //   }
      // });
    });

    return {
      ...toRefs(state),
      changeLan,
      noticeList,
      onClearNoticeClick,
      onLogoutClick,
      userInfo: computed(() => store.getters["app/userInfo"]),
      noticeTabActive,
      onNoticeTabChange,
      mobileAppSideBar,
      onShowMobileMenu,
      qr,
      qrDialog,
      showQr,

      refreshScreen,
    };
  },
});
</script>

<style lang="scss">
.mobile-admin-header {
  width: 100%;
  height: $appHeaderHeight;
  display: flex;
  align-items: center;
  box-sizing: border-box;
  justify-content: space-between;
  position: relative;
  z-index: 999;
  transition: all 0.5s;
  background: $header;
  color: $header_text;
  font-weight: bold;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.04);

  .header--left {
    display: flex;
    align-items: center;
    flex-shrink: 0;

    .fold-btn {
      margin: 0 20px;
      color: $header_text;
    }
  }

  .header--right {
    display: flex;
    align-items: center;
    .icon-notice {
      margin-right: 30px;
    }

    .search-btn {
      cursor: pointer;
      margin-right: 10px;
      font-weight: bold;
      font-size: 20px;
      padding: 10px;
      box-sizing: border-box;
      &:hover {
        transition: background-color 0.5s;
        background: rgba(#000, 0.05);
        padding: 10px;
        box-sizing: border-box;
        border-radius: 100px;
        color: $primaryColor;
      }
    }

    .user-info-section {
      display: flex;
      align-items: center;
      margin-right: 20px;
      cursor: pointer;
      .username {
        margin-left: 5px;
        font-weight: normal;
        font-size: 13px;
      }
      .user-info {
        margin-left: 10px;
        text-align: left;

        > p {
          margin-bottom: 4px;
          font-size: 14px;
        }
        .tags-section {
          display: flex;
          align-items: center;
        }
      }
    }
  }
}

.notice-popover-panel {
  height: 300px;
  overflow: auto;
  z-index: 99999999999;
  position: relative;

  .notice-item {
    padding: 10px;
    box-sizing: border-box;
    transition: all 0.5s;
    display: flex;
    &:hover {
      background: rgba(#000, 0.05);
    }
    .message-icon {
      color: $primaryColor;
      padding-top: 3px;
      flex-shrink: 0;
    }

    .notify-content {
      margin-left: 10px;
      > h2 {
        color: rgba(0, 0, 0, 0.6);
        white-space: pre-wrap;
      }
      > p {
        margin-top: 5px;
        font-size: 12px;
        color: rgba(8, 2, 2, 0.2);
      }
    }
  }
}

.userInfo-popover-panel {
  overflow: auto;
  z-index: 99999999999;
  position: relative;
  .info-item {
    display: flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    cursor: pointer;
    transition: all 0.5s;
    box-sizing: border-box;
    width: 100%;
    padding: 5px 0;

    &:hover {
      background: rgba(#000, 0.05);
    }
  }

  .panel-title {
    padding: 15px 10px;
    box-sizing: border-box;

    &:hover {
      background: #fff;
    }
  }
}

.info-popper {
  padding: 0 !important;
}

.avatar {
  display: flex;
  align-items: center;
  flex-shrink: 0;
  > img {
    width: 16px;
  }
}

.panel-btn-group {
  text-align: center;
  padding-top: 10px;
}

.header-icon {
  font-weight: bold !important;
  font-size: 22px;
  margin: 0 20px;
  > i {
    font-weight: bold !important;
  }
}

.el-tabs__nav-wrap::after {
  height: 0.5px;
}
.language-change {
  margin-right: 20px;
  display: flex;
  align-items: center;
  justify-content: center;

  img {
    width: 25px;
    margin-right: 8px;
  }
  span {
    font-size: 13px;
  }
}
.language-child {
  li {
    padding: 8px;
    cursor: pointer;
  }
  li:hover {
    background: #f5f7fa;
  }
}
</style>
