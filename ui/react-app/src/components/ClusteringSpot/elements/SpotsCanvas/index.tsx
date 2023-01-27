import { useApolloClient } from "@apollo/client";
import { LinearProgress, Paper, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import { LoadScript, GoogleMap } from "@react-google-maps/api";
import { useEffect, useState } from "react";
import {
  CDPActionType,
  CDPStateAndReducer,
  CSPActionType,
  CSPStateAndReducer,
} from "../..";
import { ClusterElement, Spot } from "../../../../generated/types";
import { DbscanResult } from "../../../../generates/types";
import { mapStyles, mapOptions } from "../../../../styles/GoogleMapStyle";
import {
  getClusterColor,
  getNotClusterColor,
  getSelectedColor,
  getUnselectedColor,
  SpotMarker,
} from "./elements/SpotMarker";
import { QueryDbscan } from "./hooks/Dbscan";

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
    zoom: number;
    center: { lat: number; lng: number };
  };
  currentSpotsProfileParams: CSPStateAndReducer;
  currentDbscanProfileParams: CDPStateAndReducer;
}

const SpotsCanvas = (props: SpotsCanvasProps) => {
  const client = useApolloClient();

  const [spots, setSpots] = useState(props.initSpots);
  const [center, setCenter] = useState(props.defaultGoogleMapParams.center);
  const [zoom, setZoom] = useState(props.defaultGoogleMapParams.zoom);

  const { currentSpotsProfile, dispatchCSP } = props.currentSpotsProfileParams;
  const { currentDbscanProfile, dispatchCDP } =
    props.currentDbscanProfileParams;

  const isSelectedMap: { [key: number]: boolean } = {};
  if (currentSpotsProfile.spotsProfile !== undefined) {
    currentSpotsProfile.spotsProfile.spots!.forEach((element) => {
      isSelectedMap[element.key] = true;
    });
  }

  const clusterdMap: { [key: number]: ClusterElement } = {};
  props.clusterElements?.forEach((element) => {
    clusterdMap[element.spot.key] = element;
  });

  const markers = spots?.map((v: Spot, i: number) => {
    console.log(props.clusterElements);
    const isSelected = isSelectedMap[v.key] !== undefined;
    const element = clusterdMap[v.key];
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

  return (
    <Grid spacing={2} justifyContent="center">
      <Item>
        <LoadScript googleMapsApiKey={googleMapsApiKey}>
          <GoogleMap
            mapContainerStyle={mapStyles}
            options={mapOptions}
            zoom={zoom}
            center={center}
            onDblClick={(ev) => {
              console.log("clicked");
            }}
          >
            {markers}
          </GoogleMap>
        </LoadScript>
      </Item>
      {props.isLoading ? <LinearProgress color="inherit" /> : ``}

      <Grid>
        {spots?.map((v: Spot, i: number) => (
          <Item key={i} sx={{ textAlign: "left", paddingLeft: 1 }}>
            <Grid container>
              <Grid>{v.key}:</Grid>
              <Grid>{v.addressRepr}</Grid>
            </Grid>
          </Item>
        ))}
      </Grid>
    </Grid>
  );
};

export default SpotsCanvas;
