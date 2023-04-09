<script>  
    let message = "";
  
    function createMakefile() {
      // @ts-ignore
      window.go.backend.App.CreateMakefile().then((/** @type {string} */ result) => {
        message = result;
      });
    }
</script>

<p>Create a Makefile by clicking on this button.</p>
<div id="input" data-wails-no-drag>
  <button class="button border-black border" on:click={createMakefile}>Create Makefile</button>
</div>
{#if message}
  <div id="result">{message}</div>
{/if}