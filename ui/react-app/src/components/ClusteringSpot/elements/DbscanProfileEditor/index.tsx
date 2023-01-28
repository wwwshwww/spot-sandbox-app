import { ApolloError } from "@apollo/client";
import { Box, Button, Paper, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import { CDPActionType, CDPStateAndReducer, CSPStateAndReducer } from "../..";
import { DbscanProfile } from "../../../../generated/types";
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
          {v.key}: {v.distanceType}
          <br />
          MinCount: {v.minCount}
          <br />
          MaxCount: {v.maxCount ? v.maxCount : "-"}
          <br />
        </Card>
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
