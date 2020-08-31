import Vue from 'vue';
import VModal from 'vue-js-modal';
import store from './store';
import App from './App.vue';

Vue.use(VModal, { dialog: true });
Vue.config.productionTip = false;

new Vue({
  store,
  render: (h) => h(App),
}).$mount('#app');
