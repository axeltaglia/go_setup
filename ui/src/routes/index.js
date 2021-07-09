import React from "react";
import async from "../components/Async";

import {
    Sliders,
    Users
} from "react-feather";

// Auth components
const SignIn = async(() => import("../pages/auth/SignIn"));
const SignUp = async(() => import("../pages/auth/SignUp"));

// Components components
const Home = async(() => import("../pages/dashboard/Home"));

const authRoutes = {
    id: "Auth",
    path: "/auth",
    icon: <Users/>,
    children: [
        {
            path: "/auth/sign-in",
            name: "Sign In",
            component: SignIn
        },
        {
            path: "/auth/sign-up",
            name: "Sign Up",
            component: SignUp
        },
    ],
    component: null
};

const dashboardsRoutes = {
    id: "Dashboard",
    path: "/dashboard",
    header: "Pages",
    icon: <Sliders/>,
    containsHome: true,
    children: [
        {
            path: "/",
            name: "Home Page",
            component: Home
        }
    ],
    component: null
};

// Routes using the Auth layout
export const authLayoutRoutes = [authRoutes];

// Routes using the Dashboard layout
export const dashboardLayoutRoutes = [
    dashboardsRoutes,
];

// Routes visible in the sidebar
export const sidebarRoutes = [
    dashboardsRoutes,
];
