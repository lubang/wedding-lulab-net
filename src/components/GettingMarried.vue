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
  </div>
</template>

<script>
import Gallery from '../components/Gallery';
import Navigation from '../components/Navigation';

export default {
  name: 'getting-married',
  components: {
    'gallery-view': Gallery,
    'navigation-view': Navigation,
  },
  data() {
    return {
    };
  },
};
</script>

