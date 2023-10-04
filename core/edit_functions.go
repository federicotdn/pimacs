package core

func (ec *execContext) styledFormat(str lispObject, message bool, objects ...lispObject) (lispObject, error) {
	if !stringp(str) {
		return ec.wrongTypeArgument(ec.s.stringp, str)
	}

	format := []rune(xStringValue(str))
	result := []rune{}
	multibyte := false
	i := 0
	n := 0

	for i < len(format) {
		c := format[i]
		i++

		if c == '%' {
			if i >= len(format) {
				return ec.signalError("Format string ends in middle of format specifier")
			}

			conversion := format[i]
			i++

			if conversion == '%' {
				result = append(result, '%')
				continue
			}

			if n >= len(objects) {
				return ec.signalError("Not enough arguments for format string")
			}

			if conversion == 's' || conversion == 'S' {
				noEscape := ec.boolVal(conversion == 's')
				out, err := ec.prin1ToString(objects[n], noEscape, ec.nil_)
				if err != nil {
					return nil, err
				}

				result = append(result, []rune(xStringValue(out))...)
			}

			n++
			continue
		}

		multibyte = multibyte || !asciiCharp(c)
		result = append(result, c)
	}

	return newString(string(result), multibyte), nil
}

func (ec *execContext) formatMessage(args ...lispObject) (lispObject, error) {
	return ec.styledFormat(args[0], true, args[1:]...)
}

func (ec *execContext) format(args ...lispObject) (lispObject, error) {
	return ec.styledFormat(args[0], false, args[1:]...)
}

func (ec *execContext) symbolsOfEditFunctions() {
	ec.defSubrM(nil, "format", (*execContext).format, 1)
	ec.defSubrM(nil, "format-message", (*execContext).formatMessage, 1)
}
