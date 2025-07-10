<template>
  <section class="bg-[#0E0E0F] px-[60px] py-[40px]">
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
              class="bg-transparent border-b border-gray-400 text-white text-xl font-bold pb-1 focus:outline-none"
            />
            <input
              v-model="editValues[song.id].genre"
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
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { onMounted, reactive, nextTick } from "vue";
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
  try {
    const { name, genre } = editValues[songId];

    console.log({ name, genre, originalName, originalGenre });

    if (name === originalName && genre === originalGenre) {
      return;
    }

    // call your update endpoint
    // await api.put(`/songs/${songId}`, { name, genre });
    // re-fetch so the store (and UI) is in sync
    // await musicStore.fetchAllArtists();
  } catch (err) {
    console.error("could not save song:", err);
  } finally {
    // always exit edit mode
    editing[songId] = false;
  }
}
</script>
