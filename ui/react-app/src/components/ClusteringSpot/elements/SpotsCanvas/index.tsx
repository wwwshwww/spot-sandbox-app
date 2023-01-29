import { useApolloClient } from "@apollo/client";
import { LinearProgress, Paper, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import { LoadScript, GoogleMap, Polyline } from "@react-google-maps/api";
import { useEffect, useState } from "react";
import {
  CDPActionType,
  CDPStateAndReducer,
  CSPActionType,
  CSPStateAndReducer,
} from "../..";
import {
  ClusterElement,
  MutationCreateSpotArgs,
  Spot,
} from "../../../../generated/types";
import { DbscanResult } from "../../../../generates/types";
import { mapStyles, mapOptions } from "../../../../styles/GoogleMapStyle";
import {
  getClusterColor,
  getNotClusterColor,
  getSelectedColor,
  getUnselectedColor,
  SpotMarker,
} from "./elements/SpotMarker";
import { MutationCreateSpot } from "./hooks/CreateSpot";
import { QueryDbscan } from "./hooks/Dbscan";
import { QueryGetAllSpot } from "./hooks/GetAllSpot";

const googleMapsApiKey = "APIKEY";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: "center",
  color: theme.palette.text.secondary,
}));

interface SpotsCanvasProps {
  isLoading: boolean;
  initSpots: Array<Spot>;
  clusterElements?: Array<ClusterElement>;
  clusterNum: number;
  defaultGoogleMapParams: {
    center: { lat: number; lng: number };
  };
  currentSpotsProfileParams: CSPStateAndReducer;
  currentDbscanProfileParams: CDPStateAndReducer;
}

const SpotsCanvas = (props: SpotsCanvasProps) => {
  const client = useApolloClient();

  const [spots, setSpots] = useState(props.initSpots);
  const [center, setCenter] = useState(props.defaultGoogleMapParams.center);

  const { currentSpotsProfile, dispatchCSP } = props.currentSpotsProfileParams;
  const { currentDbscanProfile, dispatchCDP } =
    props.currentDbscanProfileParams;

  const spotMap: Map<number, Spot> = new Map();
  spots.forEach((element) => {
    spotMap.set(element.key, element);
  });

  const isSelectedMap: Map<number, boolean> = new Map();
  if (currentSpotsProfile.spotsProfile !== undefined) {
    currentSpotsProfile.spotsProfile.spots!.forEach((element) => {
      isSelectedMap.set(element.key, true);
    });
  }

  const clusterdMap: Map<number, ClusterElement> = new Map();
  props.clusterElements?.forEach((element) => {
    clusterdMap.set(element.spot.key, element);
  });

  let lineCount = 0;
  const pathMap: Map<number, Array<number>> = new Map();
  props.clusterElements?.forEach((element) => {
    pathMap.set(
      element.spot.key,
      element.paths.map((c) => {
        lineCount++;
        return c.spot.key;
      })
    );
  });

  const markers = spots?.map((v: Spot, i: number) => {
    const isSelected = isSelectedMap.get(v.key) !== undefined;
    const element = clusterdMap.get(v.key);
    const payloadSpots = isSelected
      ? currentSpotsProfile.spotsProfile?.spots!.filter((s: Spot) => {
          return s.key !== v.key;
        })
      : currentSpotsProfile.spotsProfile?.spots!.concat(v);
    const color = isSelected
      ? element !== undefined
        ? element.assignedNumber != -1
          ? getClusterColor(
              (360 / props.clusterNum) * element!.assignedNumber + 1
            )
          : getNotClusterColor()
        : getSelectedColor()
      : getUnselectedColor();

    return (
      <SpotMarker
        key={i}
        color={color}
        labelText={v.key.toString()}
        visible={true}
        latlng={{ lat: v.lat, lng: v.lng }}
        onClick={() => {
          dispatchCSP({
            type: CSPActionType.updateSpots,
            payload: { spotsProfile: undefined, spots: payloadSpots },
          });
        }}
      />
    );
  });

  const lines: Array<JSX.Element> = new Array(lineCount);
  let cnt = 0;
  pathMap.forEach((paths, fromSpotKey) => {
    paths.forEach((destSpotKey) => {
      const from = spotMap.get(fromSpotKey);
      const dest = spotMap.get(destSpotKey);
      const coords = [
        { lat: from?.lat!, lng: from?.lng! },
        { lat: dest?.lat!, lng: dest?.lng! },
      ];
      lines[cnt] = (
        <Polyline
          key={cnt}
          path={coords}
          options={{
            strokeOpacity: 0.65,
            strokeWeight: 1,
            strokeColor: "#666666",
            icons: [
              {
                icon: {
                  path: google.maps.SymbolPath.FORWARD_CLOSED_ARROW,
                  scale: 2,
                  fillOpacity: 0.85,
                },
                offset: "50%",
              },
            ],
          }}
        />
      );
      cnt++;
    });
  });

  return (
    <Grid spacing={2} justifyContent="center">
      <Item>
        <LoadScript googleMapsApiKey={googleMapsApiKey}>
          <GoogleMap
            mapContainerStyle={mapStyles}
            options={mapOptions}
            zoom={5}
            center={center}
            onDblClick={(ev) => {
              client
                .mutate({
                  mutation: MutationCreateSpot,
                  variables: {
                    input: {
                      lat: ev.latLng?.lat(),
                      lng: ev.latLng?.lng(),
                    },
                  },
                })
                .then(() => {
                  // 複数同時利用を想定し全取得しているが、本来必要ないかもしれない
                  client
                    .query({
                      query: QueryGetAllSpot,
                      fetchPolicy: "network-only",
                    })
                    .catch((err) => {
                      throw err;
                    })
                    .then((res) => {
                      setSpots(res!.data.spots);
                    });
                });
            }}
          >
            {markers}
            {lines}
          </GoogleMap>
        </LoadScript>
      </Item>
      {props.isLoading ? <LinearProgress color="inherit" /> : ``}

      <Grid>
        {/* {spots?.map((v: Spot, i: number) => (
          <Item key={i} sx={{ textAlign: "left", paddingLeft: 1 }}>
            <Grid container>
              <Grid>{v.key}:</Grid>
              <Grid>{v.addressRepr}</Grid>
            </Grid>
          </Item>
        ))} */}
      </Grid>
    </Grid>
  );
};

export default SpotsCanvas;
