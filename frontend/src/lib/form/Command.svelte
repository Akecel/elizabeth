<script>
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    
    export let entry;
    export let index;
    
    let commands = writable(entry.commands);
    let title = writable(entry.title);
    let description = writable(entry.description);

    function addCommand() {
        commands.update((commands) => [...commands, { command: '' }]);
    }

    onMount(() => {
        commands.subscribe((value) => {
            entry.commands = value;
        });
        title.subscribe((value) => {
            entry.title = value;
        });
        description.subscribe((value) => {
            entry.description = value;
        });
    });
</script>

<h2>Makefile command {index + 1}</h2>

<div class="flex flex-row space-x-4 mb-4">
    <div class="basis-1/3">
        <label class="block text-gray-100 text-sm font-bold mb-2" for="title">
            Makefile command title
        </label>
        <input bind:value={entry.title} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight" id="title" type="test" placeholder="front/install">
    </div>
    <div class="basis-2/3">
        <label class="block text-gray-100 text-sm font-bold mb-2" for="description">
            Makefile command description
        </label>
        <input bind:value={entry.description} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight" id="description" type="test" placeholder="Install front dependancies">
    </div>
</div>

{#each entry.commands as command, index}
    <div class="mb-4">
        <label class="block text-gray-100 text-sm font-bold mb-2" for="command">
            Command {index + 1}:
        </label>
        <input bind:value={command.command} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight" id="command" type="text" placeholder="npm install">
    </div>
{/each}

<button class="button border-white border rounded-md p-2" on:click={addCommand}>Add new command</button>