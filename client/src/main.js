import "./assets/main.css";

import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

import "./assets/base.css";
import "./assets/main.css";

import FontAwesomeIcon from "./utilities/fontawesome";

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.component("FontAwesomeIcon", FontAwesomeIcon);

app.mount("#app");
