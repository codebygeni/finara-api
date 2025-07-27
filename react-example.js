// React component example for consuming the CORS-enabled API

import React, { useState, useEffect } from 'react';

const FinancialDashboard = () => {
  const [dashboardData, setDashboardData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // API base URL - adjust based on your server configuration
  const API_BASE_URL = 'http://localhost:8080';

  useEffect(() => {
    fetchDashboardData();
  }, []);

  const fetchDashboardData = async () => {
    try {
      setLoading(true);
      
      // Example API calls with CORS enabled
      const responses = await Promise.all([
        fetch(`${API_BASE_URL}/health`),
        fetch(`${API_BASE_URL}/users`),
        fetch(`${API_BASE_URL}/`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
          },
          credentials: 'include' // Important for CORS with credentials
        })
      ]);

      const [healthData, usersData, apiInfo] = await Promise.all(
        responses.map(response => {
          if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
          }
          return response.json();
        })
      );

      setDashboardData({
        health: healthData,
        users: usersData,
        apiInfo: apiInfo
      });
      
    } catch (err) {
      setError(`Failed to fetch data: ${err.message}`);
      console.error('API Error:', err);
    } finally {
      setLoading(false);
    }
  };

  const openDashboard = () => {
    // Open the HTML dashboard in a new tab
    window.open(`${API_BASE_URL}/dashboard`, '_blank');
  };

  const fetchUserDashboard = async (userId) => {
    try {
      // Fetch user-specific data
      const response = await fetch(`${API_BASE_URL}/users/${userId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      });
      
      if (response.ok) {
        const userData = await response.json();
        console.log('User data:', userData);
        
        // Open user-specific dashboard
        window.open(`${API_BASE_URL}/dashboard/${userId}`, '_blank');
      }
    } catch (err) {
      console.error('Error fetching user dashboard:', err);
    }
  };

  if (loading) return <div>Loading financial data...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div style={{ padding: '20px', fontFamily: 'Arial, sans-serif' }}>
      <h1>🚗 Financial Dashboard Integration</h1>
      
      <div style={{ marginBottom: '20px' }}>
        <button 
          onClick={openDashboard}
          style={{
            padding: '10px 20px',
            backgroundColor: '#4facfe',
            color: 'white',
            border: 'none',
            borderRadius: '5px',
            cursor: 'pointer',
            marginRight: '10px'
          }}
        >
          Open Car Goal Dashboard
        </button>
        
        <button 
          onClick={() => fetchUserDashboard('6666666666')}
          style={{
            padding: '10px 20px',
            backgroundColor: '#667eea',
            color: 'white',
            border: 'none',
            borderRadius: '5px',
            cursor: 'pointer'
          }}
        >
          Open User Dashboard
        </button>
      </div>

      {dashboardData && (
        <div>
          <h2>API Connection Status</h2>
          <div style={{ 
            backgroundColor: '#f0f9ff', 
            padding: '15px', 
            borderRadius: '8px',
            marginBottom: '20px'
          }}>
            <h3>✅ CORS Enabled Successfully!</h3>
            <p>Service: {dashboardData.apiInfo?.service}</p>
            <p>Version: {dashboardData.apiInfo?.version}</p>
            <p>Available Endpoints: {Object.keys(dashboardData.apiInfo?.endpoints || {}).length}</p>
          </div>

          <h3>Available Endpoints:</h3>
          <ul>
            {dashboardData.apiInfo?.endpoints && Object.entries(dashboardData.apiInfo.endpoints).map(([key, value]) => (
              <li key={key}>
                <strong>{key}:</strong> {value}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default FinancialDashboard;

/* 
Usage instructions:

1. Start your Go API server:
   go run main.go

2. In your React app, install any required dependencies:
   npm install

3. Import and use this component:
   import FinancialDashboard from './FinancialDashboard';
   
4. The component will automatically connect to your CORS-enabled API
   and provide buttons to open the financial dashboard.

5. Make sure your React app is running on one of the allowed origins:
   - http://localhost:3000 (Create React App)
   - http://localhost:5173 (Vite)

Environment Variables for Production:
- Update the AllowOrigins in main.go to include your production domain
- Set API_BASE_URL to your production API endpoint
*/
