import React from 'react';
import { Route, Switch } from 'react-router-dom';

import Customers from './components/Admin/Customers';
import Events from './components/Admin/Events';
import LoginAdmin from './components/Login/LoginAdmin';
import Organizers from './components/Admin/Organizers';

import './App.css';

const App = () => {
  return (
    <div>
      <Switch>
        <Route exact path='/admin' component={LoginAdmin} />
        <Route exact path='/admin/customers' component={Customers} />
        <Route exact path='/admin/organizers' component={Organizers} />
        <Route exact path='/admin/events' component={Events} />
      </Switch>
    </div>
  );
}

export default App;
