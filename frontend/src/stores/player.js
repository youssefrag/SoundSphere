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

  const playing = computed(
    () => sound.value instanceof Howl && sound.value.playing()
  );

  function newSong(song) {
    // clean up old sound
    if (sound.value instanceof Howl) sound.value.unload();

    currentSong.value = song;
    sound.value = new Howl({
      src: [song.songUrl],
      html5: true,
      onload: () => {
        // once metadata is ready, capture total length
        duration.value = formatDuration(sound.value.duration());
      },
    });

    sound.value.play();

    sound.value.on("play", () => {
      requestAnimationFrame(progress);
    });
  }

  function togglePlay(song) {
    if (
      currentSong.value.id === song.id &&
      sound.value instanceof Howl &&
      sound.value.playing()
    ) {
      sound.value.pause();
    } else if (
      currentSong.value.id === song.id &&
      sound.value instanceof Howl
    ) {
      sound.value.play();
      requestAnimationFrame(progress);
    } else {
      newSong(song);
    }
  }

  function progress() {
    if (!(sound.value instanceof Howl)) return;

    const pos = sound.value.seek();
    const total = sound.value.duration();

    seek.value = formatDuration(pos);
    duration.value = formatDuration(total);
    playerProgress.value = `${(pos / total) * 100}%`;

    // loop while playing
    if (sound.value.playing()) {
      requestAnimationFrame(progress);
    }
  }

  function updateSeek(event) {
    if (!(sound.value instanceof Howl)) return;

    const { x, width } = event.currentTarget.getBoundingClientRect();
    const clickX = event.clientX - x;
    const pct = clickX / width;
    const seconds = sound.value.duration() * pct;

    const wasPlaying = sound.value.playing();

    sound.value.seek(seconds);

    if (wasPlaying) {
      sound.value.play();
      requestAnimationFrame(progress);
    } else {
      progress();
    }
  }

  return {
    currentSong,
    playing,
    seek,
    duration,
    playerProgress,
    newSong,
    togglePlay,
    updateSeek,
  };
});
