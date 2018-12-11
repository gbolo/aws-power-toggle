import MetadataApi from '../services/api/Metadata';
import EnvironmentsApi from '../services/api/Environments';
import InstancesApi from '../services/api/Instances';

export const fetchVersion = ({ commit }) => {
  MetadataApi.getVersion()
    .then(data => commit('setVersion', data.version))
    .catch(e => commit('setError', e.response.data.error));
};

export const fetchAllEnvironmentsSummary = ({ commit }) => {
  EnvironmentsApi.fetchAllEnvironmentsSummary()
    .then(data => commit('setEnvironments', data))
    .catch(e => commit('setError', e.response.data.error));
};

export const fetchAllEnvironmentsDetails = ({ commit }) => {
  EnvironmentsApi.fetchAllEnvironmentsDetails()
    .then(data => commit('setEnvironments', data))
    .catch(e => commit('setError', e.response.data.error));
};

export const fetchEnvironmentDetails = ({ commit }, id) => {
  EnvironmentsApi.fetchEnvironmentDetails(id)
    .then(data => commit('setEnvironment', { id, data }))
    .catch(e => commit('setError', e.response.data.error));
};

export const startEnvironment = ({ dispatch, commit }, id) => {
  EnvironmentsApi.startEnvironment(id)
    .then(() => dispatch('fetchEnvironmentDetails', id))
    .catch(e => commit('setError', e.response.data.error));
};

export const stopEnvironment = ({ dispatch, commit }, id) => {
  EnvironmentsApi.stopEnvironment(id)
    .then(() => dispatch('fetchEnvironmentDetails', id))
    .catch(e => commit('setError', e.response.data.error));
};

export const startInstance = ({ dispatch, commit }, { id, envId }) => {
  InstancesApi.startInstance(id)
    .then(() => dispatch('fetchEnvironmentDetails', envId))
    .catch(e => commit('setError', e.response.data.error));
};

export const stopInstance = ({ dispatch, commit }, { id, envId }) => {
  InstancesApi.stopInstance(id)
    .then(() => dispatch('fetchEnvironmentDetails', envId))
    .catch(e => commit('setError', e.response.data.error));
};

export const clearError = ({ commit }) => {
  commit('setError', '');
};
