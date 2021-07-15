import React, {useState} from "react";
import styled from "styled-components";
import {Link} from "react-router-dom";
import {AuthContainer} from '../../containers/AuthContainer'
import Helmet from 'react-helmet';

import {
    Checkbox,
    FormControl,
    FormControlLabel,
    Input,
    InputLabel,
    Button as MuiButton,
    Paper,
    Typography
} from "@material-ui/core";
import {spacing} from "@material-ui/system";
import {withRouter} from 'react-router-dom';
import Grid from "@material-ui/core/Grid";

const Button = styled(MuiButton)(spacing);

const Wrapper = styled(Paper)`
  padding: ${props => props.theme.spacing(6)}px;

  ${props => props.theme.breakpoints.up("md")} {
    padding: ${props => props.theme.spacing(10)}px;
  }
`;

function SignIn(props) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const {
        signIn,
    } = AuthContainer.useContainer();

    const handleSubmit = async () => {
        try {
            await signIn({email, password})
            props.history.push("/")
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <Wrapper>
            <Helmet title="Sign In"/>

            <Typography component="h1" variant="h4" align="center" gutterBottom>
                Welcome!
            </Typography>
            <Typography component="h2" variant="body1" align="center">
                Sign in to your account to continue
            </Typography>
            <form>
                <FormControl margin="normal" required fullWidth>
                    <InputLabel htmlFor="email">Email Address</InputLabel>
                    <Input
                        id="email"
                        name="email"
                        autoComplete="email"
                        autoFocus
                        onChange={event => setEmail(event.target.value)}
                    />
                </FormControl>
                <FormControl margin="normal" required fullWidth>
                    <InputLabel htmlFor="password">Password</InputLabel>
                    <Input
                        id="password"
                        name="password"
                        type="password"
                        autoComplete={"password"}
                        onChange={event => setPassword(event.target.value)}
                    />
                </FormControl>
                <FormControlLabel
                    control={<Checkbox value="remember" color="primary"/>}
                    label="Remember me"
                />

                <Grid container spacing={3}>
                    <Grid item xs={6}>
                        <Button
                            component={Link}
                            color="primary"
                            variant="outlined"
                            mt={2}
                            to={"/auth/sign-up"}
                        >
                            Create an account
                        </Button>
                    </Grid>
                    <Grid item container xs={6} justify="flex-end">
                        <Button
                            component={Link}
                            variant="contained"
                            color="primary"
                            mt={2}
                            onClick={handleSubmit}
                        >
                            Sign in
                        </Button>
                    </Grid>
                </Grid>

            </form>
        </Wrapper>
    );
}

export default withRouter(SignIn);
