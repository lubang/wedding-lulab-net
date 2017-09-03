<template>
  <v-app class="app">
    <router-view></router-view>
    <footer-view></footer-view>
    <oneshot-menu></oneshot-menu>
  </v-app>
</template>

<script>
import Footer from './components/Footer';
import OneshotMenu from './components/OneshotMenu';

export default {
  name: 'app',
  components: {
    'footer-view': Footer,
    'oneshot-menu': OneshotMenu,
  },
  data() {
    return {
    };
  },
};
</script>

<style>
.app {
  max-width: 750px;
  font-family: 'NanumMyeongjo';
}
.page {
  min-height: 100vh;
  background-color: white;
}
img {
  max-width: 100%;
  max-height: 100%;
}
</style>
