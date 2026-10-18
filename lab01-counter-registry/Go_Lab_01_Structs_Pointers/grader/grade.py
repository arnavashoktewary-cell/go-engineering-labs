#!/usr/bin/env python3
"""Public functional autograder for Go Lab 01 (Python standard library only)."""
import json
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
SRC = ROOT / "src"
TESTS = [
    "TestCounter_ConstructorValid",
    "TestCounter_ConstructorInvalid",
    "TestCounter_Increment",
    "TestCounter_Add",
    "TestCounter_Snapshot",
    "TestRegistry_Register",
    "TestRegistry_Increment",
    "TestRegistry_GetSnapshot",
    "TestRegistry_Names",
    "TestRegistry_Total",
]


def main():
    if not (SRC / "go.mod").is_file():
        print("\nMissing src/go.mod. First initialize YOUR module:")
        print("  cd src && go mod init example.com/lab01")
        return 2
    try:
        p = subprocess.run(
            ["go", "test", "-json", "-count=1", "-run", "^(TestCounter_|TestRegistry_)", "."],
            cwd=SRC, text=True, capture_output=True, check=False,
        )
    except FileNotFoundError:
        print("Go is not on PATH. Check 'go version' first.")
        return 2

    status = {}
    diagnostics = {name: [] for name in TESTS}
    for line in p.stdout.splitlines():
        try:
            event = json.loads(line)
        except json.JSONDecodeError:
            continue
        name = event.get("Test")
        action = event.get("Action")
        if name in status or name in TESTS:
            if action in ("pass", "fail", "skip"):
                status[name] = action
            elif action == "output" and name in diagnostics:
                msg = event.get("Output", "").rstrip()
                if msg and not msg.lstrip().startswith(("=== RUN", "--- PASS", "--- FAIL")):
                    diagnostics[name].append(msg)

    print("\nGO LAB 01 | PUBLIC AUTOGRADER")
    print("=" * 56)
    score = 0
    for name in TESTS:
        action = status.get(name, "not run")
        passed = action == "pass"
        score += 10 if passed else 0
        label = "PASS" if passed else "FAIL"
        print(f"{label:4} | {name:<31} | {'10' if passed else ' 0'}/10")
        if not passed and diagnostics.get(name):
            for msg in diagnostics[name][:2]:
                print("       ", msg.strip()[:150])
    print("-" * 56)
    print(f"FUNCTIONAL SCORE: {score}/100")
    print("Note: local public tests only; your own tests and code review are separate.")
    if p.returncode != 0 and not status:
        print("\nCompilation/test execution failed:")
        print((p.stderr or p.stdout)[-2500:])
    elif p.stderr.strip():
        print("\nCompiler/test messages:\n", p.stderr[-1000:])
    return 0 if score == 100 and p.returncode == 0 else 1


if __name__ == "__main__":
    sys.exit(main())
