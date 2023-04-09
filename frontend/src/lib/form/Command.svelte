<script>
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    
    export let entry;
    
    let prefix = writable(entry.prefix);
    let value = writable(entry.value);
    let title = writable(entry.title);
    let description = writable(entry.description);
    
    onMount(() => {
        prefix.subscribe((value) => {
        entry.prefix = value;
        });
        value.subscribe((value) => {
        entry.value = value;
        });
        title.subscribe((value) => {
        entry.title = value;
        });
        description.subscribe((value) => {
        entry.description = value;
        });
    });
</script>

<div>
    <div class="mb-4">
        <label class="block text-gray-100 text-sm font-bold mb-2" for="prefix">
          Command Prefix
        </label>
        <input bind:value={entry.prefix} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight" id="prefix" type="text" placeholder="npm">
    </div>

    <div class="mb-4">
        <label class="block text-gray-100 text-sm font-bold mb-2" for="value">
            Command Value
        </label>
        <input bind:value={entry.value} class="shadow appearance-none  rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight" id="value" type="text" placeholder="install">
    </div>

    <div class="mb-4">
        <label class="block text-gray-100 text-sm font-bold mb-2" for="title">
            Makefile command title
        </label>
        <input bind:value={entry.title} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight" id="title" type="test" placeholder="front/install">
    </div>
    <div class="mb-4">
        <label class="block text-gray-100 text-sm font-bold mb-2" for="description">
            Makefile command description
        </label>
        <input bind:value={entry.description} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight" id="description" type="test" placeholder="Install front dependancies">
    </div>
</div>