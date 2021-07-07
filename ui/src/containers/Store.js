import React from 'react'
import {Auth} from "./AuthContainer";
import {Alert} from "./AlertContainer";

const Store = ({children}) => {
    return (
        <Alert.Provider>
            <Auth.Provider>
                {children}
            </Auth.Provider>
        </Alert.Provider>
    )
}

export default Store;