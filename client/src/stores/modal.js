import { defineStore } from "pinia";

export default defineStore("modal", {
  state: () => ({
    isOpen: false,
  }),
  actions: {
    toggleModal() {
      this.isOpen = !this.isOpen;
    },
  },
});
