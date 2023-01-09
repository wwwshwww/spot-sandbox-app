import { Box, Button } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import DbscanProfileEditor from "./elements/DbscanProfileEditor";
import SpotsCanvas from "./elements/SpotsCanvas";
import SpotsProfileEditor from "./elements/SpotsProfileEditor";

const Sample: React.FC = () => {
  return (
    <Box>
      <Grid
        sx={{ flexGrow: 1 }}
        container
        spacing={2}
        justifyContent="center"
        alignItems="flex-start"
      >
        <Grid>
          <SpotsProfileEditor />
        </Grid>
        <Grid>
          <SpotsCanvas />
        </Grid>
        <Grid>
          <DbscanProfileEditor />
        </Grid>
      </Grid>
    </Box>
  );
};

export default Sample;
