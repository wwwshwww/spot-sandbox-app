import { ApolloError } from "@apollo/client";
import { Box, Button, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
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
          {v.key}: {v.spots.length} spots
          <br />
          asdfa
        </Card>
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
