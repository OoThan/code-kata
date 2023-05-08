<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="250px" ref="formRef" :model="form">
      <el-form-item
        label="Package Name : "
        prop="package_name"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.package_name" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Duration : "
        prop="duration_end"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.duration_end" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Teacher Punishment Amount : "
        prop="teacher_punishment_amt"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.teacher_punishment_amt" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Student Punishment Amount : "
        prop="student_punishment_amt"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.student_punishment_amt" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Staff Punishment Amount : "
        prop="staff_punishment_amt"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.staff_punishment_amt" placeholder="" />
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
        package_name: "",
        teacher_punishment_amt: 0,
        student_punishment_amt: 0,
        staff_punishment_amt: 0,
        duration_start: 0,
        duration_end: 0,
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
          state.form.duration_start = parseInt(state.form.duration_start);
          state.form.duration_end = parseInt(state.form.duration_end);

          state.form.teacher_punishment_amt = parseInt(
            state.form.teacher_punishment_amt
          );
          state.form.student_punishment_amt = parseInt(
            state.form.student_punishment_amt
          );
          state.form.staff_punishment_amt = parseInt(
            state.form.staff_punishment_amt
          );

          if (state.dialogTitle == "Add") {
            http.config.addPunishment(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {};
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.config.updatePunishment(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {};
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
          package_name: props.data.package_name,
          teacher_punishment_amt: props.data.teacher_punishment_amt,
          student_punishment_amt: props.data.student_punishment_amt,
          staff_punishment_amt: props.data.staff_punishment_amt,
          duration_start: props.data.duration_start,
          duration_end: props.data.duration_end,
        };
      } else {
        state.form = {
          package_name: "",
          teacher_punishment_amt: 0,
          student_punishment_amt: 0,
          staff_punishment_amt: 0,
          duration_start: 0,
          duration_end: 0,
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
