import React, { useState } from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";
import axios from 'axios';
import { useSelector, useDispatch } from 'react-redux';
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
import { spacing } from "@material-ui/system";
import {persistAuthToken} from "../../app/auth";

const Button = styled(MuiButton)(spacing);

const Wrapper = styled(Paper)`
  padding: ${props => props.theme.spacing(6)}px;

  ${props => props.theme.breakpoints.up("md")} {
    padding: ${props => props.theme.spacing(10)}px;
  }
`;

function SignIn() {
  const dispatch = useDispatch();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async () => {
    try {
      let response = await axios.post("/login", {"username": email, "password": password})
      persistAuthToken(response.data.token)
      //dispatch(setCurrentUser(response.data.token))
    } catch (e) {
      console.log(e)
    }

  }

  return (
    <Wrapper>
      <Helmet title="Sign In" />

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
          control={<Checkbox value="remember" color="primary" />}
          label="Remember me"
        />
        <Button
          component={Link}
          fullWidth
          variant="contained"
          color="primary"
          mb={2}
          onClick={handleSubmit}
        >
          Sign in
        </Button>
      </form>
    </Wrapper>
  );
}

export default SignIn;
