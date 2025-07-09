import { defineStore } from "pinia";
import api from "@/api";

export const useMusicStore = defineStore("music", {
  state: () => ({
    music: [
      {
        id: "1",
        name: "Alex Waves",
        pictureUrl: "artist1.jpeg",
        songs: [
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
            artistId: "1",
          },
        ],
      },
      {
        name: "Salena Julie",
        pictureUrl: "/src/assets/images/artist2.jpeg",
        songs: [
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
        ],
      },
      {
        name: "David Malcom",
        pictureUrl: "/src/assets/images/artist3.jpeg",
        songs: [
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
          {
            name: "Love Me Like You Do",
            duration: "6:32",
            genre: "Hiphop",
            id: "123",
          },
        ],
      },
    ],
  }),
  actions: {
    async saveSong(name, genre, artist, songUrl) {
      const response = await api.post("/saveSong", {
        name,
        genre,
        artistEmail: artist,
        songUrl,
      });

      console.log(response);
    },
  },
});
