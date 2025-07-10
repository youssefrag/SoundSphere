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

        <div>Verified</div>
      </div>
    </div>
  </section>

  <section
    class="flex flex-col justify-center w-[100%] bg-[#0E0E0F] px-[60px] py-[40px]"
  >
    <div class="text-3xl font-extrabold text-white mb-[40px]">
      Upload Your New Song 🔥
    </div>
    <Form @submit="onSubmit">
      <div class="flex justify-between gap-8 mb-8">
        <!-- Song Name -->
        <div class="flex flex-col flex-1 gap-2">
          <label class="text-white pl-2" for="songName">Song Name</label>
          <input
            id="songName"
            v-model="songName"
            @blur="songNameBlur"
            type="text"
            placeholder="Enter song name"
            class="w-full p-[10px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] rounded-2xl text-sm font-semibold text-white placeholder-white/30 focus:outline-none focus:border-[#0DE27C]"
          />
          <p v-if="songNameError" class="text-red-500 text-sm pl-2">
            {{ songNameError }}
          </p>
        </div>

        <!-- Genre -->
        <div class="flex flex-col flex-1 gap-2">
          <label class="text-white pl-2" for="genre">Genre</label>
          <input
            id="genre"
            v-model="genre"
            @blur="genreBlur"
            type="text"
            placeholder="e.g. Hip-Hop"
            class="w-full p-[10px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] rounded-2xl text-sm font-semibold text-white placeholder-white/30 focus:outline-none focus:border-[#0DE27C]"
          />
          <p v-if="genreError" class="text-red-500 text-sm pl-2">
            {{ genreError }}
          </p>
        </div>

        <!-- File Picker -->
        <div class="flex flex-col flex-1 gap-2">
          <label class="text-white pl-2" for="upload">Upload</label>
          <input
            ref="fileInput"
            id="upload"
            type="file"
            accept="audio/*"
            @change="onFileChange"
            @blur="fileBlur"
            class="hidden"
          />
          <button
            @click="triggerFileDialog"
            type="button"
            class="inline-flex items-center py-2 px-8 bg-gray-700 rounded-xl font-bold text-lg text-white"
          >
            <font-awesome-icon
              :icon="['fas', 'upload']"
              class="text-[#FFA900] mr-2"
            />
            {{ fileName || "Select Song" }}
          </button>
          <p v-if="fileError" class="text-red-500 text-sm pl-2">
            {{ fileError }}
          </p>
        </div>
      </div>

      <button
        type="submit"
        :disabled="!meta.valid || loading"
        class="self-start text-[#0E0E0F] text-md font-semibold py-3 px-12 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C] disabled:opacity-50"
      >
        <span v-if="loading">Uploading…</span>
        <span v-else>UPLOAD SONG</span>
      </button>
    </Form>
  </section>
  <div ref="songsSection">
    <UserSongs />
  </div>
</template>

<script setup>
import VLazyImage from "v-lazy-image";
import { ref as vueRef, nextTick } from "vue";
import { Form, useForm, useField } from "vee-validate";
import * as yup from "yup";

import { useUserStore } from "@/stores/user";
import { useMusicStore } from "@/stores/music";
import { storage, storageRef, uploadBytesToPath, getURL } from "@/firebase";
import UserSongs from "@/components/UserSongs.vue";

const userStore = useUserStore();
const musicStore = useMusicStore();

// reactive refs
const fileInput = vueRef(null);
const fileName = vueRef("");
const loading = vueRef(false);

const songsSection = vueRef(null);

// validation schema
const schema = yup.object({
  songName: yup.string().required("Name is required"),
  genre: yup.string().required("Genre is required"),
  file: yup
    .mixed()
    .required("Please select a file")
    .test(
      "file-type",
      "Only audio files allowed",
      (v) => v && v.type.startsWith("audio/")
    ),
});

// form setup
const { handleSubmit, resetForm, meta } = useForm({
  validationSchema: schema,
  validateOnBlur: true,
  validateOnChange: false,
});

// fields
const {
  value: songName,
  errorMessage: songNameError,
  handleBlur: songNameBlur,
} = useField("songName");
const {
  value: genre,
  errorMessage: genreError,
  handleBlur: genreBlur,
} = useField("genre");
const {
  value: file,
  errorMessage: fileError,
  setValue: setFile,
  handleBlur: fileBlur,
} = useField("file");

// handlers
function triggerFileDialog() {
  fileInput.value?.click();
}
function onFileChange(e) {
  const selected = e.target.files?.[0] ?? null;
  setFile(selected);
  fileName.value = selected?.name || "";
}

async function uploadSong(file) {
  const path = `songs/${userStore.email}/${file.name}`;
  const songRef = storageRef(storage, path);
  await uploadBytesToPath(songRef, file);

  const url = await getURL(songRef);
  return { url, path };
}

function getAudioDuration(file) {
  return new Promise((resolve, reject) => {
    const audio = document.createElement("audio");
    audio.preload = "metadata";

    audio.onloadedmetadata = () => {
      URL.revokeObjectURL(audio.src);
      const seconds = Math.floor(audio.duration);
      resolve(seconds);
    };
    audio.onerror = (e) => reject(e);

    audio.src = URL.createObjectURL(file);
  });
}

// submit
const onSubmit = handleSubmit(async (values) => {
  loading.value = true;

  const duration = await getAudioDuration(file.value);

  try {
    const { url: songUrl, path: storagePath } = await uploadSong(file.value);

    musicStore.saveSong(
      values.songName,
      values.genre,
      userStore.email,
      songUrl,
      duration,
      storagePath
    );
  } catch (err) {
    console.log("🔴🔴🔴🔴🔴🔴🔴🔴🔴🔴🔴🔴🔴🔴🔴", err);
  } finally {
    resetForm({
      values: {
        songName: "",
        genre: "",
        file: null,
      },
    });

    fileName.value = "";

    if (fileInput.value) {
      fileInput.value.value = null;
    }

    loading.value = false;

    await nextTick();
    // then scroll!
    songsSection.value?.scrollIntoView({
      behavior: "smooth",
      block: "end",
      inline: "nearest",
    });
  }
});
</script>
