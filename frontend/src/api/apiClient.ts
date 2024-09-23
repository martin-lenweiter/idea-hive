import axios, { AxiosInstance } from 'axios';

// Create an instance of Axios with default configuration
const apiClient: AxiosInstance = axios.create({
  baseURL: process.env.REACT_APP_BASE_URL,
  // add /api to the baseURL
  timeout: 10000, // Set a timeout if desired
  headers: {
    'Content-Type': 'application/json',
  },
});

export default apiClient;