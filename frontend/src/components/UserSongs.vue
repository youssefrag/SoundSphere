<template>
  <section class="bg-[#0E0E0F] px-[60px] py-[40px]">
    <h1 class="text-white text-3xl font-extrabold mb-10">My Songs</h1>
    <div class="grid grid-cols-2 gap-4">
      <div
        v-for="song in musicStore.userSongs"
        :key="song.id"
        class="bg-[#1A1A1B] h-[160px] w-[100%] rounded-lg p-[25px] flex justify-between"
      >
        <div class="flex flex-col">
          <div class="text-white text-xl font-bold mb-3">{{ song.name }}</div>
          <div class="text-[#0DE27C] text-md font-bold mb-4">
            {{ formatDuration(song.duration) }}
          </div>
          <div class="flex gap-3">
            <div
              @click="musicStore.deleteSong(song.id)"
              class="group bg-gray-800 rounded-full h-[36px] w-[36px] flex items-center justify-center hover:bg-white transition-colors duration-200 cursor-pointer"
            >
              <font-awesome-icon
                :icon="['fas', 'trash']"
                class="text-red-500/80 group-hover:text-red-600 transition-colors duration-200"
              />
            </div>
            <div
              class="group bg-gray-800 rounded-full h-[36px] w-[36px] flex items-center justify-center hover:bg-white transition-colors duration-200 cursor-pointer"
            >
              <font-awesome-icon
                :icon="['fas', 'pencil']"
                class="text-[#FFA900]/80 group-hover:text-[#D97706] transition-colors duration-200"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { onMounted } from "vue";
import { useMusicStore } from "@/stores/music";
import { formatDuration } from "../utilities/helpers";

const musicStore = useMusicStore();

onMounted(async () => {
  await musicStore.fetchAllArtists();
});
</script>
