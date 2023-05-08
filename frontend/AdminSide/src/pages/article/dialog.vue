<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="100px" ref="formRef" :model="form">
      <el-form-item
        :label="t('common.title')"
        prop="title"
        :rules="[
          { required: true, message: t('common.require'), trigger: 'blur' },
        ]"
      >
        <el-input v-model="form.title" :placeholder="t('common.title')" />
      </el-form-item>

      <el-form-item
        :label="t('common.content')"
        prop="content"
        :rules="[
          { required: true, message: t('common.require'), trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="form.content"
          :placeholder="t('common.content')"
          type="textarea"
        />
      </el-form-item>

      <el-form-item
        :label="t('common.sort')"
        prop="sort"
        :rules="[
          { required: true, message: t('common.require'), trigger: 'blur' },
          {
            required: true,
            message: '金额必须大于 0 !',
            trigger: 'blur',
            pattern: /^(0*[0-9][0-9]*(\.[0-9]+)?|0+\.[0-9]*[0-9][0-9]*)$/,
          },
        ]"
      >
        <el-input v-model="form.sort" :placeholder="t('common.sort')" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeDialog(formRef)">{{
          t("common.close")
        }}</el-button>
        <el-button class="app-button" @click="submitDialog(formRef)">{{
          t("common.sure")
        }}</el-button>
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
import { useI18n } from "vue-i18n";

export default {
  name: "Dialog",
  props: ["show", "title", "data"],
  setup(props, context) {
    const { t } = useI18n();
    const state = reactive({
      dialogTitle: "",
      uploadPercent: 0,

      form: {
        title: "",
        content: "",
        sort: 0,
      },

      percentage: 0,
      isClickConfirm: false,
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
          if (state.isClickConfirm) return;
          state.isClickConfirm = true;
          state.form.sort = parseInt(state.form.sort);
          if (state.dialogTitle == t("common.addto")) {
            http.article
              .addArticle(state.form)
              .then((res) => {
                if (res.data.err_code == 0) {
                  closeDialog(formRef);
                  state.form = {
                    title: "",
                    content: "",
                    sort: 0,
                  };
                  ElMessage.success(res.data.err_msg);

                  formRef.resetFields();
                  context.emit("created");
                  state.isClickConfirm = false;
                } else {
                  ElMessage.error(res.data.err_msg);
                  state.isClickConfirm = false;
                }
              })
              .catch((err) => {
                state.isClickConfirm = false;
              });
          } else {
            http.article
              .editArticle(state.form)
              .then((res) => {
                if (res.data.err_code == 0) {
                  closeDialog(formRef);
                  state.form = {
                    title: "",
                    content: "",
                    sort: 0,
                  };
                  ElMessage.success(res.data.err_msg);

                  formRef.resetFields();
                  context.emit("created");
                  state.isClickConfirm = false;
                } else {
                  ElMessage.error(res.data.err_msg);
                  state.isClickConfirm = false;
                }
              })
              .catch((err) => {
                state.isClickConfirm = false;
                k;
              });
          }
        }
      });
    };

    onUpdated(() => {
      state.dialogTitle = props.title;
      if (props.data.hasOwnProperty("id")) {
        state.form = {
          id: props.data.id,
          title: props.data.title,
          content: props.data.content,
          sort: props.data.sort,
        };
      } else {
        state.form = {
          title: "",
          content: "",
          sort: 0,
        };
      }
    });
    onMounted(() => {});

    return {
      t,
      ...toRefs(state),
      closeDialog,
      submitDialog,
      formRef,
      auth,
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
