import React, {useEffect} from "react";
import styled from "styled-components";
import {NavLink as RouterNavLink} from "react-router-dom";
import {useDispatch, useSelector} from 'react-redux';

import Helmet from 'react-helmet';


import {
    Breadcrumbs as MuiBreadcrumbs,
    Card as MuiCard,
    Divider as MuiDivider,
    Grid,
    Link,
    Typography
} from "@material-ui/core";

import {spacing} from "@material-ui/system";

const NavLink = React.forwardRef((props, ref) => (
    <RouterNavLink innerRef={ref} {...props} />
));

const Card = styled(MuiCard)(spacing);

const Divider = styled(MuiDivider)(spacing);

const Breadcrumbs = styled(MuiBreadcrumbs)(spacing);

function Categories() {
    //const categories = useSelector(categories);
    const dispatch = useDispatch();

    useEffect(() => {
        console.log("dispatch")
        //dispatch(fetchCategories());
    }, [])

    return (
        <React.Fragment>
            <Helmet title="Blank"/>
            <Typography variant="h3" gutterBottom display="inline">
                Categories
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

                </Grid>
            </Grid>
        </React.Fragment>
    );
}

export default Categories;
