import http from './HTTP';

export default {
  fetchAllEnvironmentsDetails() {
    return http.get('/env/details').then(response => response.data);
  },
  fetchAllEnvironmentsSummary() {
    return http.get('/env/summary').then(response => response.data);
  },

  fetchEnvironmentDetails(id) {
    return http.get(`/env/${id}/details`).then(response => response.data);
  },
  fetchEnvironmentSummary(id) {
    return http.get(`/env/${id}/summary`).then(response => response.data);
  },

  startEnvironment(id) {
    return http.post(`/env/${id}/start`).then(response => response.data);
  },
  stopEnvironment(id) {
    return http.post(`/env/${id}/stop`).then(response => response.data);
  },
};
