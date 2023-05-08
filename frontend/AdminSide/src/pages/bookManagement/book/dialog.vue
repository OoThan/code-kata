<template>
  <el-dialog
    v-model="show"
    @close="closeDialog(formRef)"
    draggable
    :title="dialogTitle"
  >
    <el-form label-width="150px" ref="formRef" :model="form">
      <el-form-item
        label="Title"
        prop="title"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-input v-model="form.title" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Sub Category :"
        prop="sub_category_id"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-select
          :multiple="false"
          placeholder="please select category"
          style="width: 100%"
          v-model="form.sub_category_id"
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
        label="Status: "
        prop="status"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-input v-model="form.status" placeholder="" />
      </el-form-item> -->

      <el-form-item
        label="Status :"
        prop="status"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-select
          :multiple="false"
          placeholder="please select status"
          style="width: 100%"
          v-model="form.status"
          default-first-option
        >
          <el-option
            v-for="item in statusOptions"
            :key="item"
            :value="item.id"
            :label="item.title"
          />
        </el-select>
      </el-form-item>

      <el-form-item
        label="Author: "
        prop="author"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-input v-model="form.author" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Available Qty: "
        prop="available_qty"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-input v-model="form.available_qty" placeholder="" />
      </el-form-item>

      <el-form-item
        label="Description: "
        prop="summary"
        :rules="[{ required: true, message: 'Required !', trigger: 'blur' }]"
      >
        <el-input v-model="form.summary" placeholder="" type="textarea" />
      </el-form-item>

      <el-form-item label="Download Link: ">
        <el-input v-model="form.download_link" placeholder="" />
      </el-form-item>

      <!-- <el-form-item label="Publish date: ">
        <el-input v-model="form.publish_date" placeholder="" />
      </el-form-item> -->

      <el-form-item label="Publish date:">
        <el-date-picker
          v-model="form.publish_date"
          value-format="YYYY-MM-DD"
          unlink-panels
          placeholder="Select Date"
        />
      </el-form-item>

      <el-form-item label="Book Cover:">
        <el-upload
          list-type="picture-card"
          :multiple="false"
          :on-remove="handleImageRemove"
          :on-preview="handlePictureCardPreview"
          :on-success="handleAvatarSuccess"
          :file-list="fileList"
          :limit="2"
          :http-request="handleHttp"
          accept="image/*"
        >
          <el-icon>+</el-icon>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeDialog(formRef)">Close</el-button>
        <el-button class="app-button" @click="submitDialog(formRef)"
          >Sure</el-button
        >
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

export default {
  name: "Dialog",
  props: ["show", "title", "data", "categoryLists", "statusOptions"],
  setup(props, context) {
    const state = reactive({
      categoryLists: [],
      dialogTitle: "",
      uploadPercent: 0,
      form: {
        title: "",
        sub_category_id: "",
        status: 1,
        author: "",
        summary: "",
        book_image: "",
        download_link: "",
        publish_date: "",
      },
      percentage: 0,
      fileList: [],
      statusOptions: [],
    });

    const formRef = ref();

    const auth = {
      Authorization: getToken(),
    };

    const closeDialog = (formRef) => {
      formRef.resetFields();
      state.fileList = [];
      context.emit("closed");
    };

    const submitDialog = (formRef) => {
      console.log("Here is submit 111111");
      formRef.validate((valid) => {
        if (valid) {
          console.log("Here is submit 22222");
          state.form.sub_category_id = parseInt(state.form.sub_category_id);
          state.form.status = parseInt(state.form.status);
          state.form.available_qty = parseInt(state.form.available_qty);

          if (state.dialogTitle == "Add") {
            console.log("Here is submit 333333");

            http.bookManagement.addBookList(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  title: "",
                  sub_category_id: "",
                  status: 1,
                  author: "",
                  summary: "",
                  book_image: "",
                };
                ElMessage.success(res.data.err_msg);

                formRef.resetFields();
                context.emit("created");
              } else {
                ElMessage.error(res.data.err_msg);
              }
            });
          } else {
            http.bookManagement.editBookList(state.form).then((res) => {
              if (res.data.err_code == 0) {
                closeDialog(formRef);
                state.form = {
                  title: "",
                  sub_category_id: "",
                  status: 1,
                  author: "",
                  summary: "",
                  book_image: "",
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

    const handleImageUploadSuccess = (response, file, fileList) => {
      if (response.data.err_code == 0) {
        if (fileList.length > 1) {
          fileList.shift();
        }
        ElMessage.success(response.data.err_msg);
        state.form.icon_url = response.data.url;
      } else {
        ElMessage.error(res.msg);
      }
    };

    const handleImageRemove = (uploadFile, uploadFiles) => {
      state.form.book_image = "";
    };

    const handlePictureCardPreview = () => {
      document.getElementById("carousel-image").click();
    };
    const startUpload = (file, filelist) => {
      const form = new FormData();
      form.append("file", filelist[0].raw);
    };

    const handleHttp = (fileList) => {
      const form = new FormData();
      form.append("file", fileList.file);
      http.user
        .appUpload(form, (progressEvent) => {
          fileList.onProgress({
            percent: parseInt(
              Math.round((progressEvent.loaded * 100) / progressEvent.total)
            ),
          });
        })
        .then((res) => {
          if (res.data.err_code == 0) {
            fileList.onSuccess();
            console.log("Here is success");
            state.form.book_image = res.data.data.url;
            ElMessage.success(res.data.err_msg);
          }
        });
    };

    const handleProgress = (event, file, fileList) => {};

    const handleAvatarSuccess = (response, file, fileList) => {
      if (fileList.length > 1) {
        fileList.shift();
      }
    };

    onMounted(() => {});

    onUpdated(() => {
      state.dialogTitle = props.title;
      state.categoryLists = props.categoryLists;
      state.statusOptions = props.statusOptions;
      if (props.data.hasOwnProperty("id")) {
        state.form.id = props.data.id;
        state.form.title = props.data.title;
        state.form.sub_category_id = props.data.sub_category_id;
        state.form.status = props.data.status;
        state.form.author = props.data.author;
        state.form.summary = props.data.summary;
        state.form.book_image = props.data.book_image;
        state.form.available_qty = props.data.available_qty;
        state.form.borrow_qty = props.data.borrow_qty;
        state.form.download_link = props.data.download_link;
        state.form.publish_date = props.data.publish_date;

        if (props.data.book_image) {
          state.fileList.push({
            url: props.data.book_image,
          });
        }
        if (state.fileList.length > 1) {
          state.fileList.shift();
        }
        // console.log(state.fileList);
      } else {
        state.form = {
          title: "",
          sub_category_id: "",
          status: 1,
          author: "",
          summary: "",
          book_image: "",
          available_qty: "",
          download_link: "",
          publish_date: "",
        };
      }
    });

    return {
      ...toRefs(state),
      closeDialog,
      submitDialog,
      formRef,
      auth,
      handleImageUploadSuccess,
      handleImageRemove,
      handlePictureCardPreview,
      startUpload,
      handleHttp,
      handleProgress,
      handleAvatarSuccess,
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
