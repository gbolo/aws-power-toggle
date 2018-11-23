import axios from 'axios';

export default {
  getVersion() {
    return axios.get('/api/version').then(response => response.data);
  },
};
