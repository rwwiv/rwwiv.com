<script lang="ts">
	import Code from "$lib/components/Code.svelte";
	import InlineCode from "$lib/components/InlineCode.svelte";

	const socketPaths = `self.signature = os.getenv("HYPRLAND_INSTANCE_SIGNATURE")
self.runtime_dir = os.getenv("XDG_RUNTIME_DIR")

self.event_sock_path = f"{self.runtime_dir}/hypr/{self.signature}/.socket2.sock"
self.cmd_sock_path = f"{self.runtime_dir}/hypr/{self.signature}/.socket.sock"`;

	const decoratorApi = `def on(self, event_name: str) -> Callable[[EventHandler], EventHandler]:
    def decorator(func: EventHandler) -> EventHandler:
        self.handlers[event_name].append(func)
        return func
    return decorator`;

	const eventParse = `raw = line.decode().strip()
if ">>" in raw:
    event, data = raw.split(">>", 1)
    if event in self.handlers:
        for func in self.handlers[event]:
            asyncio.create_task(func(self, data))`;

	const usageExample = `hyprland = Hyprland()

@hyprland.on("monitoraddedv2")
async def on_monitor_added(ctx: Hyprland, data: str) -> None:
    # data contains monitor info from Hyprland
    await repaint_wallpaper()

hyprland.run()`;
</script>

<svelte:head>
	<title>DRAGON 3</title>
</svelte:head>

<div class="flex flex-col justify-center items-center px-4">
	<h1 class="text-5xl text-center mb-10">Hyprland IPC Micro-Library</h1>

	<div class="px-4 max-w-sm sm:max-w-xl md:max-w-3xl">
		<div id="writeup" class="mb-10 flex flex-col space-y-10">
			<div class="flex flex-col space-y-6">
				<p>
					Hyprland exposes two Unix sockets for IPC: <InlineCode
						code=".socket.sock"
						language="bash"
					/> for sending commands and <InlineCode code=".socket2.sock" language="bash" /> for receiving
					events. The event socket streams compositor events &mdash; workspace changes, window focus,
					monitor hotplug &mdash; as newline-delimited messages that any process can subscribe to.
					This micro-library wraps the event socket with a decorator-based API for Python scripts
					that need to react to compositor state.
				</p>
			</div>

			<h2 class="font-bold italic">Socket discovery</h2>
			<div class="flex flex-col space-y-2">
				<p>
					Hyprland generates a unique instance signature each session, so the socket paths aren't
					static. Both the signature and runtime directory come from environment variables:
				</p>
				<Code language="python" code={socketPaths} />
			</div>

			<h2 class="font-bold italic">Event handling</h2>
			<div class="flex flex-col space-y-2">
				<p>
					The <InlineCode code="@on()" language="python" /> decorator doesn't transform the
					function &mdash; it just appends it to a list keyed by event name. Multiple handlers can
					subscribe to the same event this way:
				</p>
				<Code language="python" code={decoratorApi} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					Events arrive as <InlineCode code="EVENT>>DATA" language="text" /> lines. Each matching
					handler gets spawned as a concurrent task, so one slow handler doesn't block others:
				</p>
				<Code language="python" code={eventParse} />
			</div>

			<h2 class="font-bold italic">Example</h2>
			<div class="flex flex-col space-y-2">
				<p>
					A wallpaper daemon might use this to repaint when monitors are added or removed:
				</p>
				<Code language="python" code={usageExample} />
			</div>

			<sub class="self-center text-sm">
				<a
					href="https://github.com/rwwiv/.dotfiles-arch/blob/main/.config/hypr/scripts/lib/hyprland_ipc.py"
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					Library source
				</a>
				and
				<a
					href="https://github.com/rwwiv/.dotfiles-arch/blob/main/.local/bin/swww-init"
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					example usage
				</a>
			</sub>
		</div>
	</div>
</div>

<style lang="postcss">
	#writeup {
		@apply font-serif;
	}

	#writeup h2 {
		@apply text-2xl;
	}

	#writeup p {
		@apply text-lg leading-relaxed;
	}

	#writeup a {
		@apply underline underline-offset-2;
		@apply dark:hover:text-blue-400 hover:text-blue-700;
	}
</style>
