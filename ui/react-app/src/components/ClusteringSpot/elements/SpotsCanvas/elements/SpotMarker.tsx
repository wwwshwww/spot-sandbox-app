import { Marker } from '@react-google-maps/api';
import React from 'react';

interface SpotMarkerProps {
  baseColor: number;
  labelText: string;
  latlng: {
    lat: number;
    lng: number;
  };
}

export const SpotMarker: React.FC<SpotMarkerProps> = (props) => {
  return (
    <Marker
      icon={{
        path: ' M 0 0 L -10 -30 A 10 12 1 0 1 10 -30 Z ',
        fillColor: 'hsl(' + props.baseColor + ', 100%, 75%)',
        fillOpacity: 0.9,
        strokeWeight: 1,
        strokeColor: 'hsl(' + props.baseColor + ', 40%, 45%)',
        scale: 1,
        labelOrigin: new google.maps.Point(0, -31),
      }}
      position={{ lat: props.latlng.lat, lng: props.latlng.lng }}
      label={{
        color: '#333333',
        text: props.labelText,
        fontSize: '11px',
        fontWeight: '700',
      }}
    />
  );
};
