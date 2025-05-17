import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

export const api = axios.create({
  baseURL: API_BASE_URL,
});

// Add a request interceptor to add the user ID header
api.interceptors.request.use((config) => {
  const user = JSON.parse(localStorage.getItem('user') || '{}');
  if (user.id) {
    config.headers['X-User-Id'] = user.id;
  }
  return config;
});

// Jobs
export const createJob = async (data: any) => {
  const response = await api.post('/jobs', data);
  return response.data;
};

export const getJobs = async () => {
  const response = await api.get('/jobs');
  return response.data;
};

export const getJobById = async (id: string) => {
  const response = await api.get(`/jobs/${id}`);
  return response.data;
};

// Applications
export const applyToJob = async (data: any) => {
  const response = await api.post('/applications', data);
  return response.data;
};

export const getApplications = async (params: any) => {
  const response = await api.get('/applications', { params });
  return response.data;
};

// Sessions
export const createSession = async (data: any) => {
  const response = await api.post('/sessions', data);
  return response.data;
};

export const getSessions = async (params: any) => {
  const response = await api.get('/sessions', { params });
  return response.data;
};

// Users
export const getCurrentUser = async () => {
  const user = JSON.parse(localStorage.getItem('user') || '{}');
  return user;
};

