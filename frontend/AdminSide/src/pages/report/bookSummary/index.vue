<template>
  <div v-loading="isLoading">
    <div class="table">
      <div class="form-content">
        <el-form>
          <el-form-item label="">
            <el-date-picker
              v-model="filterForm.dateValue"
              type="daterange"
              value-format="YYYY-MM-DD"
              unlink-panels
              range-separator="To"
              start-placeholder="Start Date"
              end-placeholder="End Date"
            />
          </el-form-item>
        </el-form>
        <div style="margin: 5px 0px 0px 10px">
          <el-button class="app-button" @click="search()">
            <font-awesome-icon
              icon="fa-solid fa-magnifying-glass"
              style="margin-right: 5px"
            />Search</el-button
          >
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
        <el-table-column label="Book ID" align="center">
          <template #default="scope">
            {{ scope.row.book_data.uuid }}
          </template>
        </el-table-column>

        <el-table-column label="Book Title" align="center">
          <template #default="scope">
            {{ scope.row.book_data.title }}
          </template>
        </el-table-column>

        <el-table-column label="Total Income Book" align="center">
          <template #default="scope">
            {{ scope.row.total_book }}
          </template>
        </el-table-column>

        <el-table-column label="Borrow Count" align="center">
          <template #default="scope">
            {{ scope.row.borrow_count }}
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
</template>
k
<script>
import { onMounted, reactive, toRefs, computed } from "vue";
import http from "@/http";
import { useStore } from "vuex";
import useTableData from "@/hooks/useTableData.js";
import { dateFormat, originalDate } from "@/utils/timeFormat.js";
import { ElMessageBox, ElMessage } from "element-plus";
import { useI18n } from "vue-i18n";

export default {
  name: "Category",
  components: {},
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
    });

    const store = useStore();

    const getTableLists = () => {
      state.isLoading = true;
      http.report.getBooksSummary(state.param).then((res) => {
        if (res.data.err_code == 0) {
          console.log(res.data.data.list);
          state.tableLists = res.data.data.list;
          state.total = res.data.data.total;
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

      if (state.filterForm.dateValue && state.filterForm.dateValue.length) {
        state.param.start_date = `${state.filterForm.dateValue[0]} 00:00:00`;
        state.param.end_date = `${state.filterForm.dateValue[1]} 23:59:59`;
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
      windowRect: computed(() => store.state.app.windowRect),
      device: computed(() => store.state.app.device),
      searchShow,
      t,
      originalDate,
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
.form-content {
  display: flex;
  //align-items: center;
}
</style>
