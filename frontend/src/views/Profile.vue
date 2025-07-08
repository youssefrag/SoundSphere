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

        <div>Verfied</div>
      </div>
    </div>
  </section>

  <section
    class="flex flex-col justify-center w-[100%] bg-[#0E0E0F] px-[60px] py-[40px]"
  >
    <div class="text-3xl font font-extrabold text-white mb-[40px]">
      Upload Your New Song 🔥
    </div>
    <Form @submit="onSubmit">
      <div class="flex justify-between gap-8 mb-8">
        <div class="flex flex-col flex-1 gap-2">
          <label class="text-white pl-2" for="songName">Song Name</label>
          <input
            id="songName"
            v-model="songName"
            type="text"
            class="w-[100%] p-[10px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] focus:border focus:border-[#0DE27C] focus:outline-none text-sm font-semibold text-white placeholder-white placeholder-white/30 rounded-2xl"
            placeholder="Enter song name"
          />
        </div>
        <div class="flex flex-col flex-1 gap-2">
          <label class="text-white pl-2" for="genre">Genre</label>
          <input
            id="genre"
            v-model="genre"
            type="text"
            class="w-[100%] p-[10px] bg-gradient-to-r from-[#0E0E0F] to-[#1e2816] border border-[#07713e] focus:border focus:border-[#0DE27C] focus:outline-none text-sm font-semibold text-white placeholder-white/30 placeholder-white rounded-2xl"
            placeholder="e.g. Hip-Hop"
          />
        </div>
        <div class="flex flex-col flex-1 gap-2">
          <label for="upload">Upload</label>
          <!-- hidden actual file input -->
          <input
            ref="fileInput"
            id="upload"
            type="file"
            accept="audio/*"
            @change="onFileChange"
            class="hidden"
          />
          <button
            @click="triggerFileDialog"
            type="button"
            class="text-white text-lg font-bold py-2 px-8 rounded-xl bg-gray-700 flex-1"
          >
            <font-awesome-icon
              :icon="['fas', 'upload']"
              class="text-[#FFA900]"
            />
            {{ fileName || "Select Song" }}
          </button>
        </div>
      </div>
      <button
        type="submit"
        class="self-start text-[#0E0E0F] text-md font-semibold py-3 px-12 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C]"
      >
        UPLOAD SONG
      </button>
    </Form>
  </section>
</template>

<script setup>
import VLazyImage from "v-lazy-image";
import { ref as vueRef } from "vue";
import { Form, useForm, useField } from "vee-validate";
import * as yup from "yup";

import { useUserStore } from "@/stores/user";

const userStore = useUserStore();

// const file = vueRef(null);
const fileName = vueRef("");
const loading = vueRef(false);
const fileInput = vueRef(null);
// const fileError = vueRef("");

const schema = yup.object({
  songName: yup.string().required("Name is required"),
  genre: yup.string().required("Genre is required"),
  file: yup
    .mixed()
    .required("Please select a file")
    .test(
      "file-type",
      "Only audio files allowed",
      (value) => value && value.type.startsWith("audio/")
    ),
});

const { handleSubmit, meta } = useForm({
  validationSchema: schema,
  validateOnBlur: true,
  validateOnChange: false,
});

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

function onFileChange(e) {
  const selected = e.target.files?.[0] ?? null;
  file.value = selected;
  fileName.value = selected.name;

  schema
    .validateAt("file", { ...meta.values, file: file.value })
    .then(() => (fileError.value = ""))
    .catch((err) => (fileError.value = err.message));
}

function triggerFileDialog() {
  fileInput.value?.click();
}

const onSubmit = handleSubmit(async (values) => {
  loading.value = true;
  try {
    // you have:
    // - values.songName
    // - values.genre
    // - file.value
    // - userStore.email / userStore.name

    // 1) await uploadSong(file.value)
    // 2) await saveSongMetadata({ ...values, url, artist: userStore.name })
    console.log("Ready to upload:", {
      ...values,
      file: file.value,
      artist: userStore.name,
    });
  } finally {
    loading.value = false;
  }
});
</script>
