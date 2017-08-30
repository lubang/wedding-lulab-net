<template>
  <v-app id="app">
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
      drawer: false,
      items: [
        { title: 'Home', icon: 'dashboard' },
        { title: 'About', icon: 'question_answer' },
      ],
      mini: false,
      right: null,
    };
  },
};
</script>
