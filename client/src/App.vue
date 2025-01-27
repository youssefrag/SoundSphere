<template>
  <div class="bg-[#e8edf3] dark:bg-[#0c1117] relative">
    <div class="sticky top-4 flex justify-end px-20 py-10 gap-9">
      <font-awesome-icon
        :icon="['fas', 'user']"
        class="cursor-pointer text-4xl text-[#121921] dark:text-[#dee5ed] transition duration-300 ease-in-out hover:scale-125 hover:text-[#334761] dark:hover:text-[#9eb2cc]"
      />
      <font-awesome-icon
        :icon="['fas', 'moon']"
        class="cursor-pointer text-4xl text-[#121921] transition duration-300 ease-in-out hover:scale-125 hover:text-[#334761] dark:hidden"
        @click="toggleDark()"
      />
      <font-awesome-icon
        :icon="['fas', 'sun']"
        class="cursor-pointer text-4xl text-[#dee5ed] transition duration-300 ease-in-out hover:scale-125 hover:text-[#9eb2cc] hidden dark:inline-block"
        @click="toggleDark()"
      />
      <font-awesome-icon
        :icon="['fas', 'bars']"
        class="cursor-pointer text-4xl text-[#121921] dark:text-[#dee5ed] transition duration-300 ease-in-out hover:scale-125 hover:text-[#334761] dark:hover:text-[#9eb2cc]"
      />
      <font-awesome-icon
        :icon="['fas', 'window-close']"
        class="cursor-pointer text-4xl text-[#121921] dark:text-[#dee5ed] transition duration-300 ease-in-out hover:scale-125 hover:text-[#334761] dark:hover:text-[#9eb2cc]"
      />
    </div>
    <nav-menu
      v-if="menuOpen"
      class="sticky top-[10rem] right-[80px] float-right"
    />
    <!-- <h1 class="text-black dark:text-white sticky top-60 right-40 float-right"> -->

    <RouterView />
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { useDark, useToggle } from "@vueuse/core";

import NavMenu from "@/components/NavMenu.vue";

const isDark = useDark();
const toggleDark = useToggle(isDark);

const [menuOpen, toggleMenu] = useToggle(false);
</script>
