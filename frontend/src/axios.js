import axios from "axios";
import { useRouter } from "vue-router";

const router = useRouter();
const token = localStorage.getItem('token');

const axiosClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  withCredentials: (token) ? true : false,
  withXSRFToken: true,
})
axiosClient.interceptors.request.use(config => {
  config.headers['Accept'] = '*/*';
  config.headers['Content-Type'] = 'application/json';
  config.headers['Access-Control-Allow-Origin'] = '*';
  config.headers['Access-Control-Allow-Methods'] = 'GET, POST, PUT, DELETE, OPTIONS';
  config.headers['Access-Control-Allow-Headers'] = 'Origin, Content-Type, Accept, Authorization';
  config.headers['Cache-Control'] = 'no-cache';
  config.headers['Pragma'] = 'no-cache';
  config.headers['Access-Control-Allow-Credentials'] = 'true';
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
}, error => Promise.reject(error));

axiosClient.interceptors.response.use((response) => {
  return response;
}, error => {
  if (error.response && error.response.status === 401 && error.response.data.error === 'invalid token') {
    localStorage.removeItem('token');
    router.push({name: 'Login'});
  }

  throw error;
})

export default axiosClient