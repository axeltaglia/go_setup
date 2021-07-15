import React, {useState} from "react";
import Helmet from 'react-helmet';

import DateFnsUtils from "@date-io/date-fns";
import {ThemeProvider as MuiThemeProvider} from "@material-ui/core/styles";
import {MuiPickersUtilsProvider} from "@material-ui/pickers";
import {StylesProvider} from "@material-ui/styles";
import {ThemeProvider} from "styled-components";

import maTheme from "./theme";
import Routes from "./routes/Routes";
import axios from 'axios';
import Store from "./containers/Store";

function App({theme}) {
    [theme] = useState(5)
    setup();
    return (
        <React.Fragment>
            <Helmet
                titleTemplate="%s | React Setup Project"
                defaultTitle="React Setup Project"
            />
            <StylesProvider injectFirst>
                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                    <MuiThemeProvider theme={maTheme[theme]}>
                        <ThemeProvider theme={maTheme[theme]}>
                            <Store>
                                <Routes/>
                            </Store>
                        </ThemeProvider>
                    </MuiThemeProvider>
                </MuiPickersUtilsProvider>
            </StylesProvider>
        </React.Fragment>
    );
}

function setup() {
    axios.interceptors.request.use(function (config) {
        config.url = "http://localhost:8080" + config.url;
        return config;
    })
}

export default App;
