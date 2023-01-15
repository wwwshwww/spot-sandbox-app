import { Button, Paper, styled } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import { LoadScript, GoogleMap } from "@react-google-maps/api";
import { useEffect, useState } from "react";
import { CSPState } from "../..";
import { Spot } from "../../../../generates/types";
import { mapStyles, mapOptions } from "../../../../styles/GoogleMapStyle";
import {
  getSelectedColor,
  getUnselectedColor,
  SpotMarker,
} from "./elements/SpotMarker";

const googleMapsApiKey = "APIKEY";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: "center",
  color: theme.palette.text.secondary,
}));

interface SpotsCanvasProps {
  initSpots: Array<Spot>;
  defaultGoogleMapParams: {
    zoom: number;
    center: { lat: number; lng: number };
  };
  currentSpotsProfile: CSPState;
}

const SpotsCanvas = (props: SpotsCanvasProps) => {
  const [spots, setSpots] = useState(props.initSpots);
  const [center, setCenter] = useState(props.defaultGoogleMapParams.center);
  const [zoom, setZoom] = useState(props.defaultGoogleMapParams.zoom);

  const selectedDict: { [key: number]: boolean } = {};
  if (props.currentSpotsProfile.spotsProfile !== undefined) {
    for (const spot of props.currentSpotsProfile.spotsProfile.spots!) {
      selectedDict[spot.key] = true;
    }
  }

  const markers = spots?.map((v: Spot, i: number) => {
    const isSelected = selectedDict[v.key] !== undefined;
    return (
      <SpotMarker
        key={i}
        color={isSelected ? getSelectedColor() : getUnselectedColor()}
        labelText={v.key.toString()}
        visible={true}
        latlng={{ lat: v.lat, lng: v.lng }}
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
