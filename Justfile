default:
	@just --list

install-tools:
	build/install-tools

check:
	build/check-go
