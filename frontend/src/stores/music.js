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

  async function saveSong(name, genre, artist, songUrl, duration) {
    const response = await api.post("/saveSong", {
      name,
      genre,
      artistEmail: artist,
      songUrl,
      duration,
    });

    console.log(response);
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

    console.log(allArtists.value);
  }

  fetchAllArtists();

  return {
    allArtists,
    userSongs,
    loading,
    error,
    saveSong,
    fetchAllArtists,
  };
});
