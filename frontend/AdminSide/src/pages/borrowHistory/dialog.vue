<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="120px" ref="formRef" :model="form">
      <el-form-item
        label="Borrow user :"
        prop="type"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-select
          :multiple="false"
          placeholder="Please select borrow user"
          style="width: 100%"
          v-model="form.type"
          default-first-option
        >
          <el-option
            v-for="item in borrowerTypes"
            :key="item"
            :value="item.id"
            :label="item.title"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="Book ID : "
        prop="book_uuid"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.book_uuid" placeholder="" />
      </el-form-item>

      <el-form-item
        label="User ID : "
        prop="user_uuid"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.user_uuid" placeholder="" />
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
        user_uuid: "",
        book_uuid: "",
        type: "",
      },
      roleList: [],
      percentage: 0,
      borrowerTypes: [
        {
          id: 1,
          title: "Teacher",
        },
        {
          id: 2,
          title: "Student",
        },
        {
          id: 3,
          title: "Staff",
        },
      ],
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
          state.form.type = parseInt(state.form.type);
          if (state.dialogTitle == "Add") {
            http.borrowHistory.addBorrow(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  user_uuid: "",
                  book_uuid: "",
                  type: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.bookManagement.editCategoryList(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  user_uuid: "",
                  book_uuid: "",
                  type: "",
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
          title: props.data.title,
          desc: props.data.desc,
        };
      } else {
        state.form = {
          user_uuid: "",
          book_uuid: "",
          type: "",
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
