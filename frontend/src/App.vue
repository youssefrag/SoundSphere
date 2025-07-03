<template>
  <div>
    <app-header />
    <RouterView />
    <Auth v-if="modalStore.isOpen" />
  </div>
</template>

<script setup>
import { RouterLink, RouterView } from "vue-router";
import AppHeader from "@/components/Header.vue";
import Auth from "@/components/Auth.vue";

import { useModalStore } from "@/stores/modal";

const modalStore = useModalStore();
</script>
