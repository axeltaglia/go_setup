import React, {useContext} from "react";
import styled from "styled-components";
import {NavLink as RouterNavLink} from "react-router-dom";

import Helmet from 'react-helmet';


import {
    Breadcrumbs as MuiBreadcrumbs,
    Card as MuiCard,
    CardContent,
    Divider as MuiDivider,
    Grid,
    Link,
    Typography
} from "@material-ui/core";

import {spacing} from "@material-ui/system";
import {Alert} from "../../containers/AlertContainer";

const NavLink = React.forwardRef((props, ref) => (
    <RouterNavLink innerRef={ref} {...props} />
));

const Card = styled(MuiCard)(spacing);

const Divider = styled(MuiDivider)(spacing);

const Breadcrumbs = styled(MuiBreadcrumbs)(spacing);

function EmptyCard() {

    return (
        <Card mb={6}>
            <CardContent>
                <Typography variant="body2" gutterBottom>
                    Empty card
                </Typography>
            </CardContent>
        </Card>
    );
}

function Blank() {
    const {
        alertSuccess,
    } = Alert.useContainer()

    return (
        <React.Fragment>
            <Helmet title="Blank"/>
            <Typography variant="h3" gutterBottom display="inline">
                Blank
            </Typography>

            <Breadcrumbs aria-label="Breadcrumb" mt={2}>
                <Link component={NavLink} exact to="/">
                    Dashboard
                </Link>
                <Link component={NavLink} exact to="/">
                    Pages
                </Link>
                <Typography>Blank</Typography>
            </Breadcrumbs>

            <Divider my={6}/>

            <Grid container spacing={6}>
                <Grid item xs={12}>
                    <button onClick={() => alertSuccess("Hola mundo")}>Show alert</button>
                </Grid>
            </Grid>
        </React.Fragment>
    );
}

export default Blank;
