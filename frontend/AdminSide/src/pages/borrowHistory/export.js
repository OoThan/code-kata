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
      id,
      status,
      type,
      user_data,
      book_data,
      created_at,
      expired_at,
      expired_day,
      punishment_amt,
    }) => {
      exportData.push({
        id: id,
        status: filterStatus(status),
        type: filterType(type),
        user_uuid: user_data.uuid,
        user_name: user_data.name,
        email: user_data.email,
        book_uuid: book_data.uuid,
        book_title: book_data.title,
        department:
          type == 1
            ? filterTeacherDepartment(user_data.department)
            : filterStaffDepartment(user_data.department),

        year: filterYear(user_data.year),
        role_no: user_data.role_no,
        borrow_time: dateFormat(created_at),
        expire_time: dateFormat(expired_at),
        expired_day: expired_day,
        punish_amount: punishment_amt,
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
      "Borrow Status",
      "Borrow Type",
      "User ID",
      "Username",
      "Email",
      "Book ID",
      "Book Title",
      "Department",
      "Year",
      "Role no",
      "Borrow Time",
      "Expire Time",
      "Expired Day",
      "Punish Amount",
    ];
    const filterVal = [
      "id",
      "status",
      "type",
      "user_uuid",
      "user_name",
      "email",
      "book_uuid",
      "book_title",
      "department",
      "year",
      "role_no",
      "borrow_time",
      "expire_time",
      "expired_day",
      "punish_amount",
    ];

    const data = formatJson(filterVal, exportData);
    excel.export_json_to_excel({
      header: tHeader,
      data,
      filename: "Books Borrow History" + date,
      autoWidth: true,
      bookType: "xlsx",
    });
    ElMessage.success("Export Success");
  });
};

const filterType = (val) => {
  switch (val) {
    case 1:
      return "Teacher";
    case 2:
      return "Student";
    case 3:
      return "Staff";

    default:
      return "--";
  }
};

const filterStatus = (val) => {
  switch (val) {
    case 1:
      return "Borrowing";
    case 2:
      return "Returned";
    case 3:
      return "Expired";

    default:
      return "--";
  }
};

const filterTeacherDepartment = (val) => {
  switch (val) {
    case 1:
      return "Faculty of Computer Systems & Technologies";
    case 2:
      return "Faculty of Computer Science";
    case 3:
      return "Faculty of Information Science";
    case 4:
      return "Department of Information Technology Supporting and Maintenance";
    case 5:
      return "Department of Physics";
    case 6:
      return "Department of Admistration and Finance";

    default:
      return "--";
  }
};

const filterStaffDepartment = (val) => {
  switch (val) {
    case 1:
      return "Others...";

    default:
      return "--";
  }
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
