import axios from 'axios';

export default {
  getVersion() {
    return axios.get('/api/v1/version').then(response => response.data);
  },
  getConfig() {
    return axios.get('/api/v1/config').then(response => response.data);
  },
  refresh() {
    return axios.post('/api/v1/refresh').then(response => response.data);
  },
};
