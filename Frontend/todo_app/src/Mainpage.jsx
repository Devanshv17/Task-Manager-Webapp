// src/MainPage.jsx

import React from 'react';
import { Link } from 'react-router-dom';

const MainPage = () => {
  return (
    <div>
      <h1>Welcome to Task Manager App</h1>
      <Link to="/login">
        <button>Start Here</button>
      </Link>
    </div>
  );
};

export default MainPage;
