import { useState, useCallback, useReducer } from 'react';
import { Box, Button, Grid, Stack } from '@mui/material';
import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

// import './App.css';
import MapContainer from './components/organisms/MapContainer';
import React from 'react';
import SpotClustering from './components/pages/SpotClustering';

const App: React.FC = () => {
  return (
    <div className="App">
      <SpotClustering />
    </div>
  );
};

export default App;
