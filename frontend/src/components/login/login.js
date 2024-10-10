import React, { useState } from 'react';

import { useNavigate } from 'react-router-dom';
import authService from '../../services/authService';

const LoginPage = ({ setIsAuthenticated }) => {
  const [username, setUsername] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    const isLoggedIn = await authService.login(username);
    if (isLoggedIn) { 
      setIsAuthenticated(true);
      navigate('/');
    } else {
      alert('Invalid username');
    }
  };

  return (
    <div>
      <h2>Sign in / Sign up</h2>
      <form onSubmit={handleLogin}>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Enter username"
          required
        />
        <button type="submit">Enter</button>
      </form>
    </div>
  );
};

export default LoginPage;
