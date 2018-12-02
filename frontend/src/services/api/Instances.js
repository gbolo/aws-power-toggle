import axios from 'axios';

export default {
  startInstance(id) {
    return axios.post(`/api/v1/instance/${id}/start`).then(response => response.data);
  },
  stopInstance(id) {
    return axios.post(`/api/v1/instance/${id}/stop`).then(response => response.data);
  },
};
