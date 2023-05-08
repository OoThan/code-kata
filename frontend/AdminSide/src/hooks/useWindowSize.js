import { toRefs, onMounted, onBeforeUnmount, reactive } from "vue";
import { debounce, getClientRect } from "@/others/utils";

export default function useWindowSize() {
  const windowRect = reactive({
    clientHeight: getClientRect().clientHeight,
    clientWidth: getClientRect().clientWidth,
  });

  const onWindowsResize = debounce(() => {
    windowRect.clientHeight = getClientRect().clientHeight;
    windowRect.clientWidth = getClientRect().clientWidth;
  }, 300);

  onMounted(() => {
    window.addEventListener("resize", onWindowsResize);
  });

  onBeforeUnmount(() => {
    window.removeEventListener("resize", onWindowsResize);
  });
  return {
    ...toRefs(windowRect),
  };
}
