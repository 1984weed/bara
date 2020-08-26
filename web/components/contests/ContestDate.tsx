import DateFnsUtils from "@date-io/date-fns"
import { Grid } from "@material-ui/core"
import { KeyboardDatePicker, KeyboardTimePicker, MuiPickersUtilsProvider } from "@material-ui/pickers"
import { convertTzDateToShiftedTzDate, convertTzDateToUtcDate } from "../../utils/dateshift"

type Props = {
    startTimeStamp: string
    onChange: (dateISOString: string) => void
}

export const ContestDate = ({ startTimeStamp, onChange }: Props) => (
    <MuiPickersUtilsProvider utils={DateFnsUtils}>
        <Grid container justify="space-around">
            <KeyboardDatePicker
                disableToolbar
                variant="inline"
                format="MM/dd/yyyy"
                margin="normal"
                id="date-picker-inline"
                label="Contest Date"
                value={convertTzDateToUtcDate(new Date(startTimeStamp))}
                onChange={date => {
                    const month = date.getMonth()
                    const day = date.getDate()

                    const newStartTime = new Date(`${startTimeStamp}`)
                    newStartTime.setMonth(month)
                    newStartTime.setDate(day)
                    onChange(convertTzDateToShiftedTzDate(newStartTime).toISOString())
                }}
            />
            <KeyboardTimePicker
                margin="normal"
                id="time-picker"
                label="Contest time"
                value={convertTzDateToUtcDate(new Date(startTimeStamp))}
                onChange={time => {
                    const newStartTime = new Date(`${startTimeStamp}`)

                    newStartTime.setHours(time.getHours())
                    newStartTime.setMinutes(time.getMinutes())

                    onChange(convertTzDateToShiftedTzDate(newStartTime).toISOString())
                }}
            />
        </Grid>
    </MuiPickersUtilsProvider>
)
