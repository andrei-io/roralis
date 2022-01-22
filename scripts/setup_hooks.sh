#!/bin/sh

cd `git rev-parse --show-toplevel`
rm .git/hooks/pre-commit
ln ./scripts/pre-commit .git/hooks