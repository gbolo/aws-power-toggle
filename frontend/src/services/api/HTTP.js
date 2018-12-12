import axios from 'axios';

const instance = axios.create({
  baseURL: '/api/v1',
});

export default instance;
