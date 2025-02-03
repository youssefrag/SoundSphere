import { defineStore } from "pinia";

export default defineStore("modal", {
  state: () => ({
    isOpen: false,
  }),
  actions: {
    toggleModal() {
      console.log("reached here");
      console.log(this.isOpen);
      this.isOpen = !this.isOpen;
      console.log(this.isOpen);
    },
  },
});
