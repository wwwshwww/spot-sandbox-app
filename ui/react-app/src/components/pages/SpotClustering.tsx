import React from 'react';
import { Grid, Button } from '@mui/material';
import MapContainer from '../organisms/MapContainer';

const SpotClustering: React.FC = () => {
  return (
    <div>
      <Grid
        sx={{ p: 2, flexGrow: 1 }}
        container
        spacing={2}
        justifyContent="center"
      >
        <Grid item>
          <MapContainer dummy={0} />
        </Grid>
        <Grid item>
          <Button>asdf</Button>
          {/* <Counter initialValue={1} decreVal={1} increVal={1}/> */}
        </Grid>
      </Grid>
    </div>
  );
};

export default SpotClustering;
