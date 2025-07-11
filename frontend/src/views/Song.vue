<template>
  <div v-if="loading" class="p-8 text-white">Loading song…</div>
  <template v-else>
    <section
      class="flex items-end justify-between w-[100%] h-[400px] top-[100px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] px-[40px] py-[40px]"
    >
      <div class="flex items-end gap-6">
        <v-lazy-image
          :src="song.artistImgUrl"
          class="h-[320px] w-[320px] object-cover object-center rounded-xl"
        />
        <div class="flex flex-col gap-4">
          <div class="text-[#0DE27C] font-bold">{{ song.artistName }}</div>
          <div class="text-[#fff] font-extrabold text-4xl">
            {{ song.name }}
          </div>
          <div class="flex gap-4">
            <div class="flex items-center gap-1">
              <div
                class="h-[28px] w-[28px] bg-slate-700 rounded-full flex justify-center items-center"
              >
                <font-awesome-icon
                  :icon="['fas', 'music']"
                  class="text-[#0DE27C] text-sm"
                />
              </div>
              <div class="text-white text-xs font-extrabold">
                {{ song.genre }}
              </div>
            </div>
            <div class="flex items-center gap-1">
              <div
                class="h-[28px] w-[28px] bg-slate-700 rounded-full flex justify-center items-center"
              >
                <font-awesome-icon
                  :icon="['fas', 'calendar']"
                  class="text-[#FFA900] text-sm"
                />
              </div>
              <div class="text-white text-xs font-extrabold">
                {{ formatDDMMYYYY(song.date) }}
              </div>
            </div>
          </div>
          <div
            class="text-white border-[1px] self-start px-3 text-sm font-bold rounded-full border-slate-500"
          >
            🔥 TRENDING
          </div>
        </div>
      </div>
      <div
        class="flex justify-center items-center h-[80px] w-[80px] bg-[#0DE27C] rounded-full pl-2"
      >
        <font-awesome-icon
          :icon="['fas', 'play']"
          class="text-white text-4xl"
        />
      </div>
    </section>
    <section class="h-[600px] w-[100%] bg-[#0E0E0F] px-[40px] py-[60px]">
      <div class="text-white text-4xl font-semibold mb-10">
        Here's What People Are Saying About {{ song.name }}
      </div>

      <div class="flex flex-col gap-4">
        <textarea
          v-model="commentText"
          cols="4"
          type="text"
          placeholder="Write Comment Here"
          class="w-[100%] h-[200px] p-[20px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] focus:border focus:border-[#0DE27C] focus:outline-none text-sm font-semibold text-white placeholder-white rounded-2xl"
        />
        <button
          @click="submitComment"
          class="mr-[20px] self-end text-[#0E0E0F] text-md font-bold py-2 px-8 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C]"
        >
          {{ submitting ? "Submitting..." : "SUBMIT COMMENT" }}
        </button>
      </div>
    </section>
  </template>
</template>

<script setup>
import VLazyImage from "v-lazy-image";
import { ref, onBeforeMount } from "vue";
import { useRoute } from "vue-router";
import api from "@/api";

import { useMusicStore } from "@/stores/music";
import { useUserStore } from "@/stores/user";
import { useModalStore } from "@/stores/modal";

import { formatDDMMYYYY } from "@/utilities/helpers";

const route = useRoute();
const songId = route.params.id;
const musicStore = useMusicStore();
const userStore = useUserStore();
const modalStore = useModalStore();

const song = ref(null);
const loading = ref(true);
const error = ref(null);

const commentText = ref("");
const submitting = ref(false);

onBeforeMount(async () => {
  try {
    song.value = await musicStore.fetchSongDetails(songId);
  } catch (e) {
    error.value = true;
  } finally {
    loading.value = false;
  }
});

async function submitComment() {
  if (!userStore.accessToken) {
    modalStore.open();
    return;
  }

  if (!commentText.value.trim()) return;

  submitting.value = true;
  try {
    console.log(commentText.value);
    console.log(userStore.email);
    console.log(songId);

    await api.post(`/songs/${songId}/comments`, {
      songId,
      content: commentText.value.trim(),
      user: userStore.email,
    });

    commentText.value = "";
    // optionally reload comments or show a toast...
  } catch (err) {
    console.error("Failed to submit comment", err);
    // you could set an error flag and show something in the UI
  } finally {
    submitting.value = false;
  }
}
</script>
