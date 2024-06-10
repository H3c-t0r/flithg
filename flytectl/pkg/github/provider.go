package github

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	go_github "github.com/google/go-github/v42/github"
	"github.com/mouuff/go-rocket-update/pkg/provider"
)

// type GitHubProvider struct {
// 	provider.Github
// }

// Github provider finds a archive file in the repository's releases to provide files
type GitHubProvider struct {
	RepositoryURL string // Repository URL, example github.com/mouuff/go-rocket-update
	ArchiveName   string // Archive name (the zip/tar.gz you upload for a release on github), example: binaries.zip

	tmpDir             string            // temporary directory this is used internally
	decompressProvider provider.Provider // provider used to decompress the downloaded archive
	archivePath        string            // path to the downloaded archive (should be in tmpDir)
}

// githubTag struct used to unmarshal response from github
// https://api.github.com/repos/ownerName/projectName/tags
type githubTag struct {
	Name string `json:"name"`
}

// githubRepositoryInfo is used to get the name of the project and the owner name
// from this fields we are able to get other links (such as the release and tags link)
type githubRepositoryInfo struct {
	RepositoryOwner string
	RepositoryName  string
}

// getRepositoryInfo parses the github repository URL
func (c *GitHubProvider) repositoryInfo() (*githubRepositoryInfo, error) {
	re := regexp.MustCompile(`github\.com/(.*?)/(.*?)$`)
	submatches := re.FindAllStringSubmatch(c.RepositoryURL, 1)
	if len(submatches) < 1 {
		return nil, errors.New("Invalid github URL:" + c.RepositoryURL)
	}
	return &githubRepositoryInfo{
		RepositoryOwner: submatches[0][1],
		RepositoryName:  submatches[0][2],
	}, nil
}

// getArchiveURL get the archive URL for the github repository
// If no tag is provided then the latest version is selected
func (c *GitHubProvider) getArchiveURL(tag string) (string, error) {
	if len(tag) == 0 {
		// Get latest version if no tag is provided
		var err error
		tag, err = c.GetLatestVersion()
		if err != nil {
			return "", err
		}
	}

	info, err := c.repositoryInfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s",
		info.RepositoryOwner,
		info.RepositoryName,
		tag,
		c.ArchiveName,
	), nil
}

// Open opens the provider
func (c *GitHubProvider) Open() (err error) {
	archiveURL, err := c.getArchiveURL("") // get archive url for latest version
	if err != nil {
		return
	}
	resp, err := http.Get(archiveURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	c.tmpDir, err = os.MkdirTemp("", "rocket-update")
	if err != nil {
		return
	}

	c.archivePath = filepath.Join(c.tmpDir, c.ArchiveName)
	archiveFile, err := os.Create(c.archivePath)
	if err != nil {
		return
	}
	_, err = io.Copy(archiveFile, resp.Body)
	archiveFile.Close()
	if err != nil {
		return
	}
	c.decompressProvider, err = provider.Decompress(c.archivePath)
	if err != nil {
		return nil
	}
	return c.decompressProvider.Open()
}

// Close closes the provider
func (c *GitHubProvider) Close() error {
	if c.decompressProvider != nil {
		c.decompressProvider.Close()
		c.decompressProvider = nil
	}

	if len(c.tmpDir) > 0 {
		os.RemoveAll(c.tmpDir)
		c.tmpDir = ""
		c.archivePath = ""
	}
	return nil
}

// GetLatestVersion gets the latest version
func (c *GitHubProvider) GetLatestVersion() (string, error) {
	tags, err := c.getReleases()
	if err != nil {
		return "", err
	}
	latest_tag := tags[0].GetTagName()
	return latest_tag, err
}

func (c *GitHubProvider) getReleases() ([]*go_github.RepositoryRelease, error) {
	g := GetGHRepoService()
	releases, _, err := g.ListReleases(context.Background(), owner, flyte, &go_github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}
	var filteredReleases []*go_github.RepositoryRelease
	for _, release := range releases {
		if strings.HasPrefix(release.GetTagName(), flytectl) {
			filteredReleases = append(filteredReleases, release)
		}
	}
	return filteredReleases, err
}

// Walk walks all the files provided
func (c *GitHubProvider) Walk(walkFn provider.WalkFunc) error {
	if c.decompressProvider == nil {
		// TODO specify error
		return provider.ErrNotOpenned
	}
	return c.decompressProvider.Walk(walkFn)
}

// Retrieve file relative to "provider" to destination
func (c *GitHubProvider) Retrieve(src string, dest string) error {
	return c.decompressProvider.Retrieve(src, dest)
}
