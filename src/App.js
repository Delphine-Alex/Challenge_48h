import React from 'react';
import { Route, Switch } from 'react-router-dom';

import Login from './components/Login/Login';

import './App.css';

const App = () => {
  return (
    <div>
      <Switch>
        <Route exact path='/' component={Login} />
      </Switch>
    </div>
  );
}

export default App;
