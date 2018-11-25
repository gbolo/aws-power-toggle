import axios from 'axios';

export default {
  getAllEnvironmentsDetails() {
    return axios.get('/api/v1/env/details').then(response => response.data);
  },
  getAllEnvironmentsSummary() {
    return axios.get('/api/v1/env/summary').then(response => response.data);
  },

  getEnvironmentDetails(id) {
    return axios.get(`/api/v1/env/${id}/details`).then(response => response.data);
  },
  getEnvironmentSummary(id) {
    return axios.get(`/api/v1/env/${id}/summary`).then(response => response.data);
  },

  startEnvironment(id) {
    return axios.post(`/api/v1/env/${id}/start`).then(response => response.data);
  },
  stopEnvironment(id) {
    return axios.post(`/api/v1/env/${id}/stop`).then(response => response.data);
  },

};
