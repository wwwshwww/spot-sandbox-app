import { Button, Paper, styled } from "@mui/material";
import Grid from '@mui/material/Unstable_Grid2'; 
import { Spot } from "../../../../generates/types";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));


interface SpotsCanvasProps {
  initSpots: Array<Spot> | undefined;
}

const SpotsCanvas: React.FC<SpotsCanvasProps> = (props) => {
  return (
    <Grid
      spacing={2}
      justifyContent="center"
    >
      <Grid>
        <Item>GoogleMapGoogleMap</Item>
      </Grid>
      <Grid>
        {
          props.initSpots?.map((v: Spot, i: number) => (
            <Item key={i}>{v.addressRepr}</Item>
          ))
        }
      </Grid>
    </Grid>
  );
}

export default SpotsCanvas