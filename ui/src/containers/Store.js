import React from 'react'
import {Auth} from "./AuthContainer";

const Store = ({children}) => {
    return (
        <Auth.Provider initialState={{firstName:  null, lastName: null}}>
            {children}
        </Auth.Provider>
    )
}

export default Store;