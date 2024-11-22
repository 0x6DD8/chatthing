# Chatthing

Chatthing is a reimplementation of the [Chatding](https://github.com/MoeDevelops/chatding) project. It is a simple chat application built using Go and HTMX, with server-sent events (SSE) for real-time updates and a web-based frontend.

## Features

- Real-time chat updates using Server-Sent Events (SSE)
- Session management with cookies
- Written in Go

## Getting Started

### Prerequisites

- Go 1.23.2 or later

### Installation

1. Clone the repository:

```sh
git clone https://github.com/0x6DD8/chatthing.git
cd chatthing
```

2. Build the project:

```sh
make all
```

3. Run the server:

```sh
./target/chatthing_linux # ./target/chatthing.exe on Windows
```

The server will start on port 5000.

### Usage

Open your web browser and navigate to `http://localhost:5000` to start using the chat application.

## License

This project is licensed under the Unlicense. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Chatding](https://github.com/MoeDevelops/chatding) - The original project that inspired this reimplementation.
- [gomponents](https://www.gomponents.com/) - A Go library for building HTML components.
- [htmx](https://htmx.org) - A library for accessing modern browser features directly from HTML.
