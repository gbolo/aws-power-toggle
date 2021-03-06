import Vue from 'vue';
import Vuex from 'vuex';
import VTooltip from 'v-tooltip';

import * as mutations from './mutations';
import * as actions from './actions';
import * as getters from './getters';

Vue.use(VTooltip);
Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    environments: [],
    version: '',
    error: '',
    environmentsLoading: {},
    instancesLoading: {},
    isLoading: false,
  },
  mutations,
  actions,
  getters,
});

export default store;
