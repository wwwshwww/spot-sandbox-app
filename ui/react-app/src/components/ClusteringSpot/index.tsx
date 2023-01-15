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
      return { spotsProfile: action.payload.spotsProfile };
    case CSPActionType.updateSpots:
      if (state.spotsProfile && action.payload.spots) {
        // TODO: アップデートクエリによる更新
        return {
          spotsProfile: {
            key: state.spotsProfile.key,
            spots: action.payload.spots,
          },
        };
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

const calcForGoogleMap = (spots: Array<Spot>) => {
  const scaleConverter = 1.7;
  let totalLat = 0;
  let totalLng = 0;
  let maxLat = -999;
  let maxLng = -999;
  let minLat = 999;
  let minLng = 999;
  for (const s of spots) {
    totalLat += s.lat;
    totalLng += s.lng;
    maxLat = s.lat > maxLat ? s.lat : maxLat;
    maxLng = s.lng > maxLng ? s.lng : maxLng;
    minLat = s.lat < minLat ? s.lat : minLat;
    minLng = s.lng < minLng ? s.lng : minLng;
  }

  const diffLat = Math.abs(maxLat-minLat);
  const diffLng = Math.abs(maxLng-minLng);
  return {
    zoom: diffLat > diffLng? scaleConverter / diffLat : scaleConverter / diffLng,
    center: {
      lat: totalLat / spots.length,
      lng: totalLng / spots.length,
    },
  };
};

export const ClusteringSpot = () => {
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
              <SpotsCanvas
                initSpots={spots!}
                defaultGoogleMapParams={calcForGoogleMap(spots!)}
                currentSpotsProfile={currentSpotsProfile}
              />
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
