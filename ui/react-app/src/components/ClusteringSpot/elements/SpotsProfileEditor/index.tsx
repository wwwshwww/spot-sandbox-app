import { ApolloError } from "@apollo/client";
import { Box, Button, IconButton, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import DeleteIcon from "@mui/icons-material/Delete";
import { CSPActionType, CSPStateAndReducer } from "../..";
import { SpotsProfile } from "../../../../generated/types";
import ScrollableList from "../../../General/ScrollableList";

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
  spotsProfilesParams: {
    isLoading: boolean;
    error?: ApolloError;
    spotsProfiles?: Array<SpotsProfile>;
  };
  initCurrent: CSPStateAndReducer;
}

const SpotsProfileEditor = (props: EditorProps) => {
  const { currentSpotsProfile, dispatchCSP } = props.initCurrent;

  const li = props.spotsProfilesParams.spotsProfiles?.map(
    (v: SpotsProfile, i: number) => (
      <span id={"sp:" + (i + 1).toString()}>
        <Grid container alignItems="center">
          <Grid padding={0}>
            <Card
              sx={
                v.key == currentSpotsProfile.spotsProfile?.key
                  ? {
                      borderColor: "#3f3f3f",
                      backgroundColor: "#ddd",
                    }
                  : {
                      borderColor: "#ccc",
                    }
              }
              width={175}
              border={1}
              borderRadius={1}
              paddingLeft={1}
              textAlign="left"
              onClick={() => {
                if (v.key == currentSpotsProfile.spotsProfile?.key) {
                  dispatchCSP({
                    type: CSPActionType.set,
                    payload: { spotsProfile: undefined, spots: undefined },
                  });
                } else {
                  dispatchCSP({
                    type: CSPActionType.set,
                    payload: { spotsProfile: v, spots: undefined },
                  });
                }
              }}
            >
              <b>No. {v.key}</b>
              <Grid padding={0}>
                <Grid padding={0}>spot count: {v.spots.length}</Grid>
              </Grid>
            </Card>
          </Grid>
          <Grid padding={0}>
            <IconButton aria-label="delete" size="small">
              <DeleteIcon />
            </IconButton>
          </Grid>
        </Grid>
      </span>
    )
  );
  const createButton = (
    <Button
      onClick={() => {
        window.location.href = "#sp:" + li?.length;
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
      {props.spotsProfilesParams.isLoading ? (
        <p>loading...</p>
      ) : (
        <ScrollableList
          title="spot profile"
          contents={li!}
          footer={createButton}
        />
      )}
    </>
  );
};

export default SpotsProfileEditor;
