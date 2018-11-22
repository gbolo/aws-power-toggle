import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;

Vue.filter(
  'capitalize',
  value => (value
    ? value
      .toString()
      .charAt(0)
      .toUpperCase() + value.slice(1)
    : ''),
);

new Vue({
  render: h => h(App),
}).$mount('#app');
