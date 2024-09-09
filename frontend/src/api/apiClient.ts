import axios, { AxiosInstance } from 'axios';

// Create an instance of Axios with default configuration
const apiClient: AxiosInstance = axios.create({
  baseURL: 'http://localhost:8080', // Replace with your API base URL
  timeout: 10000, // Set a timeout if desired
  headers: {
    'Content-Type': 'application/json',
  },
});

export default apiClient;