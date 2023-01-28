import { ApolloError } from "@apollo/client";
import { Box, Button, IconButton, Paper, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import { CDPActionType, CDPStateAndReducer, CSPStateAndReducer } from "../..";
import { DbscanProfile, DistanceType } from "../../../../generated/types";
import ScrollableList from "../../../General/ScrollableList";
import { Stack } from "@mui/system";

const Card = styled(Box)(({ theme }) => ({
  backgroundColor: "#f7f7f7",
  "&:hover": {
    marginLeft: 3,
    backgroundColor: "#eee",
  },
  "&:active": {
    backgroundColor: "#ddd",
    borderColor: "#3f3f3f",
  },
}));

interface EditorProps {
  dbscanProfilesParams: {
    isLoading: boolean;
    error?: ApolloError;
    dbscanProfiles?: Array<DbscanProfile>;
  };
  initCurrent: CDPStateAndReducer;
}

const DbscanProfileEditor = (props: EditorProps) => {
  const { currentDbscanProfile, dispatchCDP } = props.initCurrent;

  const li = props.dbscanProfilesParams.dbscanProfiles?.map(
    (v: DbscanProfile, i: number) => (
      <span id={"dp:" + (i + 1).toString()}>
        <Grid container alignItems="center">
          <Grid padding={0}>
            <Card
              sx={
                v.key == currentDbscanProfile.dbscanProfile?.key
                  ? {
                      borderColor: "#3f3f3f",
                      backgroundColor: "#ddd",
                    }
                  : {
                      borderColor: "#ccc",
                    }
              }
              border={1}
              width={175}
              borderRadius={1}
              paddingLeft={1}
              textAlign="left"
              onClick={() => {
                if (v.key == currentDbscanProfile.dbscanProfile?.key) {
                  dispatchCDP({
                    type: CDPActionType.set,
                    payload: { dbscanProfile: undefined },
                  });
                } else {
                  dispatchCDP({
                    type: CDPActionType.set,
                    payload: { dbscanProfile: v },
                  });
                }
              }}
            >
              <b>No. {v.key}</b>
              <Grid padding={0} container>
                <Grid>
                  type: <br />
                  threshold: <br />
                  count min: <br />
                  count max:
                </Grid>
                <Grid>
                  {v.distanceType}
                  <br />
                  {v.distanceType === DistanceType.TravelTime
                    ? v.minutesThreshold + " min"
                    : v.meterThreshold + " m"}
                  <br />
                  {v.minCount + " spot"}
                  <br />
                  {v.maxCount ? v.maxCount + " spot" : "-"}
                  <br />
                </Grid>
              </Grid>
            </Card>
          </Grid>
          <Grid padding={0}>
            <Stack>
            <IconButton aria-label="delete" size="small">
              <DeleteIcon />
            </IconButton>
            <IconButton aria-label="edit" size="small">
              <EditIcon />
            </IconButton>
            </Stack>
            

          </Grid>
        </Grid>
      </span>
    )
  );
  const createButton = (
    <Button
      onClick={() => {
        window.location.href = "#dp:" + li?.length;
        window.history.replaceState(
          null,
          "",
          location.pathname + location.search
        );
      }}
    >
      CREATE
    </Button>
  );

  return (
    <>
      {props.dbscanProfilesParams.isLoading ? (
        <p>loading...</p>
      ) : (
        <ScrollableList
          title="clustering profile"
          contents={li!}
          footer={createButton}
        />
      )}
    </>
  );
};

export default DbscanProfileEditor;
