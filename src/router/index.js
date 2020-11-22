import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const DEFAULT_TITLE = 'William Wernert';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: `${DEFAULT_TITLE} - Home`
    },
  },
  {
    path: '/about',
    name: 'About',
    meta: {
      title: `${DEFAULT_TITLE} - About`
    },
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/resume',
    name: 'Resume',
    meta: {
      title: `${DEFAULT_TITLE} - Resume`
    },
    component: () => import(/* webpackChunkName: "resume" */ '../views/Resume.vue'),
  },
];

const router = new VueRouter({
  routes,
  mode: 'history'
});

router.afterEach((to, from) => {
    // Use next tick to handle router history correctly
    // see: https://github.com/vuejs/vue-router/issues/914#issuecomment-384477609
    Vue.nextTick(() => {
        document.title = to.meta.title || DEFAULT_TITLE;
    });
});

export default router;
