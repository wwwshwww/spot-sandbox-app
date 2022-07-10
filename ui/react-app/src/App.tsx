import { useState, useCallback, useReducer} from 'react';
import { Box, Button, Grid, Stack } from '@mui/material';
import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

// import './App.css';
import MapContainer from './components/pages/SpotClustering'
import React from 'react';



const App: React.FC = () => {
  return (
      <div className="App">
        // TODO: move to pages
        <Grid sx={{p: 2, flexGrow: 1}} container spacing={2} justifyContent="center">
          <Grid item>
            <MapContainer dummy={0} />
          </Grid>
          <Grid item>
            <Button>asdf</Button>
          </Grid>
        </Grid>
        {/* <Counter initialValue={1} decreVal={1} increVal={1}/> */}
      </div>
  )
}

export default App;
