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

/* eslint-disable */
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

ga('create', 'UA-43684036-4', 'auto');
ga('send', 'pageview');
/* eslint-enable */
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
