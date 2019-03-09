import MetadataApi from '../services/api/Metadata';
import EnvironmentsApi from '../services/api/Environments';
import InstancesApi from '../services/api/Instances';

const errMsg = (e) => {
  let msg = 'Unexpected error occured';
  if (e && e.response && e.response.data) {
    msg = e.response.data.error || msg;
  }
  return msg;
};

export const fetchVersion = ({ commit }) => {
  MetadataApi.getVersion()
    .then(data => commit('setVersion', data.version))
    .catch(e => commit('setError', errMsg(e)));
};

export const fetchAllEnvironmentsSummary = ({ commit }) => {
  EnvironmentsApi.fetchAllEnvironmentsSummary()
    .then(data => commit('setEnvironments', data))
    .catch(e => commit('setError', errMsg(e)));
};

export const fetchAllEnvironmentsDetails = ({ commit }) => {
  commit('setIsLoading', true);
  EnvironmentsApi.fetchAllEnvironmentsDetails()
    .then(data => commit('setEnvironments', data))
    .catch(e => commit('setError', errMsg(e)))
    .finally(() => commit('setIsLoading', false));
};

export const refresh = ({ commit }) => {
  EnvironmentsApi.fetchAllEnvironmentsDetails()
    .then(data => commit('setEnvironments', data))
    .catch(e => commit('setError', e.response.data.error));
};

export const fetchEnvironmentDetails = ({ commit }, id) => {
  commit('setEnvironmentLoading', { id, flag: true });
  EnvironmentsApi.fetchEnvironmentDetails(id)
    .then(data => commit('setEnvironment', { id, data }))
    .catch(e => commit('setError', errMsg(e)))
    .finally(() => commit('setEnvironmentLoading', { id, flag: false }));
};

export const startEnvironment = ({ dispatch, commit }, id) => {
  commit('setEnvironmentLoading', { id, flag: true });
  EnvironmentsApi.startEnvironment(id)
    .then(() => dispatch('fetchEnvironmentDetails', id))
    .catch((e) => {
      commit('setError', errMsg(e));
      commit('setEnvironmentLoading', { id, flag: false });
    });
};

export const stopEnvironment = ({ dispatch, commit }, id) => {
  commit('setEnvironmentLoading', { id, flag: true });
  EnvironmentsApi.stopEnvironment(id)
    .then(() => dispatch('fetchEnvironmentDetails', id))
    .catch((e) => {
      commit('setError', errMsg(e));
      commit('setEnvironmentLoading', { id, flag: false });
    });
};

export const startInstance = ({ commit }, { id, envId }) => {
  commit('setInstanceLoading', { id, flag: true });
  InstancesApi.startInstance(id)
    .then(() => commit('setInstanceStateStatus', { id, status: 'running', envId }))
    .catch((e) => {
      commit('setError', errMsg(e));
      commit('setInstanceStateStatus', { id, status: 'stopped', envId });
    })
    .finally(() => {
      commit('setInstanceLoading', { id, flag: false });
    });
};

export const stopInstance = ({ commit }, { id, envId }) => {
  commit('setInstanceLoading', { id, flag: true });
  InstancesApi.stopInstance(id)
    .then(() => commit('setInstanceStateStatus', { id, status: 'stopped', envId }))
    .catch((e) => {
      commit('setError', errMsg(e));
      commit('setInstanceStateStatus', { id, status: 'running', envId });
    })
    .finally(() => {
      commit('setInstanceLoading', { id, flag: false });
    });
};

export const clearError = ({ commit }) => {
  commit('setError', '');
};
