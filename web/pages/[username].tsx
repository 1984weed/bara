import { NextPage } from "next"
import Error from "next/error"
import { useRouter } from "next/router"
import React, { useState, ReactNode, ChangeEvent } from "react"
import { toBase64 } from "../base64"
import Layout from "../components/Layout"
import { User } from "../graphql/types"
import { NextPageContextWithGraphql } from "../lib/with-graphql-client"
import { Avatar, makeStyles, Box, Grid, Typography, Button, TextField, Dialog, DialogTitle } from "@material-ui/core"
import { useForm, ErrorMessage } from "react-hook-form"
import PhotoIcon from "@material-ui/icons/Photo"
import { useMutation } from "graphql-hooks"

type Props = {
    session: any
    userData?: User
    pathname: string
}

const getUserQuery = `
query getUser($userName: String!) {
    user(userName: $userName) {
        id,
        displayName,
        userName,
        email,
        image,
        bio
    }
}
`

const updateUserMutation = `
mutation updateMe($userName: String, $displayName: String, $email: String, $image: String, $bio: String) {
    updateMe(input: {userName: $userName, displayName: $displayName, email: $email, image: $image, bio: $bio}) {
        displayName,
        userName,
        email,
        image,
        bio
    }
}
`

const useStyles = makeStyles(theme => ({
    avatar: {
        width: "100px",
        height: "100px",
    },
    uploadIcon: {
        position: "absolute",
        border: "5px solid white",
        borderRadius: "15px",
        background: "rgba(0, 0, 0, 0.65)",
        cursor: "pointer",
        color: "#bdc3c7",
    },
    avatarContainer: {
        position: "relative",
    },
    profileWidth: {
        width: "500px",
    },
    changePassword: {
        color: "#0070F3",
        cursor: "pointer",
        textDecoration: "underline",
    },
    button: {
        margin: theme.spacing(1),
    },
}))

type AvatarEditDialogProps = {
    open: boolean
    selectedValue?: string
    imageUrl: string
    onClose: (value: string) => void
}

const MAX_FILE_SIZE = 5242880

const readFile = (file: File): Promise<{ file: File; dataURL: string }> => {
    return new Promise((resolve, reject) => {
        const reader = new FileReader()

        // Read the image via FileReader API and save image result in state.
        reader.onload = function(e) {
            // Add the file name to the data URL
            let dataURL = e.target.result as string
            dataURL = dataURL.replace(";base64", `;name=${file.name};base64`)
            resolve({ file, dataURL })
        }

        reader.readAsDataURL(file)
    })
}

const hasExtension = (fileName: string): boolean => {
    const pattern = "(" + [".jpg", ".jpeg", ".gif", ".png"].join("|").replace(/\./g, "\\.") + ")$"
    return new RegExp(pattern, "i").test(fileName)
}

const AvatarEditDialog = ({ onClose, imageUrl, open }: AvatarEditDialogProps) => {
    const [image, setImage] = useState(imageUrl)

    const handleClose = () => {
        onClose(null)
    }

    const onDropFile = async (e: ChangeEvent<HTMLInputElement>) => {
        const files = (e.target as HTMLInputElement).files
        if (files.length !== 1) {
            return
        }
        let file = files[0]
        let fileError = {
            name: file.name,
        }
        const fileErrors = []
        // Extension check
        if (!hasExtension(file.name)) {
            fileError = Object.assign(fileError, {
                message: "Not supported extension",
            })
            fileErrors.push(fileError)
        }
        if (file.size > MAX_FILE_SIZE) {
            fileError = Object.assign(fileError, {
                message: "File size too large",
            })
            fileErrors.push(fileError)
        }

        const { dataURL } = await readFile(file)

        setImage(dataURL)
    }

    const onUploadClick = () => {
        inputElement.click()
    }

    let inputElement: HTMLInputElement

    return (
        <Dialog
            onClose={() => {
                setTimeout(() => handleClose())
            }}
            open={open}
        >
            <DialogTitle>Upload a New Avatar</DialogTitle>
            <Box display="flex" alignContent="center" justifyContent="center">
                <Box display="flex" height={200} width={200} style={{ position: "relative" }}>
                    <img src={image} height="100%" width="100%" />
                </Box>
            </Box>
            <input
                type="file"
                ref={input => (inputElement = input)}
                name="avatar"
                onChange={onDropFile}
                onClick={onUploadClick}
                accept="image/*,.png,.jpg"
                style={{ display: "none" }}
            />
            <Box>
                <Button onClick={onUploadClick}>Choose image</Button>
                <Box>
                    <Button
                        onClick={() => {
                            setTimeout(() => {
                                onClose(null)
                            })
                        }}
                    >
                        Cancel
                    </Button>
                    <Button
                        onClick={() => {
                            setTimeout(() => {
                                onClose(image)
                            })
                        }}
                    >
                        Update
                    </Button>
                </Box>
            </Box>
        </Dialog>
    )
}

const user: NextPage<Props> = ({ userData, session }: Props) => {
    const router = useRouter()
    const [isEdit, setIsEdit] = useState(true)
    const [isHoverAvatar, setIsHoverAvatar] = useState(false)
    const [openEditAvatar, setOpenEditAvatar] = useState(false)

    const [updateUser] = useMutation(updateUserMutation)
    const [userDataState, setUserData] = useState(userData)

    const classes = useStyles()
    const userProfileWidth = 500

    const { handleSubmit, register, errors, setValue, watch, reset } = useForm({
        defaultValues: {
            userName: userDataState.displayName,
            userID: userDataState.userName,
            email: userDataState.email,
            imageURL: userDataState.image,
        },
    })

    const onClickUserEditSubmit = async formData => {
        const updateData: { userName: string; displayName: string; email: string; image: string } = {
            userName: null,
            displayName: null,
            email: null,
            image: null,
        }
        if (formData.userName !== userDataState.displayName) {
            updateData.displayName = formData.userName
        }
        if (formData.userID !== userDataState.userName) {
            updateData.userName = formData.userName
        }
        if (formData.email !== userDataState.email) {
            updateData.email = formData.email
        }
        if (formData.imageURL !== userDataState.image) {
            updateData.image = formData.imageURL
        }
        const { data, error } = await updateUser({
            variables: updateData,
        })
        if(error) {
            return 
        }
        const updateUserData = data.updateMe

        reset({
            userName: updateUserData.displayName,
            userID: updateUserData.userName,
            email: updateUserData.email,
            imageURL: updateUserData.image,
        })
        setUserData(updateUserData)
        setIsEdit(false)
    }
    const handleAvatarEditClickOpen = () => {
        setOpenEditAvatar(true)
    }

    const handleAvatarEditClickClose = (imageValue: string | null) => {
        if (imageValue != null) {
            setValue("imageURL", imageValue)
        }

        setOpenEditAvatar(false)
        setIsHoverAvatar(false)
    }
    register({ name: "imageURL" }, { required: true })
    const imageURL = watch("imageURL") as string
    const canEdit = session && `${session.user}` === userData.id

    if (router == null) {
        return <></>
    }

    if (userData == null) {
        return <Error statusCode={404} />
    }
    return (
        <Layout session={session}>
            <Box display="flex" flexDirection="column" justifyContent="center" alignItems="center">
                <Box width={userProfileWidth} display="flex" alignItems="center" justifyContent="space-between">
                    <Box display="flex" alignItems="center">
                        <Box mr={1}>
                            {(() => {
                                if (isEdit) {
                                    return (
                                        <Box
                                            display="flex"
                                            className={classes.avatarContainer}
                                            onMouseEnter={() => setIsHoverAvatar(true)}
                                            onMouseLeave={() => setIsHoverAvatar(false)}
                                            onClick={handleAvatarEditClickOpen}
                                        >
                                            <Avatar className={classes.avatar} src={imageURL} />
                                            {isHoverAvatar && (
                                                <Box
                                                    width="100%"
                                                    height="100%"
                                                    display="flex"
                                                    flexDirection="column"
                                                    alignItems="center"
                                                    justifyContent="center"
                                                    className={classes.uploadIcon}
                                                >
                                                    <PhotoIcon />
                                                    <Typography variant="body1">Edit</Typography>
                                                </Box>
                                            )}
                                            <AvatarEditDialog imageUrl={imageURL} open={openEditAvatar} onClose={handleAvatarEditClickClose} />
                                        </Box>
                                    )
                                }
                                return <Avatar className={classes.avatar} src={imageURL} />
                            })()}
                        </Box>
                        <Box display="flex" flexDirection="column">
                            <Typography variant="h3" component="span">
                                {userDataState.displayName}
                            </Typography>
                            <Typography variant="body1" component="span">
                                {userDataState.userName}
                            </Typography>
                        </Box>
                    </Box>
                    {canEdit && !isEdit && (
                        <Box>
                            <Button
                                variant="outlined"
                                onClick={() => {
                                    setIsEdit(true)
                                }}
                            >
                                Edit
                            </Button>
                        </Box>
                    )}
                </Box>
                {canEdit && (
                    <Box p={2} width={userProfileWidth}>
                        {(() => {
                            if (isEdit) {
                                return (
                                    <form onSubmit={handleSubmit(onClickUserEditSubmit)}>
                                        <ProfileInfoRow
                                            labelChild={
                                                <Typography variant="body1" component="span">
                                                    Baracode ID:
                                                </Typography>
                                            }
                                            realValueChild={
                                                <>
                                                    <TextField
                                                        variant="outlined"
                                                        margin="none"
                                                        size="small"
                                                        id="userID"
                                                        name="userID"
                                                        autoComplete="Baracode ID"
                                                        fullWidth
                                                        inputRef={register({ required: true })}
                                                    />
                                                    <ErrorMessage errors={errors} name="userID" />
                                                </>
                                            }
                                        />
                                        <ProfileInfoRow
                                            labelChild={
                                                <Typography variant="body1" component="span">
                                                    Name:
                                                </Typography>
                                            }
                                            realValueChild={
                                                <>
                                                    <TextField
                                                        variant="outlined"
                                                        margin="none"
                                                        size="small"
                                                        id="userName"
                                                        name="userName"
                                                        autoComplete="Name"
                                                        fullWidth
                                                        inputRef={register({ required: true })}
                                                    />
                                                    <ErrorMessage errors={errors} name="userName" />
                                                </>
                                            }
                                        />
                                        <ProfileInfoRow
                                            labelChild={
                                                <Typography variant="body1" component="span">
                                                    Email:
                                                </Typography>
                                            }
                                            realValueChild={
                                                <>
                                                    <TextField
                                                        variant="outlined"
                                                        margin="none"
                                                        size="small"
                                                        id="email"
                                                        name="email"
                                                        autoComplete="Email"
                                                        fullWidth
                                                        inputRef={register({ required: true })}
                                                    />
                                                    <ErrorMessage errors={errors} name="email" />
                                                </>
                                            }
                                        />
                                    </form>
                                )
                            }
                            return (
                                <>
                                    <ProfileInfoRow
                                        labelChild={
                                            <Typography variant="body1" component="span">
                                                Email:
                                            </Typography>
                                        }
                                        realValueChild={
                                            <Typography variant="body1" component="span">
                                                {userDataState.email}
                                            </Typography>
                                        }
                                    />
                                    <ProfileInfoRow
                                        labelChild={
                                            <Typography variant="body1" component="span">
                                                Password:
                                            </Typography>
                                        }
                                        realValueChild={
                                            <Typography className={classes.changePassword} variant="body1" component="span">
                                                Change password
                                            </Typography>
                                        }
                                    />
                                </>
                            )
                        })()}

                        {isEdit && (
                            <Box p={2} width={userProfileWidth} display="flex" justifyContent="flex-end">
                                <Button
                                    variant="contained"
                                    onClick={() => {
                                        reset()
                                        setIsEdit(false)
                                    }}
                                    className={classes.button}
                                >
                                    Cancel
                                </Button>
                                <Button
                                    variant="contained"
                                    type="submit"
                                    className={classes.button}
                                    color="primary"
                                    onClick={handleSubmit(onClickUserEditSubmit)}
                                >
                                    Update
                                </Button>
                            </Box>
                        )}
                    </Box>
                )}
            </Box>
        </Layout>
    )
}

type RorProfileInfoProps = {
    labelChild: ReactNode
    realValueChild: ReactNode
}
const makeProfileRowStyle = makeStyles(theme => ({
    box: {
        borderBottom: `1px solid #c4c4c4`,
    },
}))

const ProfileInfoRow = ({ labelChild, realValueChild }: RorProfileInfoProps) => {
    const classes = makeProfileRowStyle()
    return (
        <Box p={2} className={classes.box} width="100%" display="flex" alignContent="space-between">
            <Box width={100} display="flex" alignItems="center">
                {labelChild}
            </Box>
            <Box marginLeft={3}>{realValueChild}</Box>
        </Box>
    )
}

user.getInitialProps = async ({ query, client }: NextPageContextWithGraphql) => {
    const result = await client.request(
        {
            query: getUserQuery,
            variables: { userName: query.username },
        },
        {}
    )

    const { data } = result

    return Promise.resolve({
        userData: data["user"] || {},
        session: "",
        pathname: "",
    })
}

export default user
