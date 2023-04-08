<script lang="ts">
	import { enhance } from "$app/forms";
	import type { ActionData, PageData } from "./$types";
    export let form :ActionData;
    import Typewriter from 'svelte-typewriter'
    let titleFinished :boolean = false;
    let contentFinished :boolean = false;
    let AddendumFinished :boolean = false;
    let checked :boolean = false;
</script>
<main>


    <div class="container min-h-[70vh] bg-black text-green-600 rounded-xl relative text-xl p-5 font-bold mx-auto">
        <div class="flex items-center">
            <span>></span>
            <form use:enhance method="POST">
                <input name="query" type="text" class="bg-black outline outline-0 container text-green-600 rounded-xl text-xl p-5 font-bold mx-auto" placeholder="" />
                <button class="display" type="submit"></button>
            </form>
        </div>
        <div class="container space-y-5">
            {#if form}
                {#if checked}
                    <span class="font-bold text-5xl"><Typewriter on:done={()=>{titleFinished = true}} interval={50}>{form?.data.Title}</Typewriter></span>
                    <div class="flex">
                        {#if titleFinished}
                            <Typewriter disabled={!checked} on:done={()=>contentFinished = true} interval={1}>
                                {form?.data.Content}                  
                            </Typewriter>
                        {/if}
                        {#if contentFinished}
                            <Typewriter disabled={!checked} on:done={()=>AddendumFinished = true} interval={1}>
                                {form?.data.Addendum}
                            </Typewriter>
                        {/if}
                        <div class="ml-auto">
                            <img class="w-fit object-cover md:max-w-[10vw] max-w-[50vw]  rounded-2xl border border-zinc-900 hover:border-white duration-300" alt="" src="{form?.data.Image}">
                        </div>
                    </div>
                {:else}
                    <span class="font-bold text-5xl">{form?.data.Title}</span> 
                    <div class="flex">
                            {form?.data.Content}                  
                            {form?.data.Addendum}
                        <div class="ml-auto">
                            <img class="w-fit object-cover md:max-w-[10vw] max-w-[50vw]  rounded-2xl border border-zinc-900 hover:border-white duration-300" alt="" src="{form?.data.Image}">
                        </div>
                    </div>
                {/if}
            {/if}
        </div>

        <div class="flex items-center right-0 p-3 bottom-0 absolute">
            <input on:click={()=>{console.log(checked)}}   bind:checked={checked} id="default-checkbox" type="checkbox"  class="w-4 h-4  text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
            <label  for="default-checkbox" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Typewriting</label>
        </div>
    </div>
    



</main>