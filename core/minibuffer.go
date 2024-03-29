package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (ec *execContext) readFromMinibuffer(prompt, initialContents, keymap, read, hist, defaultValue, inheritInputMethod lispObject) (lispObject, error) {
	if !stringp(prompt) {
		return ec.wrongTypeArgument(ec.s.stringp, prompt)
	}

	if !ec.v.nonInteractive.val {
		return ec.pimacsUnimplemented(ec.s.read, "only noninteractive read is supported")
	}

	fmt.Print(xStringValue(prompt))

	reader := bufio.NewReader(os.Stdin)
	source, _ := reader.ReadString('\n')
	source = strings.TrimRight(source, "\r\n")

	result, err := ec.readFromString(newUniOrMultibyteString(source), ec.nil_, ec.nil_)
	if err != nil {
		return nil, err
	}

	return xCar(result), nil
}

func (ec *execContext) symbolsOfMinibuffer() {
	ec.defSubr7(&ec.s.readFromMinibuffer, "read-from-minibuffer", (*execContext).readFromMinibuffer, 1)
}
