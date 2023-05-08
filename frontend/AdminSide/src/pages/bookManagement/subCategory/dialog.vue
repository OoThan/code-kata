<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="170px" ref="formRef" :model="form">
      <el-form-item
        label="Sub category name : "
        prop="sub_category_name"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.sub_category_name" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Category :"
        prop="category_id"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-select
          :multiple="false"
          placeholder="please select category"
          style="width: 100%"
          v-model="form.category_id"
          default-first-option
        >
          <el-option
            v-for="item in categoryLists"
            :key="item"
            :value="item.id"
            :label="item.title"
          />
        </el-select>
      </el-form-item>

      <!-- <el-form-item
        label="Description : "
        prop="desc"
        :rules="[{ required: true, message: 'Required!', trigger: 'blur' }]"
      >
        <el-input v-model="form.desc" placeholder="" />
      </el-form-item> -->

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
  props: ["show", "title", "data", "categoryLists"],
  setup(props, context) {
    const { t } = useI18n();
    const state = reactive({
      dialogTitle: "",
      uploadPercent: 0,

      form: {
        sub_category_name: "",
        category_id: "",
        // desc: "",
      },
      percentage: 0,
      categoryLists: [],
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
            http.bookManagement.addSubCategoryList(state.form).then((res) => {
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
            http.bookManagement.editSubCategoryList(state.form).then((res) => {
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
      state.categoryLists = props.categoryLists;

      if (props.data.hasOwnProperty("id")) {
        state.form = {
          id: props.data.id,
          sub_category_name: props.data.sub_category_name,
          category_id: props.data.category_id,
          desc: props.data.desc,
        };
      } else {
        state.form = {
          sub_category_name: "",
          category_id: "",
          // desc: "",
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
