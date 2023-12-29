// login.jsx

import React, { useState } from 'react';
import { registerUser, loginUser } from './callbacks';

const Login = () => {
  const [isRegistering, setIsRegistering] = useState(false);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const handleToggleView = () => {
    setIsRegistering(!isRegistering);
    setErrorMessage('');
  };

  const handleUserAction = () => {
    const userData = { username, password };

    const successCallback = (data) => {
      console.log(`${isRegistering ? 'Registration' : 'Login'} success:`, data);
      // Handle success, e.g., redirect to a welcome page or user's dashboard
    };

    const errorCallback = (error) => {
      console.error(`${isRegistering ? 'Registration' : 'Login'} error:`, error);
      setErrorMessage(error.error || 'An error occurred.');
    };

    if (isRegistering) {
      // Register user
      registerUser(userData, successCallback, errorCallback);
    } else {
      // Login user
      loginUser(userData, successCallback, errorCallback);
    }
  };

  return (
    <div>
      <h1>{isRegistering ? 'Register' : 'Login'}</h1>
      <form>
        <label>
          Username:
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </label>
        <br />
        <label>
          Password:
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </label>
        <br />
        <button type="button" onClick={handleUserAction}>
          {isRegistering ? 'Register' : 'Login'}
        </button>
      </form>
      {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
      <p>
        {isRegistering
          ? 'Already have an account?'
          : 'Don\'t have an account yet?'}
        <button type="button" onClick={handleToggleView}>
          {isRegistering ? 'Login' : 'Register'}
        </button>
      </p>
    </div>
  );
};

export default Login;
