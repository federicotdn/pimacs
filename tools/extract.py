from pathlib import Path
from collections import namedtuple
import os
from pyparsing import Literal, Regex, QuotedString, SkipTo, Opt

GO_KEYWORDS = [
    "break",
    "default",
    "func",
    "interface",
    "select",
    "case",
    "defer",
    "go",
    "map",
    "struct",
    "chan",
    "else",
    "goto",
    "package",
    "switch",
    "const",
    "fallthrough",
    "if",
    "range",
    "type",
    "continue",
    "for",
    "import",
    "return",
    "var",
]

DEFUN_EXPR_IDENTIFIER = Regex(r"[0-9a-zA-Z_-]+")
DEFUN_EXPR_NUMBER = Regex(r"[0-9]+")
DEFUN_EXPR = (
    Literal("DEFUN")
    + "("
    + QuotedString('"')("lname")
    + ","
    + DEFUN_EXPR_IDENTIFIER("fnname")
    + ","
    + DEFUN_EXPR_IDENTIFIER("sname")
    + ","
    + (DEFUN_EXPR_NUMBER | DEFUN_EXPR_IDENTIFIER)("minargs")
    + ","
    + (DEFUN_EXPR_NUMBER | "MANY" | "UNEVALLED")("maxargs")
    + ","
    + (DEFUN_EXPR_NUMBER | QuotedString('"', esc_char="\\", multiline=True) | "NULL")(
        "intspec"
    )
    + ","
    + "doc:"
    + "/*"
    + SkipTo("*/")("doc")
    + "*/"
    + Opt("attributes:" + (Literal("noreturn") | "const")("attributes"))
    + ")"
)


def extract_symbols(path: Path) -> tuple:
    with open(path) as f:
        contents = f.read()

    lines = contents.splitlines()
    current_defun = []
    defuns = []

    for line in lines:
        if line.startswith("DEFUN"):
            if current_defun:
                raise Exception(
                    "Previous DEFUN not closed: " + "\n".join(current_defun)
                )

            current_defun.append(line)
        elif line.startswith("{"):
            if not current_defun:
                continue

            defun = "\n".join(current_defun)
            current_defun.clear()

            parsed = DEFUN_EXPR.parse_string(defun)
            result = dict(parsed)
            result["file"] = path.stem
            defuns.append(result)
        elif current_defun:
            current_defun.append(line)

    if current_defun:
        raise Exception("Previous DEFUN not closed: " + "\n".join(current_defun))

    return (defuns,)


def to_camel_case(text: str) -> str:
    s = text.replace("-", " ").replace("_", " ")
    s = s.split()
    if len(text) == 0:
        return text
    return s[0] + "".join(i.capitalize() for i in s[1:])


def fnname_to_go(fnname: str) -> str:
    name = to_camel_case(fnname[1:])
    if name in GO_KEYWORDS:
        name += "_"
    return name + "_autogen"


def generate_go_file(defuns: list[dict]) -> str:
    GoDefun = namedtuple("GoDefun", "subr_fn numargs minargs lname fnname")
    go_defuns = {}
    s = "package elisp\n\n"

    for defun in defuns:
        lname = defun["lname"]

        try:
            maxargs = int(defun["maxargs"])
            if maxargs < 0 or maxargs > 8:
                raise ValueError(defun["maxargs"])

            subr_fn = "defSubr" + str(maxargs)
            numargs = maxargs
        except ValueError:
            if defun["maxargs"] == "MANY":
                subr_fn = "defSubrM"
                numargs = None
            elif defun["maxargs"] == "UNEVALLED":
                subr_fn = "defSubrU"
                numargs = 1
            else:
                raise ValueError(defun["maxargs"])

        try:
            minargs = int(defun["minargs"])
        except ValueError:
            minargs = "emacsConstant_" + defun["minargs"]

        fnname = fnname_to_go(defun["fnname"])
        go_defun = GoDefun(subr_fn, numargs, minargs, lname, fnname)

        if lname in go_defuns:
            # if hash(go_defuns[lname]) != hash(go_defun):
            #     raise Exception(
            #         f"Duplicated defun with different properties",
            #         go_defun,
            #         go_defuns[lname],
            #     )

            continue

        go_defuns[lname] = go_defun

    for defun in go_defuns.values():
        s += f"func (ec *execContext) {defun.fnname}("
        if defun.subr_fn == "defSubrM":
            s += "args ...lispObject"
        elif defun.subr_fn != "defSubr0":
            s += ", ".join([f"arg{i}" for i in range(defun.numargs)]) + " lispObject"
        s += ") (lispObject, error) {\n"
        s += "	return ec.nil_, nil\n"
        s += "}\n\n"

    s += "func (ec *execContext) symbolsOfEmacs() {\n"

    for defun in go_defuns.values():
        s += f'	ec.{defun.subr_fn}("{defun.lname}", ec.{defun.fnname}'
        if defun.subr_fn != "defSubr0":
            s += f", {defun.minargs}"
        s += ")\n"

    s += "}"

    return s


def main() -> None:
    base = Path(os.environ["EMACS_REPO_ROOT"])

    defuns = []

    for p in base.joinpath("src").rglob("*"):
        if not p.suffix == ".c" or p.name == "msdos.c":
            continue

        file_defuns, *_ = extract_symbols(p)
        defuns.extend(file_defuns)

    f = generate_go_file(defuns)
    print(f)


if __name__ == "__main__":
    main()
