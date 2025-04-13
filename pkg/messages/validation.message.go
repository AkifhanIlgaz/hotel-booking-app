package messages

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// MessageForTag returns a user-friendly error message for a given validation tag.
func MessageForTag(tag string, params ...string) string {
	switch tag {
	case "required":
		return "This field is required."
	case "min":
		if len(params) > 0 {
			return "This field must be at least " + params[0] + " characters long."
		}
		return "This field must have a minimum length."
	case "max":
		if len(params) > 0 {
			return "This field must be at most " + params[0] + " characters long."
		}
		return "This field must have a maximum length."
	case "email":
		return "Please enter a valid email address."
	default:
		return "Invalid tag."
	}
}
