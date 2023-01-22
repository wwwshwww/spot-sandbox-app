import { Box, Button } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import { useEffect, useReducer, useState } from "react";
import { DbscanProfile, Spot, SpotsProfile } from "../../generated/types";
import DbscanProfileEditor from "./elements/DbscanProfileEditor";
import SpotsCanvas from "./elements/SpotsCanvas";
import SpotsProfileEditor from "./elements/SpotsProfileEditor";
import {
  QueryGetAllSpotsProfile,
  SpotsProfiles,
} from "./elements/SpotsProfileEditor/hooks/GetAll";
import { useGetAll as GetSpots } from "./elements/SpotsCanvas/hooks/GetAll";
import { useApolloClient, useQuery } from "@apollo/client";
import { MutationUpdateSpotsProfile } from "./elements/SpotsProfileEditor/hooks/Update";
import {
  DbscanProfiles,
  QueryGetAllDbscanProfiles,
} from "./elements/DbscanProfileEditor/hooks/GetAll";

export interface CSPState {
  spotsProfile?: SpotsProfile;
}

export enum CSPActionType {
  set,
  updateSpots,
}

export interface CSPActionPayload {
  spotsProfile?: SpotsProfile;
  spots?: Array<Spot>;
}

export interface CSPAction {
  type: CSPActionType;
  payload: CSPActionPayload;
}

export interface CSPStateAndReducer {
  currentSpotsProfile: CSPState;
  dispatchCSP: React.Dispatch<CSPAction>;
}

export interface CDPState {
  dbscanProfile?: DbscanProfile;
}

export enum CDPActionType {
  set,
  update,
}

export interface CDPActionPayload {
  dbscanProfile?: DbscanProfile;
}

export interface CDPAction {
  type: CDPActionType;
  payload: CDPActionPayload;
}

export interface CDPStateAndReducer {
  currentDbscanProfile: CDPState;
  dispatchCDP: React.Dispatch<CDPAction>;
}

const initialCSP: CSPState = { spotsProfile: undefined };
const initialCDP: CDPState = { dbscanProfile: undefined };

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

  const diffLat = Math.abs(maxLat - minLat);
  const diffLng = Math.abs(maxLng - minLng);
  return {
    zoom:
      diffLat > diffLng ? scaleConverter / diffLat : scaleConverter / diffLng,
    center: {
      lat: totalLat / spots.length,
      lng: totalLng / spots.length,
    },
  };
};

export const ClusteringSpot = () => {
  const client = useApolloClient();

  const [currentSpotsProfile, dispatchCSP] = useReducer(
    (state: CSPState, action: CSPAction): CSPState => {
      switch (action.type) {
        case CSPActionType.set:
          return { spotsProfile: action.payload.spotsProfile };
        case CSPActionType.updateSpots:
          if (state.spotsProfile && action.payload.spots) {
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
    },
    initialCSP
  );

  const [currentDbscanProfile, dispatchCDP] = useReducer(
    (state: CDPState, action: CDPAction): CDPState => {
      switch (action.type) {
        case CDPActionType.set:
          return { dbscanProfile: action.payload.dbscanProfile };
        case CDPActionType.update:
          // TODO: implement
          return { dbscanProfile: action.payload.dbscanProfile };
        default:
          throw new Error();
      }
    },
    initialCDP
  );

  const [spotsProfiles, setSpotsProfiles] = useState<Array<SpotsProfile>>();
  const [dbscanProfiles, setDbscanProfiles] = useState<Array<DbscanProfile>>();

  const { loading: spLoading, error: spErr } = useQuery<SpotsProfiles>(
    QueryGetAllSpotsProfile,
    {
      onCompleted: (sps) => {
        setSpotsProfiles(sps.spotsProfiles);
      },
    }
  );
  const { loading: dpLoading, error: dpErr } = useQuery<DbscanProfiles>(
    QueryGetAllDbscanProfiles,
    {
      onCompleted: (dps) => {
        setDbscanProfiles(dps.dbscanProfiles);
      },
    }
  );

  const { loading: sLoading, error: sErr, spots } = GetSpots();

  useEffect(() => {
    if (currentSpotsProfile.spotsProfile) {
      client
        .mutate({
          mutation: MutationUpdateSpotsProfile,
          variables: {
            key: currentSpotsProfile.spotsProfile?.key,
            input: {
              spotKeys: currentSpotsProfile.spotsProfile?.spots.map((s) => {
                return s.key;
              }),
            },
          },
        })
        .then(() => {
          client
            .query<SpotsProfiles>({
              query: QueryGetAllSpotsProfile,
              fetchPolicy: "network-only",
            })
            .catch((err) => {
              throw err;
            })
            .then((res) => {
              setSpotsProfiles(res!.data.spotsProfiles);
            });
        });
    }
  }, [currentSpotsProfile]);

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
        <Grid>
          <SpotsProfileEditor
            spotsProfilesParams={{
              isLoading: spLoading,
              error: spErr,
              spotsProfiles,
            }}
            initCurrent={{ currentSpotsProfile, dispatchCSP }}
          />
        </Grid>
        <Grid>
          {sLoading ? (
            <p>loading...</p>
          ) : (
            <SpotsCanvas
              initSpots={spots!}
              defaultGoogleMapParams={calcForGoogleMap(spots!)}
              currentSpotsProfileParams={{ currentSpotsProfile, dispatchCSP }}
              currentDbscanProfileParams={{ currentDbscanProfile, dispatchCDP }}
            />
          )}
        </Grid>
        <Grid>
          <DbscanProfileEditor
            dbscanProfilesParams={{
              isLoading: dpLoading,
              error: dpErr,
              dbscanProfiles,
            }}
            initCurrent={{ currentDbscanProfile, dispatchCDP }}
          />
        </Grid>
      </Grid>
    </Box>
  );
};

export default ClusteringSpot;
