<template>
  <nav
    class="sticky top-0 z-50 h-[100px] bg-[#0E0E0F] w-[100%] flex items-center justify-between px-20 border-b border-gray-100"
  >
    <div class="flex items-center gap-4">
      <img alt="logo" src="../assets/Vector.png" />
      <h1 class="text-3xl font-bold text-[#fff]">SoundSphere</h1>
    </div>
    <div class="flex items-center gap-8">
      <a href="/" class="text-[#fff] font-extrabold hover:text-[#FF4D42]"
        >HOME</a
      >
      <a href="/" class="text-[#fff] font-extrabold hover:text-[#FF4D42]"
        >BROWSE</a
      >
      <a href="/" class="text-[#fff] font-extrabold hover:text-[#FF4D42]"
        >PROFILE</a
      >
    </div>
    <div class="flex items-center gap-4">
      <div
        class="h-[32px] w-[32px] bg-gray-600 rounded-full flex items-center justify-center cursor-pointer"
      >
        <font-awesome-icon :icon="['fas', 'search']" class="text-[#fff]" />
      </div>

      <button
        @click="modalStore.toggleModal"
        class="text-[#0E0E0F] text-sm py-2 px-8 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C]"
      >
        Login
      </button>
    </div>
  </nav>
</template>
☹️

<script setup>
import useModalStore from "@/stores/modal";

let modalStore = useModalStore();
</script>
