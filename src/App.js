import React from 'react';
import { Route, Switch } from 'react-router-dom';

import CreateCustomer from './components/Admin/CreateCustomer';
import CreateEvent from './components/Admin/CreateEvent';
import CreateOrganizer from './components/Admin/CreateOrganizer';
import Customers from './components/Admin/Customers';
import Events from './components/Admin/Events';
import Home from './components/Home/Home';
import Login from './components/Login/Login';
import LoginAdmin from './components/Login/LoginAdmin';
import Organizers from './components/Admin/Organizers';
import Register from './components/Register/Register';

import './App.css';

const App = () => {
  return (
    <div>
      <Switch>
        <Route exact path='/' component={Home} />
        <Route exact path='/login' component={Login} />
        <Route exact path='/register' component={Register} />
        <Route exact path='/loginadmin' component={LoginAdmin} />
        <Route exact path='/admin/customers' component={Customers} />
        <Route exact path='/admin/organizers' component={Organizers} />
        <Route exact path='/admin/events' component={Events} />
        <Route exact path='/admin/newcustomer' component={CreateCustomer} />
        <Route exact path='/admin/newevent' component={CreateEvent} />
        <Route exact path='/admin/neworganizer' component={CreateOrganizer} />
      </Switch>
    </div>
  );
}

export default App;
