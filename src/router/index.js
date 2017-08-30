import Vue from 'vue';
import Router from 'vue-router';
import GettingMarried from '@/components/GettingMarried';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'GettingMarried',
      component: GettingMarried,
    },
  ],
});
