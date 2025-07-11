import { defineStore } from "pinia";
import { ref } from "vue";

export const useModalStore = defineStore("modal", () => {
  // state
  const isOpen = ref(false);

  // actions
  function open() {
    isOpen.value = true;
  }
  function close() {
    isOpen.value = false;
  }
  function toggle() {
    isOpen.value = !isOpen.value;
  }

  return { isOpen, open, close, toggle };
});
