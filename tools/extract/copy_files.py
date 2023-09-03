from pathlib import Path
import shutil


def copy_files(
    config: dict, pimacs_base: Path, emacs_base: Path, emacs_commit: str, emacs_branch: str
) -> None:
    lisp_files = config["copy_lisp_files"]
    for path in lisp_files:
        origin = Path(emacs_base / "lisp" / path)
        dest = Path(pimacs_base / "lisp" / "emacs" / path)
        dest.parent.mkdir(parents=True, exist_ok=True)

        print(f"copy {origin} to {dest}")
        shutil.copy(origin, dest)

    with open(pimacs_base / "lisp" / "emacs" / "commit.txt", "w") as f:
        f.write(emacs_commit)
        f.write("\n")
