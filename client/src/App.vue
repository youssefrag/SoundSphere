<template>
  <Navbar />
  <RouterView />
</template>

<script setup lang="ts">
import Navbar from "@/components/Navbar.vue";
import { RouterLink, RouterView } from "vue-router";
</script>
