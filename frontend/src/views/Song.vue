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
        @click="
          playerStore.togglePlay({
            id: song.id,
            name: song.name,
            duration: song.duration,
            songUrl: song.songUrl,
            artistName: song.artistName,
          })
        "
        class="flex justify-center items-center h-[80px] w-[80px] bg-[#0DE27C] rounded-full"
      >
        <font-awesome-icon
          :icon="
            playerStore.playing && playerStore.currentSong?.id === song.id
              ? ['fas', 'pause']
              : ['fas', 'play']
          "
          class="text-white text-4xl"
        />
      </div>
    </section>
    <section class="w-[100%] bg-[#0E0E0F] px-[40px] py-[60px]">
      <div class="text-white text-4xl font-semibold mb-10">
        Here's What People Are Saying About {{ song.name }}
      </div>

      <div class="flex flex-col gap-4">
        <textarea
          v-model="commentText"
          maxlength="400"
          cols="4"
          type="text"
          placeholder="Write Comment Here..."
          class="w-[100%] h-[200px] p-[20px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] focus:border focus:border-[#0DE27C] focus:outline-none text-lg font-semibold text-white placeholder-white/70 rounded-2xl"
        />
        <button
          @click="submitComment"
          class="mr-[20px] self-end text-[#0E0E0F] text-md font-bold py-2 px-8 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C]"
        >
          {{ submitting ? "Submitting..." : "SUBMIT COMMENT" }}
        </button>
      </div>
    </section>
    <section
      class="pb-[200px] w-[100%] bg-[#0E0E0F] px-[40px] py-[60px] grid grid-cols-2 gap-4"
    >
      <div
        v-for="comment in comments"
        :key="comment.id"
        class="border border-white h-[240px] rounded-2xl px-[30px] py-[20px] flex flex-col justify-between"
      >
        <div class="flex justify-between items-center">
          <div class="flex gap-4">
            <div
              class="flex justify-center items-center h-[50px] w-[50px] overflow-hidden rounded-full"
            >
              <img
                :src="comment.userImageUrl"
                class="h-full w-full object-cover object-center"
                alt="User avatar"
              />
            </div>
            <div class="flex flex-col justify-between">
              <div class="text-white font-semibold text-lg">
                {{ comment.userName }}
              </div>
              <div class="text-[#FFA900] font-bold text-sm">ARTIST</div>
            </div>
          </div>
          <div
            v-if="userStore.email === comment.userEmail"
            @click="deleteComment(comment.id)"
            class="text-red-500 hover:text-red-400 text-md cursor-pointer transform transition-all duration-200 hover:scale-[1.2]"
          >
            <font-awesome-icon :icon="['fas', 'trash']" />
          </div>
        </div>
        <div class="text-white font-bold">{{ comment.content }}</div>
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
            {{ formatDDMMYYYY(comment.createdAt) }}
          </div>
        </div>
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
import { usePlayerStore } from "@/stores/player";

import { formatDDMMYYYY } from "@/utilities/helpers";

const route = useRoute();
const songId = route.params.id;
const musicStore = useMusicStore();
const userStore = useUserStore();
const modalStore = useModalStore();
const playerStore = usePlayerStore();

const song = ref(null);
const loading = ref(true);
const error = ref(null);

const comments = ref([]);
const commentsLoading = ref(true);
const commentsError = ref(null);

const commentText = ref("");
const submitting = ref(false);

async function fetchComments() {
  commentsLoading.value = true;
  commentsError.value = false;
  try {
    const { data } = await api.get(`/comments/${songId}`);
    comments.value = data;
  } catch {
    commentsError.value = true;
  } finally {
    commentsLoading.value = false;
  }
}

async function deleteComment(commentId) {
  try {
    await api.delete(`/comments/${commentId}`);
  } catch (err) {
    console.error(err);
  } finally {
    await fetchComments();
  }
}

onBeforeMount(async () => {
  try {
    const details = await musicStore.fetchSongDetails(songId);
    song.value = { ...details, id: Number(songId) };
  } catch (e) {
    error.value = true;
  } finally {
    loading.value = false;
  }

  await fetchComments();
});

async function submitComment() {
  if (!userStore.accessToken) {
    modalStore.open();
    return;
  }

  if (!commentText.value.trim()) return;

  submitting.value = true;
  try {
    await api.post("/comments/addNew", {
      songId,
      content: commentText.value.trim(),
      email: userStore.email,
    });

    commentText.value = "";

    await fetchComments();
  } catch (err) {
    console.error("Failed to submit comment", err);
  } finally {
    submitting.value = false;
  }
}
</script>
