import Vue from 'vue';
import registerFontAwesome from '../fa.config';
import App from './App.vue';

registerFontAwesome(Vue);

Vue.config.productionTip = false;

new Vue({
  render: h => h(App),
}).$mount('#app');
