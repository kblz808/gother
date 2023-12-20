
async function readFileIntoBuffer(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      resolve(reader.result);
    };
    reader.onerror = reject;
    reader.readAsArrayBuffer(file);
  });
}

async function init(){
  const go = new Go();
  const result = await WebAssembly.instantiateStreaming(
    fetch('main.wasm'),
    go.importObject
  );
  go.run(result.instance);

  document.getElementById('imageUpload').addEventListener('change', async function(event){
    const file = event.target.files[0];

    let buffer = await readFileIntoBuffer(file);
    buffer = new Uint8Array(buffer);
    
    decodeImage(buffer)
  })
  
}

init();

function displayImage(encodedImage) {
  const img = document.createElement("img");
  img.src = "data:image/jpeg;base64," + encodedImage;
  document.body.appendChild(img);
}