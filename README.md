# go-webview-svelte
Starter repo that wraps [svelte](https://svelte.dev/) with [webview](https://github.com/webview/webview). Frontend assets are embedded into the go binary using [pkger](https://github.com/markbates/pkger).

## building
Webview has certain depedencies (details in their readme). I've only setup this up on my linux laptop running [elementry OS](https://elementary.io/).

```bash
sudo apt install pkg-config libgtk-3-dev libwebkit2gtk-4.0-dev
```

```bash
make deps build
```

## running
```bash
./go-webview-svelte
```