import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL as string,
});

// Add a request interceptor to add the JWT token to headers
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    const branchID = localStorage.getItem('selectedBranchID');
    if (branchID) {
      config.headers['X-Branch-ID'] = branchID;
    }
    
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default api;
