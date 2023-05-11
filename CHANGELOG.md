# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Option to trim whitespace (by [mrcnski](https://github.com/mrcnski))
- Limit the length of each line being processed.
- We have a Changelog now!

### Fixed

- Installation instructions (by [mrcnski](https://github.com/mrcnski))

## [0.0.3] - 2022-11-12

### Added

- Usage information and README.

### Changed

- We now store `uint64` hashes (produced by `xxHash`) instead of the original strings we want to compare.

## [0.0.2] - 2022-08-18

### Added

- Support for `stdin` filtering.

## [0.0.1] - 2022-08-12

### Added

- Filter the input file and output only the unique lines.
