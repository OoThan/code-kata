<template>
  <div class="layout">
    <AppHeader v-if="device == 'desktop'"></AppHeader>
    <MobileAppHeader v-if="device == 'mobile'"></MobileAppHeader>

    <div class="layout-main">
      <aside>
        <AppSideBar
          :menu="sideMenu"
          :sideMenuCollapse="sideMenuCollapse == 1"
        ></AppSideBar>
      </aside>
      <div class="app-main">
        <div class="app-main-top">
          <AppBreadcrumb
            class="breadcrumb-container"
            v-if="pageSwither === '面包屑'"
          ></AppBreadcrumb>
          <AppTagsViewSwitcher
            v-if="pageSwither === '标签切换'"
          ></AppTagsViewSwitcher>
        </div>

        <main class="app-main-content">
          <el-main>
            <layout />
          </el-main>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, computed, onMounted, inject } from "vue";
import { useStore } from "vuex";
import { reactive } from "vue";
import { useRouter } from "vue-router";
import Layout from "./Layout.vue";

export default defineComponent({
  components: { Layout },
  name: "AdminLayout",
  setup() {
    const store = useStore();

    const router = useRouter();

    return {
      sideMenuCollapse: computed(() => store.state.app.menuCollapse),
      pageSwither: computed(() => store.state.app.pageIndicator),
      sideMenu: computed(() => store.getters["app/menu"]),
      device: computed(() => store.state.app.device),
    };
  },
});
</script>

<style lang="scss">
.layout {
  width: 100%;
  // height: 100%;

  .layout-main {
    display: flex;
    background: $app_background;

    aside {
      height: calc(100vh - #{$appHeaderHeight});
      @media screen and (max-width: 750px) {
        display: none;
      }
    }
  }
  .app-main {
    width: 100%;
    height: 100%;
    box-sizing: border-box;
    overflow: hidden;

    .app-main-top {
      display: flex;
      align-items: center;
      padding: 0 10px;
      flex-shrink: 0;
      height: $breadcrumbHeight;
      background: transparent;
      background-color: #fff;

      // border-bottom: 0.5px solid rgba(#000, 0.05);
    }
    .app-main-content {
      width: 100%;
      height: $adminContentHeight;
      overflow-y: auto;
      box-sizing: border-box;
      overflow-x: hidden;
    }
  }
}
.el-table .cell {
  overflow: unset !important;
}
</style>
