<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="180px" ref="formRef" :model="form">
      <!-- : "",
          : "",
          : "", -->
      <el-form-item
        label="Package No : "
        prop="package_no"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.package_no" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Amount : "
        prop="amount"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.amount" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Percent : "
        prop="percent"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.percent" placeholder="" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeDialog(formRef)"> Close </el-button>
        <el-button class="app-button" @click="submitDialog(formRef)">
          Sure
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import { onMounted, reactive, toRefs, ref, watch, onUpdated } from "vue";
import http from "@/http";
import { ElMessage } from "element-plus";
import { getToken } from "@/utils/cookie";
import axios from "axios";
import { prop } from "dom7";
import { useI18n } from "vue-i18n";
export default {
  name: "Dialog",
  props: ["show", "title", "data", "roleList"],
  setup(props, context) {
    const { t } = useI18n();
    const state = reactive({
      dialogTitle: "",
      uploadPercent: 0,

      form: {
        package_no: "",
        amount: "",
        percent: "",
      },
      roleList: [],
      percentage: 0,
    });

    const formRef = ref();

    const auth = {
      Authorization: getToken(),
    };

    const closeDialog = (formRef) => {
      formRef.resetFields();

      context.emit("closed");
    };

    const submitDialog = (formRef) => {
      formRef.validate((valid) => {
        if (valid) {
          state.form.amount = parseFloat(state.form.amount);
          state.form.percent = parseFloat(state.form.percent);
          if (state.dialogTitle == "Add") {
            http.loanManagement.addLoanPackage(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  package_no: "",
                  amount: "",
                  percent: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.loanManagement.editLoanPackage(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  package_no: "",
                  amount: "",
                  percent: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          }
        }
      });
    };

    onUpdated(() => {
      state.dialogTitle = props.title;
      state.roleList = props.roleList;
      if (props.data.hasOwnProperty("id")) {
        state.form = {
          id: props.data.id,
          package_no: props.data.package_no,
          amount: props.data.amount,
          percent: props.data.percent,
        };
      } else {
        state.form = {
          package_no: "",
          amount: "",
          percent: "",
        };
      }
    });
    onMounted(() => {});

    return {
      ...toRefs(state),
      closeDialog,
      submitDialog,
      formRef,
      auth,
      t,
    };
  },
};
</script>

<style lang="scss" scoped>
:deep(.el-select .el-input__inner) {
  height: 42px !important;
}

.lose {
  color: #f56c6c;
}

.win {
  color: #67c23a;
}
</style>
