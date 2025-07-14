<template>
  <section class="pb-[100px] bg-[#0E0E0F] px-[60px] py-[40px]">
    <h1 class="text-white text-3xl font-extrabold mb-10">My Songs</h1>
    <div class="grid grid-cols-2 gap-4">
      <div
        v-for="song in musicStore.userSongs"
        :key="song.id"
        class="bg-[#1A1A1B] h-[200px] w-[100%] rounded-lg p-[25px] flex justify-between"
      >
        <div class="flex flex-col">
          <div
            v-if="editing[song.id]"
            class="inline-grid grid-cols-[max-content] gap-3 mb-3"
          >
            <input
              v-model="editValues[song.id].name"
              :ref="(el) => (nameInputs[song.id] = el)"
              @keyup.enter.prevent="saveSong(song.id, song.name, song.genre)"
              class="bg-transparent border-b border-gray-400 text-white text-xl font-bold pb-1 focus:outline-none"
            />
            <input
              v-model="editValues[song.id].genre"
              @keyup.enter.prevent="saveSong(song.id, song.name, song.genre)"
              class="bg-transparent border-b border-gray-400 text-[#7D72FF] text-md font-semibold pb-1 focus:outline-none"
            />
          </div>
          <div v-else class="inline-grid grid-cols-[max-content] gap-3 mb-3">
            <div class="text-white text-xl font-bold">{{ song.name }}</div>
            <div class="text-[#7D72FF] text-md font-semibold">
              {{ song.genre }}
            </div>
          </div>

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
              @click="
                editing[song.id]
                  ? saveSong(song.id, song.name, song.genre)
                  : startEdit(song)
              "
              class="group bg-gray-800 rounded-full h-[36px] w-[36px] flex items-center justify-center hover:bg-white transition-colors duration-200 cursor-pointer"
            >
              <font-awesome-icon
                :icon="editing[song.id] ? ['fas', 'save'] : ['fas', 'pencil']"
                :class="
                  editing[song.id]
                    ? 'text-[#FFA900]/80 group-hover:text-[#D97706'
                    : 'text-[#FFA900]/80 group-hover:text-[#D97706]'
                "
              />
            </div>
            <div
              class="group bg-gray-800 rounded-full h-[36px] w-[36px] flex items-center justify-center hover:bg-white transition-colors duration-200 cursor-pointer"
            >
              <font-awesome-icon
                :icon="['fas', 'comment']"
                class="text-[#4636FF]/80 group-hover:text-[#382BCC] transition-colors duration-200"
              />
            </div>
          </div>
        </div>
        <div class="flex flex-col justify-between items-center">
          <div
            class="group bg-gray-800 rounded-full h-[36px] w-[36px] flex items-center justify-center hover:bg-white transition-colors duration-200 cursor-pointer"
          >
            <font-awesome-icon
              :icon="['fas', 'play']"
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
  </section>
</template>

<script setup>
import { onMounted, reactive, nextTick } from "vue";

import api from "@/api";
import { useMusicStore } from "@/stores/music";
import { formatDuration } from "../utilities/helpers";

const musicStore = useMusicStore();

const editing = reactive({});
const editValues = reactive({});
const nameInputs = reactive({});

onMounted(async () => {
  await musicStore.fetchAllArtists();
});

function startEdit(song) {
  editing[song.id] = true;

  editValues[song.id] = {
    name: song.name,
    genre: song.genre,
  };

  nextTick(() => {
    nameInputs[song.id]?.focus();
  });
}

async function saveSong(songId, originalName, originalGenre) {
  const { name, genre } = editValues[songId];

  if (name === originalName && genre === originalGenre) {
    editing[songId] = false;
    return;
  }

  await musicStore.editSong(songId, name, genre);

  editing[songId] = false;
}
</script>
