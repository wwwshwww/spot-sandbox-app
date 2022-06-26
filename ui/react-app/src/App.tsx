import { useState, useCallback, useReducer} from 'react';
import { Box, Button, Stack } from '@mui/material';
import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

// import './App.css';
import {mapOptions, mapStyles} from './mapstyle'
import React from 'react';

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

  // const locations = [
  //   {
  //     name: "Location 1",
  //     location: {
  //         lat: 41.3954,
  //         lng: 2.162
  //     },
  //   },
  //   {
  //       name: "Location 2",
  //       location: {
  //           lat: 41.3917,
  //           lng: 2.1649
  //       },
  //   },
  // ]
  
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
          fillColor: "hsl(" + h + ", 100%, 50%)",
          strokeColor: "hsl(" + h + ", 60%, 35%)",
          fillOpacity: 0.6,
          strokeWeight: 1,
          scale: 8,
        }}
        position={locations[i].latlng}
        label={{
          color: "hsl(" + h + ", 60%, 35%)",
          text: String(i),
          fontSize: "10px"
        }}
      />
      h += 360 / locations.length
    }
    return ms
  }
}

interface CounterProps {
  initialValue: number;
  increVal: number;
  decreVal: number;
}

interface CounterState {
  count: number
  arr: Array<JSX.Element>
}

class Counter extends React.Component<CounterProps, CounterState>{
  state: CounterState = {
    count: this.props.initialValue,
    arr: this.make(this.props.initialValue)
  }
  render(): React.ReactNode {
    return (
      <div>
        <div>{this.state.count}</div>
        <Stack direction={"row"}>
          <Button variant="outlined" onClick={() => this.increment(this.props.increVal)}>+{this.props.increVal}</Button>
          <Button variant="outlined" onClick={() => this.decrement(this.props.decreVal)}>-{this.props.decreVal}</Button>
        </Stack>
        <Stack>{this.state.arr}</Stack>
      </div>
    )
  }
  make(count: number): Array<JSX.Element> {
    let a: Array<JSX.Element> = new Array(count)
    let h: number = 0
    for(let i=0; i<count; i++){
      a[i] = (<a style={{backgroundColor: "hsl(" + h + ", 100%, 50%)"}}>色</a>)
      h += 360 / count
    }
    return a
  }
  
  increment = (amt: number) => {
    this.setState((state) => ({
      count: state.count + amt,
      arr: this.make(state.count + amt)
    }))
  }
  decrement = (amt: number) => {
    if ((this.state.count - amt) > 0){
      this.setState((state) => ({
        count: state.count - amt,
        arr: this.make(state.count - amt)
      }))
    }
  }
}

const App: React.FC = () => {
  return (
      <div className="App">
        <MapContainer dummy={0} />
        {/* <Counter initialValue={1} decreVal={1} increVal={1}/> */}
      </div>
  )
}

export default App;
