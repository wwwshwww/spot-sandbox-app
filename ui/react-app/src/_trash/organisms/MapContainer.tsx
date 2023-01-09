import { GoogleMap, LoadScript } from '@react-google-maps/api';
import React from 'react';

import { mapOptions, mapStyles } from '../../styles/GoogleMapStyle';
import { SpotMarker } from '../atoms/SpotMarker';

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

var KEY = "AIzaSyBg1bH_Hw4w0D-ES42SvnZJEx6wntbxtAA"

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
      <LoadScript googleMapsApiKey="AI">
        <GoogleMap
          mapContainerStyle={mapStyles}
          options={mapOptions}
          zoom={13}
          center={this.defaultCenter}
          onDblClick={(ev) => this.createMarker(ev)}
        >
          {this.state.markerInfos.map((m: MarkerInfo) => {
            return (
              <SpotMarker
                key={m.labelString}
                baseColor={m.baseColorH}
                labelText={m.labelString}
                latlng={m.position}
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
