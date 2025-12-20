<script lang="ts">
	import Dragon from "$lib/components/Dragon.svelte";

	let grayscaleCode = `func invert(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	imgSet := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			col := color.RGBA{
				255 - uint8(r/256), 255 - uint8(g/256),
				255 - uint8(b/256), uint8(a / 256),
			}
			imgSet.Set(x, y, col)
		}
	}
	return imgSet
}`;

	let screenshotCode = `# Detect monitor colorspace and convert if needed
PRESET=$(hyprctl monitors -j | jq -r '.[] | select(.focused) | .colorManagementPreset')

case "$PRESET" in
    dp3)
        /usr/bin/grim "\${ARGS[@]}" - | magick png:- -set colorspace DisplayP3 -colorspace sRGB png:-
        ;;
    *)
        /usr/bin/grim "\${ARGS[@]}" -
        ;;
esac`;

	let hyprlandIpcCode = `@hyprland.on("monitoraddedv2")
async def handle_monitor_hotplug(ctx: Hyprland, data: str) -> None:
    global wallpaper_task
    if wallpaper_task is not None:
        wallpaper_task.cancel()
    wallpaper_task = asyncio.create_task(delayed_paint())`;

	let tunnelctlCode = `# Start a tunnel interactively (remembers previous values)
$ tunnelctl up grafana
Host: dev-vm
Local port: 3000
Remote port: 3000
tunnel 'grafana': started (pid 12345, localhost:3000 -> dev-vm:3000)`;
</script>

<svelte:head>
	<title>DRAGONS</title>
</svelte:head>

<div class="flex flex-col justify-center items-center mx-auto px-6 space-y-10">
	<div class="w-[350px] sm:w-auto sm:max-w-xl space-y-2">
		<h1 class="text-4xl">
			Development Tests and Learning Exercises <span
				class="text-base italic text-green-700 dark:text-green-200"
				><br class="block md:hidden" />&mdash; Here there be dragons 🐉</span
			>
		</h1>
	</div>
	<div class="flex flex-col space-y-12 w-[350px] sm:w-auto sm:max-w-xl">
		<Dragon
			link="/dragons/image_manipulation"
			title="Image manipulation using serverless Go functions"
			code={grayscaleCode}
			language="go"
		/>
		<Dragon
			link="/dragons/screenshot_tools"
			title="grim wrapper for wide-gamut color space conversion"
			code={screenshotCode}
			language="bash"
		/>
		<Dragon
			link="/dragons/hyprland_ipc"
			title="Hyprland IPC micro-library for Python"
			code={hyprlandIpcCode}
			language="python"
		/>
		<Dragon
			link="/dragons/tunnelctl"
			title="SSH tunnel management for remote development"
			code={tunnelctlCode}
			language="bash"
		/>
	</div>
</div>
