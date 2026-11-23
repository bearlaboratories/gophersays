all:
	tinygo flash -target=esp32-coreboard-v2 ./

monitor:
	tinygo flash -target=esp32-coreboard-v2 -monitor ./
