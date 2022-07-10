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

interface SpotMarkerState {
  baseColor: number;
  labelText: string;
}

class SpotMarker extends React.Component<SpotMarkerProps, SpotMarkerState> {
  state: SpotMarkerState = {
    baseColor: this.props.baseColor,
    labelText: this.props.labelText,
  };

  render(): React.ReactNode {
    return (
      <Marker
        icon={{
          path: google.maps.SymbolPath.CIRCLE,
          fillColor: 'hsl(' + this.state.baseColor + ', 100%, 75%)',
          fillOpacity: 0.8,
          strokeWeight: 0,
          scale: 8,
        }}
        position={{ lat: this.props.latlng.lat, lng: this.props.latlng.lng }}
        label={{
          color: '#333333',
          text: this.state.labelText,
          fontSize: '10px',
        }}
      />
    );
  }
}

export default SpotMarker;
