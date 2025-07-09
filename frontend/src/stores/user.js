// src/stores/user.js
import { defineStore } from "pinia";
import { ref } from "vue";
import api from "@/api";

export const useUserStore = defineStore("user", () => {
  const accessToken = ref(null);
  const name = ref(null);
  const email = ref(null);
  const imageUrl = ref(null);

  function clearSession() {
    accessToken.value = null;
    name.value = null;
    email.value = null;
    imageUrl.value = null;
    delete api.defaults.headers.common.Authorization;
  }

  async function register(userName, userEmail, password, avatarUrl) {
    try {
      const res = await api.post("/signup", {
        name: userName,
        email: userEmail,
        password,
        imageUrl: avatarUrl,
      });
      if (res.status === 201) {
        await login(userEmail, password);
      }
    } catch (err) {
      alert(
        err.response?.status === 409
          ? "Email already in use"
          : "Registration failed"
      );
    }
  }

  async function login(userEmail, password) {
    const { data } = await api.post("/login", { email: userEmail, password });
    accessToken.value = data.token;
    email.value = data.user.email;
    name.value = data.user.name;
    imageUrl.value = data.user.imageUrl;
    api.defaults.headers.common.Authorization = `Bearer ${data.token}`;
  }

  async function logout() {
    try {
      await api.post("/logout", {}, { withCredentials: true });
    } catch (err) {
      console.warn("could not revoke refresh token", err);
    } finally {
      clearSession();
      window.location.reload(false);
    }
  }

  async function refresh() {
    try {
      const { data } = await api.post("/refresh");
      accessToken.value = data.token ?? data.access_token;
      email.value = data.user.email;
      name.value = data.user.name;
      imageUrl.value = data.user.imageUrl;
      api.defaults.headers.common.Authorization = `Bearer ${accessToken.value}`;
    } catch {
      clearSession();
    }
  }

  return {
    accessToken,
    name,
    email,
    imageUrl,

    clearSession,
    register,
    login,
    logout,
    refresh,
  };
});
