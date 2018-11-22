import axios from 'axios';

export default {
  getEnvironments() {
    return axios.get('/api/env/summary').then(response => response.data);
  },
  startEnvironment(envName) {
    return axios.post(`/api/env/startup/${envName}`);
  },
  stopEnvironment(envName) {
    return axios.post(`/api/env/powerdown/${envName}`);
  },
};
