import { createRouter, createWebHistory } from "vue-router";

const Home = () => import("@/views/Home.vue");
const Song = () => import("@/views/Song.vue");

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
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  // linkExactActiveClass: "text-[#FBBF00]",
});

export default router;
