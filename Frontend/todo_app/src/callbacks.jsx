// callbacks.jsx

import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Update with your API base URL

// Callback for user registration
export const registerUser = async (userData, successCallback, errorCallback) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/auth/register`, userData);
    successCallback(response.data);
  } catch (error) {
    errorCallback(error.response.data);
  }
};

// Callback for user login
export const loginUser = async (userData, successCallback, errorCallback) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/auth/login`, userData);
    successCallback(response.data);
  } catch (error) {
    errorCallback(error.response.data);
  }
};

// Example usage:
// import { registerUser, loginUser } from './callbacks';

// const userData = {
//   username: 'example_user',
//   password: 'example_password',
// };

// registerUser(userData,
//   (data) => {
//     console.log('Registration success:', data);
//     // Handle success, e.g., redirect to a welcome page or user's dashboard
//   },
//   (error) => {
//     console.error('Registration error:', error);
//     // Handle error, e.g., display an error message
//   }
// );

// loginUser(userData,
//   (data) => {
//     console.log('Login success:', data);
//     // Handle success, e.g., redirect to the user's dashboard
//   },
//   (error) => {
//     console.error('Login error:', error);
//     // Handle error, e.g., display an error message
//   }
// );
