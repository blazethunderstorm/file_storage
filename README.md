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

<img width="174" height="551" alt="image" src="https://github.com/user-attachments/assets/8befcb8f-87a9-42c8-b739-60dfc8223773" /><br></br>


<img width="1857" height="111" alt="image" src="https://github.com/user-attachments/assets/5777fd50-b49a-450a-aa01-ffc0722326b4" /><br></br>


<img width="1854" height="1116" alt="image" src="https://github.com/user-attachments/assets/975d9e14-ea0f-49cf-bf6b-2be885eeec86" /><br></br>


## 🧱 Project Structure

```bash

├── main.go                                   # Entry point
├── store.go                                  # Local disk store logic (write/read/delete)
├── crypto.go                                 # AES encryption/decryption helpers
├── server.go                                 # File server handling peer communication
├── store_test.go & crypto_test.go            # Test files
├── p2p/                                      # P2P transport interface (defined externally)

```

