import { Box, Grid, makeStyles, Typography } from "@material-ui/core"
import { NextPage } from "next"
import Link from "next/link"
import NextRouter, { Router, withRouter } from "next/router"
import React, { FunctionComponent } from "react"
import { Session } from "../types/Session"
import { HeaderIcon } from "./icons/HeaderIcon"
import { useSession } from "../lib/session"
const { NextAuth } = require("next-auth/client")

type Props = {
    router: Router
}

const useLinkStyles = makeStyles(theme => ({
    link: {
        padding: `0 ${theme.spacing(2)}px`,
        fontSize: "1rem",
        textDecoration: "none",
    },
}))

const StyledLink: FunctionComponent<{ className?: string; href: string; style?: any; onClick?: () => void }> = ({
    href,
    style,
    className,
    children,
    onClick,
}) => {
    const classes = useLinkStyles()
    return (
        <>
            <Link href={href}>
                <a style={{ ...style }} className={`${className} ${classes.link}`} onClick={onClick}>
                    {children}
                </a>
            </Link>
        </>
    )
}

const useStyles = makeStyles(theme => ({
    header: {
        height: "66px",
        borderBottom: `1px solid ${theme.palette.primary.main}`,
    },
    logo: {
        fontWeight: "bold",
    },
}))

const Header: NextPage<Props> = ({ router: { pathname } }: Props) => {
    const [session] = useSession()
    const classes = useStyles()
    return (
        <Grid container className={classes.header} justify="space-between" alignItems="center">
            <Grid item>
                <StyledLink href="/" className={pathname === "/" ? "is-active" : ""}>
                    Problems
                </StyledLink>
                <StyledLink href="/contests" className={pathname === "/contests" ? "is-active" : ""}>
                    Contests
                </StyledLink>
                {session?.user?.role === "admin" && (
                    <StyledLink href="/admin" className={pathname === "/admin" ? "is-active" : ""}>
                        Admin
                    </StyledLink>
                )}
            </Grid>

            <Box className="header-icon" style={{ cursor: "pointer" }}>
                <Link href="/">
                    <a style={{ display: "flex", alignItems: "center", textDecoration: "none", color: "#54DF6A" }}>
                        <HeaderIcon />
                        <Box pl={1}>
                            <Typography className={classes.logo} variant="h6" component="h1">
                                Baracode
                            </Typography>
                        </Box>
                    </a>
                </Link>
            </Box>
            {(() => {
                if (session.user) {
                    return (
                        <Box>
                            <StyledLink
                                href="#"
                                onClick={() => {
                                    fetch("/auth/signout", {method: "post"}).then(() => {
                                        NextRouter.push("/")
                                    })
                                }}
                            >
                                Sign out
                            </StyledLink>
                        </Box>
                    )
                }
                return (
                    <Box display="flex">
                        <Box>
                            <StyledLink href="/signin" style={{ fontWeight: "bold" }} className={pathname === "/signin" ? "is-active" : ""}>
                                Log in
                            </StyledLink>
                        </Box>
                        <Box>
                            <StyledLink href="/signup" style={{ fontWeight: "bold" }} className={pathname === "/signup" ? "is-active" : ""}>
                                Sign up
                            </StyledLink>
                        </Box>
                    </Box>
                )
            })()}
        </Grid>
    )
}

export default withRouter(Header)
