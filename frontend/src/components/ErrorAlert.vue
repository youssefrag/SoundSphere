<template>
  <div
    v-if="message"
    class="flex items-center bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4"
    role="alert"
  >
    <span class="flex-grow">{{ message }}</span>
    <button
      @click="$emit('clear')"
      class="ml-4 text-red-700 font-bold focus:outline-none"
    >
      &times;
    </button>
  </div>
</template>

<script setup>
const props = defineProps({
  message: {
    type: String,
    default: "",
  },
});
</script>
