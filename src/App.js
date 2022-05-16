import React from 'react';
import { Route, Switch } from 'react-router-dom';

import Customers from './components/Admin/Customers';
import Login from './components/Login/Login';

import './App.css';

const App = () => {
  return (
    <div>
      <Switch>
        <Route exact path='/' component={Login} />
        <Route exact path='/admin/customers' component={Customers} />
      </Switch>
    </div>
  );
}

export default App;
