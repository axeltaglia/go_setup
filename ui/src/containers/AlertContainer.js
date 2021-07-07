import React, {useReducer} from 'react'
import { createContainer } from "unstated-next"

const initialState = {
    open: false,
    text: ""
}

function authReducer(state = initialState, action) {
    switch (action.type) {
        case "SET_OPEN":
            return {
                ...state,
                open: action.payload,
            }
        case "SET_TEXT":
            return {
                ...state,
                text: action.payload,
            }
        default:
            return state;
    }
}

function useAlert(initialState) {
    const [state, dispatch] = useReducer(authReducer, initialState);

    const alertSuccess = (text) => {
        dispatch({type: 'SET_OPEN', payload: true})
        dispatch({type: 'SET_TEXT', payload: text})
    }

    const closeAlert = () => {
        dispatch({type: 'SET_OPEN', payload: false})
    }


    return {
        open: state.open,
        text: state.text,
        alertSuccess,
        closeAlert,
    }
}

export const Alert = createContainer(useAlert);