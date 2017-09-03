<template>
  <div>
    <div id="top" class="page">
      <img src="../assets/getting_married.png"></img>
    </div>
    <div class="page">
      <img src="../assets/invitation.png"></img>
    </div>
    <gallery-view id="photo"></gallery-view>
    <navigation-view id="bus"></navigation-view>
    <guestbook-view></guestbook-view>
  </div>
</template>

<script>
import Gallery from '../components/Gallery';
import Navigation from '../components/Navigation';
import Guestbook from '../components/Guestbook';

export default {
  name: 'getting-married',
  components: {
    'gallery-view': Gallery,
    'navigation-view': Navigation,
    'guestbook-view': Guestbook,
  },
  data() {
    return {
    };
  },
};
</script>

