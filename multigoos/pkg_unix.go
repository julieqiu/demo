// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package multigoos

func CloseOnExec(num int) (int, error) {
	return num, nil
}
