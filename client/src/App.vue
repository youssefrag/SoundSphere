<template>
  <main>
    <Navbar />
    <RouterView />
  </main>
</template>

<script setup>
import { RouterLink, RouterView } from "vue-router";

import Navbar from "@/components/Navbar.vue";
</script>
