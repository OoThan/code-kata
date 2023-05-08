import { ElMessage } from "element-plus";
import { dateFormat } from "@/utils/timeFormat.js";

const formatJson = (filterVal, jsonData) => {
  return jsonData.map((v) => filterVal.map((j) => v[j]));
};

let exportData = [];

const ExportData2Excel = (val) => {
  return new Promise((resolve, reject) => {
    if (val) {
      resolve(export2Excel(val));
    } else {
      reject("Sometimg wrong");
    }
  });
};

const export2Excel = (val) => {
  let data = [];
  data = val;
  data.forEach(({ id, title, sort, content }) => {
    exportData.push({
      id,
      title,
      sort,
      content,
    });
  });
  excelExport();
};

const excelExport = () => {
  import("@/utils/Export2Excel").then((excel) => {
    const date =
      new Date().getFullYear() +
      "-" +
      (new Date().getMonth() + 1) +
      "-" +
      new Date().getDate();
    const tHeader = ["ID", "标题", "排序", "内容"];
    const filterVal = ["id", "title", "sort", "content"];

    const data = formatJson(filterVal, exportData);
    excel.export_json_to_excel({
      header: tHeader,
      data,
      filename: "文章管理" + date,
      autoWidth: true,
      bookType: "xlsx",
    });
    ElMessage.success("Export Success");
  });
};

export default ExportData2Excel;
