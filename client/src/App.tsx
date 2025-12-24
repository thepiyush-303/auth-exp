import { useState } from 'react'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>User SignIn Login Page</h1>
      <div className="card">
        <form style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
          <div style={{ display: 'flex', flexDirection: 'column', gap: '0.5rem' }}>
        <label htmlFor="name" style={{ fontWeight: '500' }}>Name:</label>
        <input type="text" id="name" name="name" required style={{ padding: '0.75rem', borderRadius: '0.5rem', border: '1px solid #ccc' }} />
          </div>
          <div style={{ display: 'flex', flexDirection: 'column', gap: '0.5rem' }}>
        <label htmlFor="email" style={{ fontWeight: '500' }}>Email:</label>
        <input type="email" id="email" name="email" required style={{ padding: '0.75rem', borderRadius: '0.5rem', border: '1px solid #ccc' }} />
          </div>
          <div style={{ display: 'flex', flexDirection: 'column', gap: '0.5rem' }}>
        <label htmlFor="password" style={{ fontWeight: '500' }}>Password:</label>
        <input type="password" id="password" name="password" required style={{ padding: '0.75rem', borderRadius: '0.5rem', border: '1px solid #ccc' }} />
          </div>
          <button type="submit" style={{ padding: '0.75rem', marginTop: '1rem', backgroundColor: '#666', color: 'white', border: 'none', borderRadius: '0.5rem', fontWeight: '600', cursor: 'pointer' }}>Sign In</button>
          <button type="button" style={{ padding: '0.75rem', backgroundColor: '#666', color: 'white', border: 'none', borderRadius: '0.5rem', fontWeight: '600', cursor: 'pointer' }}>Login</button>
        </form>
      </div>
      {/* <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div> */}
    </>
  )
}

export default App
