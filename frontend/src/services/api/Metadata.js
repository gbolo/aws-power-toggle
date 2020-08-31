import http from './HTTP';

export default {
  getVersion() {
    return http.get('/version').then((response) => response.data);
  },
  getConfig() {
    return http.get('/config').then((response) => response.data);
  },
  refresh() {
    return http.post('/refresh').then((response) => response.data);
  },
};
