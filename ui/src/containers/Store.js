import React from 'react'
import {AuthContainer} from "./AuthContainer";
import {AlertContainer} from "./AlertContainer";

const Store = ({children}) => {
    return (
        <AlertContainer.Provider>
            <AuthContainer.Provider>
                {children}
            </AuthContainer.Provider>
        </AlertContainer.Provider>
    )
}

export default Store;