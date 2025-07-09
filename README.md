# 📁 file_storage

A **peer-to-peer encrypted file storage system** built in Go.

This project allows nodes to store and retrieve files locally and across a decentralized network of peers. It supports encryption, efficient disk storage, and dynamic file sharing using a custom transport protocol.

---

## 🚀 Features

- 🔐 **AES encryption** for secure file transfers.
- 🧠 **Content-addressable storage** using SHA-1 hashes.
- 📦 **Local disk-based file store** with structured folder hierarchy.
- 🌐 **Peer-to-peer networking** with message broadcasting and streaming.
- 📤 File sharing: store files and automatically share with connected peers.
- 📥 On-demand retrieval: get files from peers if not stored locally.

---

## 🧱 Project Structure

```bash

├── main.go              # Entry point
├── store.go             # Local disk store logic (write/read/delete)
├── crypto.go            # AES encryption/decryption helpers
├── server.go            # File server handling peer communication
├── p2p/                 # P2P transport interface (defined externally)

```

