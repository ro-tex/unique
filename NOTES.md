# Release process

1. Update `CHANGELOG.md`.
2. Commit all changes.
3. Tag the release: `git tag -a v0.0.5 -m "GoReleaser and Homebrew."`
4. Push the new tag: `git push origin v0.0.5`
5. Create a new release: `goreleaser release --clean`

// Update the Homebrew formula: `brew bump-formula-pr --strict unique`
// Update the Nix package: `nix-build -A unique`
