import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import axios from "axios";
import VueAxios from "vue-axios";

//Fontawesome
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library, type Library } from "@fortawesome/fontawesome-svg-core";
import {
  faPhone,
  faCalendarCheck,
  faCogs,
  faBuilding,
  faUtensils,
  faClock,
  faTimes,
  faUser,
  faSearch,
  faCaretDown,
} from "@fortawesome/free-solid-svg-icons";

const addIcons = (lib: Library) => {
  lib.add(faPhone);
  lib.add(faCalendarCheck);
  lib.add(faCogs);
  lib.add(faBuilding);
  lib.add(faUtensils);
  lib.add(faClock);
  lib.add(faTimes);
  lib.add(faUser);
  lib.add(faSearch);
  lib.add(faCaretDown);
};

addIcons(library);
const app = createApp(App);
axios.defaults.baseURL = "";
app.use(router);
app.use(VueAxios, axios);
app.component("font-awesome-icon", FontAwesomeIcon).mount("#app");
