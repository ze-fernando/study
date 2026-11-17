# Assembly study notes

Este repositório contém material de estudo e não possui CI intencionalmente. A validação ocorre localmente.

## Validação

Execute `scripts/validate-local.sh`. Quando `make`, `nasm` e `ld` estiverem disponíveis, o exercício em `codes` é montado e linkado. Não há cobertura, testes unitários, integração ou E2E configurados.

## Hooks

Execute `scripts/install-hooks.sh` uma vez. O hook `pre-commit` executa a validação local e `commit-msg` exige Conventional Commits (`feat:`, `fix:` ou `ci:`).
