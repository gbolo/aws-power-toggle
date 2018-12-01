import axios from 'axios';

export default {
  fetchAllEnvironmentsDetails() {
    return axios.get('/api/v1/env/details').then(response => response.data);
  },
  fetchAllEnvironmentsSummary() {
    return axios.get('/api/v1/env/summary').then(response => response.data);
  },

  fetchEnvironmentDetails(id) {
    return axios.get(`/api/v1/env/${id}/details`).then(response => response.data);
  },
  fetchEnvironmentSummary(id) {
    return axios.get(`/api/v1/env/${id}/summary`).then(response => response.data);
  },

  startEnvironment(id) {
    return axios.post(`/api/v1/env/${id}/start`).then(response => response.data);
  },
  stopEnvironment(id) {
    return axios.post(`/api/v1/env/${id}/stop`).then(response => response.data);
  },
};
