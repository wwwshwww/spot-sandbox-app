import { Box, Button, Paper, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import { useState } from "react";
import { CSPActionType, CSPStateAndReducer, CSPState } from "../..";
import { SpotsProfile } from "../../../../generates/types";
import ScrollableList from "../../../General/ScrollableList";
import useGetAll from "./hooks/GetAll";

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
  initSpotsProfiles: Array<SpotsProfile>;
  initCurrent: CSPStateAndReducer;
}

const SpotsProfileEditor: React.FC<EditorProps> = (props) => {
  const [spotsProfiles, setSpotsProfiles] = useState(props.initSpotsProfiles);
  const { currentSpotsProfile, dispatch } = props.initCurrent;

  const li = spotsProfiles.map((v: SpotsProfile) => (
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
          dispatch({
            type: CSPActionType.set,
            payload: { spotsProfile: undefined, spots: undefined },
          });
        } else {
          dispatch({
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
  ));
  const createButton = <Button>CREATE</Button>;

  return (
    <ScrollableList title="spot profile" contents={li!} footer={createButton} />
  );
};

export default SpotsProfileEditor;
