#!/usr/bin/env bash
set -euo pipefail

if command -v make >/dev/null 2>&1 && command -v nasm >/dev/null 2>&1 && command -v ld >/dev/null 2>&1; then
  make -C codes name=hello
  rm -f codes/hello codes/*.o
else
  printf '%s\n' 'Assembly toolchain unavailable; syntax/build validation skipped.'
fi
printf '%s\n' 'Local validation completed.'
