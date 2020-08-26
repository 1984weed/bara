import { Status } from "../../types/Contraints"
import { CodeLanguage } from "../../graphql/types"
import { Typography, Box } from "@material-ui/core"

type Props = {
    status: Status
    expected: string
    input: string
    time?: string
    result: string
    stdout: string
    language: CodeLanguage
    title: string
    onlyRun: boolean
}

function statusSlugToDisplay(status: Status): { label: string; color: string } {
    if (status === "success") {
        return { label: "Success", color: "accent-3" }
    }

    if (status === "fail") {
        return { label: "Wrong Answer", color: "status-error" }
    }

    return { label: "Wrong Answer", color: "status-error" }
}

function languageSlugToLabel(language: CodeLanguage): string {
    if (language === CodeLanguage.JavaScript) {
        return "JavaScript"
    }
    return "no lang"
}

const SubmittedResult: React.FunctionComponent<Props> = props => {
    const { label, color } = statusSlugToDisplay(props.status)

    return (
        <Box>
            <Box>
                <Typography variant="body1" style={{ color: color }}>
                    {label}
                </Typography>
            </Box>
            <Box>
                {!props.onlyRun && props.status === "success" && props.time != null && (
                    <Box>
                        <Typography variant="body1">
                            Runtime: {props.time} ms, {languageSlugToLabel(props.language)} online submissions for {props.title}
                        </Typography>
                    </Box>
                )}
                {(props.onlyRun || props.status === "fail") && (
                    <Box>
                        <PrintDetail title="Input" detail={props.input} />
                        {props.stdout && props.stdout != "" && <PrintDetail title="Stdout" detail={props.stdout} />}
                        <PrintDetail title="Output" detail={props.result} />
                        <PrintDetail title="Expected" detail={props.expected} />
                    </Box>
                )}
            </Box>
        </Box>
    )
}

const PrintDetail = ({ title, detail }: { title: string; detail: string }) => {
    return (
        <Box display="flex" flexDirection="row">
            <Box width="100px">{title}</Box>
            <Box style={{ whiteSpace: "pre-wrap", overflow: "scroll", maxHeight: "70px" }}>{detail}</Box>
        </Box>
    )
}

export default SubmittedResult
