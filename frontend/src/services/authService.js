const BASE_URL = 'http://localhost:8070';

const authService = {
  login: async (username) => {
    try {
      const response = await fetch(`${BASE_URL}/users`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username }),
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem('userId', data.userId);
        return true;
      } else {
        const errorData = await response.json();
        console.error(`Login failed: ${errorData.error}`);
        return false;
      }
    } catch (error) {
      console.error('Login failed:', error);
      return false;
    }
  },

  isAuthenticated: () => {
    return !!localStorage.getItem('userId');
  },

  logout: () => {
    localStorage.removeItem('userId');
  },

  getUsername: () => {
    return localStorage.getItem('userId');
  }
};

export default authService;
