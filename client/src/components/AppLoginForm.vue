<template>
  <div class="flex flex-col gap-6">
    <div>
      <label for="email" class="block mb-2 text-lg font-bold text-[#0E0E0F]"
        >Email</label
      >
      <input
        type="email"
        id="first_name"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="john@example.com"
        required
      />
    </div>
    <div>
      <label for="email" class="block mb-2 text-lg font-bold text-[#0E0E0F]"
        >Password</label
      >
      <input
        type="password"
        id="first_name"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="password"
        required
      />
    </div>
    <button
      class="text-[#fff] text-lg font-bold py-2 px-8 rounded-xl bg-[#4636FF]"
    >
      Login to your account
    </button>
  </div>
</template>

<script setup></script>
