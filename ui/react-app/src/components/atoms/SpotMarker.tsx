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
        path: 'M0,0 l-1,-8 S-10,-18 0,-20. M0,0 l1,-8 S10,-18 0,-20 ',
        fillColor: 'hsl(' + props.baseColor + ', 100%, 75%)',
        fillOpacity: 0.8,
        strokeWeight: 1,
        strokeColor: 'hsl(' + props.baseColor + ', 80%, 50%)',
        scale: 2,
        labelOrigin: new google.maps.Point(0, -15),
      }}
      position={{ lat: props.latlng.lat, lng: props.latlng.lng }}
      label={{
        color: '#333333',
        text: props.labelText,
        fontSize: '10px',
        fontWeight: '700',
      }}
    />
  );
};
