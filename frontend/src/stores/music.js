import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { getStorage, ref as storageRef, deleteObject } from "firebase/storage";

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
    error.value = null;

    try {
      const { data } = await api.delete(`songs/${songId}`, {
        headers: { Authorization: `Bearer ${userStore.accessToken}` },
      });

      const { storage_path: storagePath } = data;

      const storage = getStorage();
      const fileRef = storageRef(storage, storagePath);
      await deleteObject(fileRef);
    } catch (err) {
      error.value = err;
    } finally {
      await fetchAllArtists();
    }
  }

  async function editSong(songId, name, genre) {
    try {
      await api.put(`/songs/${songId}`, { name, genre });
    } catch (err) {
      error.value = err;
    } finally {
      await fetchAllArtists();
    }
  }

  async function fetchSongDetails(songId) {
    try {
      const { data } = await api.get(`/song-details/${songId}`);

      return data;
    } catch (error) {
      throw Error("Error fetching song details.");
    }
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
    editSong,
    fetchSongDetails,
  };
});
