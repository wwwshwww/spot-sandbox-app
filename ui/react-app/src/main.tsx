import React from 'react';
import ReactDOM from 'react-dom';
import { Routes } from 'react-router';
import { Route, BrowserRouter } from 'react-router-dom';

import App from './App';
import './index.css';

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')!,
);
