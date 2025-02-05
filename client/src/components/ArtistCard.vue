<template class="rounded-2xl">
  <div
    class="h-[488px] w-[650px] bg-center bg-cover rounded-2xl"
    :class="{
      'bg-[url(@/assets/images/artist-bg1.avif)]': index % 3 == 0,
      'bg-[url(@/assets/images/artist-bg2.avif)]': index % 3 == 1,
      'bg-[url(@/assets/images/artist-bg3.avif)]': index % 3 == 2,
    }"
  >
    <div
      class="h-[100%] w-[100%] bg-[#000] opacity-85 rounded-2xl overflow-x-auto py-[30px] px-[24px]"
    >
      <div>
        <!-- <img :src="require(artist.pictureUrl)" alt="artist-picture" /> -->
      </div>
    </div>
  </div>
</template>

<script setup>
const { artist, index } = defineProps(["index", "artist"]);
// import require from "vue";

const imagePath = () => {
  if (artist.pictureUrl !== "") {
    return require(artist.pictureUrl);
  }
};
</script>
