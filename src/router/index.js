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
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
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
