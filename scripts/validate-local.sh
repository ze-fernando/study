#!/usr/bin/env bash
set -euo pipefail

files=$(find . -name '*.go' -type f -not -path './.git/*')
if [[ -n "$files" ]]; then
  unformatted=$(gofmt -l $files)
  if [[ -n "$unformatted" ]]; then
    printf '%s\n' "$unformatted"
    exit 1
  fi
fi
printf '%s\n' 'Local validation passed: Go formatting checked; this study repository has no automated test suite.'
