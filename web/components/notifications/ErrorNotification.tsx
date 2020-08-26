import { Snackbar } from "@material-ui/core"
import React from "react"
import { Alert } from "../atoms/Alert"

export default ({ onClose, open }: { onClose: () => void, open?: boolean }) => (
    <Snackbar open={open} autoHideDuration={6000} onClose={onClose}>
        <Alert onClose={onClose} severity="error">
            This is a success message!
        </Alert>
    </Snackbar>
)
