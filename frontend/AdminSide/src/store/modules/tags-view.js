/**
 * tagsView标签切换页面
 *
 */

// visitedRoutes本地缓存（localStorage）
const getCacheVisitedRoutes = () => {
  let visitedRoutes = localStorage.getItem("visitedRoutes");
  return visitedRoutes ? JSON.parse(visitedRoutes) : [];
};
const setCacheVisitedRoutes = (visitedRoutes) => {
  localStorage.setItem("visitedRoutes", JSON.stringify(visitedRoutes));
};

export default {
  namespaced: true,
  state: () => ({
    visitedRoutes: getCacheVisitedRoutes(),
  }),
  mutations: {
    //添加
    ADD_VISITED_ROUTE(state, route) {
      let target = state.visitedRoutes.find((item) => item.path === route.path);
      // let isClick = false;
      if (target) {
        if (route.fullPath !== target.fullPath) Object.assign(target, route);
        return;
      }
      // state.visitedRoutes.push({ ...target, ...route, isClick });
      state.visitedRoutes.push({ ...target, ...route });
      setCacheVisitedRoutes(state.visitedRoutes);
    },
    //删除
    DELETE_VISITED_ROUTE(state, route) {
      state.visitedRoutes.forEach((item, index) => {
        if (item.path === route.path) state.visitedRoutes.splice(index, 1);
      });
      setCacheVisitedRoutes(state.visitedRoutes);
    },
    //删除除了本身以外的
    DELETE_OTHERS_VISITED_ROUTE(state, route) {
      state.visitedRoutes = state.visitedRoutes.filter(
        (item) => item.path == route.path
      );
      setCacheVisitedRoutes(state.visitedRoutes);
    },
    //更新当前
    UPDATE_VISITED_ROUTE(state, route) {
      state.visitedRoutes.forEach((item) => {
        if (item.path === route.path) {
          item = Object.assign(item, route);
        }
        setCacheVisitedRoutes(state.visitedRoutes);
      });
    },

    UPDATE_PARAM_ROUTE(state, route) {
      state.visitedRoutes.forEach((item) => {
        if (item.path === route.path) {
          item.default_params = route.default_params;
          setCacheVisitedRoutes(state.visitedRoutes);
        }
      });
    },

    delAllVisitedRoutes(state) {
      state.visitedRoutes = [];
      localStorage.removeItem("visitedRoutes");
    },
  },
};
