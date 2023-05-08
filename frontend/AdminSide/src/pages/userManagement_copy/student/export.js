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
  data.forEach(({ uuid, name, email, role_no, year, CreatedAt }) => {
    exportData.push({
      uuid,
      name,
      email,
      role_no,
      year: filterYear(year),
      CreatedAt: dateFormat(CreatedAt),
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
    const tHeader = [
      "ID",
      "Username",
      "Email",
      "Role Number",
      "Year",
      "Register Time",
    ];
    const filterVal = ["uuid", "name", "email", "role_no", "year", "CreatedAt"];

    const data = formatJson(filterVal, exportData);
    excel.export_json_to_excel({
      header: tHeader,
      data,
      filename: "Student" + date,
      autoWidth: true,
      bookType: "xlsx",
    });
    ElMessage.success("Export Success");
  });
};

const filterYear = (val) => {
  switch (val) {
    case 1:
      return "First Year";
    case 2:
      return "Second Year";
    case 3:
      return "Third Year";
    case 4:
      return "Fourth Year";
    case 5:
      return "Final Year";

    default:
      return "--";
  }
};

export default ExportData2Excel;
