from pathlib import Path
from collections import defaultdict
from dataclasses import dataclass

from pyparsing import (
    Literal,
    Regex,
    QuotedString,
    SkipTo,
    Opt,
    StringStart,
    StringEnd,
    delimited_list,
)

FILE_HEADER = """// Automatically generated with pimacs.tools.extract
// DO NOT MODIFY!
// Generated from GNU Emacs commit: {}, branch: {}

"""
AUTOGEN_SUFFIX = "_autogen"
PACKAGE = "core"
GO_KEYWORDS = [
    "break",
    "case",
    "chan",
    "const",
    "continue",
    "default",
    "defer",
    "else",
    "fallthrough",
    "for",
    "func",
    "go",
    "goto",
    "if",
    "import",
    "interface",
    "map",
    "package",
    "range",
    "return",
    "select",
    "struct",
    "switch",
    "type",
    "var",
]
EMACS_CONSTANTS = {
    "coding_arg_max": 13,
    "charset_arg_max": 17,
}

DEFUN_EXPR_IDENTIFIER = Regex(r"[0-9a-zA-Z_-]+")
DEFUN_EXPR_NUMBER = Regex(r"[0-9]+")
DEFUN_EXPR = (
    StringStart()
    + Literal("DEFUN")
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
    + (
        (
            Literal("(")
            + (
                delimited_list(
                    Opt("register") + "Lisp_Object" + DEFUN_EXPR_IDENTIFIER("args*"),
                    min=1,
                )
                | "void"
                | Literal("ptrdiff_t nargs, Lisp_Object *args")
                | Literal("ptrdiff_t n, Lisp_Object *args")
            )
            + ")"
        )
        | StringEnd()
    )
)


@dataclass
class EmacsDefun:
    lname: str
    fnname: str
    sname: str
    minargs: str
    maxargs: str
    intspec: str
    args: list[str]
    attributes: str
    doc: str
    path: str


@dataclass
class GoDefun:
    subr_fn: str
    numargs: str
    minargs: str
    lname: str
    fnname: str
    args: list[str]
    attributes: str
    path: str


@dataclass
class GoConstant:
    name: str
    value: str


def extract_declarations(path: Path) -> tuple[list[EmacsDefun]]:
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

            try:
                parsed = DEFUN_EXPR.parse_string(defun, parse_all=True)
            except Exception:
                print(f"----- context: ----- (file: {path})")
                print(defun)
                print("--------------------")
                raise

            result = {
                "lname": parsed["lname"],
                "fnname": parsed["fnname"],
                "sname": parsed["sname"],
                "minargs": parsed["minargs"],
                "maxargs": parsed["maxargs"],
                "intspec": parsed["intspec"],
                "args": parsed.get("args"),
                "attributes": parsed.get("attributes"),
                "doc": parsed["doc"].strip(),
                "path": str(path.name),
            }

            defuns.append(EmacsDefun(**result))
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


def fnname_to_go(fnname: str, count: int) -> str:
    return to_camel_case(fnname[1:]) + (f"_{count}" if count else "") + AUTOGEN_SUFFIX


def go_safe_identifier(identifier: str) -> str:
    return identifier if identifier not in GO_KEYWORDS else identifier + "_"


def get_go_declarations(
    defuns: list[EmacsDefun],
) -> tuple[list[GoDefun], list[GoConstant]]:
    go_defuns = defaultdict(list)
    go_constants = []

    for defun in defuns:
        try:
            maxargs = int(defun.maxargs)
            if maxargs < 0 or maxargs > 8:
                raise ValueError(defun.maxargs)

            subr_fn = "defSubr" + str(maxargs)
            numargs = maxargs
        except ValueError:
            if defun.maxargs == "MANY":
                subr_fn = "defSubrM"
                numargs = None
            elif defun.maxargs == "UNEVALLED":
                subr_fn = "defSubrU"
                numargs = 1
            else:
                raise ValueError(defun.maxargs)

        try:
            minargs = int(defun.minargs)
        except ValueError:
            minargs = "emacsConstant_" + defun.minargs + AUTOGEN_SUFFIX
            go_constants.append(
                GoConstant(minargs, str(EMACS_CONSTANTS[defun.minargs]))
            )

        lname = defun.lname
        args = None
        if defun.args is not None:
            args = [go_safe_identifier(arg) for arg in defun.args]

        count = len(go_defuns[lname])
        fnname = fnname_to_go(defun.fnname, count)
        attributes = defun.attributes
        path = defun.path

        go_defun = GoDefun(
            subr_fn, numargs, minargs, lname, fnname, args, attributes, path
        )
        go_defuns[lname].append(go_defun)

    return sum(go_defuns.values(), []), go_constants


def generate_go_file(
    defuns: list[EmacsDefun], emacs_commit: str, emacs_branch: str
) -> str:
    go_defuns, go_constants = get_go_declarations(defuns)
    s = FILE_HEADER.format(emacs_commit, emacs_branch) + f"package {PACKAGE}\n\n"

    for constant in go_constants:
        s += f"const {constant.name} = {constant.value}\n"

    if go_constants:
        s += "\n"

    for defun in go_defuns:
        s += f"func (es *emacsStubs) {defun.fnname}("
        if defun.subr_fn == "defSubrM":
            s += "args ...lispObject"
        elif defun.subr_fn != "defSubr0":
            s += ", ".join(defun.args) + " lispObject"
        s += ") (lispObject, error) {\n"
        s += f'\treturn es.stub("{defun.lname}") // Source file: {defun.path}'
        if defun.attributes:
            s += f", attributes: {defun.attributes}"
        s += "\n}\n\n"

    s += f"func (ec *execContext) symbolsOfEmacs{AUTOGEN_SUFFIX}() {{\n"
    s += "\tes := &emacsStubs{ec: ec}\n"
    s += "\tec.stubs = es\n"

    for defun in go_defuns:
        s += f'\tec.{defun.subr_fn}(nil, "{defun.lname}", es.{defun.fnname}'
        if defun.subr_fn != "defSubr0":
            s += f", {defun.minargs}"
        s += ")\n"

    s += "}\n\n"
    s += f"// Subroutines count: {len(go_defuns)}\n"
    s += f"// Constants count: {len(go_constants)}\n"

    return s


def gen_stubs(pimacs_base: Path, emacs_base: Path, emacs_commit: str, emacs_branch: str) -> None:
    all_defuns = []

    for p in sorted(emacs_base.joinpath("src").rglob("*")):
        if not p.suffix == ".c" or p.name == "msdos.c":
            continue

        defuns, *_ = extract_declarations(p)
        all_defuns.extend(sorted(defuns, key=lambda e: e.lname))

    contents = generate_go_file(all_defuns, emacs_commit, emacs_branch)
    target = pimacs_base.joinpath(f"{PACKAGE}/emacs_autogen.go")
    with open(target, "w") as f:
        f.write(contents)

    print("wrote", target)
