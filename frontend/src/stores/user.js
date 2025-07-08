import { defineStore } from "pinia";
import api from "@/api";

export const useUserStore = defineStore("user", {
  state: () => ({
    accessToken: null,
    name: null,
    email: null,
    imageUrl: null,
  }),
  actions: {
    clearSession() {
      this.accessToken = null;
      this.name = null;
      this.email = null;
      this.imageUrl = null;
      delete api.defaults.headers.common.Authorization;
    },

    async register(name, email, password, avatarUrl) {
      try {
        const response = await api.post("/signup", {
          name,
          email,
          password,
          imageUrl: avatarUrl,
        });

        if (response.status === 201) {
          this.login(email, password);
        }
      } catch (error) {
        if (error.status) {
          alert("Email already in use");
        }
      }
    },

    async login(email, password) {
      const { data } = await api.post("/login", { email, password });

      console.log(data);

      this.accessToken = data.token;
      this.email = data.user.email;
      this.name = data.user.name;
      this.imageUrl = data.user.imageUrl;

      api.defaults.headers.common.Authorization = `Bearer ${this.accessToken}`;
    },

    async logout() {
      console.log("reached here");
      try {
        await api.post("/logout", {}, { withCredentials: true });
      } catch (error) {
        console.warn("could not revode refresh token", error);
      } finally {
        this.clearSession();
      }
    },

    async refresh() {
      try {
        const { data } = await api.post("/refresh");

        this.accessToken = data.token ?? data.access_token;
        this.email = data.user.email;
        this.name = data.user.name;
        this.imageUrl = data.user.imageUrl;

        console.log(data);

        api.defaults.headers.common.Authorization = `Bearer ${this.accessToken}`;
      } catch (error) {
        this.clearSession();
      }
    },
  },
});
