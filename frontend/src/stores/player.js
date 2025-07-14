import { defineStore } from "pinia";
import { Howl } from "howler";

import { ref, computed } from "vue";

import { formatDuration } from "@/utilities/helpers";

export const usePlayerStore = defineStore("player", () => {
  const currentSong = ref({});
  const sound = ref(null);
  const seek = ref("00:00");
  const duration = ref("00:00");
  const playerProgress = ref("0%");

  const playing = computed(() => {
    return sound.value instanceof Howl && sound.value.playing();
  });

  async function newSong(song) {
    if (sound.value instanceof Howl) {
      sound.value.unload();
    }

    currentSong.value = song;

    sound.value = new Howl({
      src: [song.songUrl],
      html5: true,
    });

    sound.value.play();

    sound.value.on("play", () => {
      requestAnimationFrame(progress);
    });
  }

  function togglePlay(song) {
    if (
      currentSong.value?.id === song.id &&
      sound.value instanceof Howl &&
      sound.value.playing()
    ) {
      sound.value.pause();
    } else if (
      currentSong.value?.id === song.id &&
      sound.value instanceof Howl
    ) {
      sound.value.play();
    } else {
      newSong(song);
    }
  }

  function progress() {
    if (!sound.value) return;

    seek.value = formatDuration(sound.value.seek());
    duration.value = formatDuration(sound.value.duration());
    playerProgress.value = `${
      (sound.value.seek() / sound.value.duration()) * 100
    }%`;

    if (sound.value.playing()) {
      requestAnimationFrame(progress);
    }
  }

  function updateSeek(event) {
    if (!sound.value) return;

    const { x, width } = event.currentTarget.getBoundingClientRect();
    const clickX = event.clientX - x;
    const percentage = clickX / width;
    const seconds = sound.value.duration() * percentage;

    sound.value.seek(seconds);
    sound.value.once("seek", progress);
  }

  return {
    currentSong,
    sound,
    playing,
    seek,
    newSong,
    togglePlay,
    progress,
    updateSeek,
  };
});
