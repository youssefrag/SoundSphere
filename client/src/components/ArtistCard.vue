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
          :class="{
            'bg-[url(@/assets/images/artist1.jpeg)]': index % 3 == 0,
            'bg-[url(@/assets/images/artist2.jpeg)]': index % 3 == 1,
            'bg-[url(@/assets/images/artist3.jpeg)]': index % 3 == 2,
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
          <div class="text-white font-bold">{{ song.duration }}</div>

          <div
            class="h-[36px] w-[36px] bg-gray-800 opacity-100 flex justify-center items-center rounded-full"
          >
            <font-awesome-icon
              :icon="['fas', 'play']"
              class="text-[#0DE27C] text-xl"
            />
          </div>

          <div
            class="h-[36px] w-[36px] bg-gray-800 opacity-100 flex justify-center items-center rounded-full"
          >
            <font-awesome-icon
              :icon="['fas', 'comment']"
              class="text-[#0DE27C] text-xl"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const { artist, index } = defineProps(["index", "artist"]);

const { songs } = artist;

console.log(songs);
</script>
