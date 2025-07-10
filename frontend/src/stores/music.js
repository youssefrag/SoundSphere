import { defineStore } from "pinia";
import { ref, computed } from "vue";
import api from "@/api";
import { useUserStore } from "./user";

export const useMusicStore = defineStore("music", () => {
  const allArtists = ref([]);
  const loading = ref(false);
  const error = ref(null);

  const userStore = useUserStore();

  const userSongs = computed(() => {
    const bucket = allArtists.value.find((a) => a.email === userStore.email);
    return bucket ? bucket.songs : [];
  });

  async function saveSong(name, genre, artist, songUrl, duration, storagePath) {
    await api.post("/saveSong", {
      name,
      genre,
      artistEmail: artist,
      songUrl,
      duration,
      storagePath,
    });

    await fetchAllArtists();
  }

  async function fetchAllArtists() {
    loading.value = true;
    error.value = null;

    try {
      const { data } = await api.get("/allSongs");
      allArtists.value = data;
    } catch (err) {
      error.value = err;
    } finally {
      loading.value = false;
    }
  }

  async function deleteSong(songId) {
    try {
      const { data } = await api.delete(`songs/${songId}`);
    } catch (err) {}
  }

  fetchAllArtists();

  return {
    allArtists,
    userSongs,
    loading,
    error,
    saveSong,
    fetchAllArtists,
    deleteSong,
  };
});
