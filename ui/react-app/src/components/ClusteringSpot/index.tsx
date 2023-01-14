import { Box, Button } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import { useEffect, useReducer } from "react";
import { Spot, SpotsProfile } from "../../generates/types";
import DbscanProfileEditor from "./elements/DbscanProfileEditor";
import SpotsCanvas from "./elements/SpotsCanvas";
import SpotsProfileEditor from "./elements/SpotsProfileEditor";
import { useGetAll as GetSpotsProfiles } from "./elements/SpotsProfileEditor/hooks/GetAll";
import { useGetAll as GetSpots } from "./elements/SpotsCanvas/hooks/GetAll";

export interface CSPState {
  spotsProfile: SpotsProfile | undefined;
}

export enum CSPActionType {
  set,
  updateSpots,
}

export interface CSPActionPayload {
  spotsProfile: SpotsProfile | undefined;
  spots: Array<Spot> | undefined;
}

export interface CSPAction {
  type: CSPActionType;
  payload: CSPActionPayload;
}

function reducer(state: CSPState, action: CSPAction): CSPState {
  switch (action.type) {
    case CSPActionType.set:
      state.spotsProfile = action.payload.spotsProfile;
      return state;
    case CSPActionType.updateSpots:
      if (state.spotsProfile && action.payload.spots) {
        // TODO: アップデートクエリによる更新
        state.spotsProfile.spots = action.payload.spots;
        return state;
      } else {
        throw new Error();
      }
    default:
      throw new Error();
  }
}

export interface CSPStateAndReducer {
  currentSpotsProfile: CSPState;
  dispatch: React.Dispatch<CSPAction>;
}

const initialCSP: CSPState = { spotsProfile: undefined };

const getCurrentSpotsProfileState = (): CSPStateAndReducer => {
  const [currentSpotsProfile, dispatch] = useReducer(reducer, initialCSP);
  return { currentSpotsProfile, dispatch };
};

export const ClusteringSpot: React.FC = () => {
  const {
    loading: spLoading,
    error: spErr,
    spotsProfiles,
  } = GetSpotsProfiles();
  const { currentSpotsProfile, dispatch: cspDisp } =
    getCurrentSpotsProfileState();

  const { loading: sLoading, error: sErr, spots } = GetSpots();

  return (
    <Box>
      <Grid
        sx={{ flexGrow: 1 }}
        container
        spacing={1}
        // rowSpacing={0}
        justifyContent="center"
        alignItems="flex-start"
      >
        {spLoading || sLoading ? (
          <p>loading...</p>
        ) : (
          <>
            <Grid>
              <SpotsProfileEditor
                initSpotsProfiles={spotsProfiles!}
                initCurrent={{ currentSpotsProfile, dispatch: cspDisp }}
              />
            </Grid>
            <Grid>
              <SpotsCanvas initSpots={spots} />
            </Grid>
            <Grid>
              <DbscanProfileEditor />
            </Grid>
          </>
        )}
      </Grid>
    </Box>
  );
};

export default ClusteringSpot;
