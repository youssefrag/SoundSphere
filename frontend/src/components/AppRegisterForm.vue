<template>
  <div class="flex flex-col gap-6">
    <div>
      <label for="name" class="block mb-2 text-lg font-bold text-[#0E0E0F]"
        >Name</label
      >
      <input
        type="text"
        id="name"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="John Smith"
        required
      />
    </div>
    <div>
      <label for="email" class="block mb-2 text-lg font-bold text-[#0E0E0F]"
        >Email</label
      >
      <input
        type="email"
        id="email"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="john@example.com"
        required
      />
    </div>
    <div>
      <label for="password" class="block mb-2 text-lg font-bold text-[#0E0E0F]"
        >Password</label
      >
      <input
        type="password"
        id="password"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="password"
        required
      />
    </div>
    <button
      class="text-white text-lg font-bold py-2 px-8 rounded-xl bg-gray-700"
    >
      <font-awesome-icon :icon="['fas', 'upload']" class="text-[#FFA900]" />
      Upload Picture
    </button>
    <button
      class="text-[#fff] text-lg font-bold py-2 px-8 rounded-xl bg-[#055a32]"
    >
      Register your account
    </button>
  </div>
</template>
<script setup>
import { reactive } from "vue";

const formData = reactive({
  name: "",
  email: "",
  password: "",
  confirmPassword: "",
});

const handleRegister = () => {
  console.log(formData);
};
</script>
