<template>
  <div>
    <nav
      class="sticky z-10 bg-[#faf5f8] dark:bg-[#0a0508] text-[#140b10] border-b flex items-center px-14 py-5 justify-between"
    >
      <h3
        class="text-[#140b10] dark:text-[#f4ebf0] font-bold text-5xl cursor-pointer"
      >
        Sound<span class="text-[#7d4561] dark:text-[#b17192]">S</span>phere
      </h3>
      <div class="flex gap-8">
        <button
          class="text-[1rem] px-[2rem] py-[0.5rem] rounded-[0.5rem] text-[#140b10] dark:text-[#f4ebf0] bg-[#aad5af] dark:bg-[#2a552f] bg-opacity-30 dark:bg-opacity-30"
        >
          Manage Music
        </button>
        <button
          class="text-[1rem] px-[2rem] py-[0.5rem] rounded-[0.5rem] text-[#140b10] dark:text-[#f4ebf0] bg-[#b2628b] dark:bg-[#9d4d77]"
        >
          Login/Register
        </button>
        <font-awesome-icon
          :icon="['fas', 'moon']"
          class="cursor-pointer text-4xl text-[#140b10] transition duration-300 ease-in-out hover:scale-125 hover:text-[#b2628b] dark:hidden"
          @click="toggleDark()"
        />
        <font-awesome-icon
          :icon="['fas', 'sun']"
          class="cursor-pointer text-4xl text-[#f4ebf0] transition duration-300 ease-in-out hover:scale-125 hover:text-[#9d4d77] hidden dark:inline-block"
          @click="toggleDark()"
        />
      </div>
    </nav>
  </div>
</template>

<script setup>
import { useDark, useToggle } from "@vueuse/core";

const isDark = useDark();
const toggleDark = useToggle(isDark);

console.log(isDark.value);
</script>
