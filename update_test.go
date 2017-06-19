package main_test

import (
	"os"
	"os/exec"
	"path/filepath"

	cli "github.com/heroku/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("update.go", func() {
	Describe("UpdateCLI()", func() {
		It("downloads a new CLI", func() {
			Skip("not working on circle")
			os.Remove(cli.UpdateLockPath)
			manifest := cli.GetUpdateManifest("dev", "")
			Expect(manifest.Channel).To(Equal("dev"))
			dest := filepath.Join("tmp", "newcli")
			cli.DownloadCLI("dev", dest, manifest)
			out, err := exec.Command(filepath.Join(dest, "bin", BinaryName), "version").Output()
			must(err)
			Expect(out).To(HavePrefix(BinaryName + "-cli/"))
		})
	})

	Describe("IsUpdateNeeded()", func() {
		It("disabled with env var", func() {
			os.Remove(cli.AutoupdateFile)
			needsUpdate := cli.IsUpdateNeeded()
			Expect(needsUpdate).To(Equal(true))
			os.Setenv("SFDX_AUTOUPDATE_DISABLE", "true")
			needsUpdate = cli.IsUpdateNeeded()
			Expect(needsUpdate).To(Equal(false))
		})
	})
})
