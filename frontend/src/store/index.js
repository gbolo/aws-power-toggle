import Vue from 'vue';
import Vuex from 'vuex';

import * as mutations from './mutations';
import * as actions from './actions';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    environments: [],
    version: '',
    error: '',
  },
  mutations,
  actions,
});

export default store;
