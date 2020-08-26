import { Snackbar } from "@material-ui/core"
import React from "react"
import { Alert } from "../atoms/Alert"

export default ({ label, open, onClose }: { label?: string; open: boolean; onClose: () => void }) => (
    <Snackbar open={open} autoHideDuration={6000} onClose={onClose}>
        <Alert onClose={onClose} severity="success">
            This is a success message!
        </Alert>
    </Snackbar>
)