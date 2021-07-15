import React from 'react';
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import Dashboard from "../layouts/Dashboard";
import Home from "../pages/dashboard/Home";
import SignIn from "../pages/auth/SignIn";
import SignUp from "../pages/auth/SignUp";
import PrivateRoute from "./PrivateRoute";
import Auth from "../layouts/Auth";

const Routes = () => {
    return (
        <Router>
            <Switch>
                <PrivateRoute exact path={'/'} component={Home} layout={Dashboard} />
                <Route exact path={'/auth/sign-in'} render={props => (<Auth><SignIn {...props} /></Auth>)}/>
                <Route exact path={'/auth/sign-up'} render={props => (<Auth><SignUp {...props} /></Auth>)}/>
            </Switch>
        </Router>
    )
}

export default Routes;