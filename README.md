# Go-Share

Go-Share is a lightweight local file sharing server that lets you share files between devices on the same network through a simple web interface.

---

## ✅ Features
- 📁 Browse & download files from shared folders
- 🌐 Accessible on LAN (e.g., phone ↔️ PC)
- 📄 Auto-indexing based on folder config
- 🌓 Responsive web UI with clean design
- 📦 Structured as a clean Go monolith

---

## 📦 Usage

### 1. Configure Folders

Update the `config.json` file with the paths you want to expose.

```json
{
  // Choose one of the platform-specific path sections below
  // and adjust them to your actual directories.

  // === Linux Example ===
  "videos": "/home/yourname/Videos",
  "music": "/home/yourname/Music",
  "docs": "/home/yourname/Documents",
  "downloads": "/home/yourname/Downloads"

  // === Windows Example ===
  // "videos": "C:\\Users\\YourName\\Videos",
  // "music": "C:\\Users\\YourName\\Music",
  // "docs": "C:\\Users\\YourName\\Documents",
  // "downloads": "C:\\Users\\YourName\\Downloads"

  // === macOS Example ===
  // "videos": "/Users/yourname/Movies",
  // "music": "/Users/yourname/Music",
  // "docs": "/Users/yourname/Documents",
  // "downloads": "/Users/yourname/Downloads"
}
```

### 2. Run the app

```bash
go run main.go
```

### 3. Connect devices

Ensure your phone/laptop are on the same Wi-Fi network. Open the shown LAN IP in a browser (e.g., http://192.168.1.5:8080).

--- 

## ✅ TODO Progress

- [x] Share files via UI
- [x] Dynamic folder config
- [x] Responsive frontend
- [x] Better UX Styling
- [ ] Proper packaging
- [ ] Share text notes
- [ ] Passcode protection
- [ ] QR code for quick link
- [ ] Drag-and-drop upload
- [ ] Export as desktop/mobile app (using Capacitor or Tauri)
