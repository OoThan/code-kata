<template>
  <AdminLayout
    ref="adminLayout"
    v-if="$route.meta.layout && $route.meta.layout == 'admin'"
  >
    <router-view class="router-view" v-slot="{ Component }">
      <keep-alive v-if="pageKeepAlive == 1">
        <component :is="Component" />
      </keep-alive>
      <component :is="Component" v-else />
    </router-view>
  </AdminLayout>
  <router-view class="router-view" v-else></router-view>
</template>

<script>
import { useStore } from "vuex";
import {
  defineComponent,
  computed,
  watch,
  onMounted,
  reactive,
  toRefs,
  ref,
} from "vue";
import http from "@/http";

import useWindowSize from "@/hooks/useWindowSize";

export default defineComponent({
  name: "App",
  setup() {
    const store = useStore();

    const { clientWidth, clientHeight } = useWindowSize();
    clientWidth.value <= 750 && store.commit("app/SET_DEVICE", "mobile");
    clientWidth.value > 750 && store.commit("app/SET_DEVICE", "desktop");
    watch([clientWidth, clientHeight], (newVal) => {
      newVal[0] <= 750 && store.commit("app/SET_DEVICE", "mobile");
      newVal[0] > 750 && store.commit("app/SET_DEVICE", "desktop");
      store.commit("app/SET_WINDOW_RECT", {
        clientWidth: newVal[0],
        clientHeight: newVal[1],
      });
    });

    onMounted(() => {
      localStorage.setItem("pageIndicator", "标签切换"); //for breadcrumb
      store.state.app.pageIndicator = "标签切换";
    });
    return {
      pageKeepAlive: computed(() => store.state.app.pageKeepAlive),
    };
  },
});
</script>
<style>
@import url("https://fonts.googleapis.com/css2?family=Quicksand:wght@500;600;700&display=swap");
* {
  font-family: "Quicksand", sans-serif !important;
}
</style>
