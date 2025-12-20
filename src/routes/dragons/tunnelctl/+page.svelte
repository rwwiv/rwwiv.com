<script lang="ts">
	import Code from "$lib/components/Code.svelte";
	import InlineCode from "$lib/components/InlineCode.svelte";

	const sshpParse = `if [[ "$1" == *+* ]]; then
    local session="\${1%+*}"
    local host="\${1#*+}"
    shift
    ssh -t "$host" "shpool attach '$session' || shpool create '$session'" "$@"
else
    ssh "$@"
fi`;

	const tunnelMeta = `# Load previous values as defaults
if [[ -f "$metafile" ]]; then
    source "$metafile"
    default_host="$host"
    default_local_port="$local_port"
    default_remote_port="$remote_port"
fi`;

	const tunnelPrompt = `host=$(gum input --prompt "Host: " --value "$default_host")
local_port=$(gum input --prompt "Local port: " --value "$default_local_port")
remote_port=$(gum input --prompt "Remote port: " --value "$default_remote_port")`;

	const tunnelStart = `ssh -N -L "\${local_port}:localhost:\${remote_port}" "$host" &
echo "$!" > "$(_tunnel_pidfile "$name")"
echo "host=$host\\nlocal_port=$local_port\\nremote_port=$remote_port" > "$metafile"`;
</script>

<svelte:head>
	<title>DRAGON 4</title>
</svelte:head>

<div class="flex flex-col justify-center items-center px-4">
	<h1 class="text-5xl text-center mb-10">SSH Tunnel Management</h1>

	<div class="px-4 max-w-sm sm:max-w-xl md:max-w-3xl">
		<div id="writeup" class="mb-10 flex flex-col space-y-10">
			<div class="flex flex-col space-y-6">
				<p>
					I do most of my <a
						href="https://github.com/grafana/grafana"
						target="_blank"
						rel="noreferrer">Grafana</a
					> development on a headless VM running on my home server &mdash; always on, available over <a
						href="https://tailscale.com"
						target="_blank"
						rel="noreferrer">Tailscale</a
					> when remote, and the heavy builds run on plugged-in hardware instead of draining laptop battery.
					The workflow is: connect to the VM, start the dev server, tunnel port 3000 back to hit
					<InlineCode code="localhost:3000" language="bash" /> in a local browser.
				</p>
				<p>
					Two small tools make this easier: <InlineCode code="sshp" language="bash" /> for persistent
					shell sessions, and <InlineCode code="tunnelctl" language="bash" /> for managing tunnels by
					name.
				</p>
			</div>

			<hr
				class="w-[150px] h-[1px] rounded-full border-t-0 bg-neutral-900 dark:bg-neutral-100 opacity-100 dark:opacity-50 self-center"
			/>

			<h1 class="font-bold italic">shpool</h1>
			<div class="flex flex-col space-y-6">
				<p>
					For persistent shell sessions, I use <a
						href="https://github.com/shell-pool/shpool"
						target="_blank"
						rel="noreferrer">shpool</a
					> &mdash; to borrow its tagline, think tmux, then aim lower. Sessions survive disconnects,
					so I can close my laptop, reconnect later, and pick up where I left off.
				</p>
				<p>
					To make connecting easier, I've wrapped it in a convenience function. The syntax <InlineCode
						code="sshp grafana+dev-vm"
						language="bash"
					/> connects to host <InlineCode code="dev-vm" language="bash" /> and attaches to (or creates)
					the session named <InlineCode code="grafana" language="bash" />. If no <InlineCode
						code="+"
						language="bash"
					/> is present, it falls through to regular ssh:
				</p>
				<Code language="bash" code={sshpParse} />
			</div>

			<hr
				class="w-[150px] h-[1px] rounded-full border-t-0 bg-neutral-900 dark:bg-neutral-100 opacity-100 dark:opacity-50 self-center"
			/>

			<h1 class="font-bold italic">tunnelctl</h1>
			<div class="flex flex-col space-y-6">
				<p>
					Manages named SSH tunnels with persistent configuration. Running <InlineCode
						code="tunnelctl up grafana"
						language="bash"
					/> prompts for connection details the first time, then remembers them for subsequent runs.
				</p>
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					When starting a tunnel, we first check if a metadata file exists for this name. If so, we
					source it to get the previous values:
				</p>
				<Code language="bash" code={tunnelMeta} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					We then prompt for each parameter using <a
						href="https://github.com/charmbracelet/gum"
						target="_blank"
						rel="noreferrer">gum</a
					>, pre-filling with the defaults. On subsequent runs, hitting enter accepts the previous config:
				</p>
				<Code language="bash" code={tunnelPrompt} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					Finally, we start the tunnel in the background, save the PID for later cleanup, and write
					the current values back to the metadata file:
				</p>
				<Code language="bash" code={tunnelStart} />
			</div>

			<sub class="self-center text-sm">
				<a
					href="https://github.com/rwwiv/.dotfiles-arch/blob/main/.local/bin/tunnelctl"
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					Source on GitHub
				</a>
			</sub>
		</div>
	</div>
</div>

<style lang="postcss">
	#writeup {
		@apply font-serif;
	}

	#writeup h1 {
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
