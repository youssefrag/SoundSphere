import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { getStorage, ref as storageRef, deleteObject } from "firebase/storage";

import api from "@/api";
import { useUserStore } from "./user";

export const useMusicStore = defineStore("music", () => {
  const allArtists = ref([]);
  const loading = ref(false);
  const error = ref(null);
  const searchTerm = ref("");

  const userStore = useUserStore();

  const userSongs = computed(() => {
    const bucket = allArtists.value.find((a) => a.email === userStore.email);
    return bucket ? bucket.songs : [];
  });

  const allSongs = computed(() =>
    allArtists.value.flatMap((artist) =>
      artist.songs.map((song) => ({
        ...song,
        artistEmail: artist.email,
        artistName: artist.name,
      }))
    )
  );

  const filteredSongs = computed(() => {
    if (!searchTerm.value) return [];
    const term = searchTerm.value.toLowerCase();
    return allSongs.value.filter(
      (song) =>
        song.name.toLowerCase().startsWith(term) ||
        song.genre.toLowerCase().startsWith(term) ||
        song.artistName.toLowerCase().startsWith(term)
    );
  });

  function clearSearchTerm() {
    searchTerm.value = "";
  }

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
    filteredSongs,
    searchTerm,
    loading,
    error,
    saveSong,
    fetchAllArtists,
    deleteSong,
    editSong,
    fetchSongDetails,
    clearSearchTerm,
  };
});
