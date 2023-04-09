<script>
  import Command from './form/Command.svelte'

  let message = "";

  let entries = [{ prefix: '', value: '', title: '', description: '' }];

  function addEntry() {
    entries = [...entries, { prefix: '', value: '', title: '', description: '' }];
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

<hr class="my-8">

<div class="m-8 py-4">
  {#each entries as entry, i}
    <Command {entry} />
  {/each}
</div>

<button class="button border-white border rounded-md p-2" on:click={addEntry}>Add a command</button>

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
</style>