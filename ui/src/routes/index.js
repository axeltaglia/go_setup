import React from "react";
import async from "../components/Async";

import {
  Sliders,
  Users
} from "react-feather";

// Auth components
const SignIn = async(() => import("../pages/auth/SignIn"));

// Components components
const Blank = async(() => import("../pages/dashboard/Blank"));
const Categories = async(() => import("../pages/dashboard/Categories"));

const authRoutes = {
  id: "Auth",
  path: "/auth",
  icon: <Users />,
  children: [
    {
      path: "/auth/sign-in",
      name: "Sign In",
      component: SignIn
    }
  ],
  component: null
};

const dashboardsRoutes = {
  id: "Dashboard",
  path: "/dashboard",
  header: "Pages",
  icon: <Sliders />,
  containsHome: true,
  children: [
    {
      path: "/",
      name: "Blank Page",
      component: Blank
    },
    {
      path: "/categories",
      name: "Categories",
      component: Categories
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
