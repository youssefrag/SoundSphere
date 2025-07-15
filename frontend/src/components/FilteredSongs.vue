<template>
  <div class="grid grid-cols-2 gap-6 px-[100px]">
    <div
      v-for="song in musicStore.filteredSongs"
      :key="song.id"
      class="bg-[#1A1A1B] h-[200px] w-[100%] rounded-lg p-[25px] flex justify-between"
    >
      <div class="inline-grid grid-cols-[max-content] gap-3">
        <div class="text-white text-xl font-bold">{{ song.name }}</div>
        <div class="text-[#7D72FF] text-md font-semibold">
          {{ song.genre }}
        </div>
        <div class="text-[#0DE27C] text-md font-bold">
          {{ formatDuration(song.duration) }}
        </div>
      </div>
      <div class="flex flex-col justify-between items-center">
        <div
          @click="
            playerStore.togglePlay({
              id: song.id,
              name: song.name,
              duration: song.duration,
              songUrl: song.songUrl,
              artistName: song.artistName,
            })
          "
          class="group bg-gray-800 rounded-full h-[36px] w-[36px] flex items-center justify-center hover:bg-white transition-colors duration-200 cursor-pointer"
        >
          <font-awesome-icon
            :icon="
              playerStore.playing && playerStore.currentSong?.id === song.id
                ? ['fas', 'pause']
                : ['fas', 'play']
            "
            class="text-[#0DE27c]/80 group-hover:text-[#099E56] transition-colors duration-200"
          />
        </div>
        <div
          class="group flex items-center gap-2 text-[#A49CFF] border-b-2 border-transparent hover:border-[#A49CFF] transition-colors duration-300 ease-in-out cursor-pointer"
        >
          <router-link :to="`/song/${song.id}`">
            <div class="font-bold">Song page</div>
          </router-link>
          <font-awesome-icon
            class="font-bold opacity-0 group-hover:opacity-100 transition-opacity duration-300 ease-in-out"
            :icon="['fas', 'arrow-right']"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useMusicStore } from "@/stores/music";
import { usePlayerStore } from "@/stores/player";
import { formatDuration } from "@/utilities/helpers";

const musicStore = useMusicStore();
const playerStore = usePlayerStore();

console.log(musicStore.filteredSongs);
</script>

<style></style>
