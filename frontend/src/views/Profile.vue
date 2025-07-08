<template>
  <section
    class="flex flex-col justify-center w-[100%] top-[100px] bg-[#0E0E0F] px-[40px] py-[40px]"
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

    <div
      class="w-full flex flex-col items-center transform -translate-y-[125px]"
    >
      <v-lazy-image
        v-if="userStore.imageUrl"
        :src="userStore.imageUrl"
        class="h-[250px] w-[250px] object-cover rounded-full border-2 border-white"
      />
      <v-lazy-image
        v-else
        src="https://wallpapers.com/png/generic-person-icon-profile-ulmsmhnz0kqafcqn-2.html"
        class="h-[250px] w-[250px] object-cover rounded-full border-2 border-white"
      />
    </div>
  </section>
</template>

<script setup>
import VLazyImage from "v-lazy-image";

import { useUserStore } from "@/stores/user";

const userStore = useUserStore();

console.log(userStore.imageUrl);
</script>
