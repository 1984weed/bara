export const convertTzDateToShiftedTzDate = (pickerDate: Date) => {
	let tzDate = new Date()
	tzDate.setTime(pickerDate.getTime() - tzDate.getTimezoneOffset() * 60000)

	return tzDate
}

export const convertTzDateToUtcDate = (tzDate: Date) => {
	let utcDate = new Date()
	utcDate.setTime(tzDate.getTime() + tzDate.getTimezoneOffset() * 60000)

	return utcDate
}