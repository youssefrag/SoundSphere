<template>
  <!-- Player -->
  <div class="fixed bottom-0 left-0 bg-white px-4 py-2 w-full">
    <!-- Track Info -->
    <div class="text-center" v-if="currentSong.name">
      <span class="song-title font-bold">{{ currentSong.name }}</span>
      by
      <span class="song-artist">{{ currentSong.artistName }}</span>
    </div>
    <div class="flex flex-nowrap gap-4 items-center">
      <div @click="togglePlay(currentSong)" class="cursor-pointer">
        <font-awesome-icon
          :icon="playing ? ['fas', 'pause'] : ['fas', 'play']"
          class="text-white text-4xl"
        />
      </div>

      <div class="player-currenttime">{{ seek }}</div>
    </div>
  </div>
</template>

<script setup>
import { usePlayerStore } from "@/stores/player";

import { storeToRefs } from "pinia";

const playerStore = usePlayerStore();

const { currentSong, playing, seek } = storeToRefs(playerStore);
const { togglePlay } = playerStore;

console.log(currentSong);
</script>
