<template class="rounded-2xl">
  <div
    class="h-[488px] w-[650px] bg-center bg-cover rounded-2xl"
    :class="{
      'bg-[url(@/assets/images/artist-bg1.avif)]': index % 3 == 0,
      'bg-[url(@/assets/images/artist-bg2.avif)]': index % 3 == 1,
      'bg-[url(@/assets/images/artist-bg3.avif)]': index % 3 == 2,
    }"
  >
    <div
      class="h-[100%] w-[100%] bg-[#000] opacity-85 rounded-2xl overflow-x-auto py-[30px] px-[24px]"
    >
      <div class="flex items-center gap-5">
        <div
          class="bg-[#fff] h-[80px] w-[80px] opacity-100 bg-cover bg-center rounded-xl"
          :style="{
            backgroundImage: `url(${artist.imageUrl || DEFAULT_AVATAR})`,
          }"
        ></div>
        <div class="text-white text-xl font-bold">{{ artist.name }}</div>
      </div>
      <div class="text-white text-2xl font-bold mt-7 mb-3">All Songs</div>
      <div
        v-for="(song, index) in songs"
        :key="index"
        class="py-5 flex justify-between items-center"
      >
        <router-link
          :to="{ name: 'song', params: { id: song.id } }"
          class="text-white font-bold"
          >{{ song.name }}</router-link
        >
        <div class="flex items-center gap-3">
          <div class="text-white font-bold">
            {{ formatDuration(song.duration) }}
          </div>

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
            class="h-[36px] w-[36px] bg-gray-800 opacity-100 flex justify-center items-center rounded-full cursor-pointer"
          >
            <font-awesome-icon
              :icon="
                playerStore.playing && playerStore.currentSong?.id === song.id
                  ? ['fas', 'pause']
                  : ['fas', 'play']
              "
              class="text-[#0DE27C] text-xl"
            />
          </div>

          <router-link
            :to="{ name: 'song', params: { id: song.id } }"
            class="h-[36px] w-[36px] bg-gray-800 opacity-100 flex justify-center items-center rounded-full"
          >
            <font-awesome-icon
              :icon="['fas', 'comment']"
              class="text-[#0DE27C] text-xl"
            />
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { usePlayerStore } from "@/stores/player";

import { formatDuration } from "@/utilities/helpers";
const { artist, index } = defineProps(["index", "artist"]);

const playerStore = usePlayerStore();

const DEFAULT_AVATAR =
  "https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D";

const { songs } = artist;
</script>
