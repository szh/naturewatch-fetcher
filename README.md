# Hi there! 👋

This is a simple project I put together for fun to allow me to pull photos and
videos from my wildlife camera (built using
[github.com/interactionresearchstudio/NaturewatchCameraServer](https://github.com/interactionresearchstudio/NaturewatchCameraServer),
full tutorial available on their [website](https://mynaturewatch.net/make)) to
my Raspberry Pi NAS.

To use this, you need to connect your wildlife camera to a local WiFi network
and run this Go program on a Raspberry Pi (or any other computer) on the same
network. You can build and run the code yourself, or if you're using a Raspberry
Pi you can download the pre-built binary from the
[releases](https://github.com/szh/naturewatch-fetcher/releases) here on GitHub.

## Configuration

There are a few configuration options you need to set in an `app.env` file in
the same directory as the binary, and you can find an example of this file
[here](app.env). The configuration options are as follows:

- NATUREWATCH_URL The URL of the NatureWatch server on your local network. This
    can be an IP address or a hostname but must include the protocol (generally
    `http://`).

- FETCH_INTERVAL_SECONDS The number of seconds to wait between fetches. If this
    is 0 or negative, the process will exit after the first fetch.

- OUTPUT_PATH The output path on the local filesystem where photos and videos
    will be saved. This path must already exist. However, the `photos` and
    `videos` subdirectories will be created if they do not exist.

## Building

Just run `go build` in the root of the repository to build the binary. You can
then run the binary with `./naturewatch-fetcher`.

## Contributing

If you have any suggestions or improvements, feel free to open an issue or a
pull request.

## License

This poroject is licensed under the GPL-3.0 license. You can find the full text
of the license in the [LICENSE](LICENSE) file.
