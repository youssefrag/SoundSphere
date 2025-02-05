import { defineStore } from "pinia";

export default defineStore("music", {
  state: () => ({
    music: [
      {
        name: "Alex Waves",
        pictureUrl: "@/assets/images/artist1.jpeg",
        songs: [
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
        ],
      },
      {
        name: "Salena Julie",
        pictureUrl: "@/assets/images/artist2.jpeg",
        songs: [
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
        ],
      },
      {
        name: "David Malcom",
        pictureUrl: "@/assets/images/artist3.jpeg",
        songs: [
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
          { name: "Love Me Like You Do", duration: "6:32", genre: "Hiphop" },
        ],
      },
    ],
  }),
});
