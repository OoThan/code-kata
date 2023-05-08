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
        prop="name"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.name" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Email : "
        prop="email"
        :rules="[
          { required: true, message: 'Required!', trigger: 'blur' },
          {
            required: true,
            message: 'Please enter valid email!',
            trigger: 'blur',
            pattern: /^[\w\-]+@[a-zA-Z\d\-]+(\.[a-zA-Z]{2,8}){1,2}$/,
          },
        ]"
      >
        <el-input v-model="form.email" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Department :"
        prop="department"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-select
          :multiple="false"
          placeholder="Please select department"
          style="width: 100%"
          v-model="form.department"
          default-first-option
        >
          <el-option
            v-for="item in departmentLists"
            :key="item"
            :value="item.id"
            :label="item.title"
          />
        </el-select>
      </el-form-item>

      <el-form-item
        v-if="dialogTitle == 'Add'"
        label="Password"
        prop="password"
        :rules="[
          { required: true, message: 'Required!', trigger: 'blur' },
          {
            required: true,
            message:
              'Password must contain at least 1 Upper and lower letter, 1 digit and 1 special character!',
            trigger: 'blur',
            pattern:
              /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$/,
          },
        ]"
      >
        <el-input v-model="form.password" placeholder="" type="password" />
      </el-form-item>

      <el-form-item label="Password" v-else>
        <el-input v-model="form.password" placeholder="" type="password" />
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
  props: ["show", "title", "data", "roleList", "departmentLists"],
  setup(props, context) {
    const { t } = useI18n();
    const state = reactive({
      dialogTitle: "",
      departmentLists: [],
      uploadPercent: 0,

      form: {
        name: "",
        email: "",
        password: "",
        department: "",
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
          state.form.department = parseInt(state.form.department);
          if (state.dialogTitle == "Add") {
            http.userManagement.addTeacher(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  name: "",
                  email: "",
                  password: "",
                  department: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.userManagement.editTeacher(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  name: "",
                  email: "",
                  password: "",
                  department: "",
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
      state.departmentLists = props.departmentLists;
      if (props.data.hasOwnProperty("id")) {
        state.form = {
          id: props.data.id,
          uuid: props.data.uuid,
          name: props.data.name,
          email: props.data.email,
          password: props.data.password,
          department: props.data.department,
        };
      } else {
        state.form = {
          name: "",
          email: "",
          password: "",
          department: "",
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
