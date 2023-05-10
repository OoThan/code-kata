<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="180px" ref="formRef" :model="form">
      <el-form-item
        label="Username : "
        prop="username"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.username" placeholder="" />
      </el-form-item>

      <el-form-item
        label="NRC : "
        prop="user_nrc"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.user_nrc" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Phone No : "
        prop="user_phone_number"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.user_phone_number" placeholder="" />
      </el-form-item>

      <el-form-item label="Reference User Name : ">
        <el-input v-model="form.reference_user_name" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Street : "
        prop="street"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.street" placeholder="" />
      </el-form-item>

      <el-form-item
        label="City : "
        prop="city"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.city" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Region : "
        prop="region"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.region" placeholder="" />
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
        username: "",
        user_nrc: "",
        user_phone_number: "",
        reference_user_name: "",
        street: "",
        city: "",
        region: "",
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
            http.userManagement.addUser(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  username: "",
                  user_nrc: "",
                  user_phone_number: "",
                  reference_user_name: "",
                  street: "",
                  city: "",
                  region: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.userManagement.editUser(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  username: "",
                  user_nrc: "",
                  user_phone_number: "",
                  reference_user_name: "",
                  street: "",
                  city: "",
                  region: "",
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
          user_nrc: props.data.user_nrc,
          user_phone_number: props.data.user_phone_number,
          reference_user_name: props.data.reference_user_name,
          street: props.data.street,
          city: props.data.city,
          region: props.data.region,
        };
      } else {
        state.form = {
          username: "",
          user_nrc: "",
          user_phone_number: "",
          reference_user_name: "",
          street: "",
          city: "",
          region: "",
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
