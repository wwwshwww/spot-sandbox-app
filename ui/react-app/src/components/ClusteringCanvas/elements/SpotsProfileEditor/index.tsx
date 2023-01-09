import { Box, Button, Paper, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import { borderBottom } from "@mui/system";
import { useEffect } from "react";
import { Scalars, SpotsProfile, Spot } from "../../../../generates/types";
import ScrollableList from "../../../General/ScrollableList";
import useFetchList from "./hooks/useFetchList";

interface CurrentSPState {
  spotsProfile: SpotsProfile | undefined;
}

enum CurrentSPActionType {
  updateSpots,
}

interface CurrentSPActionPayload {
  spots: Array<Spot> | undefined;
}

interface CurrentSPAction {
  type: CurrentSPActionType;
  payload: CurrentSPActionPayload;
}

function reducer(
  state: CurrentSPState,
  action: CurrentSPAction
): CurrentSPState {
  switch (action.type) {
    case CurrentSPActionType.updateSpots:
      if (state.spotsProfile && action.payload.spots) {
        state.spotsProfile.spots = action.payload.spots;
        return state;
      } else {
        throw new Error();
      }
    default:
      throw new Error();
  }
}

const SpotsProfileEditor: React.FC = () => {
  const { loading, error, spotsProfiles } = useFetchList();
  if (error) {
    return <p>error :\</p>;
  }

  const li = spotsProfiles?.map((v: SpotsProfile) => {
    return (
      <Box sx={{ borderBottom: 1 }}>
        {v.key}: {v.spots.length} spots
      </Box>
    );
  });
  const createButton = <Button disabled={loading}>CREATE</Button>;

  return (
    <ScrollableList
      title="spot profile"
      contents={li!}
      footer={createButton}
    />
  );
};

export default SpotsProfileEditor;
