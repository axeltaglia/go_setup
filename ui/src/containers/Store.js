import React from 'react'
import {Auth} from "./AuthContainer";
import {Alert} from "./AlertContainer";

const Store = ({children}) => {
    return (
        <Alert.Provider initialState={{open: false, text: ""}}>
            <Auth.Provider initialState={{firstName: null, lastName: null}}>
                {children}
            </Auth.Provider>
        </Alert.Provider>
    )
}

export default Store;