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
      dark fab fixed bottom left small
      v-model="fab"
    >
      <v-icon>fa-heart</v-icon>
    </v-btn>

    <v-btn
      id="oneshot-menu-story"
      class="cyan"
      dark fab fixed bottom left small
    >
      <v-icon>fa-story</v-icon>
    </v-btn>
    <v-btn
      id="oneshot-menu-bus"
      class="light-blue"
      dark fab fixed bottom left small
    >
      <v-icon>fa-invitation</v-icon>
    </v-btn>
    <v-btn
      id="oneshot-menu-gallery"
      class="red"
      dark fab fixed bottom left small
    >
      <v-icon>fa-delete</v-icon>
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
};
</script>

<style>
#oneshot-menu {
  z-index: 100;
}
#oneshot-menu-bus {
  margin-left: 30px;
  margin-bottom: 70px;
}
#oneshot-menu-invitation {
  margin-left: 70px;
  margin-bottom: 30px;
}
#oneshot-menu-gallery {
  margin-left: 90px;
  margin-bottom: 0px;
}
</style>
