import React from 'react';
import ReactDOM from 'react-dom';
import { Routes } from 'react-router';
import { Route, BrowserRouter } from 'react-router-dom';

import { 
  ApolloClient,
  ApolloProvider,
  HttpLink,
  InMemoryCache,
  gql
} from '@apollo/client';

import App from './pages/_app';
import './styles/index.css';

const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: new HttpLink({
    uri: 'http://localhost:8080/query',
  }),
});

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={
          <ApolloProvider client={client}>
            <App />
          </ApolloProvider>
        } />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')!,
);
