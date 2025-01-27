<template>
  <!-- <div>
    <h1>Example Navbar</h1>
    <font-awesome-icon :icon="['fas', 'phone']" />
  </div> -->
  <nav class="sticky z-10 bg-[#faf5f8] dark:bg-[#0a0508] h-20 border-b">
    <button @click="toggleDark()" class="dark:text-white">Change color</button>
  </nav>
</template>

<script setup>
import { useDark, useToggle } from "@vueuse/core";

const isDark = useDark();
const toggleDark = useToggle(isDark);

console.log(isDark.value);
</script>
