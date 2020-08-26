import Button from "@material-ui/core/Button"
import React from "react"
import GitHubIcon from "@material-ui/icons/GitHub"

type Props = {
    label: string
}

export default ({ label }: Props) => (
    <Button
        type="button"
        fullWidth
        style={{
            backgroundColor: "#333333",
            textTransform: "none",
            color: "#fff",
        }}
        startIcon={<GitHubIcon />}
        href="/auth/oauth/github"
    >
        {label}
    </Button>
)
