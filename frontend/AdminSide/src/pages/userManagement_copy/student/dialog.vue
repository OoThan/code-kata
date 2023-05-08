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
        label="Role no : "
        prop="role_no"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.role_no" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Year :"
        prop="year"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-select
          :multiple="false"
          placeholder="Please select year"
          style="width: 100%"
          v-model="form.year"
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
  props: ["show", "title", "data", "roleList", "yearLists"],
  setup(props, context) {
    const { t } = useI18n();
    const state = reactive({
      dialogTitle: "",
      uploadPercent: 0,
      yearLists: [],
      form: {
        name: "",
        email: "",
        password: "",
        role_no: "",
        year: "",
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
          state.form.year = parseInt(state.form.year);
          if (state.dialogTitle == "Add") {
            http.userManagement.addStudent(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  name: "",
                  email: "",
                  password: "",
                  role_no: "",
                  year: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.userManagement.editStudent(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  name: "",
                  email: "",
                  password: "",
                  role_no: "",
                  year: "",
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
      state.yearLists = props.yearLists;
      if (props.data.hasOwnProperty("id")) {
        state.form = {
          id: props.data.id,
          uuid: props.data.uuid,
          name: props.data.name,
          email: props.data.email,
          password: props.data.password,
          role_no: props.data.role_no,
          year: props.data.year,
        };
      } else {
        state.form = {
          name: "",
          email: "",
          password: "",
          role_no: "",
          year: "",
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
