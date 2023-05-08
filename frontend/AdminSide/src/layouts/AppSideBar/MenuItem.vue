<template>
  <template v-for="item in menuList">
    <el-sub-menu
      v-if="item.type == 'group' && item.permission"
      :key="item.url"
      :index="item.name"
      style="user-select: none"
    >
      <template v-slot:title>
        <font-awesome-icon :icon="item.icon" style="margin-right: 10px" />

        <span>{{ item.name }}</span>
      </template>

      <MenuItem :menuList="item.children"></MenuItem>
    </el-sub-menu>

    <el-menu-item
      v-if="item.type == 'page' && item.permission"
      :key="item.id"
      :index="item.url"
      style="user-select: none"
      @click="onSelectItem(item)"
    >
      <template v-if="item.icon">
        <font-awesome-icon :icon="item.icon" />
      </template>

      <template #title>
        <span>{{ item.name }}</span>
      </template>
    </el-menu-item>
  </template>
</template>

<script>
import { defineComponent } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useStore } from "vuex";
import { useI18n } from "vue-i18n";

export default defineComponent({
  name: "MenuItem",
  props: ["menuList"],
  setup() {
    const { t } = useI18n();
    const router = useRouter();
    const route = useRoute();
    const store = useStore();

    //点击菜单
    const onSelectItem = (item) => {
      if (route.path === item.url) return;

      router.push({
        path: item.url,
      });
    };

    return {
      t,
      onSelectItem,
    };
  },
});
</script>

<style scss scoped>
.el-menu-item {
  height: 35px;
}
</style>
