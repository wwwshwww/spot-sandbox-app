import { useState, useCallback} from 'react';
import { GoogleMap, LoadScript } from '@react-google-maps/api';

import './App.css';

const MapContainer = () => {
  
  const mapStyles = {
    top: "50px",
    margin: "auto",
    height: "500px",
    width: "500px", 
  };
  
  const defaultCenter = {
    lat: 41.3851, lng: 2.1734
  }
  
  return (
     <LoadScript
       googleMapsApiKey='APIKEY'>
        <GoogleMap
          mapContainerStyle={mapStyles}
          zoom={13}
          center={defaultCenter}
        />
     </LoadScript>
  )
}

const App: React.FC = () => {
  const [ count, setCount ] = useState<number>(0)

  const handleIncrement = useCallback(() => {
      setCount(prev => prev + 1)
  }, [])

  const handleDecrement = useCallback(() => {
      setCount(prev => prev - 1)
  }, []);

  return (
      <div className="App">
          <div>{ count }</div>
          <div>
              <button onClick={handleIncrement}>+1</button>
              <button onClick={handleDecrement}>-1</button>
          </div>
          <MapContainer/>
      </div>
  )
}

export default App;
