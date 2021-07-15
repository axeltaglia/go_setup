import React from "react";
import styled from "styled-components";

import {
    Container,
    Grid,
    Hidden,
    Typography as MuiTypography,
} from "@material-ui/core";

import { spacing } from "@material-ui/system";
import {withRouter} from "react-router-dom";

const Typography = styled(MuiTypography)(spacing);

const IntroductionBase = styled.div`
  padding: 3vw 5vw;
`;

const IntroductionContent = styled.div`
  padding: ${props => props.theme.spacing(6)}px;
  line-height: 150%;
`;

const IntroductionImage = styled.img`
  margin: ${props => props.theme.spacing(6)}px;
  max-width: 100%;
  height: auto;
  display: block;
  box-shadow: 0 6px 18px 0 rgba(18,38,63,.1);
`;

const IntroductionSubtitle = styled(Typography)`
  font-size: ${props => props.theme.typography.h5.fontSize};
  font-weight: ${props => props.theme.typography.fontWeightRegular};
  color: ${props => props.theme.palette.grey[800]};
  font-family: ${props => props.theme.typography.fontFamily};
  margin-bottom: ${props => props.theme.spacing(4)}px;
`;

function Introduction() {
    return (
        <IntroductionBase>
            <Container>
                <Grid container alignItems="center" justify="center">
                    <Grid item xs={12} xl={6}>
                        <IntroductionContent>
                            <Typography variant="h1" gutterBottom>
                                Institutional site
                            </Typography>
                            <IntroductionSubtitle>
                                This is a institutional site. The backend was build in Go and the frontend in React JS and designed with Material design. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                                <br />
                                Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
                            </IntroductionSubtitle>
                        </IntroductionContent>
                    </Grid>
                    <Hidden lgDown>
                        <Grid item xs={12} xl={6}>
                            <IntroductionImage
                                alt="Material App - React Admin Template"
                                src={`/images/golang.png`}
                            />
                        </Grid>
                    </Hidden>
                </Grid>
            </Container>
        </IntroductionBase>
    );
}

function Home() {
    return (
        <React.Fragment>
            <Introduction />
        </React.Fragment>
    );
}

export default withRouter(Home);
