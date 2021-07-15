import React, {useReducer} from 'react'
import {createContainer} from "unstated-next"
import axios from "axios";

const authState = {
    token: window.localStorage.getItem("tokenGo"),
}

function authReducer(state = authState, action) {
    switch (action.type) {
        case "LOGIN_USER":
            return {
                token: action.payload,
            }
        case "LOGOUT_USER":
            return {
                token: null,
            }
        default:
            return state;
    }
}

function useAuth(initialState = authState) {
    const [auth, dispatch] = useReducer(authReducer, initialState);

    const signUp = async (signUpRequest) => {
        let response = await axios.post("/signUp", signUpRequest)
        dispatch({type: 'LOGIN_USER', payload: response.data.token})
    }

    const signIn = async (signInRequest) => {
        const response = await axios.post("/signIn", signInRequest)
        dispatch({type: 'LOGIN_USER', payload: response.data.token})
        window.localStorage.setItem("tokenGo", response.data.token);
    }

    const logout = () => {
        dispatch({type: 'LOGOUT_USER'})
        window.localStorage.removeItem("tokenGo");
    }

    const isLoggedIn = () => {
        return auth && auth.token !== null;
    }

    return {
        token: auth ? auth.token : null,
        signUp,
        signIn,
        logout,
        isLoggedIn
    }
}

export const AuthContainer = createContainer(useAuth);