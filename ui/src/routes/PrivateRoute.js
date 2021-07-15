import React from 'react';
import {AuthContainer} from '../containers/AuthContainer'
import {Route} from "react-router-dom";
import Auth from "../layouts/Auth";
import SignIn from "../pages/auth/SignIn";

const PrivateRoute = ({component: Component, layout: Layout, ...rest}) => {
    const {
        token,
    } = AuthContainer.useContainer();
    return (
        <Route {...rest} render={(props) => (
            token
                ? <Layout><Component {...props} /></Layout>
                :
                <Auth><SignIn {...props} /></Auth>
        )}/>
    )
};


export default PrivateRoute;