import React, {useReducer} from 'react'
import { createContainer } from "unstated-next"
import axios from "axios";

const authState = {
    token: null,
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
        let response = await axios.post("/signIn", signInRequest)
        dispatch({type: 'LOGIN_USER', payload: response.data.token})
    }

    const logout = () => {
        dispatch({type: 'LOGOUT_USER'})
    }

    const isLoggedIn = () => {
        return auth.token !== null;
    }

    return {
        user: auth,
        signUp,
        signIn,
        logout,
        isLoggedIn
    }
}

export const Auth = createContainer(useAuth);