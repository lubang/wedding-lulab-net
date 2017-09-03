<template>
  <v-speed-dial
    v-model="fab"
    :top="top"
    :bottom="bottom"
    :right="right"
    :left="left"
    :direction="direction"
    :hover="hover"
    :transition="transition"
  >
    <v-btn
      id="oneshot-menu"
      class="deep-orange"
      slot="activator"
      dark fab fixed bottom left
      v-model="fab"
    >
      <v-icon>fa-heart</v-icon>
    </v-btn>
    <v-btn
      id="oneshot-menu-story"
      class="cyan"
      dark fab fixed bottom left
      v-on:click="gotoTop()"
    >
      <v-icon>fa-vcard-o</v-icon>
    </v-btn>
    <v-btn
      id="oneshot-menu-bus"
      class="light-blue"
      dark fab fixed bottom left
      v-on:click="gotoBus()"
    >
      <v-icon>fa-bus</v-icon>
    </v-btn>
    <v-btn
      id="oneshot-menu-gallery"
      class="red"
      dark fab fixed bottom left
      v-on:click="gotoPhoto()"
    >
      <v-icon>fa-photo</v-icon>
    </v-btn>
  </v-speed-dial>
</template>

<script>
export default {
  data: () => ({
    direction: 'top',
    fab: false,
    fling: false,
    hover: false,
    tabs: null,
    top: false,
    right: true,
    bottom: true,
    left: false,
    transition: 'slide-y-reverse-transition',
  }),

  watch: {
    top(val) {
      this.bottom = !val;
    },
    right(val) {
      this.left = !val;
    },
    bottom(val) {
      this.top = !val;
    },
    left(val) {
      this.right = !val;
    },
  },

  computed: {
    activeFab() {
      switch (this.tabs) {
        case 'one': return { class: 'purple', icon: 'account_circle' };
        case 'two': return { class: 'red', icon: 'edit' };
        case 'three': return { class: 'green', icon: 'keyboard_arrow_up' };
        default: return {};
      }
    },
  },

  methods: {
    gotoTop() {
      document.getElementById('top').scrollIntoView();
    },
    gotoPhoto() {
      document.getElementById('photo').scrollIntoView();
    },
    gotoBus() {
      document.getElementById('bus').scrollIntoView();
    },
  },
};
</script>

<style scoped>
#oneshot-menu {
  z-index: 100;
}
#oneshot-menu-story {
  margin-bottom: 250px;
}
#oneshot-menu-bus {
  margin-bottom: 170px;
}
#oneshot-menu-gallery {
  margin-bottom: 90px;
}
</style>
