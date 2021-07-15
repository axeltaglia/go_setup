import React from 'react'
import MuiAlert from "@material-ui/lab/Alert";
import Snackbar from "@material-ui/core/Snackbar";
import {AlertContainer} from "../containers/AlertContainer";

const AlertMessage = () => {
    const {
        open,
        text,
        closeAlert
    } = AlertContainer.useContainer()

    const handleClose = () => {
        closeAlert()
    }

    return (
        <Snackbar open={open} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{vertical: "bottom", horizontal: "right"}} >
            <MuiAlert elevation={6} variant="filled" onClose={handleClose} severity="success">
                {text}
            </MuiAlert>
        </Snackbar>
    )
}

export default AlertMessage;