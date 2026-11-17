# Go study exercises

Este repositório contém exercícios de estudo e não possui CI intencionalmente. A validação ocorre localmente.

## Validação

Execute `scripts/validate-local.sh` para verificar a formatação Go. Não há cobertura, testes unitários, integração ou E2E configurados para estes exercícios.

## Hooks

Execute `scripts/install-hooks.sh` uma vez. O hook `pre-commit` bloqueia arquivos Go não formatados e `commit-msg` exige Conventional Commits (`feat:`, `fix:` ou `ci:`).
