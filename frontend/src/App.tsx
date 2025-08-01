import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import { LandingPage } from './interface/pages/LandingPage';
import { GamePage } from './interface/pages/GamePage';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/user/:userId" element={<GamePage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
