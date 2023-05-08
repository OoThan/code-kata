import { reactive } from "vue";
import { getType } from "@/others/utils";
import { ElMessage } from "element-plus";
import { useRoute } from "vue-router";

export default function useTableData(api) {
  if (typeof api !== "function" && getType(api) !== "Promise") {
    throw new Error("请传入Promise类型的api请求");
  }

  const PAGE_SIZES = [10, 20, 50, 100, 200, 500, 1000];

  const DEFAULT_TABLE_DATA = {
    list: [],
    total: 0,
    isLoading: false,
    page: 1,
    page_size: PAGE_SIZES[0],
  };

  let tableData = reactive({
    list: DEFAULT_TABLE_DATA.list,
    total: DEFAULT_TABLE_DATA.total,
  });

  let tableParams = {
    page: DEFAULT_TABLE_DATA.page,
    page_size: DEFAULT_TABLE_DATA.page_size,
  };

  // let tableParams = {};

  // const route = useRoute();
  // if (localStorage.getItem("visitedRoutes")) {
  //   let visitedRoutes = JSON.parse(localStorage.getItem("visitedRoutes"));
  //   let path = visitedRoutes.find((item) => item.path === route.path);
  //   if (path.query.page) {
  //     tableParams = path.query;
  //   } else {
  //     tableParams = {
  //       page: DEFAULT_TABLE_DATA.page,
  //       page_size: DEFAULT_TABLE_DATA.page_size,
  //     };
  //   }
  // }

  const fetchData = async () => {
    tableData.isLoading = true;
    try {
      let res = await api(tableParams);
      if (res.data.err_code == 0) {
        tableData.isLoading = false;

        tableData.list = res.data.data;
        tableData.total = res.data.data.total;
      } else {
        tableData.isLoading = false;
        ElMessage.error(res.data.err_msg);
      }
    } catch (e) {
      tableData.isLoading = false;
      ElMessage.error(e);
    }
  };

  fetchData();

  const reset = () => {
    tableData.list = DEFAULT_TABLE_DATA.list;
    tableData.total = DEFAULT_TABLE_DATA.total;
    tableData.isLoading = DEFAULT_TABLE_DATA.isLoading;
    tableParams = {
      page: 1,
      page_size: 10,
    };
    fetchData();
  };

  const refresh = () => {
    fetchData();
  };

  const setParams = (params = {}) => {
    tableParams = { ...params };
    fetchData();
  };

  return {
    PAGE_SIZES,
    tableData,
    tableParams,
    setParams,
    reset,
    refresh,
  };
}
