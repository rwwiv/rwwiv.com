import Vue from 'vue';
import VueHead from 'vue-head'
import './plugins/fontawesome'
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/tailwind.css'

import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { FontAwesomeIcon, FontAwesomeLayers } from '@fortawesome/vue-fontawesome'

library.add(fas)
library.add(far)
Vue.component('fa', FontAwesomeIcon)
Vue.component('fa-layers', FontAwesomeLayers)

Vue.config.productionTip = false;

const EventBus = new Vue();
 
Vue.use(VueHead)

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');

export { EventBus }
