import { Box, Text, Paragraph } from 'grommet';
import { Status } from '../../types/Contraints';
import { CodeLanguage } from '../../graphql/types';

type Props = {
    status: Status; 
    expected: string;
    time: string;
    result: string;
    language: CodeLanguage;
    title: string;
}

function statusSlugToLabel(status: Status): string {
    if(status === "success") {
        return "Success"
    }

    if(status === "fail") {
        return "Wrong Answer"
    }

    return "Time out"
}

function languageSlugToLabel(language: CodeLanguage): string {
    if(language === CodeLanguage.JavaScript) {
        return "JavaScript"
    }
    return "no lang"
}

const SubmittedResult: React.FunctionComponent<Props> = props => 
    <Box align='center'>
        <Text color='accent-3'>{statusSlugToLabel(props.status)}</Text>
        <Paragraph>
            {props.status === "success" && <Text>Runtime: {props.time} ms, {languageSlugToLabel(props.language)} online submissions for {props.title}</Text>}
            {props.status === "fail" && <Text>
                It failed
                <PrintDetail title='Input' detail='[9,9,9,9]' />
                <PrintDetail title='Output' detail={props.result} />
                <PrintDetail title='Expected' detail={props.expected} />
            </Text>}
        </Paragraph>
    </Box>  

const PrintDetail = ({title, detail}: {title: string; detail: string}) =>
    <Box direction='row'>
        <Box margin='xxsmall' width='50px'>{title}</Box>
        <Box margin='xxsmall' background={{color: 'light-2'}} flex='grow'>{detail}</Box>
    </Box>

export default SubmittedResult