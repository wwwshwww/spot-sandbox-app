import { Box, Button } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import Item from "./elements/Item";

interface ScrollableListProps {
  title: string;
  contents: Array<JSX.Element>;
  footer?: JSX.Element;
}

const ScrollableList: React.FC<ScrollableListProps> = (props) => {
  return (
    <Grid container direction='column' justifyContent='center'>
      <Grid>
        <Item>
          <Grid
            sx={{
              textTransform: 'uppercase',
            }}
          >
            {props.title}
          </Grid>
          <Grid minWidth={200}>
            <Box sx={{ border: 1, borderRadius: 1, borderColor: "#ddd" }}>
              <Grid
                maxHeight={250}
                paddingRight={1}
                sx={{
                  overflowY: 'scroll',
                }}
              >
                {props.contents?.map((v: JSX.Element, i: number) => (
                  <Grid key={i}>{v}</Grid>
                ))}
              </Grid>
            </Box>
          </Grid>
          {props.footer!}
        </Item>
      </Grid>
    </Grid>
  );
};

export default ScrollableList;
