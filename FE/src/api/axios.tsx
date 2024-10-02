import axios, { AxiosInstance } from 'axios';

const BASE_URL: string = 'http://localhost:8081';

// Create a default Axios instance
const axiosInstance: AxiosInstance = axios.create({
    baseURL: BASE_URL,
});

export default axiosInstance;

// Create a private Axios instance with additional headers and credentials
export const axiosPrivate: AxiosInstance = axios.create({
    baseURL: BASE_URL,
    headers: { 'Content-Type': 'application/json' },
    withCredentials: true,
});