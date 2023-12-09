import App from './App'
import React from 'react'
import ReactDOM from 'react-dom/client'
import Layout from './components/Layout'

// Frontend at http://127.0.0.1:1420/
// Backend at http://127.0.0.1:8080/api/cars

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Layout>
      <App />
    </Layout>
  </React.StrictMode>
)
