package iso

import (
	vmwcommon "github.com/hashicorp/packer/builder/vmware/common"
)

type RemoteDriver interface {
	vmwcommon.Driver

	// UploadISO uploads a local ISO to the remote side and returns the
	// new path that should be used in the VMX along with an error if it
	// exists.
	UploadISO(string, string, string) (string, error)

	// Adds a VM to inventory specified by the path to the VMX given.
	Register(string) error

	// Checks if the VM is destroyed.
	IsDestroyed() (bool, error)

	// Uploads a local file to remote side.
	upload(dst, src string) error

	// Reload VM on remote side.
	ReloadVM() error
}
