<template>
  <el-result
    icon="error"
    title="404"
    subTitle="Page does not exist!"
    style="width: 100%; height: 100vh"
  >
    <template #extra>
      <el-button type="primary" @click="onNavigateBack" class="app-button"
        >return</el-button
      >
    </template>
  </el-result>
</template>

<script>
import { defineComponent } from "vue";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "404",

  setup() {
    const router = useRouter();
    const onNavigateBack = () => {
      router.back();
    };

    return {
      onNavigateBack,
    };
  },
});
</script>
<style lang="scss" scoped>
.con {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100vh;

  img {
    width: 40vw;
    margin-top: 20px;
    margin-bottom: 20px;
  }
}
</style>
