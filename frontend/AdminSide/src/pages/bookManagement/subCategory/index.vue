<template>
  <div v-loading="isLoading">
    <div class="form" v-show="isShowSearch">
      <div class="form-content">
        <el-form label-position="top" label-width="300px">
          <el-form-item label="Id">
            <el-input placeholder="" v-model="filterForm.id" />
          </el-form-item>

          <el-form-item label="Sub category name">
            <el-input placeholder="" v-model="filterForm.sub_category_name" />
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
        <el-table-column prop="id" label="ID" align="center" width="120" />

        <el-table-column
          prop="sub_category_name"
          label="Sub Category Name"
          align="center"
        />

        <el-table-column label="Category" align="center">
          <template #default="scope">
            {{ filterCategory(scope.row.category_id) }}
          </template>
        </el-table-column>

        <!-- <el-table-column label="Description" align="center">
          <template #default="scope">
            {{ scope.row.desc }}
          </template>
        </el-table-column> -->
        <!-- <el-table-column prop="desc" label="Description" align="center" /> -->

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

            <!-- <el-tooltip
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
            </el-tooltip> -->
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
  />
</template>
k
<script>
import { onMounted, reactive, toRefs, computed } from "vue";
import Dialog from "./dialog.vue";
import http from "@/http";
import { useStore } from "vuex";
import useTableData from "@/hooks/useTableData.js";
import { dateFormat } from "@/utils/timeFormat.js";
import { ElMessageBox, ElMessage } from "element-plus";
import { useI18n } from "vue-i18n";

export default {
  name: "Category",
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
      cate_param: {
        page: 1,
        page_size: 500,
      },
      isShowSearch: false,
      tableLists: [],
      total: 0,
      cetegoryLists: [],
    });

    const store = useStore();

    const getTableLists = () => {
      state.isLoading = true;
      http.bookManagement.getSubCategoryList(state.param).then((res) => {
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
      http.bookManagement.getCategoryList(state.cate_param).then((res) => {
        if (res.data.err_code == 0) {
          // console.log(res.data.data.list);
          state.cetegoryLists = res.data.data.list.map((item) => {
            return {
              id: item.id,
              title: item.category_name,
            };
          });

          console.log("Here is category lists", state.cetegoryLists);
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

      if (state.filterForm.id) {
        state.param.id = parseInt(state.filterForm.id);
      }
      if (state.filterForm.sub_category_name) {
        state.param.sub_category_name = state.filterForm.sub_category_name;
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
          http.bookManagement.deleteCategoryList({ id: id }).then((res) => {
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

    onMounted(() => {
      getCategoryLists();
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
      searchShow,
      t,
      filterCategory,
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
