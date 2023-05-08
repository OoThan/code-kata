<template>
  <div v-loading="isLoading">
    <div class="form" v-show="isShowSearch">
      <div class="form-content">
        <el-form label-position="top" label-width="300px">
          <el-form-item label="ID">
            <el-input placeholder="ID" v-model="filterForm.uuid" />
          </el-form-item>

          <el-form-item label="Title">
            <el-input placeholder="Title" v-model="filterForm.title" />
          </el-form-item>

          <el-form-item label="Category">
            <el-select
              :multiple="false"
              placeholder="Category"
              style="width: 100%"
              v-model="filterForm.category_id"
              default-first-option
            >
              <el-option
                v-for="item in cetegoryLists"
                :key="item"
                :value="item.id"
                :label="item.title"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="Status">
            <el-select
              :multiple="false"
              placeholder="Status"
              style="width: 100%"
              v-model="filterForm.status"
              default-first-option
            >
              <el-option
                v-for="item in statusOptions"
                :key="item"
                :value="item.id"
                :label="item.title"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="Author">
            <el-input placeholder="Author" v-model="filterForm.author" />
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
        <el-table-column prop="uuid" label="ID" align="center" width="110" />
        <el-table-column
          prop="title"
          label="Title"
          align="center"
          width="250"
        />
        <el-table-column
          prop="author"
          label="Author"
          align="center"
          width="250"
        />
        <el-table-column label="Category" align="center" width="250">
          <template #default="scope">
            {{ filterCategory(scope.row.sub_category_id) }}
          </template>
        </el-table-column>

        <el-table-column
          prop="available_qty"
          label="Total Book"
          align="center"
          width="130"
        />

        <el-table-column label="Available Qty" align="center" width="130">
          <template #default="scope">
            {{ scope.row.available_qty - scope.row.borrow_qty }}
          </template>
        </el-table-column>
        <el-table-column
          prop="borrow_qty"
          label="Borrowing Qty"
          align="center"
          width="130"
        />

        <el-table-column label="Image" align="center" width="130">
          <template #default="scope">
            <el-image
              style="width: 35px; height: 35px"
              :src="scope.row.book_image"
            />
          </template>
        </el-table-column>

        <el-table-column
          prop="status"
          label="Status"
          align="center"
          width="150"
        >
          <template #default="scope">
            <!-- {{ filterStatus(scope.row.status) }} -->

            <el-tag
              class="ml-2"
              type="success"
              v-if="scope.row.status == 1"
              effect="dark"
              >Available</el-tag
            >
            <el-tag
              class="ml-2"
              type="danger"
              v-if="scope.row.status == 2"
              effect="dark"
              >Damage or Lost</el-tag
            >

            <el-tag
              class="ml-2"
              type="warning"
              v-if="scope.row.status == 3"
              effect="dark"
              >All borrow</el-tag
            >
          </template>
        </el-table-column>

        <el-table-column label="Download Link" align="center" width="180">
          <template #default="scope">
            {{ scope.row.download_link ? scope.row.download_link : "--" }}
          </template>
        </el-table-column>

        <el-table-column label="Publish Date" align="center" width="180">
          <template #default="scope">
            {{ scope.row.publish_date || "--" }}
          </template>
        </el-table-column>

        <el-table-column label="Create Time" align="center" width="180">
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
    :categoryLists="cetegoryLists"
    :statusOptions="statusOptions"
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
  name: "Book",
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
      cate_param: {
        page: 1,
        page_size: 500,
      },
      dialog: {
        dialogTitle: "",
        dialogData: {},
      },
      imagesSource: "/images/logo.png",
      isShowSearch: false,
      tableLists: [],
      cetegoryLists: [],
      total: 0,

      statusOptions: [
        {
          id: 1,
          title: "Available",
        },
        {
          id: 2,
          title: "Lost or Damage",
        },
      ],
    });

    const store = useStore();

    const filterStatus = (val) => {
      switch (val) {
        case 1:
          return "Available";
        case 2:
          return "Lost or Damage";
        default:
          return "--";
      }
    };

    const getTableLists = () => {
      state.isLoading = true;
      http.bookManagement.getBookList(state.param).then((res) => {
        if (res.data.err_code == 0) {
          console.log(res.data.data.list);
          state.tableLists = res.data.data.list;
          state.total = res.data.data.total;
          state.isLoading = false;
        }
      });
    };

    const getCategoryLists = () => {
      state.isLoading = true;
      http.bookManagement.getSubCategoryList(state.cate_param).then((res) => {
        if (res.data.err_code == 0) {
          console.log(res.data.data.list);
          state.cetegoryLists = res.data.data.list.map((item) => {
            return {
              id: item.id,
              title: item.sub_category_name,
            };
          });
        }
      });
    };

    const filterCategory = (val) => {
      if (state.cetegoryLists) {
        let myCategory = state.cetegoryLists.filter((item) => {
          return item.id == val;
        });

        return myCategory[0]?.title;
      }
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
        state.param.uuid = state.filterForm.uuid;
      }
      if (state.filterForm.title) {
        state.param.title = state.filterForm.title;
      }
      if (state.filterForm.category_id) {
        state.param.category_id = parseInt(state.filterForm.category_id);
      }
      if (state.filterForm.status) {
        state.param.status = parseInt(state.filterForm.status);
      }
      if (state.filterForm.author) {
        state.param.author = state.filterForm.author;
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
          http.bookManagement.deleteBookList({ id: id }).then((res) => {
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
      http.bookManagement.getBookList(state.param).then((res) => {
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
      getCategoryLists();
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
      searchShow,
      t,
      filterCategory,
      filterStatus,
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
