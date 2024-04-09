package main

import (
	"github.com/dop251/goja"

	"github.com/dop251/goja_nodejs/require"
)

func main() {
	registry := new(require.Registry)
	runtime := goja.New()
	req := registry.Enable(runtime)
	runtime.RunString(`
        var m = require("./spot.js");
        m.test();
    `)

	m, err := req.Require("./spot.js")
	_, _ = m, err
	runtime.RunString(`
	window.onSpotifyWebPlaybackSDKReady = () => {
		const token = 'BQCaPy-sf8SxRbJtGmh25mLy1vuaXR451gko6okDNYzDgEl4H_ww8vHxK4IzsMs1YXfyiVzKo8i6SHLw8rvwRbFg2XTZP73L1inklfeSawCKD93WR-FXOriAtvBOfbsc_Myd41MML2g10jwUuAN5tpS6WhyyepP8_3gY-dolpH0aOZRURO-Hb7WezueHt_IPzRwOkQF5';
		const player = new Spotify.Player({
			name: 'playback',
			getOAuthToken: cb => {cb(token); },
			volume: 0.5
	});

	player.addListener('ready', ({ device_id }) => {
		console.log('Ready with Device ID', device_id);
	});

	player.addListener('not_ready', ({ device_id }) => {
		console.log('Device ID has gone offline', device_id);
	});

	player.addListener('initialization_error', ({ message }) => {
		console.error(message);
	});

	player.addListener('authentication_error', ({ message }) => {
		console.error(message);
	});

	player.addListener('account_error', ({ message }) => {
		console.error(message);
	});

	document.getElementById('togglePlay').onclick = function() {
	  player.togglePlay();
	};

	player.connect();
`)
}
