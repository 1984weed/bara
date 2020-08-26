export enum LogLevel {
	INFO = "info",
	WARN = "warn",
	ERROR = "error",
}
const levels = Object.keys(LogLevel)

export const logger = {
	logLevel: "info",
	info: function(message: any, data: any) {
		logger.log(LogLevel.INFO, message, data)
	},
	warn: function(message: any, data: any) {
		logger.log(LogLevel.WARN, message, data)
	},
	error: function(message: any, data: any) {
		logger.log(LogLevel.ERROR, message, data)
	},
    log: function(level: LogLevel, message: any, data: any) {
        if (levels.indexOf(level) <= levels.indexOf(logger.logLevel)) {
            if (typeof message !== "string") {
                message = JSON.stringify(message)
				console.log(`${level}: ${message}`)
				return 
            }
			console.log(`${level}: ${message} ${data || ""}`)
        }
    },
}
