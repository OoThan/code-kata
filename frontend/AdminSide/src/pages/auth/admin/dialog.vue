<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="120px" ref="formRef" :model="form">
      <el-form-item
        label="Username : "
        prop="username"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.username" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Email : "
        prop="email"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.email" placeholder="" />
      </el-form-item>

      <el-form-item
        v-if="dialogTitle == 'Add'"
        label="Password"
        prop="password"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.password" placeholder="" type="password" />
      </el-form-item>

      <el-form-item label="Password" v-else>
        <el-input v-model="form.password" placeholder="" type="password" />
      </el-form-item>

      <!-- <el-form-item :label="t('table.state')">
        <el-radio-group v-model.number="form.status">
          <el-radio :label="1">{{ t("common.normal") }}</el-radio>
          <el-radio :label="2">{{ t("common.hide") }}</el-radio>
        </el-radio-group>
      </el-form-item> -->
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
        username: "",
        email: "",
        password: "",
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
          if (state.dialogTitle == "Add") {
            http.admin.addAdminList(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  name: "",
                  email: "",
                  password: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.admin.editAdminList(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  name: "",
                  email: "",
                  password: "",
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
          username: props.data.username,
          email: props.data.email,
          password: props.data.password,
          status: props.data.status,
        };
      } else {
        state.form = {
          name: "",
          email: "",
          password: "",
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
