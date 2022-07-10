import React from "react";
import { Box, Button, Grid, Stack } from '@mui/material';
import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

import {mapOptions, mapStyles} from '../../styles/GoogleMapStyle'
import SpotMarker from '../atoms/SpotMarker'

type Location = {
  index: number
  latlng: {
    lat: number
    lng: number
  }
}

interface MapContainerState {
  locations: Array<Location>
  markers: Array<JSX.Element>
}

interface MapContainerProps {
  dummy: number
}

class MapContainer extends React.Component<MapContainerProps, MapContainerState> {

  defaultCenter = {
    lat: 41.3851, 
    lng: 2.1734
  }

  state: MapContainerState = {
    locations: [],
    markers: []
  }
  
  render(): React.ReactNode {
    return (
      <LoadScript
        googleMapsApiKey='APIKEY'>
        <GoogleMap
          mapContainerStyle={mapStyles}
          options={mapOptions}
          zoom={13}
          center={this.defaultCenter}
          onDblClick={ev => this.createMarker(ev)}
        >
          {this.state.markers}
        </GoogleMap>
      </LoadScript>
    )
  }

  createMarker(ev: google.maps.MapMouseEvent) {
    if (ev.latLng != null){
      const loc: Location = {
        index: this.state.locations.length,
        latlng: {
          lat: ev.latLng.lat(),
          lng: ev.latLng.lng()
        }
      }
      
      this.setState((state) => {
        let locs = state.locations.concat([loc]);
        return {
          locations: locs,
          markers: this.fill(locs)
        }
      })
    }
  }

  fill(locations: Array<Location>): Array<JSX.Element> {
    const ms: Array<JSX.Element> = new Array(locations.length)
    let h: number = 0
    for (let i in locations){
      ms[i] = <Marker
        icon={{
          path: google.maps.SymbolPath.CIRCLE,
          fillColor: "hsl(" + h + ", 100%, 75%)",
          fillOpacity: 0.8,
          strokeWeight: 0,
          scale: 8,
        }}
        position={locations[i].latlng}
        label={{
          color: "#333333",
          text: String(i),
          fontSize: "10px"
        }}
      />
      // ms[i] = <SpotMarker latlng={locations[i].latlng} labelText={i} baseColor={h} /> 
      h += 360 / locations.length
    }
    return ms
  }
}

export default MapContainer