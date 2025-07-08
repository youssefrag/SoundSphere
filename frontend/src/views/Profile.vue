<template>
  <section
    class="flex items-end justify-center w-[100%] top-[100px] bg-[#0E0E0F] px-[40px] py-[40px]"
  >
    <div
      class="w-[100%] h-[300px] rounded-xl bg-gradient-to-r from-[#15104c] via-black to-[#032d19] flex flex-col justify-end items-start p-[20px]"
    >
      <div
        class="items-self-end flex justify-center items-center h-[60px] w-[60px] bg-[#0DE27C] rounded-full pl-2"
      >
        <font-awesome-icon
          :icon="['fas', 'play']"
          class="text-white text-3xl"
        />
      </div>
    </div>

    <div class="absolute left-0 bottom-0 w-full transform translate-y-1/2">
      <v-lazy-image :src="userStore.imageUrl" />
    </div>
  </section>
</template>

<script setup>
import VLazyImage from "v-lazy-image";

import { useUserStore } from "@/stores/user";

const userStore = useUserStore();

console.log(userStore.imageUrl);
</script>
