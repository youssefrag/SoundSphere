<template>
  <!-- Player -->
  <div class="fixed bottom-0 left-0 bg-white px-4 py-2 w-full">
    <!-- Track Info -->
    <div class="text-center" v-if="currentSong.name">
      <span class="song-title font-bold text-[#15104c]">{{
        currentSong.name
      }}</span>
      by
      <span class="song-artist font-bold text-[#15104c]">{{
        currentSong.artistName
      }}</span>
    </div>
    <div class="flex flex-nowrap gap-4 items-center">
      <div @click="togglePlay(currentSong)" class="cursor-pointer">
        <font-awesome-icon
          :icon="playing ? ['fas', 'pause'] : ['fas', 'play']"
          class="text-[#07713e] text-4xl"
        />
      </div>

      <div class="player-currenttime">{{ seek }}</div>
      <div
        @click.prevent="updateSeek($event)"
        class="w-full h-2 rounded bg-gray-200 relative cursor-pointer"
      >
        <span
          class="absolute -top-2.5 -ml-2.5 text-gray-800 text-lg"
          :style="{ left: playerProgress }"
        >
          <font-awesome-icon :icon="['fas', 'circle']" />
        </span>
        <span
          class="block h-2 rounded bg-gradient-to-r from-[#07713e] to-[#07713e]"
          :style="{ width: playerProgress }"
        ></span>
      </div>
      <div class="player-duration">{{ duration }}</div>
    </div>
  </div>
</template>

<script setup>
import { usePlayerStore } from "@/stores/player";

import { storeToRefs } from "pinia";

const playerStore = usePlayerStore();

const { currentSong, playing, seek, playerProgress, duration } =
  storeToRefs(playerStore);
const { togglePlay, updateSeek } = playerStore;
</script>
