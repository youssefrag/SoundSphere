<template>
  <nav
    class="sticky top-0 z-50 h-[100px] bg-[#0E0E0F] w-[100%] flex items-center justify-between px-20 border-b border-gray-100"
  >
    <router-link to="/" class="flex items-center gap-4 md-down:hidden">
      <img alt="logo" src="../assets/Vector.png" />
      <h1 class="text-3xl font-bold text-[#fff]">SoundSphere</h1>
    </router-link>
    <div class="flex items-center gap-8">
      <!-- <a href="/" class="text-[#fff] font-extrabold hover:text-[#FF4D42]"
        >HOME</a
      > -->
      <router-link
        to="/"
        class="text-[#fff] font-extrabold hover:text-[#FF4D42]"
        >Home</router-link
      >
      <router-link
        v-if="userStore.accessToken"
        to="/profile"
        class="text-[#fff] font-extrabold hover:text-[#FF4D42]"
        >Profile</router-link
      >
      <div
        @click="modalStore.toggle"
        v-else
        class="text-[#fff] font-extrabold hover:text-[#FF4D42] cursor-pointer"
      >
        Profile
      </div>
    </div>
    <div class="flex items-center gap-4">
      <div
        @click="onSearchClick"
        class="h-[32px] w-[32px] bg-gray-600 rounded-full flex items-center justify-center cursor-pointer"
      >
        <font-awesome-icon :icon="['fas', 'search']" class="text-[#fff]" />
      </div>

      <button
        @click="userStore.login('davidsmith@gmail.com', 'SuperSecret123')"
        v-if="!userStore.accessToken"
        class="text-white text-sm py-2 px-6 rounded-3xl bg-[#4636FF] flex gap-2 items-center"
      >
        <font-awesome-icon :icon="['fas', 'user']" class="text-white" />
        <div class="font-bold">Demo Account Login</div>
      </button>

      <button
        v-if="!userStore.accessToken"
        @click="modalStore.toggle"
        class="text-[#0E0E0F] text-sm py-2 px-8 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C]"
      >
        Login
      </button>
      <router-link to="/profile">
        <button
          v-if="userStore.accessToken"
          class="text-white text-sm py-2 px-6 rounded-3xl bg-[#4636FF] flex gap-2 items-center"
        >
          <font-awesome-icon :icon="['fas', 'user']" class="text-white" />
          <div class="font-bold">{{ userStore.name.split(" ")[0] }}</div>
        </button>
      </router-link>
      <button
        v-if="userStore.accessToken"
        @click="userStore.logout"
        class="text-[#0E0E0F] text-sm py-2 px-8 rounded-3xl bg-gradient-to-r from-[#98C970] to-[#0DE27C]"
      >
        Logout
      </button>
    </div>
  </nav>
</template>

<script setup>
import { useRouter } from "vue-router";

import { useModalStore } from "@/stores/modal";
import { useUserStore } from "@/stores/user";

const router = useRouter();

const modalStore = useModalStore();
const userStore = useUserStore();

function onSearchClick() {
  router.push({
    path: "/",
    query: { focus: "search", t: Date.now() },
  });
}
</script>
