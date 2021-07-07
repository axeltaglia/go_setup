import React, {useState} from "react";
import styled from "styled-components";
import {Link, withRouter} from "react-router-dom";

import Helmet from 'react-helmet';

import {
    FormControl,
    Input,
    InputLabel,
    Button as MuiButton,
    Paper,
    Typography
} from "@material-ui/core";
import { spacing } from "@material-ui/system";
import {Auth} from "../../containers/AuthContainer";
import {Alert} from "../../containers/AlertContainer";

const Button = styled(MuiButton)(spacing);

const Wrapper = styled(Paper)`
  padding: ${props => props.theme.spacing(6)}px;

  ${props => props.theme.breakpoints.up("md")} {
    padding: ${props => props.theme.spacing(10)}px;
  }
`;

function SignUp(props) {
    const [name, setName] = useState('');
    const [lastName, setLastName] = useState('');
    const [occupation, setOccupation] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const {
        isLoggedIn,
        signUp,
    } = Auth.useContainer();

    const {
        alertSuccess
    } = Alert.useContainer();

    /*
    useEffect(() => {
        if(isLoggedIn()) {
            props.history.push("/")
        }
    }, [])

     */


    const handleSubmit = async () => {
        try {
            await signUp({name, lastName, occupation, email, password})
            alertSuccess("User created!")
            props.history.push("/")
        } catch (e) {
            alertSuccess("Something went wrong")
            console.log(e)
        }
    }

    return (
        <Wrapper>
            <Helmet title="Sign Up" />
            <Typography component="h1" variant="h4" align="center" gutterBottom>
                Get started
            </Typography>
            <Typography component="h2" variant="body1" align="center">
                Start creating the best possible user experience for you customers
            </Typography>
            <form>
                <FormControl margin="normal" required fullWidth>
                    <InputLabel htmlFor="name">Name</InputLabel>
                    <Input id="name" name="name" autoFocus onChange={event => setName(event.target.value)} />
                </FormControl>
                <FormControl margin="normal" required fullWidth>
                    <InputLabel htmlFor="lastName">Last name</InputLabel>
                    <Input id="lastName" name="lastName" autoFocus onChange={event => setLastName(event.target.value)} />
                </FormControl>
                <FormControl margin="normal" required fullWidth>
                    <InputLabel htmlFor="occupation">Occupation</InputLabel>
                    <Input id="occupation" name="occupation" onChange={event => setOccupation(event.target.value)} />
                </FormControl>
                <FormControl margin="normal" required fullWidth>
                    <InputLabel htmlFor="email">Email Address</InputLabel>
                    <Input id="email" name="email" autoComplete="email" onChange={event => setEmail(event.target.value)} />
                </FormControl>
                <FormControl margin="normal" required fullWidth>
                    <InputLabel htmlFor="password">Password</InputLabel>
                    <Input name="password" type="password" id="password" autoComplete="current-password" onChange={event => setPassword(event.target.value)} />
                </FormControl>
                <Button
                    component={Link}
                    fullWidth
                    variant="contained"
                    color="primary"
                    mt={2}
                    onClick={handleSubmit}
                >
                    Sign up
                </Button>
            </form>
        </Wrapper>
    );
}

export default withRouter(SignUp);
