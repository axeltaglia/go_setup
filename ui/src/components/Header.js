import React, {useState} from "react";
import styled, {withTheme} from "styled-components";
import {connect} from "react-redux";

import {
    AppBar as MuiAppBar,
    Grid,
    Hidden,
    IconButton as MuiIconButton,
    Toolbar
} from "@material-ui/core";

import {Menu as MenuIcon, Power} from "@material-ui/icons";
import Menu from "@material-ui/core/Menu";
import MenuItem from "@material-ui/core/MenuItem";
import Link from "@material-ui/core/Link";
import {NavLink as RouterNavLink} from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';

const AppBar = styled(MuiAppBar)`
  background: ${props => props.theme.header.background};
  color: ${props => props.theme.header.color};
  box-shadow: ${props => props.theme.shadows[1]};
`;

const IconButton = styled(MuiIconButton)`
  svg {
    width: 22px;
    height: 22px;
  }
`;

const NavLink = React.forwardRef((props, ref) => (
    <RouterNavLink innerRef={ref} {...props} />
));

function UserMenu() {
    const dispatch = useDispatch();
    const [anchorMenu, setAnchorMenu] = useState(null);

    const toggleMenu = event => {
        setAnchorMenu(event.currentTarget);
    };

    const closeMenu = () => {
        setAnchorMenu(null);
    };

    const handleSignOut = () => {
        //dispatch(signOutCurrentUser())
        closeMenu()
    };

    return (
        <React.Fragment>
            <IconButton
                aria-owns={Boolean(anchorMenu) ? "menu-appbar" : undefined}
                aria-haspopup="true"
                onClick={toggleMenu}
                color="inherit"
            >
                <Power />
            </IconButton>
            <Menu
                id="menu-appbar"
                anchorEl={anchorMenu}
                open={Boolean(anchorMenu)}
                onClose={closeMenu}
            >
                <MenuItem onClick={closeMenu}>
                    <Link
                        button
                        dense
                        component={NavLink}
                        exact
                        to={'/auth/sign-in'}
                        activeClassName="active"
                    >
                        Sign In
                    </Link>
                </MenuItem>
                <MenuItem onClick={handleSignOut}>
                    Sign out
                </MenuItem>
            </Menu>
        </React.Fragment>
    );
}


const Header = ({onDrawerToggle}) => (
    <React.Fragment>
        <AppBar position="sticky" elevation={0}>
            <Toolbar>
                <Grid container alignItems="center">
                    <Hidden mdUp>
                        <Grid item>
                            <IconButton
                                color="inherit"
                                aria-label="Open drawer"
                                onClick={onDrawerToggle}
                            >
                                <MenuIcon/>
                            </IconButton>
                        </Grid>
                    </Hidden>
                    <Grid item xs/>
                    <Grid item>
                        <UserMenu />
                    </Grid>
                </Grid>
            </Toolbar>
        </AppBar>
    </React.Fragment>
);

export default connect()(withTheme(Header));
