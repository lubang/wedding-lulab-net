import Vue from 'vue';

export default {
  save(reservation) {
    return Vue.resource('/api/reservations{/id}').save(reservation);
  },
};
