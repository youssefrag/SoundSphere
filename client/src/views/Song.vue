<template>
  <section
    class="w-[100%] h-[400px] absolute top-[100px] bg-gradient-to-r from-[#0E0E0F] to-[#1c1666] px-[60px] py-[20px]"
  >
    <div class="h-[320px] w-[320px]" :class="bgClass"></div>
    <!-- <div class="h-[320px] w-[320px]">
      <img :src="imageUrl" alt="artist" class="h-[100%] w-auto" />
    </div> -->
  </section>
</template>

<script setup>
import { computed } from "vue";

const song = {
  name: "Midnight Drive",
  artist: "Alex Waves",
  genre: "Hip Hop",
  pictureUrl: "artist1.jpeg",
};

const imageUrl = `http://localhost:5173/src/assets/images/${song.pictureUrl}`;
const bgClass = `bg-[url(http://localhost:5173/src/assets/images/${song.pictureUrl})]`;

console.log(imageUrl);
</script>
