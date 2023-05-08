<template>
  <div v-loading="tableList.tableData.isLoading">
    <div class="form" v-show="isShowSearch">
      <div class="form-content">
        <el-form label-position="top" label-width="300px">
          <el-form-item label="Id">
            <el-input placeholder="Id" v-model="filterForm.id" />
          </el-form-item>

          <el-form-item label="Category name">
            <el-input placeholder="" v-model="filterForm.category_name" />
          </el-form-item>

          <el-form-item :label="t('common.sort')">
            <el-input
              :placeholder="t('common.sort')"
              v-model="filterForm.sort"
            />
          </el-form-item>

          <div style="margin-top: 34px" class="buttonBox">
            <el-button class="app-button" @click="search()">
              <font-awesome-icon
                icon="fa-solid fa-magnifying-glass"
                style="margin-right: 5px"
              />{{ t("common.search") }}</el-button
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
            {{ t("common.addto") }}</el-button
          >
          <el-button
            @click="bulkDelete"
            type="danger"
            style="margin-bottom: 25px"
            :disabled="selectedLists.length == 0"
          >
            <font-awesome-icon
              icon="fa-solid fa-trash"
              style="margin-right: 5px"
            />
            {{ t("common.delete") }}
          </el-button>
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
        :data="tableList.tableData.list.articles"
        :row-style="{ height: '55px' }"
        :header-row-style="{ height: '55px' }"
        :border="true"
        @selection-change="onTableSelectionChange"
        height="70vh"
        style="width: 100%"
      >
        <el-table-column
          type="selection"
          show-overflow-tooltip
          header-align="center"
          align="center"
        />
        <el-table-column prop="id" label="ID" align="center" />
        <el-table-column
          prop="title"
          :label="t('common.title')"
          align="center"
        />
        <el-table-column prop="sort" :label="t('common.sort')" align="center" />
        <el-table-column
          :label="t('common.operate')"
          align="center"
          width="120"
        >
          <template #default="scope">
            <el-tooltip
              class="box-item"
              :content="t('common.edit')"
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
              :content="t('common.delete')"
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
          :total="tableList.tableData.total"
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
import ExportData2Excel from "./export.js";
import { useI18n } from "vue-i18n";

export default {
  name: "Article",
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
      selectedExportLists: [],
      selectedLists: [],
    });

    const tableList = useTableData(http.article.getArticle);

    const store = useStore();

    const addNew = () => {
      state.dialog.dialogTitle = t("common.addto");
      state.dialog.dialogData = {};
      state.showDialog = true;
    };

    const handleEdit = (row) => {
      state.dialog.dialogTitle = t("common.edit");
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
      if (state.filterForm.title) {
        state.param.title = state.filterForm.title;
      }
      if (state.filterForm.sort) {
        state.param.sort = state.filterForm.sort;
      }
    };

    const onSizeChange = (val) => {
      state.param.page_size = val;
      tableList.setParams(state.param);
    };
    const onCurrentChange = (val) => {
      state.param.page = val;
      tableList.setParams(state.param);
    };

    const deleteHandler = (id) => {
      ElMessageBox.confirm("确定删除此项？", "温馨提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        draggable: true,
      })
        .then(() => {
          http.article.deleteArticle({ ids: [id] }).then((res) => {
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
            message: "删除取消",
          });
        });
    };

    const search = () => {
      filter();
      state.param.page_size = 10;
      tableList.setParams(state.param);
    };

    const onTableSelectionChange = (val) => {
      state.selectedLists = val.map((item) => item.id);
      state.selectedExportLists = val;
    };

    const bulkDelete = () => {
      ElMessageBox.confirm("确定删除此项？", "温馨提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        draggable: true,
      })
        .then(() => {
          http.article
            .deleteArticle({ ids: state.selectedLists })
            .then((res) => {
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
            message: "删除取消",
          });
        });
    };
    const searchShow = () => {
      if (state.isShowSearch) {
        state.isShowSearch = false;
      } else {
        state.isShowSearch = true;
      }
    };

    const handlerExport = () => {
      tableList.tableData.isLoading = true;
      let data = [];
      if (state.selectedExportLists.length != 0) {
        data = state.selectedExportLists;
        exporttttt(data);
      } else {
        state.param.page_size = tableList.tableData.total;
        http.article.getArticle(state.param).then((res) => {
          if (res.data.err_code === 0) {
            // log
            data = res.data.data.articles;
            exporttttt(data);
          }
        });
      }
    };

    const exporttttt = (data) => {
      ExportData2Excel(data)
        .then((res) => {
          state.param.page_size = 10;
          search();
          tableList.tableData.isLoading = false;
        })
        .catch((err) => {
          state.param.page_size = 10;
          search();
          tableList.tableData.isLoading = false;
        });
    };

    onMounted(() => {});

    return {
      t,
      ...toRefs(state),
      addNew,
      handleEdit,
      closeDialog,
      dateFormat,
      tableList,
      search,
      onSizeChange,
      onCurrentChange,
      deleteHandler,
      windowRect: computed(() => store.state.app.windowRect),
      device: computed(() => store.state.app.device),
      onTableSelectionChange,
      bulkDelete,
      searchShow,
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
