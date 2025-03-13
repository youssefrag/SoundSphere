<template>
  <!-- Auth Modal -->
  <div
    class="fixed z-10 inset-0 overflow-y-auto"
    id="modal"
    :class="hiddenClass"
  >
    <div
      class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0"
    >
      <div class="fixed inset-0 transition-opacity">
        <div class="absolute inset-0 bg-gray-800 opacity-75"></div>
      </div>

      <span class="hidden sm:inline-block sm:align-middle sm:h-screen"
        >&#8203;</span
      >

      <div
        class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
      >
        <div class="py-4 text-left px-6">
          <!--Title-->
          <div class="flex justify-between items-center pb-4">
            <p class="text-2xl font-bold">Your Account</p>
            <!-- Modal Close Button -->
            <div
              class="modal-close cursor-pointer z-50"
              @click="modalStore.toggleModal"
            >
              <i class="fas fa-times text-xl"></i>
            </div>
          </div>

          <div class="flex justify-between gap-3 mb-6">
            <div
              class="flex-1 rounded py-3 px-4 transition cursor-pointer text-center font-bold"
              @click="tab = 'login'"
              :class="{
                'bg-[#055a32] text-white': tab === 'login',
                'hover:text-[#055a32]': tab === 'register',
              }"
            >
              Login
            </div>
            <div
              class="flex-1 rounded py-3 px-4 transition cursor-pointer text-center font-bold"
              @click="tab = 'register'"
              :class="{
                'bg-[#055a32] text-white': tab === 'register',
                'hover:text-[#055a32]': tab === 'login',
              }"
            >
              Register
            </div>
          </div>

          <app-login-form v-if="tab === 'login'" />
          <app-register-form v-else />
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import useModalStore from "@/stores/modal";
import AppLoginForm from "./AppLoginForm.vue";
import AppRegisterForm from "./AppRegisterForm.vue";
import { ref } from "vue";

let modalStore = useModalStore();

let tab = ref("login");
</script>
