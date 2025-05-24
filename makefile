APP_NAME=traytimer
VERSION=1.0.0
ARCH=amd64
BUILD_DIR=build
DEB_DIR=$(BUILD_DIR)/$(APP_NAME)_$(VERSION)
BIN_PATH=$(DEB_DIR)/usr/bin
DESKTOP_PATH=$(DEB_DIR)/usr/share/applications
ICON_PATH=$(DEB_DIR)/usr/share/icons/hicolor/64x64/apps

all: clean build package

build:
	@echo "-> Building $(APP_NAME)..."
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)

package:
	@echo "-> Preparing package structure..."
	mkdir -p $(BIN_PATH)
	mkdir -p $(DESKTOP_PATH)
	mkdir -p $(ICON_PATH)
	mkdir -p $(DEB_DIR)/DEBIAN

	# Copy binary
	cp $(APP_NAME) $(BIN_PATH)/

	# Control file
	echo "Package: $(APP_NAME)" > $(DEB_DIR)/DEBIAN/control
	echo "Version: $(VERSION)" >> $(DEB_DIR)/DEBIAN/control
	echo "Section: utils" >> $(DEB_DIR)/DEBIAN/control
	echo "Priority: optional" >> $(DEB_DIR)/DEBIAN/control
	echo "Architecture: $(ARCH)" >> $(DEB_DIR)/DEBIAN/control
	echo "Maintainer: Intekhab Rizvi <me@intekhab.in>" >> $(DEB_DIR)/DEBIAN/control
	echo "Description: A simple system tray timer and stopwatch app." >> $(DEB_DIR)/DEBIAN/control

	# .desktop file
	echo "[Desktop Entry]" > $(DESKTOP_PATH)/$(APP_NAME).desktop
	echo "Name=Tray Timer" >> $(DESKTOP_PATH)/$(APP_NAME).desktop
	echo "Exec=$(APP_NAME) --duration 15m" >> $(DESKTOP_PATH)/$(APP_NAME).desktop
	echo "Icon=$(APP_NAME)" >> $(DESKTOP_PATH)/$(APP_NAME).desktop
	echo "Type=Application" >> $(DESKTOP_PATH)/$(APP_NAME).desktop
	echo "Categories=Utility;" >> $(DESKTOP_PATH)/$(APP_NAME).desktop
	echo "Terminal=false" >> $(DESKTOP_PATH)/$(APP_NAME).desktop
	echo "StartupNotify=false" >> $(DESKTOP_PATH)/$(APP_NAME).desktop

	# Icon (replace with your PNG name if needed)
	cp icon.png $(ICON_PATH)/$(APP_NAME).png

	@echo "-> Building .deb..."
	dpkg-deb --build $(DEB_DIR)

	@echo "-> Done: $(BUILD_DIR)/$(APP_NAME)_$(VERSION).deb"

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(APP_NAME)
