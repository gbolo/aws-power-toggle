import http from './HTTP';

export default {
  startInstance(id) {
    return http.post(`/instance/${id}/start`).then((response) => response.data);
  },
  stopInstance(id) {
    return http.post(`/instance/${id}/stop`).then((response) => response.data);
  },
};
