import { defineStore } from "pinia";
import api from "@/api";

export const useUserStore = defineStore("user", {
  state: () => ({
    accessToken: null,
    name: null,
    email: null,
  }),
  actions: {
    clearSession() {
      this.accessToken = null;
      this.name = null;
      this.email = null;
      delete api.defaults.headers.common.Authorization;
    },

    async register(name, email, password) {
      const response = await api.post("/signup", { name, email, password });

      console.log(response);

      if (response.status === 201) {
        this.login(email, password);
      }
    },

    async login(email, password) {
      const { data } = await api.post("/login", { email, password });

      this.accessToken = data.token;
      this.email = data.user.email;
      this.name = data.user.name;

      api.defaults.headers.common.Authorization = `Bearer ${this.accessToken}`;
    },
    async refresh() {
      try {
        const { data } = await api.post("/refresh");

        this.accessToken = data.token ?? data.access_token;
        this.email = data.user.email;
        this.name = data.user.name;

        api.defaults.headers.common.Authorization = `Bearer ${this.accessToken}`;
      } catch (error) {
        this.clearSession();
      }
    },
  },
});
