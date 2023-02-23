<script>
    import Console from "../components/Console.svelte";
    
    const onKeyPress = e => {
        if (e.charCode === 13) addInput();
    };

    function addInput()
    {
    var inputValue = document.getElementById("inputVal").value;

    var consoleDiv = document.getElementById("konsoleDiv")
    var console = document.getElementById("konsole");
    var li = document.createElement("li");

    // inputValue    const cmdArgs = inputValue.split(' ');
    // var command = cmdArgs[0];

    switch(inputValue)
    {
        case "clear":
        console.innerHTML = null;
            li.appendChild(document.createTextNode(inputValue));
            break;

        default:
            console.innerHTML += "Instruction: " +  inputValue + " not found.";     
            break;
    }
    
    console?.appendChild(li);


    
    consoleDiv.scrollTop = consoleDiv.scrollHeight;
}
</script>
<div class="flex justify-center mt-3">
    <img class="h-52 mt-5 opacity-60" src="https://upload.wikimedia.org/wikipedia/commons/2/2a/SCP_Foundation_Logo_v2_Fixed.png">
</div>
<div class="m-6">
    <div class="flex mt-5 max-h-64 h-64 overflow-x-scroll mb-3" id="consoleDiv">
        <ul id="console">
        </ul> 
     </div>
     <div>
         <input id="inputVal" on:keypress={onKeyPress} type="text" placeholder="Type here" class="input input-bordered input-sm w-full hover:border-green-500 focus:border-green-500" style="color:lime;" />
     </div>
</div>
<style>
    #console
    {
        font-family: monospace;
        color:rgb(0, 255, 0);    
    }
</style>
