<template>
  <div>
    <section
      class="h-[44rem] bg-[url(@/assets/images/home-bg.jpeg)] bg-cover bg-top"
    >
      <section
        class="h-[100%] bg-[#0E0E0F] opacity-90 text-white flex flex-col items-center py-[8rem] gap-10"
      >
        <h1 class="text-white text-7xl font-bold tracking-wide">
          WHERE YOUR MUSIC SHAPES THE
        </h1>
        <span
          class="text-transparent bg-clip-text bg-gradient-to-r from-[#4636FF] to-[#0DE27C] text-7xl font-bold tracking-wider mb-10"
          >SPHERE</span
        >
        <div
          class="w-[824px] h-[84px] bg-gray-400 rounded-full pr-3 pl-10 flex items-center justify-between"
        >
          <input
            type="text"
            placeholder="Seach for music here"
            class="w-[80%] bg-gray-400 border-none focus:border-none focus:outline-none text-3xl font-black text-gray-700 placeholder-gray-700"
          />
          <button
            class="text-[#0E0E0F] text-sm py-2 px-8 rounded-full bg-gradient-to-r from-[#98C970] to-[#0DE27C] h-[80%]"
          >
            Search
          </button>
        </div>
      </section>
    </section>
    <section class="h-[757px] bg-[#0E0E0F] py-[6rem]">
      <div class="text-[#fff] text-center text-5xl mb-[4rem] font-semibold">
        Trending & Latest Music
      </div>
      <div class="flex items-center h-[30rem] px-10 gap-8">
        <font-awesome-icon
          :icon="['fas', 'arrow-left']"
          class="text-[#FFA900] text-4xl"
          @click="scrollLeft"
        />

        <div
          class="flex-1 w-[100%] flex gap-5 pb-[2rem] overflow-y-auto [&::-webkit-scrollbar]:h-2 [&::-webkit-scrollbar-track]:bg-[#565657] [&::-webkit-scrollbar-thumb]:bg-[#4636FF] [&::-webkit-scrollbar-track]:rounded-full [&::-webkit-scrollbar-thumb]:rounded-full"
          ref="carousel"
        >
          <div v-for="(artist, index) in music" :key="index">
            <artist-card :artist="artist" :index="index" />
          </div>
        </div>
        <font-awesome-icon
          :icon="['fas', 'arrow-right']"
          class="text-[#FFA900] text-4xl"
          @click="scrollRight"
        />
      </div>
    </section>
  </div>
</template>

<script setup>
import { useScroll } from "@vueuse/core";
import { useTemplateRef } from "vue";

import useMusicStore from "@/stores/music";
import ArtistCard from "@/components/ArtistCard.vue";

let { music } = useMusicStore();

const carousel = useTemplateRef("carousel");
let { x, y } = useScroll(carousel, { behavior: "smooth" });

const scrollRight = () => {
  x.value += 400;
};

const scrollLeft = () => {
  x.value -= 400;
};
</script>
