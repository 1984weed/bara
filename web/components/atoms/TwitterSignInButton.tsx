import Button from "@material-ui/core/Button"
import React from "react"
import TwitterIcon from "@material-ui/icons/Twitter"

type Props = {
    label: string
}

export default ({ label }: Props) => (
    <Button
        type="button"
        fullWidth
        style={{
            backgroundColor: "#38A1F3",
            textTransform: "none",
            color: "#fff",
        }}
        startIcon={<TwitterIcon />}
        href="/auth/oauth/twitter"
    >
        {label}
    </Button>
)
