import React, {useState} from 'react'
import { createContainer } from "unstated-next"

function useAlert() {
    const [open, setOpen] = useState(false)
    const [text, setText] = useState("")

    const alertSuccess = (text) => {
        setText(text)
        setOpen(true)
    }

    const closeAlert = () => {
        setText("")
        setOpen(false)
    }

    return {
        open,
        text,
        alertSuccess,
        closeAlert,
    }
}

export const Alert = createContainer(useAlert);