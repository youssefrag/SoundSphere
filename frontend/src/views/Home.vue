<template>
  <div>
    <section
      class="h-[44rem] bg-[url(@/assets/images/home-bg.jpeg)] bg-cover bg-top md-down:h-[35rem] sm-down:h-[28rem]"
    >
      <section
        class="h-[100%] bg-[#0E0E0F] opacity-90 text-white flex flex-col items-center py-[8rem] gap-10 sm-down:gap-5 sm-down:py-[6rem]"
      >
        <h1
          class="text-white text-6xl font-bold tracking-wide text-center md-down:text-5xl sm-down:text-2xl"
        >
          WHERE YOUR MUSIC SHAPES THE
        </h1>
        <span
          class="text-transparent bg-clip-text bg-gradient-to-r from-[#4636FF] to-[#0DE27C] text-7xl font-bold tracking-wider mb-10 md-down:text-5xl sm-down:text-3xl"
          >SPHERE</span
        >
        <div
          class="w-[824px] h-[84px] bg-gray-400 rounded-full pr-3 pl-10 flex items-center justify-between md-down:w-[600px] md-down:h-[60px] sm-down:w-[400px] sm-down:h-[40px]"
        >
          <input
            ref="searchInput"
            maxlength="15"
            v-model="musicStore.searchTerm"
            type="text"
            placeholder="Seach for music here"
            class="w-[80%] bg-gray-400 border-none focus:border-none focus:outline-none text-3xl font-black text-gray-700 placeholder-gray-700 md-down:w-[60%] md-down:text-xl"
          />
        </div>
      </section>
    </section>
    <section
      v-if="musicStore.searchTerm === ''"
      class="pb-[200px] bg-[#0E0E0F] py-[6rem]"
    >
      <div
        class="text-[#fff] text-center text-5xl mb-[4rem] font-semibold md-down:text-4xl sm-down:text-3xl"
      >
        Trending & Latest Music
      </div>
      <div class="flex items-center h-[30rem] px-10 gap-8">
        <font-awesome-icon
          :icon="['fas', 'arrow-left']"
          class="text-[#FFA900] text-4xl cursor-pointer"
          @click="scrollLeft"
        />

        <div
          class="flex-1 w-[100%] flex gap-5 pb-[2rem] overflow-y-auto [&::-webkit-scrollbar]:h-2 [&::-webkit-scrollbar-track]:bg-[#565657] [&::-webkit-scrollbar-thumb]:bg-[#4636FF] [&::-webkit-scrollbar-track]:rounded-full [&::-webkit-scrollbar-thumb]:rounded-full"
          ref="carousel"
        >
          <div v-for="(artist, index) in allArtists" :key="index">
            <artist-card :artist="artist" :index="index" />
          </div>
        </div>
        <font-awesome-icon
          :icon="['fas', 'arrow-right']"
          class="text-[#FFA900] text-4xl cursor-pointer"
          @click="scrollRight"
        />
      </div>
    </section>
    <section
      ref="resultsSection"
      v-else
      class="pb-[200px] bg-[#0E0E0F] py-[6rem]"
    >
      <div class="flex items-center justify-center mb-[4rem] gap-4">
        <div
          class="text-[#fff] text-center text-5xl font-semibold md-down:text-4xl"
        >
          Search results for "{{ musicStore.searchTerm }}"
        </div>
        <div
          @click="musicStore.clearSearchTerm"
          class="h-[52px] w-[52px] bg-gray-800 opacity-100 flex justify-center items-center rounded-full cursor-pointer md-down:h-[45px] md-down:w-[45px]"
        >
          <font-awesome-icon
            :icon="['fas', 'x']"
            class="text-[#4636FF] text-3xl md-down:text-2xl"
          />
        </div>
      </div>
      <div
        v-if="musicStore.filteredSongs.length === 0"
        class="text-center text-2xl font-bold text-[#FFA900]"
      >
        No results!
      </div>
      <FilteredSongs v-else />
    </section>
  </div>
</template>

<script setup>
import { useScroll } from "@vueuse/core";
import {
  useTemplateRef,
  onMounted,
  computed,
  ref,
  watch,
  nextTick,
  onUnmounted,
} from "vue";
import { useRoute, useRouter } from "vue-router";

import { useMusicStore } from "@/stores/music";
import ArtistCard from "@/components/ArtistCard.vue";
import FilteredSongs from "@/components/FilteredSongs.vue";

const musicStore = useMusicStore();

const route = useRoute();
const router = useRouter();
const searchInput = ref(null);

const loadingArtists = ref(true);

const allArtists = computed(() => musicStore.allArtists);

onMounted(async () => {
  await musicStore.fetchAllArtists(); // whatever your store action is
  loadingArtists.value = false;
});

onUnmounted(() => {
  musicStore.clearSearchTerm();
});

const carousel = useTemplateRef("carousel");
let { x, y } = useScroll(carousel, { behavior: "smooth" });

const scrollRight = () => {
  x.value += 400;
};

const scrollLeft = () => {
  x.value -= 400;
};

// Logic for when we navigate using search bar
watch(
  () => route.query.focus,
  async (focus) => {
    if (focus === "search") {
      await nextTick(); // wait for DOM
      searchInput.value?.focus(); // focus the input
      // remove the query so it won't re‑trigger
      const { focus: _, t, ...keep } = route.query;
      router.replace({ query: keep });
    }
  },
  { immediate: true }
);

const resultsSection = ref(null);

const ABOVE = 500;

watch(
  () => musicStore.searchTerm,
  (newTerm, oldTerm) => {
    if (oldTerm === "" && newTerm.length > 0) {
      nextTick(() => {
        const el = resultsSection.value;
        if (!el) return;

        // get element's position relative to whole page...
        const topOfEl = el.getBoundingClientRect().top + window.pageYOffset;
        // …then subtract your offset
        window.scrollTo({
          top: topOfEl - ABOVE,
          behavior: "smooth",
        });
      });
    }
  }
);
</script>
