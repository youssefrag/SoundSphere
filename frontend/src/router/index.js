import { useUserStore } from "@/stores/user";
import { createRouter, createWebHistory } from "vue-router";

const Home = () => import("@/views/Home.vue");
const Song = () => import("@/views/Song.vue");
const Profile = () => import("@/views/Profile.vue");

const routes = [
  {
    name: "home",
    path: "/",
    component: Home,
  },
  {
    name: "song",
    path: "/song/:id",
    component: Song,
  },
  {
    name: "ptofile",
    path: "/profile",
    component: Profile,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach(async (to, fromJSON, next) => {
  const userStore = useUserStore();

  if (!to.meta.requiresAuth) {
    return next();
  }

  if (!userStore.accessToken) {
    await userStore.refresh();
  }

  const token = userStore.accessToken;

  if (!token) {
    return next({ name: "home" });
  }

  next();
});

export default router;
