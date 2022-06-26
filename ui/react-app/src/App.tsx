import { useState, useCallback, useReducer} from 'react';
import { Box, Button, Stack } from '@mui/material';
import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';

// import './App.css';
import React from 'react';

const MapContainer = () => {
  
  const mapStyles = {
    height: "500px",
    width: "500px", 
  };
  
  const defaultCenter = {
    lat: 41.3851, 
    lng: 2.1734
  }

  const locations = [
    {
      name: "Location 1",
      location: {
          lat: 41.3954,
          lng: 2.162
      },
  },
  {
      name: "Location 2",
      location: {
          lat: 41.3917,
          lng: 2.1649
      },
  },
  ]
  
  return (
     <LoadScript
       googleMapsApiKey='APIKEY'>
        <GoogleMap
          mapContainerStyle={mapStyles}
          zoom={13}
          center={defaultCenter}
        >
          <Marker
            icon={{
              path: google.maps.SymbolPath.CIRCLE,
              fillColor: "hsl(100, 100%, 50%)",
              strokeColor: "hsl(100, 60%, 50%)",
              fillOpacity: 1,
              strokeWeight: 2,
              scale: 8,
            }}
            position={locations[0]["location"]}
          />
        </GoogleMap>
     </LoadScript>
  )
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
        <MapContainer/>
        {/* <Counter initialValue={1} decreVal={1} increVal={1}/> */}
      </div>
  )
}

export default App;
