import argparse
import json
import subprocess
from pathlib import Path

import stubs

PIMACS_REPO_PATH = (Path(__file__) / "../../..").resolve()
CONFIG_FILE = "config.json"


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument('--emacs-repo-path', required=True)
    subparsers = parser.add_subparsers(required=True, dest="subparser")

    subparsers.add_parser("gen-stubs", help="Generate Emacs subroutine/var stubs")
    subparsers.add_parser("copy", help="Copy selected Emacs files")

    args = parser.parse_args()

    with open(CONFIG_FILE) as f:
        config = json.load(f)

    emacs_commit = subprocess.check_output(
        "git rev-parse HEAD", text=True, shell=True, cwd=Path(args.emacs_repo_path)
    ).strip()
    emacs_branch = subprocess.check_output(
        "git branch --show-current", text=True, shell=True, cwd=Path(args.emacs_repo_path)
    ).strip()

    if config["emacs_commit"] != emacs_commit:
        raise Exception(
            f"Target Emacs revision does not match repository revision: "
            f'want {config["emacs_commit"]}, have {emacs_commit}'
        )
    if config["emacs_branch"] != emacs_branch:
        raise Exception(
            f"Target Emacs branch does not match repository branch: "
            f'want \'{config["emacs_branch"]}\', have \'{emacs_branch}\''
        )

    match args.subparser:
        case "gen-stubs":
            stubs.gen_stubs(
                PIMACS_REPO_PATH, Path(args.emacs_repo_path), emacs_commit, emacs_branch
            )
        case "copy":
            ...
        case _:
            raise Exception("Unknown command")


if __name__ == "__main__":
    main()
