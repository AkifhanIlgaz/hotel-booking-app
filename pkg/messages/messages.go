package messages

const (
	SomethingWentWrong         string = "Something went wrong. Please try again later."
	InvalidJSONOrMissingFields string = "The information you submitted is incomplete or invalid. Please check the form."
	EmailAlreadyRegistered     string = "This email address is already in use. Try logging in instead."
	UserNotFound               string = "No user found with the given information. Please double-check your details."
	WrongPassword              string = "Incorrect password. Please try again."
	SuccessfullyRegistered     string = "Registration successful! You can now log in."
	SuccessfullyLoggedIn       string = "Logged in successfully."
	SuccessfullyLoggedOut      string = "Logged out successfully."
	TokenExpired               string = "Your session has expired. Please log in again."
	TokenNotFound              string = "Token not found. Please log in again."
	SentOTPCode                string = "Check your inbox! We’ve just sent you a code to reset your password."
	InvalidToken               string = "Invalid Token" // ! change
	InvalidAuthHeader          string = "Invalid Authorization Header"
	InvalidOTP                 string = "Invalid OTP. Please try again."
)
