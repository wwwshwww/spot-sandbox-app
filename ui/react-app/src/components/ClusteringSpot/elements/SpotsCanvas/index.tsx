import { Button, Paper, styled } from "@mui/material";
import Grid from '@mui/material/Unstable_Grid2'; 

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));

const SpotsCanvas: React.FC = () => {
  return (
    <Grid
      container
      spacing={2}
      justifyContent="center"
    >
      <Grid>
        <Item>GoogleMapGoogleMap</Item>
      </Grid>
    </Grid>
  );
}

export default SpotsCanvas