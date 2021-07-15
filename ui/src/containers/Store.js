import React from 'react'
import {AuthContainer} from "./AuthContainer";
import {Alert} from "./AlertContainer";

const Store = ({children}) => {
    return (
        <Alert.Provider>
            <AuthContainer.Provider>
                {children}
            </AuthContainer.Provider>
        </Alert.Provider>
    )
}

export default Store;