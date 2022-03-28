import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import axios from "axios";
import VueAxios from "vue-axios";

//Fontawesome
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faPhone,
  faCalendarCheck,
  faCogs,
  faBuilding,
  faUtensils,
  faClock,
  faTimes,
} from "@fortawesome/free-solid-svg-icons";

library.add(faPhone);
library.add(faCalendarCheck);
library.add(faCogs);
library.add(faBuilding);
library.add(faUtensils);
library.add(faClock);
library.add(faTimes);
const app = createApp(App);
axios.defaults.baseURL = "";
app.use(router);
app.use(VueAxios, axios);
app.component("font-awesome-icon", FontAwesomeIcon).mount("#app");
