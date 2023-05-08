<template>
  <el-menu
    unique-opened
    mode="vertical"
    active-text-color="#000"
    :collapse="sideMenuCollapse"
    class="admin-aside-menu"
    text-color="#000"
    collapse-transition
    :default-active="$route.path"
    @select="onMenuSelect"
  >
    <MenuItem :menuList="menu"></MenuItem>
  </el-menu>
</template>

<script>
/**
 * 侧边菜单导航栏
 *
 */
import MenuItem from "./MenuItem.vue";
import { defineComponent } from "vue";
export default defineComponent({
  name: "AdminSideBar",
  props: {
    sideMenuCollapse: {
      type: Boolean,
      default: false,
    },
    menu: {
      type: Array,
      default: () => [],
    },
  },
  components: {
    MenuItem,
  },
  setup(props, context) {
    const onMenuSelect = (index) => {
      context.emit("menuSelect", index);
    };
    return {
      onMenuSelect,
    };
  },
});
</script>

<style lang="scss">
.admin-aside-menu {
  height: 100%;
  background-color: $sidebarBgColor;
  font-size: 14px;
  width: 100%;
  overflow: auto;
  user-select: none;
  overflow-x: hidden;
  //border-right: 2px solid #999;

  &:not(.el-menu--collapse) {
    width: $sideMenuWidth;
  }

  .el-sub-menu__title {
    color: $white !important;
  }

  .el-sub-menu__title:hover {
    color: $white !important;
    background: $sidbarHoverBg;
  }
  .el-sub-menu__title,
  .el-menu-item {
    text-align: left !important;
    transition: all 0.05s;
    color: $white !important;
    background-color: $sidbarItemBg;
    line-height: 50px;
    height: 50px;

    svg {
      color: $title !important;
      margin-right: 0 !important;
      font-size: 15px;
    }

    span {
      color: $title !important;
      margin-left: 10px;
    }

    &:hover {
      background: $sidbarHoverBg !important;
      transition: background-color 0.2s;

      span {
        color: $white !important;
      }

      svg {
        color: $white !important;
      }
    }

    &.is-active {
      font-weight: bold;
      background: $sidbarActiveBg !important;
      box-sizing: border-box;
      span {
        color: $white !important;
      }

      svg {
        color: $white !important;
      }
    }
  }
  .el-menu-item {
    //background: $sidbarHoverBg;
  }

  .el-submenu__title {
    span {
      margin-left: 8px;
    }
  }
}

/**折叠*/
.el-menu--popup {
  .el-menu-item {
    text-align: left !important;
    transition: all 0.1s;
    svg {
      color: #000 !important;
      color: $sidebarBgColor !important;
      padding-right: 2px;
    }

    span {
      color: #000 !important;
      color: $sidebarBgColor !important;
      margin-left: 10px;
    }

    &:hover,
    &.is-active {
      background: $sidebarBgColor !important;

      span {
        color: $white !important;
      }

      svg {
        color: $white !important;
      }
    }
  }
}
</style>
