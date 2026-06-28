# Veil

> **Linux-native OverlayFS workspace manager for isolated development environments.**

Veil is a Linux-native command-line tool that orchestrates **OverlayFS**, **copy-on-write filesystem semantics**, and **inotify-based filesystem observation** to provision isolated development workspaces. Instead of replicating project directories, Veil composes a writable workspace by mounting an immutable project directory as the lower layer and redirecting runtime modifications into an isolated writable layer.

The project serves as a practical implementation of Linux filesystem primitives while demonstrating production-style service architecture, filesystem orchestration, and event-driven monitoring. This approach minimizes disk utilization, preserves the integrity of the original project, and demonstrates practical usage of Linux filesystem primitives.

---

# Installation


Install Veil:

```bash
curl -fsSL https://raw.githubusercontent.com/lakshya-8000cr/veil/main/install.sh | bash
```

<a href="https://asciinema.org/a/1259398" target="_blank"><img src="https://asciinema.org/a/1259398.svg" /></a>


Navigate to the project you want to isolate.

```bash
cd ~/projects/my-project
```

Create a workspace.

```bash
veil spawn
```

Mount the workspace.

```bash
veil mount my-project
```

Open the merged workspace.

```bash
code ~/.veil/workspaces/my-project/merged
```

---

# Commands

```bash
veil spawn
veil mount <workspace>
veil unmount <workspace>
veil apply <workspace>
veil destroy <workspace>
veil inspect <workspace>
veil list
veil watch <workspace>
```

---

# Architecture

```text
                    Original Project
                 (Immutable Lower Layer)
                           │
                           ▼
                    Linux OverlayFS
            ┌──────────────┴──────────────┐
            │                             │
       Upper Layer                  Work Layer
   (Copy-on-Write Data)       (Kernel Metadata)
            │                             │
            └──────────────┬──────────────┘
                           ▼
                    Merged Workspace
                     (Developer View)
                           │
                           ▼
                        Veil CLI
```

---

# Workspace Lifecycle

```text
spawn
   │
   ▼
Workspace Metadata Initialization
   │
   ▼
mount
   │
   ▼
OverlayFS Composition
   │
   ▼
Developer Modifications
   │
   ▼
Copy-on-Write Redirection
   │
   ▼
apply
   │
   ▼
Streaming Workspace Synchronization
   │
   ▼
destroy
```

---

# Service-Oriented Architecture

Veil intentionally separates workspace orchestration from Linux-specific implementation details through dedicated service layers.

```text
                 CLI Commands
                      │
                      ▼
              Workspace Service
                      │
        ┌─────────────┼─────────────┐
        ▼             ▼             ▼
 Overlay Service   Filesystem    Watch Service
                   Service
        │             │             │
        └─────────────┼─────────────┘
                      ▼
                 Linux Kernel
```

Each layer owns a single responsibility.

* **Workspace Service** – lifecycle orchestration and workspace metadata management.
* **Overlay Service** – OverlayFS mount composition and teardown.
* **Filesystem Service** – streaming synchronization and filesystem operations.
* **Watch Service** – event-driven filesystem observation using Linux inotify.

This separation keeps business logic independent from kernel interaction while improving maintainability and extensibility.

---

# Engineering Highlights

### OverlayFS-Based Workspace Composition

Rather than duplicating project directories, Veil composes isolated workspaces using OverlayFS. The original project is mounted as the immutable lower layer while all runtime mutations are redirected into a writable upper layer through Linux copy-on-write semantics.

---

### Metadata-Driven Workspace Reconstruction

Workspace state is persisted through lightweight metadata, allowing lifecycle operations to reconstruct environments without relying on shell session state or process-local memory.

---

### Canonical Path Resolution

Workspace metadata stores normalized absolute paths instead of relative paths, ensuring deterministic workspace reconstruction regardless of the caller's working directory.

---

### Streaming File Synchronization

Workspace synchronization is implemented using streamed I/O (`io.Copy`) instead of loading entire files into memory.

This minimizes memory consumption while allowing synchronization to scale efficiently with large project assets.

---

### Event-Driven Filesystem Observation

Workspace monitoring is implemented using Linux **inotify** (via `fsnotify`) with recursive directory registration, allowing Veil to observe filesystem activity across nested project structures.

---

### Layered Abstraction

Linux-specific functionality remains encapsulated behind dedicated service boundaries, allowing workspace management logic to evolve independently from OverlayFS and filesystem implementations.

---

# Performance Characteristics

### Traditional Workspace Duplication

```text
Project
   │
Copy Entire Directory
   │
Workspace

Disk Usage
O(Project Size)

Provision Time
Dependent on Project Size
```

### Veil Workspace Provisioning

```text
Project
   │
OverlayFS Mount
   │
Workspace

Disk Usage
Only Modified Files

Provision Time
Metadata Initialization + OverlayFS Mount
```

---

# Linux Technologies

* OverlayFS
* Linux Virtual Filesystem (VFS)
* Copy-on-Write (CoW)
* Mount / Unmount
* inotify
* Recursive Filesystem Monitoring
* Streaming File I/O
* Go
* Cobra CLI

---

# License

MIT License
