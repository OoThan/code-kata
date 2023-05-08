<template>
  <div v-loading="isLoading">
    <div class="form" v-show="isShowSearch">
      <div class="form-content">
        <el-form label-position="top" label-width="300px">
          <el-form-item label="Id">
            <el-input placeholder="id" v-model="filterForm.uuid" />
          </el-form-item>

          <el-form-item label="Username">
            <el-input placeholder="Username" v-model="filterForm.name" />
          </el-form-item>

          <el-form-item label="Email">
            <el-input placeholder="Email" v-model="filterForm.email" />
          </el-form-item>

          <el-form-item label="Role number">
            <el-input placeholder="Role numbeer" v-model="filterForm.role_no" />
          </el-form-item>

          <el-form-item label="Year">
            <el-select
              :multiple="false"
              placeholder="Year"
              style="width: 100%"
              v-model="filterForm.year"
              default-first-option
            >
              <el-option
                v-for="item in yearLists"
                :key="item"
                :value="item.id"
                :label="item.title"
              />
            </el-select>
          </el-form-item>

          <div style="margin-top: 34px" class="buttonBox">
            <el-button class="app-button" @click="search()">
              <font-awesome-icon
                icon="fa-solid fa-magnifying-glass"
                style="margin-right: 5px"
              />Search</el-button
            >
          </div>
        </el-form>
      </div>
    </div>

    <div class="table">
      <div class="top-panel">
        <div class="left">
          <el-button
            @click="addNew"
            class="app-button"
            style="margin-bottom: 25px"
            ><font-awesome-icon
              icon="fa-solid fa-plus"
              style="margin-right: 5px"
            />
            Add</el-button
          >
        </div>
        <div class="right">
          <el-button
            @click="handlerExport"
            style="margin-bottom: 25px"
            type="warning"
          >
            <font-awesome-icon icon="fa-solid fa-download" />
          </el-button>
          <el-button
            @click="searchShow"
            style="margin-bottom: 25px"
            type="info"
          >
            <font-awesome-icon icon="fa-solid fa-magnifying-glass" />
          </el-button>
        </div>
      </div>

      <el-table
        :data="tableLists"
        :row-style="{ height: '55px' }"
        :header-row-style="{ height: '55px' }"
        :border="true"
        height="65vh"
        style="width: 100%"
      >
        <el-table-column prop="uuid" label="ID" align="center" width="100px" />
        <el-table-column
          prop="name"
          label="Username"
          align="center"
          width="300px"
        />
        <el-table-column
          prop="email"
          label="Email"
          align="center"
          width="300px"
        />
        <el-table-column
          prop="role_no"
          label="Role no"
          align="center"
          width="300px"
        />

        <el-table-column label="Year" align="center" width="200px">
          <template #default="scope">
            {{ filterYear(scope.row.year) }}
          </template>
        </el-table-column>

        <el-table-column label="Register Time" align="center" width="250px">
          <template #default="scope">
            {{ dateFormat(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column
          label="Operate"
          align="center"
          width="120"
          fixed="right"
        >
          <template #default="scope">
            <el-tooltip
              class="box-item"
              content="Edit"
              placement="top"
              effect="customized"
            >
              <el-button
                type="primary"
                style="margin-left: 10px; margin-bottom: 5px"
                circle
                @click="handleEdit(scope.row)"
              >
                <font-awesome-icon icon="fa-solid fa-file-pen" />
              </el-button>
            </el-tooltip>

            <el-tooltip
              class="box-item"
              content="Delete"
              placement="top"
              effect="customized"
            >
              <el-button
                type="danger"
                circle
                style="margin-bottom: 5px"
                @click="deleteHandler(scope.row.id)"
              >
                <font-awesome-icon icon="fa-solid fa-trash" />
              </el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <div class="table-pager">
        <el-pagination
          @size-change="onSizeChange"
          @current-change="onCurrentChange"
          :page-size="param.page_size"
          :page-sizes="pageList"
          :current-page="param.page"
          :layout="`total,prev,pager,next,jumper,${
            device === 'mobile' ? '' : 'sizes'
          }`"
          :total="total"
          :small="device === 'mobile' ? true : false"
        >
        </el-pagination>
      </div>
    </div>
  </div>

  <Dialog
    :show="showDialog"
    @closed="closeDialog"
    @created="search"
    :title="dialog.dialogTitle"
    :data="dialog.dialogData"
    :roleList="roleList"
    :yearLists="yearLists"
  />
</template>

<script>
import { onMounted, reactive, toRefs, computed } from "vue";
import Dialog from "./dialog.vue";
import http from "@/http";
import { useStore } from "vuex";
import useTableData from "@/hooks/useTableData.js";
import { dateFormat } from "@/utils/timeFormat.js";
import { ElMessageBox, ElMessage } from "element-plus";
import { useI18n } from "vue-i18n";
import ExportData2Excel from "./export.js";

export default {
  name: "Admin",
  components: {
    Dialog,
  },
  setup() {
    const { t } = useI18n();

    const state = reactive({
      isLoading: true,
      showDialog: false,
      filterForm: {},
      pageList: [10, 20, 60, 80, 100],
      param: {
        page: 1,
        page_size: 10,
      },
      dialog: {
        dialogTitle: "",
        dialogData: {},
      },
      isShowSearch: false,
      tableLists: [],
      roleList: [],
      total: 0,
      yearLists: [
        {
          id: 1,
          title: "First Year",
        },
        {
          id: 2,
          title: "Second year",
        },
        {
          id: 3,
          title: "Third Year",
        },
        {
          id: 4,
          title: "Fourth Year",
        },
        {
          id: 5,
          title: "Final Year",
        },
      ],
    });

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

    const store = useStore();

    const getTableLists = () => {
      state.isLoading = true;
      http.userManagement.getStudents(state.param).then((res) => {
        if (res.data.err_code == 0) {
          console.log(res.data.data.list);
          state.tableLists = res.data.data.list;
          state.total = res.data.data.total;
          state.roleList = res.data.data.roles;
          state.isLoading = false;
        }
      });
    };

    const addNew = () => {
      state.dialog.dialogTitle = "Add";
      state.dialog.dialogData = {};
      state.showDialog = true;
    };

    const handleEdit = (row) => {
      state.dialog.dialogTitle = "Edit";
      state.dialog.dialogData = row;
      state.showDialog = true;
    };

    const closeDialog = () => {
      state.showDialog = false;
    };

    const filter = () => {
      state.param = {
        page: 1,
      };

      if (state.filterForm.uuid) {
        state.param.uuid = parseInt(state.filterForm.uuid);
      }

      if (state.filterForm.name) {
        state.param.name = state.filterForm.name;
      }

      if (state.filterForm.email) {
        state.param.email = state.filterForm.email;
      }

      if (state.filterForm.role_no) {
        state.param.role_no = state.filterForm.role_no;
      }

      if (state.filterForm.year) {
        state.param.year = parseInt(state.filterForm.year);
      }
    };

    const onSizeChange = (val) => {
      state.param.page_size = val;
      getTableLists();
    };
    const onCurrentChange = (val) => {
      state.param.page = val;
      getTableLists();
    };

    const deleteHandler = (id) => {
      ElMessageBox.confirm("Are sure want to delete?", "Warning", {
        confirmButtonText: "Sure",
        cancelButtonText: "cancel",
        type: "warning",
        draggable: true,
      })
        .then(() => {
          http.auth.deleteStudent({ id: id }).then((res) => {
            if (res.data.err_code == 0) {
              ElMessage.success(res.data.err_msg);
              search();
            } else {
              ElMessage.error(res.data.err_msg);
            }
          });
        })
        .catch(() => {
          ElMessage({
            type: "info",
            message: "You have canceled",
          });
        });
    };

    const search = () => {
      filter();
      state.param.page_size = 10;
      getTableLists();
    };

    const searchShow = () => {
      if (state.isShowSearch) {
        state.isShowSearch = false;
      } else {
        state.isShowSearch = true;
      }
    };

    const handlerExport = () => {
      state.isLoading = true;
      let data = [];

      state.param.page_size = state.total;
      http.userManagement.getStudents(state.param).then((res) => {
        if (res.data.err_code === 0) {
          data = res.data.data.list;
          exporttttt(data);
        }
      });
    };

    const exporttttt = (data) => {
      ExportData2Excel(data)
        .then((res) => {
          state.param.page_size = 10;
          search();
          state.isLoading = false;
        })
        .catch((err) => {
          state.param.page_size = 10;
          search();
          state.isLoading = false;
        });
    };

    onMounted(() => {
      getTableLists();
    });

    return {
      ...toRefs(state),
      addNew,
      handleEdit,
      closeDialog,
      dateFormat,
      search,
      onSizeChange,
      onCurrentChange,
      deleteHandler,
      windowRect: computed(() => store.state.app.windowRect),
      device: computed(() => store.state.app.device),
      // changeStatus,
      searchShow,
      t,
      filterYear,
      handlerExport,
    };
  },
};
</script>

<style lang="scss" scoped>
.top-panel {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
