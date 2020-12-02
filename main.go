package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/mod/semver"
)

const ()

func cdIntoRepo(repo string) error {
	if repo == "" {
		return nil
	}

	err := os.Chdir(repo)
	if err != nil {
		return fmt.Errorf("changing working directory '%s': %w", repo, err)
	}
	return nil
}

func checkoutVersion(v string) error {
	if !semver.IsValid(v) {
		return fmt.Errorf("invalid semver, for example v1, v1.2, v1.2.3, or v1.2.3-pre")
	}

	major := semver.Major(v) == v
	majorminor := semver.MajorMinor(v) == v

	cmd := exec.Command("git", "tag")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("`git tag` failed: %w", err)
	}

	match := false
	k := "v0.0.0"
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		x := scanner.Text()
		if !semver.IsValid(x) {
			continue
		}
		if major {
			if semver.Major(v) != semver.Major(x) {
				continue
			}
		} else if majorminor {
			if semver.MajorMinor(v) != semver.MajorMinor(x) {
				continue
			}
		} else if v != x {
			continue
		}

		if semver.Compare(k, x) < 0 {
			match = true
			k = x
		}
	}

	if !match {
		return fmt.Errorf("version not found")
	}

	cmd = exec.Command("git", "checkout", k)
	out, err = cmd.CombinedOutput()
	fmt.Printf("%s", out)
	if err != nil {
		return fmt.Errorf("git checkout %s: %w", k, err)
	}

	return nil
}

func mainx() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("incorrect number of arguments given")
	}

	argParts := strings.SplitN(os.Args[1], "@", -1)
	if len(argParts) != 2 {
		return fmt.Errorf("invalid argument: %s", os.Args[1])
	}

	err := cdIntoRepo(argParts[0])
	if err != nil {
		return err
	}

	err = checkoutVersion(argParts[1])
	if err != nil {
		return fmt.Errorf("getting version '%s': %w", argParts[1], err)
	}

	return nil
}

func main() {
	err := mainx()
	if err != nil {
		fmt.Printf("Usage: getversion <git-repo>@v<semver>\n")
		fmt.Printf("Version: %s, github.com/shortmoose/getversion\n\n", version)
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
