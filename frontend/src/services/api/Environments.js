import axios from 'axios';

export default {
  getEnvironments() {
    return axios.get('/api/env/summary').then(response => response.data);
  },
  startEnvironment(envName) {
    return axios.post(`/api/env/startup/${envName}`).then(response => response.data);
  },
  stopEnvironment(envName) {
    return axios.post(`/api/env/powerdown/${envName}`).then(response => response.data);
  },
};
