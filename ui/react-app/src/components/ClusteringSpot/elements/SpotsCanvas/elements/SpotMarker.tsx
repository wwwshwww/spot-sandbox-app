import { Marker } from "@react-google-maps/api";
import React from "react";

interface SpotMarkerProps {
  color: {
    fill: string;
    fillOpacity: number;
    stroke: string;
    strokeOpacity: number;

  };
  labelText: string;
  visible: boolean;
  latlng: {
    lat: number;
    lng: number;
  };
  onClick: ((e: google.maps.MapMouseEvent) => void) | undefined;
}

export const getUnselectedColor = () => {
  return {
    fill: "hsl(0, 0%, 75%)",
    fillOpacity: 0.5,
    stroke: "hsl(0, 0%, 45%)",
    strokeOpacity: 0.5,
  };
};

export const getSelectedColor = () => {
  return {
    fill: "hsl(0, 100%, 75%)",
    fillOpacity: 0.9,
    stroke: "hsl(0, 40%, 45%)",
    strokeOpacity: 1,
  };
};

export const SpotMarker = (props: SpotMarkerProps) => {
  return (
    <Marker
      onClick={props.onClick}
      icon={{
        path: ' M 0 0 L -10 -30 A 10 12 1 0 1 10 -30 Z ',
        fillColor: props.color.fill,
        // fillColor: 'hsl(' + baseColor + ', 100%, 75%)',
        fillOpacity: props.color.fillOpacity,
        strokeWeight: 1,
        strokeColor: props.color.stroke,
        strokeOpacity: props.color.strokeOpacity,
        // strokeColor: 'hsl(' + baseColor + ', 40%, 45%)',
        scale: 1,
        labelOrigin: new google.maps.Point(0, -31),
      }}
      position={{ lat: props.latlng.lat, lng: props.latlng.lng }}
      visible={props.visible}
      label={{
        color: "#333333",
        text: props.labelText,
        fontSize: '11px',
        fontWeight: '700',
      }}
    />
  );
};
