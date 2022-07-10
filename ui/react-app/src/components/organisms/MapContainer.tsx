import React from 'react';
import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

import { mapOptions, mapStyles } from '../../styles/GoogleMapStyle';

type MarkerInfo = {
  position: {
    lat: number;
    lng: number;
  };
  baseColorH: number;
  labelString: string;
};

interface MapContainerState {
  markerInfos: Array<MarkerInfo>;
}

interface MapContainerProps {
  dummy: number;
}

class MapContainer extends React.Component<
  MapContainerProps,
  MapContainerState
> {
  defaultCenter = {
    lat: 41.3851,
    lng: 2.1734,
  };

  state: MapContainerState = {
    markerInfos: new Array<MarkerInfo>(),
  };

  render(): React.ReactNode {
    return (
      <LoadScript googleMapsApiKey="APIKEY">
        <GoogleMap
          mapContainerStyle={mapStyles}
          options={mapOptions}
          zoom={13}
          center={this.defaultCenter}
          onDblClick={(ev) => this.createMarker(ev)}
        >
          {this.state.markerInfos.map((m: MarkerInfo) => {
            return (
              <Marker
                icon={{
                  path: google.maps.SymbolPath.CIRCLE,
                  fillColor: 'hsl(' + m.baseColorH + ', 100%, 75%)',
                  fillOpacity: 0.8,
                  strokeWeight: 0,
                  scale: 8,
                }}
                position={m.position}
                label={{
                  color: '#333333',
                  text: m.labelString,
                  fontSize: '10px',
                }}
              />
            );
          })}
        </GoogleMap>
      </LoadScript>
    );
  }

  createMarker(ev: google.maps.MapMouseEvent) {
    console.log(this.state);
    if (ev.latLng != null) {
      let newMarker: MarkerInfo = {
        position: {
          lat: ev.latLng.lat(),
          lng: ev.latLng.lng(),
        },
        baseColorH: 0,
        labelString: '',
      };
      this.setState((state) => {
        const infos = state.markerInfos.concat([newMarker]);
        let h: number = 0;
        for (let i in infos) {
          infos[i].baseColorH = h;
          infos[i].labelString = i;
          h += 360 / infos.length;
          console.log(i);
        }
        return {
          markerInfos: infos,
        };
      });
    }
  }
}

export default MapContainer;
