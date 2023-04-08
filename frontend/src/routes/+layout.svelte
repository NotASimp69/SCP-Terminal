<script>
  import "../app.css";
  import Navbar from "../components/Navbar.svelte";
</script>

<main class="bg-zinc-800 min-h-screen space-y-10">
  <Navbar />  
  <slot />
</main>