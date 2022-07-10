import { useState, useCallback, useReducer} from 'react';
import { Box, Button, Stack } from '@mui/material';
import React from 'react';


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
  
export default Counter