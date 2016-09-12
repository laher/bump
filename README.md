# bump

Bump a given version up to the next number.

 * Handles semver, amongst others.
 * Basically, any version numbering scheme where:
   * There's a delimiter between each part (1.2.3 -> 1.2.4)
   * There's a numeric component in each part
 * Handles bumping major/minor/patch (1.2.3 -> 1.3.0)
 * Passes over prefixes (v1.2.3 -> v1.2.4)
 * Handles suffixes (bumps to the next version to ensure new version is higher than old) (v1.2.3-rc -> v1.2.4)
 * Handles alternate delimiters. Default is '.' (1-rc1 -> 1-rc2)


## Installation

	**install go, including 'set GOPATH'**
	go install github.com/laher/bump/...

## Usage

### Basic usage

	$ bump v1.0.1
 	v1.0.2

	$ bump -part=1 v1.0.1
 	v1.1.0

	$ bump -part=1 v1.0.1-prerelease
 	v1.1.0

	$ bump -part=0 -ltr 1.0.1-prerelease
 	2.0.0

### e.g. Work with external apps, e.g. git tags

	$ git describe --tags --abbrev=0
	v-1.2.3
	$ bump `git describe --tags --abbrev=0`
	v-1.2.4

### e.g. Use pipes

	$ cat version.txt
	1.2.2
	$ cat version.txt | bump -stdin
	1.2.3
	$ cat version.txt | bump -stdin > version.txt

### e.g. Alternative delimiter
	
	$ bump -delimiter=- v1.0.1-hotfix1
	v1.0.1-hotfix2
