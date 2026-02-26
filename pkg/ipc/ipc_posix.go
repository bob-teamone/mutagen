//go:build !windows

package ipc

import (
	"context"
	"net"
	"syscall"
)

// DialContext attempts to establish an IPC connection, timing out if the
// provided context expires.
func DialContext(ctx context.Context, path string) (net.Conn, error) {
	// Create a zero-valued dialer, which will have the same dialing behavior as
	// the raw dialing functions.
	dialer := &net.Dialer{}

	// Perform dialing.
	return dialer.DialContext(ctx, "unix", path)
}

// NewListener creates a new IPC listener.
func NewListener(path string) (net.Listener, error) {
	// Narrow the umask to 0177 before creating the socket so that the kernel
	// creates it with mode 0600 (owner-only read/write) from the outset.
	// This eliminates the TOCTOU window that would otherwise exist between
	// net.Listen and an os.Chmod call. syscall.Umask is process-wide, not
	// goroutine-local, but the critical section is a single syscall, so the
	// window during which the umask is narrowed is effectively instantaneous.
	oldUmask := syscall.Umask(0177)
	listener, err := net.Listen("unix", path)
	syscall.Umask(oldUmask)
	if err != nil {
		return nil, err
	}

	// Success.
	return listener, nil
}
