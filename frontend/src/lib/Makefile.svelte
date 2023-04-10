<script>
  import Command from './form/Command.svelte'

  let message = "";

  let entries = [{title: '', description: '', commands: [{ command: '' }] }];

  function addEntry() {
    entries = [...entries, {title: '', description: '', commands: [{ command: '' }] }];
  }

  function generateMakefile() {
    // @ts-ignore
    window.go.backend.App.GenerateMakefile(entries).then((/** @type {string} */ result) => {
      message = result;
    });
  }
</script>

<h1>Makefile Generator</h1>
<p>Create complete Makefiles in a few clicks.</p>

<hr class="mt-8">

<div class="pb-8">
  {#each entries as entry, i}
    <Command {entry}/>
  {/each}
</div>

<button class="button border-white border rounded-md p-2" on:click={addEntry}>Add a Makefile command</button>

<hr class="my-8">

<button class="button border-white border rounded-md p-2" on:click={generateMakefile}>Generate your Makefile</button>

{#if message}
  <div class="my-2" id="result">{message}</div>
{/if}

<style>
  h1 {
    font-size: 2em;
    font-weight: 600;
    margin: 0;
  }

  button:hover {
    transition: 0.2s;
    background-color: #FFFFFF;
    color: #0C0D1D;
  }
</style>