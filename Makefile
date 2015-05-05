DEBUG_FLAG = $(if $(DEBUG),-debug)

upload:
	./scripts/upload.sh
