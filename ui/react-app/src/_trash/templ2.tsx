import { Button } from "@mui/material";
import React, { useEffect, useReducer, useState } from "react"

interface MyState{
  count: number
}

interface MyAction{
  type: string
}

const initialState = {count: 0};

function reducer(state: MyState, action: MyAction) {
  switch (action.type) {
    case 'increment':
      return {count: state.count + 1};
    case 'decrement':
      return {count: state.count - 1};
    default:
      throw new Error();
  }
}

interface MyHoji{
  state: MyState
  dispatch: React.Dispatch<MyAction>
}

function Hoji(): MyHoji {
  useEffect(()=>{
    console.log("eff")
  })

  const [state, dispatch] = useReducer(reducer, initialState);
  console.log("rendered", state)

  return (
    {state, dispatch}
  );
}

const Coun: React.FC<MyHoji> = (prop: MyHoji) => {
  const hoji = prop
  return(
    <div>
      <p>{hoji.state.count}</p>
       <Button onClick={() => hoji.dispatch({type:"increment"})}>s</Button>
    </div>
  )
}


const Counter: React.FC = () => {
  const h = Hoji()
  const c1 = Coun(h)
  const c2 = Coun(h)
  return(
    <div>
      {c1}
      {c2}
    </div>
  )
}

export default Counter;