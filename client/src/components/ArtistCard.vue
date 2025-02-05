<template class="rounded-2xl">
  <div
    class="h-[488px] w-[650px] bg-center bg-cover rounded-2xl"
    :class="{
      'bg-[url(@/assets/images/artist-bg1.png)]': index % 3 == 0,
      'bg-[url(@/assets/images/artist-bg2.png)]': index % 3 == 1,
      'bg-[url(@/assets/images/artist-bg3.png)]': index % 3 == 2,
    }"
  >
    <div class="h-[100%] w-[100%] bg-[#000] opacity-85 rounded-2xl"></div>
  </div>
</template>

<script setup>
const { artist, index } = defineProps(["index", "artist"]);

console.log({ index, artist });
console.log(index % 3);
</script>
