# LitGit 📦

Welcome to **LitGit**, a fun and casual command-line tool to interact with Git repositories. It’s built to make your Git experience smoother and add some "vibes" to your workflow. 🚀

Документация на русском языке здесь : [Документация](./docs/RU_DOCS.md)

---

## What is LitGit? 🤔

LitGit is a set of commands that lets you easily interact with Git, from basic operations like cloning and pushing to more fun commands like "yeet" and "yoink"! Whether you're a developer or just someone who wants to mess with Git in a cool way, LitGit’s got you covered. 🙌

### Features:
- **Clone a repo** with `yoink`.
- **Push your changes** with `yeet`.
- **Check your Git status** with `vibes`.
- **Reset your commit** with `cap`.
- **And much more!** ⚡

This is not your average Git CLI; it's built to add some flair to your dev life! 🌟

---

## Commands ⚙️

Here’s a list of all the cool commands you can use:

- `yuh (or yh) snatch [url]` – Clone a repository from the given URL.
- `yh yeet` – Push your changes to the remote repo.
- `yh yoink` – Pull the latest changes from the remote.
- `yh vibes` – Check the working tree status.
- `yh cap [commit_hash]` – Reset your commit to a specific commit hash.
- `yh slay` – Stage your files (git add).
- `yh bet` – Commit your changes.
- `yh rizz` – Initialize a Git repo in your current directory.

---

## How to Use LitGit? 🛠️

1. Clone the repo:

    ```bash
    git clone https://github.com/yourusername/litGit.git
    ```

2. Install the dependencies (make sure you have Go installed):

    ```bash
    go mod tidy
    ```

3. Build the script to make the binaries:

    ```bash
    ./makeScript.sh
    ```

----
## Or Using the Prebuilt Binary
1. Go to the [Releases](https://github.com/shuklarituparn/litGit/releases) page.

- Download the binary suitable for your operating system:
    - **Linux**: `yh-linux-amd64`
    - **Windows**: `yh-windows-amd64.exe`
    - **macOS**: `yh-darwin-amd64`


- After downloading, move the binary to a directory in your system's `PATH`:
    - **Linux/macOS**:
      ```bash
      sudo mv yh-linux-amd64 /usr/local/bin/yh    # For Linux
      sudo mv yh-darwin-amd64 /usr/local/bin/yh    # For macOS
      chmod +x /usr/local/bin/litGit
      ```
    - **Windows**:
        1. Rename `yh-windows-amd64.exe` to a shorter name if needed.
        2. Add the directory containing `yh-windows-amd64.exe` to your `PATH` environment variable.
- Run it from anywhere using:


   ```bash
      yh <command>
   ```


---

## License 📝

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.

