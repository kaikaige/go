package form

type BadRequest string

func (e BadRequest) Error() string { return string(e) }

type NotFound string

func (e NotFound) Error() string { return string(e) }

type FormValidate string

func (e FormValidate) Error() string { return string(e) }
