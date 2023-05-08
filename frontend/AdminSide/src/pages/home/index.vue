<template>
  <div class="home-container" v-loading="isLoading">
    <el-row :gutter="20">
      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con pupple">
                <font-awesome-icon icon="fa-solid fa-people-group" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.total_admin ? data.total_admin : "0" }}
              </h2>
              <h2 class="content">Admin Users</h2>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <!-- First row -->
    <el-row :gutter="20">
      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con pupple">
                <font-awesome-icon icon="fa-solid fa-people-group" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.total_teacher ? data.total_teacher : "0" }}
              </h2>
              <h2 class="content">Teachers</h2>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con pupple">
                <font-awesome-icon icon="fa-solid fa-people-group" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.total_student ? data.total_student : "0" }}
              </h2>
              <h2 class="content">Students</h2>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con pupple">
                <font-awesome-icon icon="fa-solid fa-people-group" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.total_staff ? data.total_staff : "0" }}
              </h2>
              <h2 class="content">Staffs</h2>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <!-- End first row -->

    <!-- First row -->
    <el-row :gutter="20">
      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con skyblue">
                <font-awesome-icon icon="fa-solid fa-volleyball" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.total_category ? data.total_category : "0" }}
              </h2>
              <h2 class="content">Total Categories</h2>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con skyblue">
                <font-awesome-icon icon="fa-solid fa-book" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.total_book ? data.total_book : "0" }}
              </h2>
              <h2 class="content">Total Books</h2>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con skyblue">
                <font-awesome-icon
                  icon="fa-solid fa-arrows-left-right-to-line"
                />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.total_borrow ? data.total_borrow : "0" }}
              </h2>
              <h2 class="content">Total Borrow</h2>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <!-- End first row -->

    <!-- First row -->
    <el-row :gutter="20">
      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con red">
                <font-awesome-icon icon="fa-solid fa-arrows-down-to-line" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.under_borrow ? data.under_borrow : "0" }}
              </h2>
              <h2 class="content">Under Borrowing</h2>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con green">
                <font-awesome-icon icon="fa-solid fa-arrows-up-to-line" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.have_returned ? data.have_returned : "0" }}
              </h2>
              <h2 class="content">Have Returned</h2>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :lg="8" :md="12" :xs="24">
        <el-card :body-style="{ padding: '0px' }" class="home-card">
          <div class="body">
            <div class="left">
              <div class="icon-con red">
                <font-awesome-icon icon="fa-solid fa-arrows-down-to-line" />
              </div>
            </div>
            <div class="right">
              <h2 class="count">
                {{ data.expired_count ? data.expired_count : "0" }}
              </h2>
              <h2 class="content">Expired</h2>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <!-- End first row -->

    <!-- Third row end  -->
  </div>
</template>

<script>
import { onMounted, reactive, toRefs, ref } from "vue";
import { Chart, Grid, Line, Tooltip, Area, Marker } from "vue3-charts";
import Web3 from "web3";
import http from "@/http";
import { useI18n } from "vue-i18n";

export default {
  name: "Home",
  components: { Chart, Grid, Line, Tooltip, Area },

  setup() {
    const state = reactive({
      data: {},
      isLoading: false,
    });
    const { t } = useI18n();

    const getDashboardData = () => {
      state.isLoading = true;
      http.user.getDashboard().then((res) => {
        if (res.data.err_code == 0) {
          state.data = res.data.data;
          state.isLoading = false;
        } else {
          state.data = {};
          state.isLoading = false;
        }
      });
    };

    onMounted(() => {
      getDashboardData();
    });

    return {
      ...toRefs(state),
      t,
    };
  },
};
</script>

<style lang="scss" scoped>
:deep(.el-select .el-input__inner) {
  height: 42px !important;
}

.base-card {
  padding-bottom: 15px;
}

.lose {
  color: #f56c6c;
}

.win {
  color: #000000;
}
.home-card {
  margin-bottom: 20px;
  padding: 20px;
  color: #fff;
  border-radius: 10px;
  box-shadow: none;
  border: none;
  background: rgb(235, 238, 249);
  .body {
    display: flex;
    .left {
      .icon-con {
        background: $sidbarActiveBg;
        border-radius: 10px;
        font-size: 30px;
        width: 65px;
        height: 65px;
        display: flex;
        align-items: center;
        justify-content: center;
      }
    }

    .right {
      color: $title;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      margin-left: 15px;
      .count {
        font-size: 30px;
        font-weight: bold;
      }
      .content {
        font-size: 20px;
        font-weight: 100;
      }
    }
  }
}
.red {
  background-color: #f05050 !important;
}
.pupple {
  background-color: rgb(118, 147, 246) !important;
}
.skyblue {
  background-color: #23b7e5 !important;
}
.green {
  background-color: #27c24c !important;
}
</style>
