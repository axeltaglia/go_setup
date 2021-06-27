import React from "react";
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import {dashboardLayoutRoutes, authLayoutRoutes} from "./index";

import DashboardLayout from "../layouts/Dashboard";
import AuthLayout from "../layouts/Auth";
import Page404 from "../pages/auth/Page404";

const childRoutes = (Layout, routes) =>
    routes.map(({children, path, component: Component}, index) =>
        children ? (
            // Route item with children
            children.map(({path, component: Component}, index) => (
                <Route
                    key={index}
                    path={path}
                    exact
                    render={props => (
                        <Layout>
                            <Component {...props} />
                        </Layout>
                    )}
                />
            ))
        ) : (
            // Route item without children
            <Route
                key={index}
                path={path}
                exact
                render={props => (
                    <Layout>
                        <Component {...props} />
                    </Layout>
                )}
            />
        )
    );

const Routes = () => (
    <Router>
        <Switch>
            {childRoutes(DashboardLayout, dashboardLayoutRoutes)}
            {childRoutes(AuthLayout, authLayoutRoutes)}
            <Route
                render={() => (
                    <AuthLayout>
                        <Page404/>
                    </AuthLayout>
                )}
            />
        </Switch>
    </Router>
);

export default Routes;
