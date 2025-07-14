<template>
  <div>
    <app-header />
    <RouterView />
    <Auth v-if="modalStore.isOpen" />
    <Player />
  </div>
</template>

<script setup>
import { RouterLink, RouterView } from "vue-router";
import AppHeader from "@/components/Header.vue";
import Auth from "@/components/Auth.vue";

import Player from "./components/Player.vue";

import { useModalStore } from "@/stores/modal";

const modalStore = useModalStore();
</script>
