<template>
  <el-breadcrumb separator="/">
    <template v-for="(item, index) in $route.meta.breadcrumb" :key="item">
      <el-breadcrumb-item
        :to="{ path: $route.meta.backlink[index] }"
        v-if="$route.meta.backlink.length > 0"
      >
        <span>{{ t(`bot.${item}`) }}</span>
      </el-breadcrumb-item>
      <el-breadcrumb-item v-else>
        <span>{{ t(`bot.${item}`) }}</span>
      </el-breadcrumb-item>
    </template>
  </el-breadcrumb>
</template>

<script>
import { useI18n } from "vue-i18n";

export default {
  setup() {
    const { t } = useI18n();

    return {
      t,
    };
  },
};
</script>

<style lang="scss" scoped>
.app-breadcrumb {
  display: inline-block;
  font-size: 14px;
  line-height: 30px;
  margin-left: 8px;

  .el-breadcrumb__item {
    span {
      color: rgba(#000, 0.4);
    }

    &:last-child {
      span {
        color: #000;
      }
    }
  }
}
</style>
