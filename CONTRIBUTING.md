# Contributing to Titan

Thanks for your interest in contributing! This guide explains how to open a pull request.

## Opening a pull request

1. **Fork** the repository and clone your fork locally.
2. **Create a branch** from `master` with a descriptive name, e.g. `fix-rpop-empty-list`.
3. **Make your change.** Keep pull requests small and focused on a single concern.
4. **Build and test** locally before pushing:

   ```sh
   make build
   make test
   make lint
   ```

5. **Commit** with a clear message describing what changed and why.
6. **Push** your branch to your fork and open a pull request against `master`.
7. **Describe your change** in the pull request: what it does, why it is needed, and how you tested it. Link any related issues (e.g. `Fixes #123`).
8. **Respond to review feedback** by pushing additional commits to the same branch.

## Reporting issues

If you find a bug or want to propose a feature, please [open an issue](https://github.com/distributedio/titan/issues) first so it can be discussed before you start work.
