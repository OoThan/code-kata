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
  data.forEach(
    ({
      uuid,
      title,
      author,
      available_qty,
      borrow_qty,
      status,
      download_link,
      publish_date,
      CreatedAt,
    }) => {
      exportData.push({
        uuid,
        title,
        author,
        available_qty,
        borrow_qty,
        status: filterStatus(status),
        download_link,
        publish_date,
        CreatedAt: dateFormat(CreatedAt),
      });
    }
  );
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
    const tHeader = [
      "ID",
      "Title",
      "Author",
      "Total Book",
      "Borrowing Qty",
      "Status",
      "Download Link",
      "Publish Date",
      "Create Time",
    ];
    const filterVal = [
      "uuid",
      "title",
      "author",
      "available_qty",
      "borrow_qty",
      "status",
      "download_link",
      "publish_date",
      "CreatedAt",
    ];

    const data = formatJson(filterVal, exportData);
    excel.export_json_to_excel({
      header: tHeader,
      data,
      filename: "Books" + date,
      autoWidth: true,
      bookType: "xlsx",
    });
    ElMessage.success("Export Success");
  });
};

const filterStatus = (val) => {
  switch (val) {
    case 1:
      return "Available";
    case 2:
      return "Lost or Damage";
    case 3:
      return "All borrow";

    default:
      return "--";
  }
};

export default ExportData2Excel;
