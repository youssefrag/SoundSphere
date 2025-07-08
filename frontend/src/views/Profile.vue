<template>
  <section
    class="flex flex-col justify-center w-[100%] bg-[#0E0E0F] px-[40px] pt-[40px]"
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
      class="w-full flex flex-col items-center transform -translate-y-[125px] mb-[-100px]"
    >
      <v-lazy-image
        v-if="userStore.imageUrl"
        :src="userStore.imageUrl"
        class="h-[250px] w-[250px] object-cover rounded-full border-2 border-white"
      />
      <v-lazy-image
        v-else
        src="https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
        class="h-[250px] w-[250px] object-cover rounded-full border-2 border-white"
      />

      <div class="text-white font-extrabold text-3xl mt-4">
        {{ userStore.name }}
      </div>
      <div class="text-white text mt-4">Music Artist</div>
      <div class="font-bold mt-4 text-[#007bff] flex items-center gap-2">
        <font-awesome-icon :icon="['fas', 'check']" />

        <div>Verfied</div>
      </div>
    </div>
  </section>

  <section
    class="flex flex-col justify-center w-[100%] bg-[#0E0E0F] px-[60px] py-[40px]"
  >
    <div class="text-3xl font font-extrabold text-white mb-[40px]">
      Upload Your New Song 🔥
    </div>
    <div class="flex justify-between gap-8 mb-8">
      <div class="flex flex-col flex-1 gap-2">
        <label class="text-white pl-2" for="songName">Song Name</label>
        <input
          id="songName"
          type="text"
          class="w-[100%] p-[10px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] focus:border focus:border-[#0DE27C] focus:outline-none text-sm font-semibold text-white placeholder-white rounded-2xl"
        />
      </div>
      <div class="flex flex-col flex-1 gap-2">
        <label class="text-white pl-2" for="genre">Genre</label>
        <input
          id="genre"
          type="text"
          class="w-[100%] p-[10px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] focus:border focus:border-[#0DE27C] focus:outline-none text-sm font-semibold text-white placeholder-white rounded-2xl"
        />
      </div>
      <div class="flex flex-col flex-1 gap-2">
        <label for="upload">Upload</label>
        <button
          type="button"
          class="text-white text-lg font-bold py-2 px-8 rounded-xl bg-gray-700 flex-1"
        >
          <font-awesome-icon :icon="['fas', 'upload']" class="text-[#FFA900]" />
          {{ fileName || "Select Song" }}
        </button>
      </div>
    </div>
    <button
      class="self-start text-[#0E0E0F] text-md font-semibold py-3 px-12 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C]"
    >
      UPLOAD SONG
    </button>
  </section>
</template>

<script setup>
import VLazyImage from "v-lazy-image";

import { useUserStore } from "@/stores/user";

const userStore = useUserStore();

console.log(userStore.imageUrl);
</script>
